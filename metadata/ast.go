package metadata

import (
	"fmt"
	"go/ast"
	"sort"
	"strings"
)

// This takes a function declaration and returns the owner and Service
// the Service summarises the inputs and outputs
func analyseFunc(fun *ast.FuncDecl, messages map[string]*Message) (owner string, srv *Service) {
	if !isMethodValid(fun) {
		return
	}

	srv = &Service{
		Name:       fun.Name.String(),
		InputNames: make([]string, 0),
		InputTypes: make([]string, 0),
		Output:     make([]string, 0),
		IsMethod:   fun.Recv != nil && len(fun.Recv.List) > 0,
	}
	//log.Printf("analyseFunc %s, isMethod %t\n", srv.Name, srv.IsMethod)
	// context is the first parameter and DB is the second parameter
	// TD: Handle the parameters on the function
	for i := 0; i < len(fun.Type.Params.List); i++ {
		p := fun.Type.Params.List[i]
		if exprToStr(p.Type) == "context.Context" {
			srv.HasContext = true
			continue
		}
		if exprToStr(p.Type) == "DB" {
			continue
		}

		for _, n := range p.Names {
			srv.InputNames = append(srv.InputNames, n.Name)
			srv.InputTypes = append(srv.InputTypes, adjustType(exprToStr(p.Type), messages))
		}
	}

	// error is the last result
	// Handle where the results go
	for i := 0; i < len(fun.Type.Results.List)-1; i++ {
		p := fun.Type.Results.List[i]
		srv.Output = append(srv.Output, adjustType(exprToStr(p.Type), messages))
	}

	owner = getOwner(fun)
	// TD The definition of a method is it has a receiver name
	// TD So the attributes of the receiver have to be unpacked
	if srv.IsMethod {
		receiverName := fun.Recv.List[0].Names[0].Name
		methodAttributes := make([]string, 0)
		deDup := make(map[string]struct{})
		// TD Only pick up the attributes that are used in the function to simplify the proto message
		for _, stmt := range fun.Body.List {
			ast.Inspect(stmt, func(n ast.Node) bool {
				if selector, ok := n.(*ast.SelectorExpr); ok && fmt.Sprintf("%s", selector.X) == receiverName && firstIsUpper(selector.Sel.Name) {
					name := selector.Sel.Name
					//fmt.Sprintf("Name:%s\n", name)
					if _, ok := deDup[name]; !ok {
						methodAttributes = append(methodAttributes, name)
						deDup[name] = struct{}{}
					}
					return false
				}
				// Ignore passing by reference
				if _, ok := n.(*ast.UnaryExpr); ok {
					return false
				}
				return true
			})
		}

		receiverType := strings.TrimPrefix(exprToStr(fun.Recv.List[0].Type), "*")
		receiver := messages[receiverType]
		// TD: Special code for specific methods
		switch srv.Name {
		case "Delete":
			receiver.PkNames = make([]string, 0)
			receiver.PkNames = append(receiver.PkNames, methodAttributes...)
		case "Insert": // TD:Return Receiver so ID is known
			//	log.Printf("Input %v", fun.Recv.List[0].Type)
			srv.ReturnReceiver = true
			//srv.HasCustomOutput()
			srv.Output = append(srv.Output, fmt.Sprintf("*%s", receiverType))
			fmt.Println(receiverName, receiverName, receiverType)
			//if receiverName == "ur" {
			//	fmt.Println(srv.Output)
			//	panic(1)
			//}
			//srv.Output = append([]string{adjustType(exprToStr(fun.Recv.List[0].Type), messages)}, srv.Output...)
		}
		sort.Strings(methodAttributes)
		methodTypes := make([]string, 0)
		for _, n := range methodAttributes {
			methodTypes = append(methodTypes, receiver.attributeTypeByName(n))
		}

		if srv.RelationshipMethod() {
			for _, attr := range methodAttributes {
				receiver.IndexNames[attr] = struct{}{}
			}
		}

		srv.InputMethodNames = methodAttributes
		srv.InputMethodTypes = methodTypes
	}
	srv.Owner = owner
	fmt.Println(fun.Name.String(), srv.Output)

	return
}

func parseMessages(pkg *ast.Package) map[string]*Message {
	messages := make(map[string]*Message)
	for _, file := range pkg.Files {
		if file.Scope != nil {
			for name, obj := range file.Scope.Objects {
				if isErrType(name) {
					continue
				}
				if typ, ok := obj.Decl.(*ast.TypeSpec); ok {
					switch t := typ.Type.(type) {
					case *ast.Ident:
						messages[name] = createAliasMessage(name, t)
					case *ast.StructType:
						messages[name] = createStructMessage(name, t)
					case *ast.ArrayType:
						messages[name] = createArrayMessage(name, t)
					}
				}
			}
		}
	}
	for _, file := range pkg.Files {
		for _, n := range file.Decls {
			if fun, ok := n.(*ast.FuncDecl); ok {
				if r, ok := isTextUnmarshaler(fun); ok {
					r = strings.TrimPrefix(r, "*")
					if m, ok := messages[r]; ok {
						m.HasTextUnmarshaler = true
					}
				} else if r, ok := isParseFromString(fun); ok {
					r = strings.TrimPrefix(r, "*")
					if m, ok := messages[r]; ok {
						m.HasParser = true
					}
				}
			}
		}
	}
	for _, m := range messages {
		m.IndexNames = make(map[string]struct{})
		m.adjustType(messages)
	}
	return messages
}

func getOwner(fun *ast.FuncDecl) string {
	if fun.Recv != nil && len(fun.Recv.List) > 0 {
		return canonicalType(exprToStr(fun.Recv.List[0].Type))
	}

	if len(fun.Type.Results.List) > 1 {
		typ := canonicalType(exprToStr(fun.Type.Results.List[0].Type))
		if firstIsUpper(canonicalType(typ)) {
			return typ
		}
	}
	return "Services"
}

func isMethodValid(fun *ast.FuncDecl) bool {
	if fun.Name == nil {
		return false
	}

	if !fun.Name.IsExported() || fun.Name.Name == "Save" {
		return false
	}

	if fun.Type.Params == nil || len(fun.Type.Params.List) == 0 ||
		fun.Type.Results == nil || len(fun.Type.Results.List) == 0 {
		return false
	}

	if exprToStr(fun.Type.Results.List[len(fun.Type.Results.List)-1].Type) != "error" {
		return false
	}

	for _, param := range fun.Type.Params.List {
		if t, ok := param.Type.(*ast.Ident); ok && t.Name == "DB" {
			return true
		}
	}

	return false
}

func isTextUnmarshaler(fun *ast.FuncDecl) (receiver string, ok bool) {
	if fun.Name.Name != "UnmarshalText" {
		return
	}

	if fun.Recv == nil || len(fun.Recv.List) != 1 ||
		fun.Type.Params == nil || len(fun.Type.Params.List) != 1 ||
		fun.Type.Results == nil || len(fun.Type.Results.List) != 1 {
		return
	}

	if exprToStr(fun.Type.Params.List[0].Type) != "[]byte" {
		return
	}

	if exprToStr(fun.Type.Results.List[0].Type) != "error" {
		return
	}

	receiver = exprToStr(fun.Recv.List[0].Type)

	if !strings.HasPrefix(receiver, "*") {
		return
	}

	return receiver, true
}

func isParseFromString(fun *ast.FuncDecl) (receiver string, ok bool) {
	if fun.Name.Name != "Parse" {
		return
	}

	if fun.Recv == nil || len(fun.Recv.List) != 1 ||
		fun.Type.Params == nil || len(fun.Type.Params.List) != 1 ||
		fun.Type.Results == nil || len(fun.Type.Results.List) != 1 {
		return
	}

	if exprToStr(fun.Type.Params.List[0].Type) != "string" {
		return
	}

	if exprToStr(fun.Type.Results.List[0].Type) != "error" {
		return
	}

	receiver = exprToStr(fun.Recv.List[0].Type)

	if !strings.HasPrefix(receiver, "*") {
		return
	}

	return receiver, true
}

func canonicalType(typ string) string {
	return strings.TrimPrefix(strings.TrimPrefix(typ, "[]"), "*")
}

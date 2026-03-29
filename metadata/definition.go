package metadata

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"sort"
	"strings"
)

type Definition struct {
	GoModule             string
	DatabaseDriverModule string
	DatabaseDriverName   string
	ModelsPath           string
	Packages             []*Package
	IgnorePackages       map[string]struct{}
}

func (d *Definition) ProtoImports() []string {
	r := make([]string, 0)
	if d.importTimestamp() {
		r = append(r, `import "google/protobuf/timestamp.proto";`)
	}
	if d.importWrappers() {
		r = append(r, `import "google/protobuf/wrappers.proto";`)
	}
	return r
}

func (d *Definition) Messages() map[string]*Message {
	if len(d.Packages) > 0 {
		return d.Packages[0].Messages
	}
	return nil
}

func (d *Definition) importTimestamp() bool {
	for _, m := range d.Messages() {
		if m.importTimestamp() {
			return true
		}
	}

	return false
}

func (d *Definition) importWrappers() bool {
	for _, m := range d.Messages() {
		if m.importWrappers() {
			return true
		}
	}

	return false
}

type Package struct {
	Package    string
	GoModule   string
	SrcPath    string
	SrcPackage string
	Services   []*Service
	Messages   map[string]*Message
}

func (p *Package) ProtoImports() []string {
	r := make([]string, 0)
	if p.importEmpty() {
		r = append(r, `import "google/protobuf/empty.proto";`)
	}
	if p.importTimestamp() {
		r = append(r, `import "google/protobuf/timestamp.proto";`)
	}
	if p.importWrappers() {
		r = append(r, `import "google/protobuf/wrappers.proto";`)
	}
	if p.importTypes() {
		r = append(r, `import "typespb.proto";`)
	}
	return r
}

func (p *Package) importEmpty() bool {
	for _, s := range p.Services {
		if s.EmptyInput() || s.EmptyOutput() {
			return true
		}
	}
	return false
}

func (p *Package) importTimestamp() bool {
	for _, s := range p.Services {
		for _, n := range s.InputTypes {
			if n == "time.Time" || strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.InputMethodTypes {
			if n == "time.Time" || strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.Output {
			if n == "time.Time" || strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
	}
	return false
}

func (p *Package) importWrappers() bool {
	for _, s := range p.Services {
		for _, n := range s.InputTypes {
			if strings.HasPrefix(n, "sql.Null") && !strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.InputMethodTypes {
			if strings.HasPrefix(n, "sql.Null") && !strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.Output {
			if strings.HasPrefix(n, "sql.Null") && !strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
	}
	return false
}

func (p *Package) importTypes() bool {
	for _, s := range p.Services {
		if s.HasCustomParams() || s.HasCustomOutput() {
			return true
		}
	}
	return false
}

// ParsePackages finds all the packages to process
func ParsePackages(src, module string, ignorePackages map[string]struct{}) ([]*Package, error) {
	fset := token.NewFileSet()
	// First of all build an Abstract Syntax Tree for all the files in the directory - should be one package
	pkgs, err := parser.ParseDir(fset, src, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	for ignorePkg := range ignorePackages {
		delete(pkgs, ignorePkg)
	}
	if total := len(pkgs); total != 1 {
		return nil, fmt.Errorf("too many packages: %d", total)
	}

	var pkgName string
	var pkg *ast.Package
	// TD: Get the package name from the pkgs main, we only have one or error above
	for pkgName, pkg = range pkgs {
		break
	}

	messages := parseMessages(pkg)

	owners := make(map[string][]*Service)
	for _, file := range pkg.Files {
		//log.Println("file.Name", file.Name)
		// Analyse the declarations
		skipOwner := false
		for _, n := range file.Decls {
			if fun, ok := n.(*ast.FuncDecl); ok {
				//p := getOwner(fun)
				//// Skip packages we don't want to publish an API for
				//if _, inMap := ignorePackages[p]; inMap {
				//	continue
				//}
				//log.Println("Function:", fun.Name.Name)
				owner, srv := analyseFunc(fun, messages)
				if _, skipOwner = ignorePackages[owner]; skipOwner {
					fmt.Println("Ignoring", owner)
					break
				}

				if owner == "Services" {
					log.Println("Found Services, stream?", fun.Name.Name, "Ignoring")
					//break
					//log.Println(fun.Name.Name)
					//continue
					//os.Exit(1)
				}
				if srv != nil && owner != "Services" {
					srv.Messages = messages
					if _, ok := owners[owner]; !ok {
						owners[owner] = make([]*Service, 0)
					}
					owners[owner] = append(owners[owner], srv)
				}
			}
		}
	}
	var result []*Package
	for owner, services := range owners {
		//log.Printf("owner:%s\n ", owner)
		// TD 29.11.22 - remove strings.Compare([i].Name,[j].Name) < 0 from sort
		sort.SliceStable(services, func(i, j int) bool {
			return services[i].Name < services[j].Name
		})
		p := Package{
			Package:    owner,
			SrcPath:    src,
			SrcPackage: pkgName,
			GoModule:   module,
			Messages:   messages,
			Services:   services,
		}
		//log.Printf(fmt.Sprintf("Owner: %s, Package:%s, srcPath:%s, Messages:%v, Services :%v", owner, p.Package, p.SrcPath, p.Messages, p.Services))
		setReaderEntity(&p)
		result = append(result, &p)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return strings.Compare(result[i].Package, result[j].Package) < 0
	})
	return result, nil
}

// setReaderEntity
func setReaderEntity(p *Package) {
	m, ok := p.Messages[p.Package]
	if !ok {
		return
	}
	//log.Printf("setReaderEntity Package %s", p.Package)
	for _, s := range p.Services {
		if len(s.Output) != 1 { // TD:Ignore services with no output
			continue
		}
		if m.Name != strings.TrimPrefix(s.Output[0], "*") { // Ignore services that aren't a struct (e.g. *Datum)
			continue
		}
		if !strings.HasPrefix(s.Name, p.Package+"By") { // services that are getters e.g. DatumBy...
			continue
		}
		if len(m.PkNames) != len(s.InputNames) {
			continue
		}
		var incompatibleInterface bool
		//log.Println("Incompatible Interface Check", m.PkNames, s.InputNames)
		for _, pk := range m.PkNames {
			pkLower := strings.ToLower(pk)
			var found bool
			for _, in := range s.InputNames {
				if strings.ToLower(in) == pkLower {
					found = true
					break
				}
			}
			if !found {
				incompatibleInterface = true
				break
			}
		}
		if incompatibleInterface {
			//log.Printf("**Incompatible Interface Package:%s, Service:%s Output:%s", p.Package, s.Name, s.Output[0])
			continue
		}
		//log.Printf("Compatible Interface Package:%s, Service:%s Output:%s", p.Package, s.Name, s.Output[0])
		m.ReaderService = s
		return
	}
}

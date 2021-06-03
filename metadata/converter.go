package metadata

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
	"unicode"
)

func exprToStr(e ast.Expr) string {
	switch exp := e.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprToStr(exp.X), exp.Sel.Name)
	case *ast.Ident:
		return exp.String()
	case *ast.StarExpr:
		return "*" + exprToStr(exp.X)
	case *ast.ArrayType:
		return "[]" + exprToStr(exp.Elt)
	default:
		panic(fmt.Sprintf("invalid type %T", exp))
	}
}

func toProtoType(typ string) string {
	if strings.HasPrefix(typ, "*") {
		return toProtoType(typ[1:])
	}
	if strings.HasPrefix(typ, "[]") && typ != "[]byte" {
		return "repeated " + toProtoType(typ[2:])
	}
	switch typ {
	case "json.RawMessage", "[]byte":
		return "bytes"
	case "sql.NullBool":
		return ".google.protobuf.BoolValue"
	case "sql.NullInt32":
		return ".google.protobuf.Int32Value"
	case "int":
		return "int64"
	case "int16":
		return "int32"
	case "uint16":
		return "uint32"
	case "sql.NullInt64":
		return ".google.protobuf.Int64Value"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "sql.NullFloat64":
		return ".google.protobuf.DoubleValue"
	case "sql.NullString":
		return ".google.protobuf.StringValue"
	case "sql.NullTime", "time.Time", "pq.NullTime", "mysql.NullTime", "xoutil.SqTime":
		return ".google.protobuf.Timestamp"
	case "uuid.UUID", "net.HardwareAddr", "net.IP":
		return "string"
	default:
		if firstIsUpper(typ) {
			return "typespb." + typ
		}

		if strings.Contains(typ, "encoding.TextUnmarshaler") {
			return "string"
		}
		return typ
	}
}

func UpperFirstCharacter(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return str
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

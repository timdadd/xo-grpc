package metadata

import (
	"fmt"
	"strings"
)

const (
	textUnmarshalerTypePrefix = "encoding.TextUnmarshaler."
	parserTypePrefix          = "encoding.Parser."
)

type Service struct {
	Owner            string
	Name             string
	InputNames       []string
	InputTypes       []string
	Output           []string
	IsMethod         bool
	InputMethodNames []string
	InputMethodTypes []string
	HasContext       bool

	Messages map[string]*Message
}

func (s *Service) SimplePK() string {
	if m, ok := s.Messages[s.Owner]; ok && len(m.PkNames) == 1 {
		return m.PkNames[0]
	}
	return ""
}

func (s *Service) PKJoin(sep string) string {
	if m, ok := s.Messages[s.Owner]; ok {
		return strings.Join(m.PkNames, sep)
	}
	return ""
}

func (s *Service) PKParams(prefix string) string {
	if m, ok := s.Messages[s.Owner]; ok {
		params := make([]string, len(m.PkNames))
		for i, n := range m.PkNames {
			switch m.AttributeTypeByName(n) {
			case "int":
				params[i] = "int(" + prefix + n + ")"
			case "int16":
				params[i] = "int16(" + prefix + n + ")"
			default:
				params[i] = prefix + n
			}
		}
		return strings.Join(params, ", ")
	}
	return ""
}

func (s *Service) MethodInputType() string {
	switch {
	case s.EmptyInput():
		return "emptypb.Empty"
	default:
		return fmt.Sprintf("pb.%sRequest", s.Name)
	}
}

func (s *Service) MethodOutputType() string {
	switch {
	case s.EmptyOutput():
		return "emptypb.Empty"
	case s.HasCustomOutput():
		return fmt.Sprintf("typespb.%s", strings.TrimPrefix(s.Output[0], "*"))
	default:
		return fmt.Sprintf("pb.%sResponse", s.Name)
	}
}

func (s *Service) ReturnCallDatabase() string {
	if !s.EmptyOutput() {
		return "result,"
	}
	return ""
}

func (s *Service) ParamsCallDatabase() string {
	if s.EmptyInput() {
		return ""
	}
	return ", " + strings.Join(s.InputNames, ", ")
}

func (s *Service) InputGrpc() []string {
	res := make([]string, 0)
	if s.EmptyInput() {
		return res
	}

	for i, attr := range s.InputMethodNames {
		res = append(res, bindToGo("req", fmt.Sprintf("m.%s", attr), attr, s.InputMethodTypes[i], false)...)
	}

	if s.HasCustomParams() {
		typ := s.InputTypes[0]
		in := s.InputNames[0]
		res = append(res, fmt.Sprintf("var %s %s", in, typ))
		m := s.Messages[typ]
		for i, name := range m.AttrNames {
			attrName := UpperFirstCharacter(name)
			res = append(res, bindToGo("req", fmt.Sprintf("%s.%s", in, attrName), attrName, m.AttrTypes[i], false)...)
		}
	} else {
		for i, n := range s.InputNames {
			res = append(res, bindToGo("req", n, UpperFirstCharacter(n), s.InputTypes[i], true)...)
		}
	}

	return res
}

func (s *Service) OutputGrpc() []string {
	res := make([]string, 0)

	if s.EmptyOutput() {
		return res
	}

	if s.HasArrayOutput() {
		res = append(res, "for _, r := range result {")
		typ := strings.TrimPrefix(strings.TrimPrefix(s.Output[0], "[]"), "*")
		res = append(res, fmt.Sprintf("var item typespb.%s", typ))
		m := s.Messages[typ]
		for i, attr := range m.AttrNames {
			res = append(res, bindToProto("r", "item", UpperFirstCharacter(attr), m.AttrTypes[i])...)
		}
		res = append(res, "res.Value = append(res.Value, &item)")
		res = append(res, "}")
		return res
	}

	if s.HasCustomOutput() {
		for _, n := range s.Output {
			m := s.Messages[strings.TrimPrefix(n, "*")]
			for i, attr := range m.AttrNames {
				res = append(res, bindToProto("result", "res", UpperFirstCharacter(attr), m.AttrTypes[i])...)
			}
		}
		return res
	}
	if !s.EmptyOutput() {
		res = append(res, "res.Value = result")
		return res
	}

	return res
}

func (s *Service) RpcSignature() string {
	var b strings.Builder
	b.WriteString(s.Name)
	b.WriteString("(")
	switch {
	case s.EmptyInput():
		b.WriteString("google.protobuf.Empty")
	default:
		b.WriteString(fmt.Sprintf("%sRequest", s.Name))
	}
	b.WriteString(") returns (")
	switch {
	case s.EmptyOutput():
		b.WriteString("google.protobuf.Empty")
	case s.HasCustomOutput():
		b.WriteString("typespb." + strings.TrimPrefix(s.Output[0], "*"))
	default:
		b.WriteString(fmt.Sprintf("%sResponse", s.Name))
	}
	b.WriteString(")")
	return b.String()
}

func (s *Service) HasCustomParams() bool {
	if s.EmptyInput() || len(s.InputTypes) == 0 {
		return false
	}

	return customType(s.InputTypes[0])
}

func (s *Service) HasArrayParams() bool {
	if s.EmptyInput() || len(s.InputTypes) == 0 {
		return false
	}

	return strings.HasPrefix(s.InputTypes[0], "[]") && s.InputTypes[0] != "[]byte"
}

func (s *Service) HasCustomOutput() bool {
	if s.EmptyOutput() {
		return false
	}

	return customType(s.Output[0])
}

func (s *Service) HasArrayOutput() bool {
	if s.EmptyOutput() {
		return false
	}
	return strings.HasPrefix(s.Output[0], "[]") && s.Output[0] != "[]byte"
}

func (s *Service) ProtoInputs() string {
	var pk string
	if s.IsMethod && s.Name == "Update" {
		pk = s.SimplePK()
	}
	var b strings.Builder
	var count int
	for i, name := range s.InputNames {
		count = count + 1
		fmt.Fprintf(&b, "\n    %s %s = %d;", toProtoType(s.InputTypes[i]), name, count)
	}
	for i, name := range s.InputMethodNames {
		count = count + 1
		if pk == name {
			fmt.Fprint(&b, "\n    // Output only.")
		}
		fmt.Fprintf(&b, "\n    %s %s = %d;", toProtoType(s.InputMethodTypes[i]), name, count)
	}
	return b.String()
}

func (s *Service) EmptyInput() bool {
	if len(s.InputTypes) > 0 || len(s.InputMethodTypes) > 0 {
		return false
	}
	return true
}

func (s *Service) EmptyOutput() bool {
	return len(s.Output) == 0
}

func (s *Service) ProtoOutputs() string {
	var b strings.Builder
	for i, name := range s.Output {
		fmt.Fprintf(&b, "    %s value = %d;\n", toProtoType(name), i+1)

	}
	return b.String()
}

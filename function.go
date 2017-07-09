package jot

import (
	"fmt"
	"io"
	"strings"
)

type FunctionSpec struct {
	Doc         string
	Name        string
	Code        []string
	ReturnTypes []TypeSpec
	Parameters  []*ParameterSpec
}

type ParameterSpec struct {
	Name string
	Typ  TypeSpec
}

func (s *ParameterSpec) Write(w io.Writer) error {
	io.WriteString(w, fmt.Sprintf("%s ", s.Name))
	s.Typ.Write(w)

	return nil
}

func Function(name string) *FunctionSpec {
	return &FunctionSpec{Name: name}
}

func (s *FunctionSpec) SetDoc(doc string) *FunctionSpec {
	s.Doc = doc
	return s
}

func (s *FunctionSpec) AddCode(code string) *FunctionSpec {
	s.Code = append(s.Code, code)
	return s
}

func (s *FunctionSpec) AddCodeFmt(code string, args ...interface{}) *FunctionSpec {
	for i, arg := range args {
		code = strings.Replace(code, fmt.Sprintf("{%d}", i), fmt.Sprint(arg), -1)
	}
	s.AddCode(code)
	return s
}

func (s *FunctionSpec) AddParameter(name string, typ TypeSpec) *FunctionSpec {
	s.Parameters = append(s.Parameters, &ParameterSpec{Name: name, Typ: typ})
	return s
}

func (s *FunctionSpec) AddReturnType(typ TypeSpec) *FunctionSpec {
	if typ == nil {
		panic("nil typ")
	}
	s.ReturnTypes = append(s.ReturnTypes, typ)
	return s
}

func (s *FunctionSpec) Write(w io.Writer) error {
	WriteDoc(w, s.Doc)
	io.WriteString(w, fmt.Sprintf("func %s(", s.Name))
	s.writeParameters(w)
	io.WriteString(w, ")")
	s.writeReturnTypes(w)
	io.WriteString(w, " {")
	io.WriteString(w, "\n")
	s.writeCode(w)
	io.WriteString(w, "}")
	io.WriteString(w, "\n")
	return nil
}

func (s *FunctionSpec) writeCode(w io.Writer) error {
	for _, c := range s.Code {
		io.WriteString(w, c)
		io.WriteString(w, "\n")
	}

	return nil
}

func (s *FunctionSpec) writeSignature(w io.Writer) error {
	io.WriteString(w, fmt.Sprintf("%s(", s.Name))
	s.writeParameters(w)
	io.WriteString(w, ")")
	s.writeReturnTypes(w)

	return nil
}

func (s *FunctionSpec) writeParameters(w io.Writer) error {
	for i, p := range s.Parameters {
		if i > 0 {
			io.WriteString(w, ", ")
		}
		p.Write(w)
	}

	return nil
}

func (g *FunctionSpec) writeReturnTypes(w io.Writer) error {
	if len(g.ReturnTypes) > 0 {
		io.WriteString(w, " ")
	}
	if len(g.ReturnTypes) > 1 {
		io.WriteString(w, "(")
	}
	for i, r := range g.ReturnTypes {
		if i > 0 {
			io.WriteString(w, ", ")
		}
		r.Write(w)
	}
	if len(g.ReturnTypes) > 1 {
		io.WriteString(w, ")")
	}

	return nil
}

package jot

import (
	"fmt"
	"io"
)

type InterfaceSpec struct {
	Doc     string
	Name    string
	Methods []*FunctionSpec
}

func Interface(name string) *InterfaceSpec {
	return &InterfaceSpec{Name: name}
}

func (s *InterfaceSpec) AddMethod(fn *FunctionSpec) *InterfaceSpec {
	s.Methods = append(s.Methods, fn)
	return s
}

func (s *InterfaceSpec) SetDoc(doc string) *InterfaceSpec {
	s.Doc = doc
	return s
}

func (s *InterfaceSpec) Write(w io.Writer) error {
	WriteDoc(w, s.Doc)
	io.WriteString(w, fmt.Sprintf("type %s interface {", s.Name))
	if len(s.Methods) > 0 {
		io.WriteString(w, "\n")
		for _, f := range s.Methods {
			io.WriteString(w, fmt.Sprintf("%s(", f.Name))
			f.writeParameters(w)
			io.WriteString(w, fmt.Sprintf(")"))
			f.writeReturnTypes(w)
			io.WriteString(w, "\n")
		}
	}
	io.WriteString(w, "}")
	io.WriteString(w, "\n")

	return nil
}

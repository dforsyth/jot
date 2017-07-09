package jot

import (
	"fmt"
	"io"
)

type VariableSpec struct {
	Doc  string
	Name string
	Typ  TypeSpec
}

func Variable(name string, typ TypeSpec) *VariableSpec {
	return &VariableSpec{Name: name, Typ: typ}
}

func (s *VariableSpec) SetDoc(doc string) *VariableSpec {
	s.Doc = doc
	return s
}

func (s *VariableSpec) Write(w io.Writer) error {
	WriteDoc(w, s.Doc)
	io.WriteString(w, fmt.Sprintf("var %s ", s.Name))
	s.Typ.Write(w)

	return nil
}

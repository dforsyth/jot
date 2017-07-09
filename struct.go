package jot

import (
	"fmt"
	"io"
)

type StructSpec struct {
	Doc     string
	Name    string
	Fields  []*FieldSpec
	Methods []*MethodSpec
}

type FieldSpec struct {
	Doc  string
	Name string
	Typ  TypeSpec
	Tags map[string]string
}

type MethodSpec struct {
	// TODO: Change the way methods are created to support individual docs?
	Func     *FunctionSpec
	RecvName string
	Ptr      bool
}

func Field(name string, typ TypeSpec) *FieldSpec {
	return &FieldSpec{Name: name, Typ: typ, Tags: make(map[string]string)}
}

func (s *FieldSpec) SetTag(key, value string) *FieldSpec {
	s.Tags[key] = value
	return s
}

func (s *FieldSpec) SetDoc(doc string) *FieldSpec {
	s.Doc = doc
	return s
}

func (s *FieldSpec) Write(w io.Writer) error {
	WriteDoc(w, s.Doc)
	io.WriteString(w, fmt.Sprintf("%s ", s.Name))
	s.Typ.Write(w)
	if len(s.Tags) > 0 {
		tags := "`"
		for key, value := range s.Tags {
			if len(tags) > len("`") {
				tags += " "
			}
			tags += key + ":\"" + value + "\""
		}
		tags += "`"
		io.WriteString(w, " "+tags)
	}

	return nil
}

func Struct(name string) *StructSpec {
	return &StructSpec{Name: name}
}

func (s *StructSpec) AddField(f *FieldSpec) *StructSpec {
	s.Fields = append(s.Fields, f)
	return s
}

func (s *StructSpec) AddMethod(f *FunctionSpec) *StructSpec {
	s.Methods = append(s.Methods, &MethodSpec{Func: f})
	return s
}

func (s *StructSpec) AddMethodRecv(recvName string, ptr bool, f *FunctionSpec) *StructSpec {
	s.Methods = append(s.Methods, &MethodSpec{Func: f, RecvName: recvName, Ptr: ptr})
	return s
}

func (s *StructSpec) SetDoc(doc string) *StructSpec {
	s.Doc = doc
	return s
}

func (s *StructSpec) TypeSpec() TypeSpec {
	return &BaseTypeSpec{name: s.Name}
}

func (s *StructSpec) Write(w io.Writer) error {
	WriteDoc(w, s.Doc)

	io.WriteString(w, fmt.Sprintf("type %s struct {", s.Name))
	if len(s.Fields) > 0 {
		io.WriteString(w, "\n")
		for _, f := range s.Fields {
			f.Write(w)
			io.WriteString(w, "\n")
		}
	}
	io.WriteString(w, "}")

	io.WriteString(w, "\n")

	if len(s.Methods) > 0 {
		for i, m := range s.Methods {
			if i > 0 {
				io.WriteString(w, "\n")
			}
			io.WriteString(w, "func (")

			if m.RecvName != "" {
				io.WriteString(w, fmt.Sprintf("%s ", m.RecvName))
			}

			typ := s.TypeSpec()
			if m.Ptr {
				typ = Ptr(typ)
			}
			typ.Write(w)
			io.WriteString(w, ") ")

			m.Func.writeSignature(w)
			io.WriteString(w, " {")
			io.WriteString(w, "\n")
			m.Func.writeCode(w)
			io.WriteString(w, "}")
			io.WriteString(w, "\n")
		}
	}

	return nil
}

// write struct

// write methods

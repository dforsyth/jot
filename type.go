package jot

import (
	"fmt"
	"go/ast"
	"go/types"
	"io"
	"reflect"
)

type BaseTypeSpec struct {
	name string
	imp  *ImportSpec
}

type TypeSpec interface {
	Spec
	Name() string
	Import() *ImportSpec
	From(*ImportSpec) TypeSpec
}

type PtrSpec struct {
	BaseTypeSpec
	Ptr TypeSpec
}

func (s *PtrSpec) Write(w io.Writer) error {
	if _, err := io.WriteString(w, "*"); err != nil {
		return err
	}
	return s.Ptr.Write(w)
}

type ArraySpec struct {
	BaseTypeSpec
	Elt TypeSpec
}

func (s *ArraySpec) Write(w io.Writer) error {
	if _, err := io.WriteString(w, "[]"); err != nil {
		return err
	}
	return s.Elt.Write(w)
}

type MapSpec struct {
	BaseTypeSpec
	Key   TypeSpec
	Value TypeSpec
}

func (s *MapSpec) Write(w io.Writer) error {
	if _, err := io.WriteString(w, "map["); err != nil {
		return err
	}
	if err := s.Key.Write(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "]"); err != nil {
		return err
	}
	return s.Value.Write(w)
}

type ChanDir int

const (
	None ChanDir = iota
	Send ChanDir = iota
	Recv ChanDir = iota
)

type ChanSpec struct {
	BaseTypeSpec
	dir   ChanDir
	value TypeSpec
}

func (s *ChanSpec) Write(w io.Writer) error {
	if s.dir == Send {
		if _, err := io.WriteString(w, "<-"); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "chan"); err != nil {
		return err
	}
	if s.dir == Recv {
		if _, err := io.WriteString(w, "<-"); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, " "); err != nil {
		return err
	}
	return s.value.Write(w)
}

type FuncSpec struct {
	BaseTypeSpec
	params  []TypeSpec
	returns []TypeSpec
}

func (s *FuncSpec) Write(w io.Writer) error {
	io.WriteString(w, "func(")
	for _, p := range s.params {
		p.Write(w)
	}
	io.WriteString(w, ")")
	if len(s.returns) > 0 {
		io.WriteString(w, " ")
		if len(s.returns) > 1 {
			io.WriteString(w, "(")
		}

		for _, r := range s.returns {
			r.Write(w)
		}

		if len(s.returns) > 1 {
			io.WriteString(w, ")")
		}
	}
	return nil
}

func Ptr(typ TypeSpec) *PtrSpec {
	return &PtrSpec{Ptr: typ}
}

func Array(typ TypeSpec) *ArraySpec {
	return &ArraySpec{Elt: typ}
}

func Map(key TypeSpec, value TypeSpec) *MapSpec {
	return &MapSpec{Key: key, Value: value}
}

func Chan(value TypeSpec) *ChanSpec {
	return &ChanSpec{value: value}
}

func SendChan(value TypeSpec) *ChanSpec {
	return &ChanSpec{dir: Send, value: value}
}

func RecvChan(value TypeSpec) *ChanSpec {
	return &ChanSpec{dir: Recv, value: value}
}

func Func() *FuncSpec {
	return &FuncSpec{}
}

func (s *FuncSpec) AddReturnType(typ TypeSpec) *FuncSpec {
	s.returns = append(s.returns, typ)
	return s
}

func (s *FuncSpec) AddParameter(typ TypeSpec) *FuncSpec {
	s.params = append(s.params, typ)
	return s
}

func Type(typ reflect.Type) TypeSpec {
	return &BaseTypeSpec{name: typ.String()}
}

func TypeASTObject(obj *ast.Object) TypeSpec {
	return &BaseTypeSpec{name: obj.Name}
}

func TypeTypesObject(obj types.Object) TypeSpec {
	return &BaseTypeSpec{name: obj.Type().String()}
}

func TypeASTSelector(sel *ast.SelectorExpr) TypeSpec {
	// XXX: LOL, hope so!
	x := sel.X.(*ast.Ident).Name
	f := sel.Sel.Name
	return TypeString(fmt.Sprintf("%s.%s", x, f))
}

func TypeString(name string) TypeSpec {
	return &BaseTypeSpec{name: name}
}

func (s *BaseTypeSpec) From(imp *ImportSpec) TypeSpec {
	s.imp = imp
	return s
}

func (s *BaseTypeSpec) Import() *ImportSpec {
	return s.imp
}

func (s *BaseTypeSpec) Name() string {
	return s.name
}

func (s *BaseTypeSpec) Write(w io.Writer) error {
	_, err := io.WriteString(w, s.Name())
	return err
}

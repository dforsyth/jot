package jot

import (
	"bytes"
	"fmt"
	"golang.org/x/tools/imports"
	"io"
)

type FileSpec struct {
	PackageName string
	Imports     []*ImportSpec

	specs []Spec
}

type ConstSpec struct {
	Name string
	Expr string
}

func (s *ConstSpec) Write(w io.Writer) error {
	_, err := io.WriteString(w, fmt.Sprintf("const %s = %s\n", s.Name, s.Expr))
	return err
}

func File(packageName string) *FileSpec {
	return &FileSpec{PackageName: packageName}
}

func (s *FileSpec) AddImport(imp *ImportSpec) *FileSpec {
	s.Imports = append(s.Imports, imp)
	return s
}

func (s *FileSpec) AddSpec(spec Spec) *FileSpec {
	s.specs = append(s.specs, spec)
	return s
}

func (s *FileSpec) AddInterface(iface *InterfaceSpec) *FileSpec {
	s.specs = append(s.specs, iface)
	return s
}

func (s *FileSpec) AddStruct(strct *StructSpec) *FileSpec {
	s.specs = append(s.specs, strct)
	return s
}

func (s *FileSpec) AddFunction(fn *FunctionSpec) *FileSpec {
	s.specs = append(s.specs, fn)
	return s
}

func (s *FileSpec) AddVariable(v *VariableSpec) *FileSpec {
	s.specs = append(s.specs, v)
	return s
}

func (s *FileSpec) AddConst(name string, expr string) *FileSpec {
	s.specs = append(s.specs, &ConstSpec{Name: name, Expr: expr})
	return s
}

func (s *FileSpec) ProcessForImports() error {
	for _, spec := range s.specs {
		switch checked := spec.(type) {
		case (*InterfaceSpec):
			for _, m := range checked.Methods {
				for _, f := range m.Parameters {
					if f.Typ.Import() != nil {
						s.Imports = append(s.Imports, f.Typ.Import())
					}
				}
				for _, r := range m.ReturnTypes {
					if r.Import() != nil {
						s.Imports = append(s.Imports, r.Import())
					}
				}
			}
		case (*StructSpec):
			for _, f := range checked.Fields {
				if f.Typ.Import() != nil {
					s.Imports = append(s.Imports, f.Typ.Import())
				}
			}
			for _, m := range checked.Methods {
				for _, f := range m.Func.Parameters {
					if f.Typ.Import() != nil {
						s.Imports = append(s.Imports, f.Typ.Import())
					}
				}
				for _, r := range m.Func.ReturnTypes {
					if r == nil {
						fmt.Printf("r is %s\n", r)
					}
					if r.Import() != nil {
						s.Imports = append(s.Imports, r.Import())
					}
				}
			}
		case (*FunctionSpec):
			for _, p := range checked.Parameters {
				if p.Typ.Import() != nil {
					s.Imports = append(s.Imports, p.Typ.Import())
				}
			}
			for _, r := range checked.ReturnTypes {
				if r.Import() != nil {
					s.Imports = append(s.Imports, r.Import())
				}
			}
		case (*VariableSpec):
			if checked.Typ.Import() != nil {
				s.Imports = append(s.Imports, checked.Typ.Import())
			}
		}
	}

	return nil
}

func (s *FileSpec) Write(w io.Writer) error {
	// Preprocess the elements ahead of time to find manual imports.
	if err := s.ProcessForImports(); err != nil {
		return err
	}

	io.WriteString(w, fmt.Sprintf("package %s", s.PackageName))
	io.WriteString(w, "\n")

	if len(s.Imports) > 0 {
		io.WriteString(w, "import (")
		io.WriteString(w, "\n")
		for _, i := range s.Imports {
			i.Write(w)
			io.WriteString(w, "\n")
		}
		io.WriteString(w, ")")
		io.WriteString(w, "\n")
	}

	for _, spec := range s.specs {
		spec.Write(w)
		io.WriteString(w, "\n")
	}

	return nil
}

func (s *FileSpec) Generate(w io.Writer) error {
	b := bytes.NewBuffer(nil)

	s.Write(b)

	formatted, err := imports.Process("", b.Bytes(), nil)
	if err != nil {
		return err
	}

	w.Write(formatted)

	return nil
}

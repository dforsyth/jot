package jot

import (
	"fmt"
	"io"
)

type ImportSpec struct {
	Path string
}

func Import(path string) *ImportSpec {
	return &ImportSpec{Path: path}
}

func (s *ImportSpec) Write(w io.Writer) error {
	io.WriteString(w, fmt.Sprintf("\"%s\"", s.Path))

	return nil
}

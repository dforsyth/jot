package jot

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFunction(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Function("testFunction")
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "func testFunction() {\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

func TestFunctionParameter(t *testing.T) {
	b := new(bytes.Buffer)
	typeSpec := Type(reflect.TypeOf((*string)(nil)).Elem())

	spec := Function("testFunction").AddParameter("s1", typeSpec)
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "func testFunction(s1 string) {\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")

	spec.AddParameter("s2", typeSpec)
	b.Reset()
	err = spec.Write(b)

	output = b.String()
	expect = "func testFunction(s1 string, s2 string) {\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

func TestFunctionReturns(t *testing.T) {
	b := new(bytes.Buffer)
	typeSpec := Type(reflect.TypeOf((*string)(nil)).Elem())

	spec := Function("testFunction").AddReturnType(typeSpec)
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "func testFunction() string {\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")

	spec.AddReturnType(typeSpec)
	b.Reset()
	err = spec.Write(b)
	assert.Nil(t, err)

	output = b.String()
	expect = "func testFunction() (string, string) {\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

func TestFunctionCode(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Function("testFunction").AddCode("println(1 + 2)")
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "func testFunction() {\nprintln(1 + 2)\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")

	spec = Function("testFunction").AddCodeFmt("println({0} + {1})", 1, 2)
	b.Reset()
	err = spec.Write(b)
	assert.Nil(t, err)

	output = b.String()
	assert.Equal(t, expect, output, "Output mismatch.")
}

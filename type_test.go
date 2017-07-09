package jot

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeString(t *testing.T) {
	b := new(bytes.Buffer)
	spec := TypeString("testType")
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "testType"
	assert.Equal(t, expect, output, "Output mismatch.")
}

func TestTypeFrom(t *testing.T) {
	b := new(bytes.Buffer)
	spec := TypeString("testType").From(Import("imaginarypkg"))
	file := File("type_test").AddVariable(Variable("test", spec))
	err := file.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "package type_test\nimport (\n\"imaginarypkg\"\n)\nvar test testType\n"
	assert.Equal(t, expect, output, "Output mismatch.")

	// Make sure what we write can be generated.
	b.Reset()
	err = file.Generate(b)
	assert.Nil(t, err)
}

func TestPtr(t *testing.T) {

}

func TestArray(t *testing.T) {

}

func TestMap(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Map(TypeString("key"), TypeString("value"))
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "map[key]value"
	assert.Equal(t, expect, output, "Output mismatch.")

	// Make sure what we write can be generated.
	b.Reset()
	err = File("test_map").AddVariable(Variable("TestMap", spec)).Generate(b)
	assert.Nil(t, err)
}

func TestChan(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Chan(TypeString("test"))
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "chan test"
	assert.Equal(t, expect, output, "Output mismatch.")

	// Make sure what we write can be generated.
	b.Reset()
	err = File("test_chan").AddVariable(Variable("TestChan", spec)).Generate(b)
	assert.Nil(t, err)
}

func TestFunc(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Func().AddReturnType(TypeString("a")).AddParameter(TypeString("b"))
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "func(b) a"
	assert.Equal(t, expect, output, "Output mismatch.")
}

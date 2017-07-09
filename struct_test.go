package jot

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStruct(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Struct("TestStruct")
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "type TestStruct struct {}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

func TestStructField(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Struct("TestStruct").AddField(Field("field", TypeString("sometype")))
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "type TestStruct struct {\nfield sometype\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

func TestStructMethod(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Struct("TestStruct").AddMethod(Function("Method"))
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "type TestStruct struct {}\nfunc (TestStruct) Method() {\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

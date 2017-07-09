package jot

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Interface("TestInterface")
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "type TestInterface interface {}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

func TestInterfaceMethod(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Interface("TestInterface").AddMethod(Function("Method"))
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "type TestInterface interface {\nMethod()\n}\n"
	assert.Equal(t, expect, output, "Output mismatch.")
}

package jot

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariable(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Variable("testVar", TypeString("testType"))
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "var testVar testType"
	assert.Equal(t, expect, output, "Output mismatch.")
}

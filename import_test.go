package jot

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImport(t *testing.T) {
	b := new(bytes.Buffer)
	spec := Import("fmt")
	err := spec.Write(b)
	assert.Nil(t, err)

	output := b.String()
	expect := "\"fmt\""
	assert.Equal(t, expect, output, "Output mismatch.")
}

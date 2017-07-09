package jot

import (
	"io"
)

// Base Spec interface. Anything that implements this can be added to the
// tree and generated.
type Spec interface {
	Write(io.Writer) error
}

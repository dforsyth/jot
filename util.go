package jot

import (
	"fmt"
	"io"
	"strings"
)

func WriteComment(w io.Writer, comment string) error {
	_, err := io.WriteString(w, fmt.Sprintf(JotCommentPrefixFmt, comment))
	return err
}

func WriteDoc(w io.Writer, doc string) error {
	if doc == "" {
		return nil
	}

	for _, ln := range strings.Split(doc, "\n") {
		if err := WriteComment(w, ln); err != nil {
			return err
		}
		io.WriteString(w, "\n")
	}
	return nil
}

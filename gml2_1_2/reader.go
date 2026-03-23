package gml2_1_2

import (
	"io"

	intl "github.com/Kuroki-g/go-gml/gml2_1_2/internal"
)

// Reader is a GML 2.1.2 streaming geometry reader.
// Use NewReader to create a Reader from an io.Reader.
type Reader = intl.Reader

// NewReader creates a Reader that streams geometry elements from r.
// For non-UTF-8 encoded files (e.g. Shift-JIS), call SetCharsetReader after creation.
func NewReader(r io.Reader) *Reader {
	return intl.NewReader(r)
}

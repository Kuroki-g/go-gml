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
// On the first call to Next, the document is pre-scanned to populate the xlink:href
// resolver cache, enabling forward references to be resolved correctly.
func NewReader(r io.ReadSeeker) *Reader {
	return intl.NewReader(r)
}

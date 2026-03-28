package internal

import (
	"io"

	core "github.com/Kuroki-g/go-gml/core"
)

var _ core.Decoder = (*Reader)(nil)

// Decode decodes a single GML geometry subtree from src.
// Implements core.Decoder.
func (r *Reader) Decode(src io.ReadSeeker) (core.Geometry, error) {
	sub := NewReader(src)
	return sub.Next()
}

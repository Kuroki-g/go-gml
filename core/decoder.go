package core

import "io"

// Decoder decodes a single GML geometry subtree from r.
type Decoder interface {
	Decode(r io.ReadSeeker) (Geometry, error)
}

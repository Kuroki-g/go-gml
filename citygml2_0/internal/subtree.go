package internal

import (
	"bytes"
	"encoding/xml"
	"io"
)

// ExtractSubtree re-encodes the subtree starting at se (already consumed)
// through its matching EndElement into a bytes.Reader suitable for Decode.
func ExtractSubtree(dec *xml.Decoder, se xml.StartElement) (io.ReadSeeker, error) {
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	if err := enc.EncodeToken(se); err != nil {
		return nil, err
	}
	depth := 1
	for depth > 0 {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		if err := enc.EncodeToken(tok); err != nil {
			return nil, err
		}
		switch tok.(type) {
		case xml.StartElement:
			depth++
		case xml.EndElement:
			depth--
		}
	}
	if err := enc.Flush(); err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}

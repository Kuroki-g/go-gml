package internal

import (
	"bytes"
	"io"
)

// gmlNS32xsd is the non-standard GML 3.2 namespace URI used in older 国土数値情報 files.
// Technically this is a schema URL, not a namespace URI, but many pre-2025 files use it.
const gmlNS32xsd = "http://schemas.opengis.net/gml/3.2.1/gml.xsd"

var (
	nsNormOld = []byte(gmlNS32xsd)
	nsNormNew = []byte(gmlNS32)
)

// nsNormReader wraps an io.Reader and replaces occurrences of gmlNS32xsd with
// the canonical GML 3.2 namespace URI (gmlNS32), so that v3 struct XML tags
// match correctly during DecodeElement.
type nsNormReader struct {
	r    io.Reader
	tail []byte
	buf  []byte
	done bool
}

const nsNormChunk = 32 * 1024

func newNSNormReader(r io.Reader) *nsNormReader {
	return &nsNormReader{r: r}
}

func (n *nsNormReader) Read(p []byte) (int, error) {
	if len(n.buf) == 0 && !n.done {
		chunk := make([]byte, nsNormChunk)
		nr, err := n.r.Read(chunk)
		data := append(n.tail, chunk[:nr]...)
		if err == io.EOF {
			n.done = true
			n.buf = bytes.ReplaceAll(data, nsNormOld, nsNormNew)
			n.tail = nil
		} else if err != nil {
			return 0, err
		} else {
			carry := len(nsNormOld) - 1
			safe := len(data) - carry
			if safe < 0 {
				safe = 0
			}
			n.buf = bytes.ReplaceAll(data[:safe], nsNormOld, nsNormNew)
			n.tail = append([]byte(nil), data[safe:]...)
		}
	}
	copied := copy(p, n.buf)
	n.buf = n.buf[copied:]
	if copied == 0 && n.done {
		return 0, io.EOF
	}
	return copied, nil
}

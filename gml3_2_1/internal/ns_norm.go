package internal

import (
	"bytes"
	"io"
)

// gmlNS32xsd is the non-standard GML 3.2 namespace URI used in older 国土数値情報 files.
// Technically this is a schema URL, not a namespace URI, but many pre-2025 files use it.
const gmlNS32xsd = "http://schemas.opengis.net/gml/3.2.1/gml.xsd"

// gmlNS32schemas is a shorter non-standard variant seen in L01 (地価公示) files.
// e.g. xmlns:gml="http://schemas.opengis.net/gml/3.2"
const gmlNS32schemas = "http://schemas.opengis.net/gml/3.2"

var nsNormPairs = [][2][]byte{
	{[]byte(gmlNS32xsd), []byte(gmlNS32)},
	{[]byte(gmlNS32schemas), []byte(gmlNS32)},
}

// nsNormReader wraps an io.Reader and replaces non-standard GML 3.2 namespace
// variants with the canonical URI (gmlNS32), so that v3 struct XML tags match
// correctly during DecodeElement.
// Pairs are applied longest-first to avoid partial matches (gmlNS32xsd contains
// gmlNS32schemas as a prefix, so xsd must be replaced before schemas).
type nsNormReader struct {
	r    io.Reader
	tail []byte
	buf  []byte
	done bool
}

const nsNormChunk = 32 * 1024

// nsNormMaxOld is the length of the longest old pattern, used to size the carry buffer.
var nsNormMaxOld = func() int {
	max := 0
	for _, p := range nsNormPairs {
		if len(p[0]) > max {
			max = len(p[0])
		}
	}
	return max
}()

func newNSNormReader(r io.Reader) *nsNormReader {
	return &nsNormReader{r: r}
}

func replaceAllPairs(data []byte) []byte {
	for _, p := range nsNormPairs {
		data = bytes.ReplaceAll(data, p[0], p[1])
	}
	return data
}

func (n *nsNormReader) Read(p []byte) (int, error) {
	if len(n.buf) == 0 && !n.done {
		chunk := make([]byte, nsNormChunk)
		nr, err := n.r.Read(chunk)
		data := append(n.tail, chunk[:nr]...)
		if err == io.EOF {
			n.done = true
			n.buf = replaceAllPairs(data)
			n.tail = nil
		} else if err != nil {
			return 0, err
		} else {
			carry := nsNormMaxOld - 1
			safe := len(data) - carry
			if safe < 0 {
				safe = 0
			}
			n.buf = replaceAllPairs(data[:safe])
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

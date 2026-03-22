package gml

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
//
// 背景: GML 3.2.1 XSD の targetNamespace は "http://www.opengis.net/gml/3.2" であり、
// "http://schemas.opengis.net/gml/3.2.1/gml.xsd" はスキーマファイルの所在 URL に過ぎない。
// 国土数値情報 N03-2022/2024 等の政府データはこの URL を xmlns:gml= に誤って記述しており、
// GML XSD 非準拠のデータバグである。厳密な GML プロセッサなら未知 namespace として無視すべき挙動だが、
// この SDK は実データ優先の設計方針に基づきデータ側のバグを寛容に吸収する。
//
// A carry buffer of len(old)-1 bytes is retained between reads to handle
// replacements that would otherwise span chunk boundaries.
type nsNormReader struct {
	r    io.Reader
	tail []byte // suffix of previous chunk (len <= len(nsNormOld)-1)
	buf  []byte // pending output bytes ready to be consumed
	done bool
}

const nsNormChunk = 32 * 1024

func newNSNormReader(r io.Reader) *nsNormReader {
	return &nsNormReader{r: r}
}

func (n *nsNormReader) Read(p []byte) (int, error) {
	// Refill buf if empty and underlying reader is not exhausted.
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
			// Only process bytes up to len(data)-len(old)+1 so that a match
			// straddling the current chunk boundary is not missed.
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

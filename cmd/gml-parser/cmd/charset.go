package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/Kuroki-g/go-gml/gml"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(strings.ReplaceAll(charset, "-", "_")) {
	case "shift_jis", "sjis", "windows_31j", "cp932":
		return transform.NewReader(input, japanese.ShiftJIS.NewDecoder()), nil
	default:
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}
}

// newGMLReader creates a core.Reader for the given GML version with Shift-JIS charset support.
// Older 国土数値情報 files declare encoding="shift_jis" in their XML header.
// version must be one of "3.2.1", "3.1.1", "2.1.2".
// onUnknown, if non-nil, is called for each unrecognised element name encountered.
func newGMLReader(r io.ReadSeeker, version string, onUnknown func(string)) (gml.Reader, error) {
	switch version {
	case "3.2.1":
		reader := gml.NewReader321(r)
		reader.SetCharsetReader(charsetReader)
		reader.OnUnknownElement = onUnknown
		return reader, nil
	case "3.1.1":
		reader := gml.NewReader311(r)
		reader.SetCharsetReader(charsetReader)
		reader.OnUnknownElement = onUnknown
		return reader, nil
	case "2.1.2":
		reader := gml.NewReader212(r)
		reader.SetCharsetReader(charsetReader)
		reader.OnUnknownElement = onUnknown
		return reader, nil
	default:
		return nil, fmt.Errorf("unsupported GML version: %s (valid: 3.2.1, 3.1.1, 2.1.2)", version)
	}
}

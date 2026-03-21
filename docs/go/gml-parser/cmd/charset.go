package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/Kuroki-g/go-gml/pkg/gml"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// newGMLReader creates a gml.Reader with Shift-JIS charset support.
// Older 国土数値情報 files declare encoding="shift_jis" in their XML header.
func newGMLReader(r io.Reader) *gml.Reader {
	reader := gml.NewReader(r)
	reader.SetCharsetReader(func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(strings.ReplaceAll(charset, "-", "_")) {
		case "shift_jis", "sjis", "windows_31j", "cp932":
			return transform.NewReader(input, japanese.ShiftJIS.NewDecoder()), nil
		default:
			return nil, fmt.Errorf("unsupported charset: %s", charset)
		}
	})
	return reader
}

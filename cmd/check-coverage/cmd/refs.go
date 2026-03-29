package cmd

import (
	"os"
	"path/filepath"
	"regexp"
)

var (
	fieldAccessRe = regexp.MustCompile(`\.([A-Z][A-Za-z0-9]+)\b`)
	identRe       = regexp.MustCompile(`\b([A-Z][A-Za-z0-9]+)\b`)
)

// Refs holds field-access and type-identifier references found in internal/.
type Refs struct {
	Fields map[string]bool // ".FieldName" accesses
	Types  map[string]bool // bare "TypeName" identifiers (for struct name filter)
}

// collectFieldRefs scans all .go files in dir and returns a Refs.
func collectFieldRefs(dir string) (Refs, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return Refs{}, err
	}

	r := Refs{
		Fields: make(map[string]bool),
		Types:  make(map[string]bool),
	}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return Refs{}, err
		}
		for _, m := range fieldAccessRe.FindAllSubmatch(data, -1) {
			r.Fields[string(m[1])] = true
		}
		for _, m := range identRe.FindAllSubmatch(data, -1) {
			r.Types[string(m[1])] = true
		}
	}
	return r, nil
}

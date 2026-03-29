package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

// Finding represents one unhandled PropertyType pointer field.
type Finding struct {
	Module  string
	Struct  string
	Field   string
	XMLElem string
}

// checkModule analyzes one parser module directory for unhandled dispatch fields.
func checkModule(dir string) ([]Finding, error) {
	genDir := filepath.Join(dir, "generated")
	intDir := filepath.Join(dir, "internal")

	if _, err := os.Stat(genDir); err != nil {
		return nil, nil
	}
	if _, err := os.Stat(intDir); err != nil {
		return nil, nil
	}

	propFields, referenced, err := parseGeneratedStructs(genDir)
	if err != nil {
		return nil, err
	}

	refs, err := collectFieldRefs(intDir)
	if err != nil {
		return nil, err
	}

	var findings []Finding
	for name, fields := range propFields {
		if !referenced[name] {
			continue // dead struct: never used as a field type
		}
		if !refs.Types[name] {
			continue // struct type never mentioned in internal/: out of scope
		}
		for _, f := range fields {
			if isSkippedField(f.Name) {
				continue
			}
			if !refs.Fields[f.Name] {
				findings = append(findings, Finding{
					Module:  dir,
					Struct:  name,
					Field:   f.Name,
					XMLElem: f.XMLElem,
				})
			}
		}
	}
	return findings, nil
}

// isSkippedField returns true for abstract elements and xlink/gml attribute
// fields that don't require a dispatch branch in decode functions.
func isSkippedField(name string) bool {
	if strings.HasPrefix(name, "Abstract") {
		return true
	}
	switch name {
	case "Href", "RemoteSchema", "TypeField", "Role", "Arcrole",
		"Title", "Show", "Actuate", "NilReason", "Owns":
		return true
	}
	return false
}

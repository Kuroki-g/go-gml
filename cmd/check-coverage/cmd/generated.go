package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

// FieldInfo describes one XML-tagged field of a PropertyType struct.
type FieldInfo struct {
	Name    string // Go field name (e.g., "Href", "Solid")
	XMLName string // XML local name or attribute name
	IsAttr  bool   // true if xml:",attr" tag
}

// parsePropertyTypes parses all .go files in dir and returns a map of
// PropertyType struct name → all XML-tagged, non-abstract fields.
//
// "PropertyType" structs are identified by the "PropertyType" suffix.
// All XML-tagged fields are collected: inline element pointers, string
// attributes (Href, SrsName, etc.), and slice fields. Abstract substitution
// group head fields (Abstract* prefix) are excluded because they cannot be
// instantiated by a concrete document.
func parsePropertyTypes(dir string) (map[string][]FieldInfo, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]FieldInfo)
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				gd, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, spec := range gd.Specs {
					ts, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if !strings.HasSuffix(ts.Name.Name, "PropertyType") {
						continue
					}
					st, ok := ts.Type.(*ast.StructType)
					if !ok {
						continue
					}
					fields := collectXMLFields(st)
					if len(fields) > 0 {
						result[ts.Name.Name] = fields
					}
				}
			}
		}
	}
	return result, nil
}

// collectXMLFields returns all XML-tagged, non-abstract fields of a struct.
// This includes pointer fields (inline elements), string attributes, and
// slice fields — anything the XML decoder may populate.
func collectXMLFields(st *ast.StructType) []FieldInfo {
	var fields []FieldInfo
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 || f.Tag == nil {
			continue
		}
		name := f.Names[0].Name
		if strings.HasPrefix(name, "Abstract") {
			// Abstract substitution group heads: no concrete XML element maps to these.
			continue
		}
		xmlName, isAttr := extractXMLInfo(f.Tag.Value)
		if xmlName == "" {
			continue
		}
		fields = append(fields, FieldInfo{
			Name:    name,
			XMLName: xmlName,
			IsAttr:  isAttr,
		})
	}
	return fields
}

var xmlTagRe = regexp.MustCompile(`xml:"([^"]*)"`)

// extractXMLInfo returns the XML local name and whether the field is an
// attribute from a struct tag string.
func extractXMLInfo(tag string) (name string, isAttr bool) {
	m := xmlTagRe.FindStringSubmatch(tag)
	if m == nil {
		return "", false
	}
	v := m[1]
	isAttr = strings.Contains(v, ",attr")
	base := strings.SplitN(v, ",", 2)[0]
	parts := strings.Fields(base)
	if len(parts) == 0 {
		return "", false
	}
	return parts[len(parts)-1], isAttr
}

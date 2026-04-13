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
// association struct name → all XML-tagged, non-abstract fields.
//
// Association structs are identified by having an "Href" field (i.e. they
// embed AssociationAttributeGroup from the XSD). This covers *PropertyType,
// *MemberType, *RefType, and *AssociationType suffixes.
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
					st, ok := ts.Type.(*ast.StructType)
					if !ok {
						continue
					}
					if !structHasHrefField(st) {
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

// structHasHrefField reports whether a struct has a field named "Href",
// which indicates it embeds AssociationAttributeGroup from the GML XSD.
func structHasHrefField(st *ast.StructType) bool {
	for _, f := range st.Fields.List {
		for _, n := range f.Names {
			if n.Name == "Href" {
				return true
			}
		}
	}
	return false
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

// parseContainerFields scans all .go files in dir and returns a map of
// field name → PropertyType name for slice fields whose element type is a
// known PropertyType (e.g. "PointMember" → "PointPropertyType").
//
// Only fields whose name maps to exactly one PropertyType across all structs
// are included. Ambiguous field names (same name, different element types in
// different structs) are excluded to avoid false positives.
func parseContainerFields(dir string, knownTypes map[string][]FieldInfo) (map[string]string, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return nil, err
	}
	// allSliceElemTypes collects ALL element type names seen for each slice field
	// name across all structs (regardless of whether the type is a known
	// PropertyType). This prevents false positives when the same field name is
	// used with different element types in different structs.
	allSliceElemTypes := make(map[string]map[string]bool)
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
					st, ok := ts.Type.(*ast.StructType)
					if !ok {
						continue
					}
					for _, f := range st.Fields.List {
						if len(f.Names) == 0 {
							continue
						}
						arr, ok := f.Type.(*ast.ArrayType)
						if !ok {
							continue
						}
						typeName := baseTypeName(arr.Elt, "")
						if typeName == "" {
							continue
						}
						fieldName := f.Names[0].Name
						if allSliceElemTypes[fieldName] == nil {
							allSliceElemTypes[fieldName] = make(map[string]bool)
						}
						allSliceElemTypes[fieldName][typeName] = true
					}
				}
			}
		}
	}
	// Keep only fields where all occurrences use the same element type AND that
	// type is a known PropertyType. Fields shared across structs with different
	// element types are excluded to avoid false positives.
	result := make(map[string]string)
	for fieldName, elemTypes := range allSliceElemTypes {
		if len(elemTypes) != 1 {
			continue // ambiguous: different element types across structs
		}
		for typeName := range elemTypes {
			if _, known := knownTypes[typeName]; known {
				result[fieldName] = typeName
			}
		}
	}
	return result, nil
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

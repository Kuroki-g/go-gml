package resolve

import (
	"strings"

	"xsd2go-lite2/internal/parse"
)

// ResolveAll resolves all loaded schemas and returns ComplexTypes for targetNS.
func (r *Resolver) ResolveAll(targetNS string) []*parse.ComplexType {
	for _, ct := range r.allComplexTypes {
		if len(ct.Fields) == 0 {
			r.resolveComplexType(ct, nil)
		}
	}
	var result []*parse.ComplexType
	for key, ct := range r.allComplexTypes {
		if strings.HasPrefix(key, targetNS+" ") {
			result = append(result, ct)
		}
	}
	return result
}

// resolveComplexType builds ct.Fields (depth-first, cycle-guarded by visiting).
func (r *Resolver) resolveComplexType(ct *parse.ComplexType, visiting map[string]bool) {
	if visiting == nil {
		visiting = make(map[string]bool)
	}
	key := ct.Source + " " + ct.Name
	if visiting[key] {
		return
	}
	visiting[key] = true
	defer delete(visiting, key)

	if len(ct.Fields) > 0 {
		return
	}

	if ct.Derivation != nil {
		// Inheritance (extension/restriction) is implemented in Step 4.
		return
	}

	// Direct definition: no inheritance.
	var fields []parse.Field

	if ct.Content != nil {
		fields = append(fields, r.resolveContentModel(ct.Content, ct.Source, false, false)...)
	}
	if ct.CharData != nil {
		fields = append(fields, r.resolveCharData(ct.CharData, ct.Source)...)
	}
	for _, ad := range ct.Attrs {
		fields = append(fields, r.resolveAttrDecl(ad, ct.Source)...)
	}
	for _, agRef := range ct.AttrGroups {
		fields = append(fields, r.resolveAttrGroupRef(agRef, ct.Source)...)
	}

	ct.Fields = fields
}

// resolveCharData converts a CharDataDecl to a chardata Field.
func (r *Resolver) resolveCharData(cd *parse.CharDataDecl, schemaNS string) []parse.Field {
	goType := r.resolveTypeName(cd.TypeRef, schemaNS)
	if goType == "" || parse.IsStructType(goType) {
		goType = "string"
	}
	return []parse.Field{{
		GoName: "Value",
		XMLTag: ",chardata",
		GoType: goType,
		IsChar: true,
	}}
}

// deduplicateFields removes fields with duplicate GoNames.
// When an attribute and an element conflict, the later-defined (own) field wins.
// Used only for extension (Step 4); defined here for reuse.
func deduplicateFields(fields []parse.Field) []parse.Field {
	type entry struct {
		idx    int
		isAttr bool
	}
	if len(fields) == 0 {
		return fields
	}
	seen := make(map[string]entry)
	result := make([]parse.Field, 0, len(fields))
	for _, f := range fields {
		e, exists := seen[f.GoName]
		if !exists {
			seen[f.GoName] = entry{len(result), f.IsAttr}
			result = append(result, f)
		} else if f.IsAttr != e.isAttr {
			result[e.idx] = f
			seen[f.GoName] = entry{e.idx, f.IsAttr}
		}
	}
	return result
}

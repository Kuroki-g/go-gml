package resolve

import (
	"fmt"
	"os"

	"xsd2go-lite2/internal/parse"
)

// resolveBaseExtension resolves the base type for extension and returns all its Fields.
func (r *Resolver) resolveBaseExtension(base, schemaNS string, visiting map[string]bool) []parse.Field {
	baseNS, baseName := r.resolveQName(base, schemaNS)
	baseCT, ok := r.allComplexTypes[baseNS+" "+baseName]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: extension base type not found: %s\n", base)
		return nil
	}
	r.resolveComplexType(baseCT, visiting)
	result := make([]parse.Field, len(baseCT.Fields))
	copy(result, baseCT.Fields)
	return result
}

// resolveBaseRestrictionAttrs collects attribute Fields from the base type,
// filtered by prohibitedSet, then merged with own attribute declarations.
// own attrs (non-prohibited) override base attrs with the same GoName.
func (r *Resolver) resolveBaseRestrictionAttrs(base, schemaNS string, visiting map[string]bool, prohibitedSet map[string]bool, ownAttrs []parse.AttrDecl, ownAttrGroups []string) []parse.Field {
	baseNS, baseName := r.resolveQName(base, schemaNS)
	baseCT, ok := r.allComplexTypes[baseNS+" "+baseName]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: restriction base type not found: %s\n", base)
		var fields []parse.Field
		for _, ad := range ownAttrs {
			if ad.Use == parse.UseProhibited {
				continue
			}
			fields = append(fields, r.resolveAttrDecl(ad, schemaNS)...)
		}
		for _, agRef := range ownAttrGroups {
			r.resolveAttrGroupRef(agRef, schemaNS) // populate ag.Fields only
		}
		return deduplicateFields(fields)
	}
	r.resolveComplexType(baseCT, visiting)

	var fields []parse.Field
	for _, f := range baseCT.Fields {
		if !f.IsAttr || prohibitedSet[f.GoName] {
			continue
		}
		fields = append(fields, f)
	}
	for _, ad := range ownAttrs {
		if ad.Use == parse.UseProhibited {
			continue
		}
		fields = append(fields, r.resolveAttrDecl(ad, schemaNS)...)
	}
	for _, agRef := range ownAttrGroups {
		r.resolveAttrGroupRef(agRef, schemaNS) // populate ag.Fields only
	}
	return deduplicateFields(fields)
}

// collectProhibited builds a set of GoNames for use="prohibited" attribute declarations.
func (r *Resolver) collectProhibited(attrs []parse.AttrDecl, schemaNS string) map[string]bool {
	set := make(map[string]bool)
	for _, ad := range attrs {
		if ad.Use != parse.UseProhibited {
			continue
		}
		if ad.Ref != "" {
			_, name := r.resolveQName(ad.Ref, schemaNS)
			set[parse.GoName(name)] = true
		} else {
			set[parse.GoName(ad.Name)] = true
		}
	}
	return set
}

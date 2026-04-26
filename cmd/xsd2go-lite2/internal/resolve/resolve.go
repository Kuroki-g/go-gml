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
		switch ct.Derivation.Kind {
		case "extension":
			// simpleContent base is always a SimpleType; no struct fields to inherit.
			var baseFields []parse.Field
			if ct.ContentKind != parse.ContentKindSimple {
				baseFields = r.resolveBaseExtension(ct.Derivation.Base, ct.Source, visiting)
				ct.Embeds = r.copyBaseEmbeds(ct.Derivation.Base, ct.Source)
			}
			var ownFields []parse.Field
			if ct.Content != nil {
				ownFields = append(ownFields, r.resolveContentModel(ct.Content, ct.Source, false, false)...)
			}
			if ct.CharData != nil {
				ownFields = append(ownFields, r.resolveCharData(ct.CharData, ct.Source)...)
			}
			for _, ad := range ct.Attrs {
				ownFields = append(ownFields, r.resolveAttrDecl(ad, ct.Source)...)
			}
			for _, agRef := range ct.AttrGroups {
				r.resolveAttrGroupRef(agRef, ct.Source) // populate ag.Fields only
				agNS, agName := r.resolveQName(agRef, ct.Source)
				ct.Embeds = appendEmbed(ct.Embeds, parse.EmbedRef{XSDName: agName, NS: agNS, Kind: "attributeGroup"})
			}
			ct.Fields = deduplicateFields(append(baseFields, ownFields...))
			r.addGroupEmbeds(ct)
		case "restriction":
			prohibited := r.collectProhibited(ct.Attrs, ct.Source)
			if ct.Content != nil {
				ct.Fields = append(ct.Fields, r.resolveContentModel(ct.Content, ct.Source, false, false)...)
			}
			if ct.CharData != nil {
				ct.Fields = append(ct.Fields, r.resolveCharData(ct.CharData, ct.Source)...)
			}
			attrFields := r.resolveBaseRestrictionAttrs(ct.Derivation.Base, ct.Source, visiting, prohibited, ct.Attrs, ct.AttrGroups)
			ct.Fields = append(ct.Fields, attrFields...)
			// Base is now resolved; copy its Embeds then add own.
			ct.Embeds = r.copyBaseEmbeds(ct.Derivation.Base, ct.Source)
			for _, agRef := range ct.AttrGroups {
				agNS, agName := r.resolveQName(agRef, ct.Source)
				ct.Embeds = appendEmbed(ct.Embeds, parse.EmbedRef{XSDName: agName, NS: agNS, Kind: "attributeGroup"})
			}
			r.addGroupEmbeds(ct)
		}
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
		r.resolveAttrGroupRef(agRef, ct.Source) // populate ag.Fields only
		agNS, agName := r.resolveQName(agRef, ct.Source)
		ct.Embeds = appendEmbed(ct.Embeds, parse.EmbedRef{XSDName: agName, NS: agNS, Kind: "attributeGroup"})
	}

	ct.Fields = fields
	r.addGroupEmbeds(ct)
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

// deduplicateFields removes fields with duplicate GoNames, keeping the last occurrence.
// Derived type declarations override base type declarations (last wins).
// Used for extension; defined here for reuse.
func deduplicateFields(fields []parse.Field) []parse.Field {
	if len(fields) == 0 {
		return fields
	}
	seen := make(map[string]int)
	result := make([]parse.Field, 0, len(fields))
	for _, f := range fields {
		if idx, exists := seen[f.GoName]; exists {
			result[idx] = f
		} else {
			seen[f.GoName] = len(result)
			result = append(result, f)
		}
	}
	return result
}

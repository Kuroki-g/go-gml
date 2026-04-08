package parse

import "strings"

func convertSimpleType(st xsdSimpleType) SimpleType {
	r := SimpleType{Name: st.Name}
	switch {
	case st.List != nil:
		r.Kind = "list"
		r.GoType = "string"
	case st.Union != nil:
		r.Kind = "union"
		r.GoType = "string"
	case st.Restriction != nil:
		r.Kind = "restriction"
		r.GoType = XsBuiltinGoType(st.Restriction.Base)
	default:
		r.Kind = "alias"
		r.GoType = "string"
	}
	return r
}

func convertComplexType(ct xsdComplexType, ns string) ComplexType {
	result := ComplexType{
		Name:     ct.Name,
		Abstract: ct.Abstract == "true",
		Mixed:    ct.Mixed == "true",
		Source:   ns,
	}
	if ct.Annotation != nil {
		result.Doc = strings.TrimSpace(ct.Annotation.Documentation)
	}

	switch {
	case ct.SimpleContent != nil:
		convertSimpleContentType(ct.SimpleContent, &result)

	case ct.ComplexContent != nil:
		convertComplexContentType(ct.ComplexContent, &result)

	default:
		// Direct sequence/all/choice
		if ct.Sequence != nil {
			result.RawFields = append(result.RawFields, collectSequenceFields(ct.Sequence)...)
		}
		if ct.All != nil {
			result.RawFields = append(result.RawFields, collectSequenceFields(ct.All)...)
		}
		if ct.Choice != nil {
			result.RawFields = append(result.RawFields, collectSequenceFields(ct.Choice)...)
		}
		result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(ct.AttrGroups)...)
		result.RawFields = append(result.RawFields, convertAttributes(ct.Attributes)...)
	}

	return result
}

func convertSimpleContentType(sc *xsdSimpleContent, result *ComplexType) {
	if ext := sc.Extension; ext != nil {
		// chardata field for the value
		result.RawFields = append(result.RawFields, RawField{
			LocalName: "__chardata__",
			TypeRef:   ext.Base,
			IsChar:    true,
		})
		result.RawFields = append(result.RawFields, collectExtensionFields(ext)...)
		for _, ag := range ext.AttrGroups {
			result.AttrGroups = append(result.AttrGroups, ag.Ref)
		}
	} else if res := sc.Restriction; res != nil {
		result.RawFields = append(result.RawFields, RawField{
			LocalName: "__chardata__",
			TypeRef:   res.Base,
			IsChar:    true,
		})
		result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(res.AttrGroups)...)
		result.RawFields = append(result.RawFields, convertAttributes(res.Attributes)...)
	}
}

func convertComplexContentType(cc *xsdComplexContent, result *ComplexType) {
	if ext := cc.Extension; ext != nil {
		// Base type placeholder (resolved later)
		result.RawFields = append(result.RawFields, RawField{
			Ref: "__base__:" + ext.Base,
		})
		result.RawFields = append(result.RawFields, collectExtensionFields(ext)...)
		result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(ext.AttrGroups)...)
	} else if res := cc.Restriction; res != nil {
		// complexContent/restriction replaces the parent's element content model
		// but inherits its attribute declarations. Only the restriction's own
		// element fields are included; attribute fields are inherited from the base.
		if res.Sequence != nil {
			result.RawFields = append(result.RawFields, collectSequenceFields(res.Sequence)...)
		}
		result.RawFields = append(result.RawFields, convertAttributes(res.Attributes)...)
		result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(res.AttrGroups)...)
		// Inherit only attribute fields from the base type.
		result.RawFields = append(result.RawFields, RawField{
			Ref: "__base_attrs__:" + res.Base,
		})
	}
}

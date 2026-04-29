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
		r.Base = st.Restriction.Base
		for _, e := range st.Restriction.Enumerations {
			r.Enumerations = append(r.Enumerations, e.Value)
		}
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
		result.ContentKind = ContentKindSimple
		convertSimpleContentType(ct.SimpleContent, &result)
	case ct.ComplexContent != nil:
		result.ContentKind = ContentKindComplex
		convertComplexContentType(ct.ComplexContent, &result)
	default:
		cm := firstCompositor(ct.Sequence, ct.All, ct.Choice)
		if cm == nil {
			cm = directGroupsContentModel(ct.Groups)
		}
		result.Content = cm
		result.AttrGroups = collectAttrGroupRefs(ct.AttrGroups)
		result.Attrs = convertAttributes(ct.Attributes)
	}
	return result
}

func convertSimpleContentType(sc *xsdSimpleContent, result *ComplexType) {
	if ext := sc.Extension; ext != nil {
		result.Derivation = &Derivation{Kind: "extension", Base: ext.Base}
		result.CharData = &CharDataDecl{TypeRef: ext.Base}
		result.Attrs = convertAttributes(ext.Attributes)
		result.AttrGroups = collectAttrGroupRefs(ext.AttrGroups)
	} else if res := sc.Restriction; res != nil {
		result.Derivation = &Derivation{Kind: "restriction", Base: res.Base}
		result.CharData = &CharDataDecl{TypeRef: res.Base}
		result.Attrs = convertAttributes(res.Attributes)
		result.AttrGroups = collectAttrGroupRefs(res.AttrGroups)
	}
}

func convertComplexContentType(cc *xsdComplexContent, result *ComplexType) {
	if ext := cc.Extension; ext != nil {
		result.Derivation = &Derivation{Kind: "extension", Base: ext.Base}
		cm := firstCompositor(ext.Sequence, ext.All, ext.Choice)
		if cm == nil {
			cm = directGroupsContentModel(ext.Groups)
		}
		result.Content = cm
		result.Attrs = convertAttributes(ext.Attributes)
		result.AttrGroups = collectAttrGroupRefs(ext.AttrGroups)
	} else if res := cc.Restriction; res != nil {
		result.Derivation = &Derivation{Kind: "restriction", Base: res.Base}
		cm := firstCompositor(res.Sequence, res.All, res.Choice)
		if cm == nil {
			cm = directGroupsContentModel(res.Groups)
		}
		result.Content = cm
		result.Attrs = convertAttributes(res.Attributes)
		result.AttrGroups = collectAttrGroupRefs(res.AttrGroups)
	}
}

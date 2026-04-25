package parse

import (
	"log"
	"strings"
)

// buildContentModel converts an xsdSequence to a ContentModel.
// kind is "sequence", "choice", or "all".
func buildContentModel(seq *xsdSequence, kind string) ContentModel {
	cm := ContentModel{
		Kind:      kind,
		MinOccurs: seq.MinOccurs,
		MaxOccurs: seq.MaxOccurs,
	}
	for _, el := range seq.Elements {
		cm.Elems = append(cm.Elems, buildElemDecl(el))
	}
	for _, gr := range seq.Groups {
		cm.Groups = append(cm.Groups, GroupRef{
			Ref:       gr.Ref,
			MinOccurs: gr.MinOccurs,
			MaxOccurs: gr.MaxOccurs,
		})
	}
	for i := range seq.Sequences {
		cm.Children = append(cm.Children, buildContentModel(&seq.Sequences[i], "sequence"))
	}
	for i := range seq.Choices {
		cm.Children = append(cm.Children, buildContentModel(&seq.Choices[i], "choice"))
	}
	if len(seq.Anys) > 0 {
		log.Printf("xsd2go-lite2: xs:any in %s compositor (ignored)", kind)
	}
	return cm
}

// firstCompositor returns a ContentModel built from the first non-nil compositor.
func firstCompositor(seq, all, choice *xsdSequence) *ContentModel {
	switch {
	case seq != nil:
		cm := buildContentModel(seq, "sequence")
		return &cm
	case all != nil:
		cm := buildContentModel(all, "all")
		return &cm
	case choice != nil:
		cm := buildContentModel(choice, "choice")
		return &cm
	}
	return nil
}

func buildElemDecl(el xsdElement) ElemDecl {
	ed := ElemDecl{
		MinOccurs: el.MinOccurs,
		MaxOccurs: el.MaxOccurs,
		Default:   el.Default,
		Fixed:     el.Fixed,
	}
	if el.Annotation != nil {
		ed.Doc = strings.TrimSpace(el.Annotation.Documentation)
	}
	if el.Ref != "" {
		ed.Ref = el.Ref
	} else {
		ed.Name = el.Name
		ed.TypeRef = el.Type
		if ed.TypeRef == "" {
			ed.TypeRef = "string"
		}
	}
	return ed
}

// convertAttributes converts []xsdAttribute to []AttrDecl.
func convertAttributes(attrs []xsdAttribute) []AttrDecl {
	var result []AttrDecl
	for _, a := range attrs {
		doc := ""
		if a.Annotation != nil {
			doc = strings.TrimSpace(a.Annotation.Documentation)
		}
		ad := AttrDecl{
			Use:     a.Use,
			Default: a.Default,
			Fixed:   a.Fixed,
			Doc:     doc,
		}
		if a.Ref != "" {
			ad.Ref = a.Ref
		} else {
			ad.Name = a.Name
			ad.TypeRef = a.Type
		}
		result = append(result, ad)
	}
	return result
}

// collectAttrGroupRefs extracts attributeGroup ref QNames.
func collectAttrGroupRefs(refs []xsdAttrGroupRef) []string {
	var result []string
	for _, r := range refs {
		result = append(result, r.Ref)
	}
	return result
}

// collectAttrGroupDecl builds an AttrGroup from an xsdAttrGroup definition.
func collectAttrGroupDecl(ag xsdAttrGroup) AttrGroup {
	result := AttrGroup{Name: ag.Name}
	result.Attrs = convertAttributes(ag.Attributes)
	for _, nested := range ag.AttrGroups {
		result.NestedGroups = append(result.NestedGroups, nested.Ref)
	}
	return result
}

// collectGroupContent builds a ContentModel from an xsdGroup definition.
func collectGroupContent(g xsdGroup) *ContentModel {
	return firstCompositor(g.Sequence, g.All, g.Choice)
}

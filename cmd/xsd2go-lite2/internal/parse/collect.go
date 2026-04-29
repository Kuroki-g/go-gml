package parse

import (
	"log"
	"strings"
)

// buildCtx carries the effective compositor context propagated top-down during parse.
// Fields represent accumulated state: once true/set, they remain so for all descendants.
type buildCtx struct {
	optional  bool // any ancestor compositor is choice or has minOccurs="0"
	unbounded bool // any ancestor compositor has maxOccurs="unbounded"
}

func (c buildCtx) derive(kind, minOccurs, maxOccurs string) buildCtx {
	return buildCtx{
		optional:  c.optional || kind == "choice" || minOccurs == "0",
		unbounded: c.unbounded || maxOccurs == "unbounded",
	}
}

func (c buildCtx) effectiveMin(own string) string {
	if c.optional || own == "0" {
		return "0"
	}
	return own
}

func (c buildCtx) effectiveMax(own string) string {
	if c.unbounded || own == "unbounded" {
		return "unbounded"
	}
	return own
}

// buildContentModel converts an xsdSequence to a ContentModel.
// kind is "sequence", "choice", or "all".
func buildContentModel(seq *xsdSequence, kind string, ctx buildCtx) ContentModel {
	childCtx := ctx.derive(kind, seq.MinOccurs, seq.MaxOccurs)
	cm := ContentModel{
		Kind:      kind,
		MinOccurs: seq.MinOccurs,
		MaxOccurs: seq.MaxOccurs,
	}
	for _, el := range seq.Elements {
		cm.Elems = append(cm.Elems, buildElemDecl(el, childCtx))
	}
	for _, gr := range seq.Groups {
		cm.Groups = append(cm.Groups, GroupRef{
			Ref:       gr.Ref,
			MinOccurs: childCtx.effectiveMin(gr.MinOccurs),
			MaxOccurs: childCtx.effectiveMax(gr.MaxOccurs),
		})
	}
	for i := range seq.Sequences {
		cm.Children = append(cm.Children, buildContentModel(&seq.Sequences[i], "sequence", childCtx))
	}
	for i := range seq.Choices {
		cm.Children = append(cm.Children, buildContentModel(&seq.Choices[i], "choice", childCtx))
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
		cm := buildContentModel(seq, "sequence", buildCtx{})
		return &cm
	case all != nil:
		cm := buildContentModel(all, "all", buildCtx{})
		return &cm
	case choice != nil:
		cm := buildContentModel(choice, "choice", buildCtx{})
		return &cm
	}
	return nil
}

func buildElemDecl(el xsdElement, ctx buildCtx) ElemDecl {
	ed := ElemDecl{
		MinOccurs: ctx.effectiveMin(el.MinOccurs),
		MaxOccurs: ctx.effectiveMax(el.MaxOccurs),
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

// directGroupsContentModel wraps direct <group ref> particles (under extension/restriction/complexType)
// into a synthetic sequence ContentModel. Returns nil if groups is empty.
func directGroupsContentModel(groups []xsdGroupRef) *ContentModel {
	if len(groups) == 0 {
		return nil
	}
	ctx := buildCtx{}
	cm := ContentModel{Kind: "sequence"}
	for _, gr := range groups {
		cm.Groups = append(cm.Groups, GroupRef{
			Ref:       gr.Ref,
			MinOccurs: ctx.effectiveMin(gr.MinOccurs),
			MaxOccurs: ctx.effectiveMax(gr.MaxOccurs),
		})
	}
	return &cm
}

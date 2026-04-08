package xsdlite

import "strings"

func collectExtensionFields(ext *xsdExtension) []rawField {
	var fields []rawField
	if ext.Sequence != nil {
		fields = append(fields, collectSequenceFields(ext.Sequence)...)
	}
	if ext.All != nil {
		fields = append(fields, collectSequenceFields(ext.All)...)
	}
	if ext.Choice != nil {
		fields = append(fields, collectSequenceFields(ext.Choice)...)
	}
	fields = append(fields, convertAttributes(ext.Attributes)...)
	return fields
}

func collectSequenceFields(seq *xsdSequence) []rawField {
	if seq == nil {
		return nil
	}
	var fields []rawField
	for _, el := range seq.Elements {
		var rf rawField
		if el.Ref != "" {
			rf = rawField{
				Ref:       el.Ref,
				MinOccurs: el.MinOccurs,
				MaxOccurs: el.MaxOccurs,
			}
		} else {
			t := el.Type
			if t == "" && el.ComplexType != nil {
				t = "__inline__"
			} else if t == "" {
				t = "string"
			}
			rf = rawField{
				LocalName: el.Name,
				TypeRef:   t,
				MinOccurs: el.MinOccurs,
				MaxOccurs: el.MaxOccurs,
			}
		}
		if el.Annotation != nil {
			rf.Doc = strings.TrimSpace(el.Annotation.Documentation)
		}
		// If this sequence/choice container has maxOccurs="unbounded", propagate
		// it to direct child elements so they become slices.
		if seq.MaxOccurs == "unbounded" && rf.MaxOccurs != "unbounded" {
			rf.MaxOccurs = "unbounded"
		}
		fields = append(fields, rf)
	}
	// Recurse into nested sequences/choices.
	// If the container itself has maxOccurs="unbounded", propagate it to all
	// child fields so that elements like gml:pos inside
	// <xs:choice maxOccurs="unbounded"> become slices.
	for i := range seq.Sequences {
		childFields := collectSequenceFields(&seq.Sequences[i])
		if seq.Sequences[i].MaxOccurs == "unbounded" {
			for j := range childFields {
				childFields[j].MaxOccurs = "unbounded"
			}
		}
		fields = append(fields, childFields...)
	}
	for i := range seq.Choices {
		childFields := collectSequenceFields(&seq.Choices[i])
		if seq.Choices[i].MaxOccurs == "unbounded" {
			for j := range childFields {
				childFields[j].MaxOccurs = "unbounded"
			}
		}
		fields = append(fields, childFields...)
	}
	// Group refs
	for _, gr := range seq.Groups {
		fields = append(fields, rawField{
			GroupRef:  gr.Ref,
			MinOccurs: gr.MinOccurs,
			MaxOccurs: gr.MaxOccurs,
		})
	}
	// Inline attributes in sequence (rare but valid)
	fields = append(fields, convertAttributes(seq.Attributes)...)
	return fields
}

func convertAttributes(attrs []xsdAttribute) []rawField {
	var fields []rawField
	for _, a := range attrs {
		doc := ""
		if a.Annotation != nil {
			doc = strings.TrimSpace(a.Annotation.Documentation)
		}
		if a.Ref != "" {
			fields = append(fields, rawField{
				AttrRef: a.Ref,
				IsAttr:  true,
				Doc:     doc,
			})
		} else {
			fields = append(fields, rawField{
				LocalName: a.Name,
				TypeRef:   a.Type,
				IsAttr:    true,
				MinOccurs: a.Use, // "required" | "optional" | ""
				Doc:       doc,
			})
		}
	}
	return fields
}

func collectAttrGroupFields(ag xsdAttrGroup) []rawField {
	var fields []rawField
	fields = append(fields, convertAttributes(ag.Attributes)...)
	for _, nested := range ag.AttrGroups {
		fields = append(fields, rawField{
			AttrRef: "__group__:" + nested.Ref,
			IsAttr:  true,
		})
	}
	return fields
}

func collectGroupFields(g xsdGroup) []rawField {
	var seq *xsdSequence
	switch {
	case g.Sequence != nil:
		seq = g.Sequence
	case g.Choice != nil:
		seq = g.Choice
	case g.All != nil:
		seq = g.All
	}
	return collectSequenceFields(seq)
}

func collectAttrGroupRefs(refs []xsdAttrGroupRef) []string {
	var result []string
	for _, r := range refs {
		result = append(result, r.Ref)
	}
	return result
}

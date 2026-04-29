package resolve

import (
	"fmt"
	"os"
	"strings"

	"xsd2go-lite2/internal/parse"
)

// resolveContentModel recursively converts a ContentModel to []Field.
//
// optional: some ancestor compositor is choice or has minOccurs=0 (only relevant
// for group content, since inline ElemDecl.MinOccurs is pre-computed in parse).
// unbounded: some ancestor compositor has maxOccurs=unbounded (same caveat).
func (r *Resolver) resolveContentModel(cm *parse.ContentModel, schemaNS string, optional, unbounded bool) []parse.Field {
	childUnbounded := unbounded || cm.MaxOccurs == "unbounded"
	childOptional := optional || cm.MinOccurs == "0" || cm.Kind == "choice"

	var fields []parse.Field
	for _, ed := range cm.Elems {
		fields = append(fields, r.resolveElemDecl(ed, schemaNS, childOptional, childUnbounded)...)
	}
	for _, gr := range cm.Groups {
		fields = append(fields, r.resolveGroupRef(gr, schemaNS, childOptional, childUnbounded)...)
	}
	for i := range cm.Children {
		fields = append(fields, r.resolveContentModel(&cm.Children[i], schemaNS, childOptional, childUnbounded)...)
	}
	return fields
}

// resolveElemDecl converts an ElemDecl to zero or more Fields.
func (r *Resolver) resolveElemDecl(ed parse.ElemDecl, schemaNS string, optional, unbounded bool) []parse.Field {
	if ed.Ref != "" {
		return r.resolveElemRef(ed, schemaNS, optional, unbounded)
	}

	goType := r.resolveTypeName(ed.TypeRef, schemaNS)
	if goType == "" {
		goType = "string"
	}

	var typeNS string
	if !parse.IsBuiltinGoType(goType) && ed.TypeRef != "" {
		typeNS, _ = r.resolveQName(ed.TypeRef, schemaNS)
	}

	isSlice := unbounded || ed.MaxOccurs == "unbounded" || (ed.MaxOccurs != "" && ed.MaxOccurs != "1" && ed.MaxOccurs != "0")
	isOmit := optional || ed.MinOccurs == "0"

	f := parse.Field{
		GoName: parse.GoName(ed.Name),
		XMLTag: buildXMLTag(schemaNS, ed.Name, false),
		GoType: goType,
		TypeNS: typeNS,
		Omit:   isOmit,
		Slice:  isSlice,
		Doc:    ed.Doc,
	}

	if isSlice {
		f.GoType = "[]" + strings.TrimPrefix(goType, "[]")
	} else if !parse.IsBuiltinGoType(goType) {
		f.GoType = "*" + strings.TrimPrefix(goType, "*")
	} else if isOmit {
		f.GoType = "*" + goType
	}

	return []parse.Field{f}
}

// resolveElemRef resolves an element ref QName, expanding substitutionGroup members.
func (r *Resolver) resolveElemRef(ed parse.ElemDecl, schemaNS string, optional, unbounded bool) []parse.Field {
	refNS, refName := r.resolveQName(ed.Ref, schemaNS)
	refKey := refNS + " " + refName

	elem, ok := r.allElements[refKey]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: element ref not found: %s\n", ed.Ref)
		return nil
	}

	resolvedType := r.resolveTypeName(elem.TypeRef, refNS)
	if resolvedType == "" {
		resolvedType = "string"
	}

	doc := ed.Doc
	if doc == "" {
		doc = elem.Doc
	}

	isSlice := unbounded || ed.MaxOccurs == "unbounded" || (ed.MaxOccurs != "" && ed.MaxOccurs != "1" && ed.MaxOccurs != "0")
	isOmit := optional || ed.MinOccurs == "0"

	f := parse.Field{
		GoName: parse.GoName(refName),
		XMLTag: buildXMLTag(refNS, refName, false),
		GoType: resolvedType,
		TypeNS: refNS,
		Omit:   isOmit,
		Slice:  isSlice,
		Doc:    doc,
	}
	if isSlice {
		f.GoType = "[]" + strings.TrimPrefix(f.GoType, "[]")
	} else {
		f.GoType = "*" + strings.TrimPrefix(f.GoType, "*")
	}

	result := []parse.Field{f}
	// Expand substitutionGroup: add a parallel field for each concrete member.
	for _, member := range r.sortedSubstMembers(refKey) {
		memberElem, ok := r.allElements[member.NS+" "+member.Name]
		if !ok {
			continue
		}
		memberGoType := r.resolveTypeName(memberElem.TypeRef, member.NS)
		if memberGoType == "" {
			memberGoType = "string"
		}
		var memberTypeNS string
		if !parse.IsBuiltinGoType(memberGoType) && memberElem.TypeRef != "" {
			memberTypeNS, _ = r.resolveQName(memberElem.TypeRef, member.NS)
		}
		if isSlice {
			memberGoType = "[]" + strings.TrimPrefix(memberGoType, "[]")
		} else {
			memberGoType = "*" + strings.TrimPrefix(memberGoType, "*")
		}
		result = append(result, parse.Field{
			GoName: parse.GoName(member.Name),
			XMLTag: buildXMLTag(member.NS, member.Name, false),
			GoType: memberGoType,
			TypeNS: memberTypeNS,
			Omit:   !isSlice,
			Slice:  isSlice,
		})
	}
	return result
}

// resolveGroupRef expands a named model group reference into Fields.
// gr.MinOccurs/MaxOccurs are effective values pre-computed in the parse phase.
func (r *Resolver) resolveGroupRef(gr parse.GroupRef, schemaNS string, parentOptional, parentUnbounded bool) []parse.Field {
	groupNS, groupName := r.resolveQName(gr.Ref, schemaNS)
	groupKey := groupNS + " " + groupName
	group, ok := r.allGroups[groupKey]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: group ref not found: %s\n", gr.Ref)
		return nil
	}
	if group.Content == nil {
		return nil
	}
	// Combine caller context with this group ref's own effective values.
	optional := parentOptional || gr.MinOccurs == "0"
	unbounded := parentUnbounded || gr.MaxOccurs == "unbounded"
	return r.resolveContentModel(group.Content, groupNS, optional, unbounded)
}

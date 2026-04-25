package resolve

import (
	"fmt"
	"os"
	"strings"

	"xsd2go-lite2/internal/parse"
)

// resolveContentModel recursively converts a ContentModel to []Field.
//
// inChoice: some ancestor compositor is "choice"; all descendants become optional.
// parentUnbounded: OR semantics — if any ancestor compositor has maxOccurs=unbounded,
// all descendant Fields become slices regardless of their own maxOccurs.
func (r *Resolver) resolveContentModel(cm *parse.ContentModel, schemaNS string, inChoice, parentUnbounded bool) []parse.Field {
	// OR semantics: compositor unbounded propagates down to all descendant Fields.
	unbounded := parentUnbounded || cm.MaxOccurs == "unbounded"

	var fields []parse.Field
	for _, ed := range cm.Elems {
		fields = append(fields, r.resolveElemDecl(ed, schemaNS, inChoice, unbounded)...)
	}
	for _, gr := range cm.Groups {
		fields = append(fields, r.resolveGroupRef(gr, schemaNS, inChoice, unbounded)...)
	}
	for i := range cm.Children {
		child := &cm.Children[i]
		childInChoice := inChoice || child.Kind == "choice"
		fields = append(fields, r.resolveContentModel(child, schemaNS, childInChoice, unbounded)...)
	}
	return fields
}

// resolveElemDecl converts an ElemDecl to zero or more Fields.
func (r *Resolver) resolveElemDecl(ed parse.ElemDecl, schemaNS string, inChoice, parentUnbounded bool) []parse.Field {
	if ed.Ref != "" {
		return r.resolveElemRef(ed, schemaNS, inChoice, parentUnbounded)
	}

	goType := r.resolveTypeName(ed.TypeRef, schemaNS)
	if goType == "" {
		goType = "string"
	}

	var typeNS string
	if !parse.IsBuiltinGoType(goType) && ed.TypeRef != "" {
		typeNS, _ = r.resolveQName(ed.TypeRef, schemaNS)
	}

	// OR semantics: either the ancestor compositor or this element's own maxOccurs makes it a slice.
	isSlice := parentUnbounded || ed.MaxOccurs == "unbounded" || (ed.MaxOccurs != "" && ed.MaxOccurs != "1" && ed.MaxOccurs != "0")
	isOmit := inChoice || ed.MinOccurs == "0"

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
func (r *Resolver) resolveElemRef(ed parse.ElemDecl, schemaNS string, inChoice, parentUnbounded bool) []parse.Field {
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

	// OR semantics: either the ancestor compositor or this element's own maxOccurs makes it a slice.
	isSlice := parentUnbounded || ed.MaxOccurs == "unbounded" || (ed.MaxOccurs != "" && ed.MaxOccurs != "1" && ed.MaxOccurs != "0")
	isOmit := inChoice || ed.MinOccurs == "0"

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

// resolveGroupRef expands a model group reference into Fields.
func (r *Resolver) resolveGroupRef(gr parse.GroupRef, schemaNS string, inChoice, parentUnbounded bool) []parse.Field {
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
	// The group ref itself may also have maxOccurs=unbounded.
	grUnbounded := parentUnbounded || gr.MaxOccurs == "unbounded"
	return r.resolveContentModel(group.Content, groupNS, inChoice, grUnbounded)
}

package resolve

import (
	"fmt"
	"os"
	"strings"

	"xsd2go-lite/internal/parse"
)

// resolveElemRef handles element ref="ns:name" including substitutionGroup expansion.
func (r *Resolver) resolveElemRef(rf parse.RawField, schemaNS string) []parse.Field {
	refNS, refName := r.resolveQName(rf.Ref, schemaNS)
	refKey := refNS + " " + refName
	elemType := r.allElements[refKey]
	if elemType == "" {
		fmt.Fprintf(os.Stderr, "warn: element ref not found: %s\n", rf.Ref)
		elemType = "string"
	}
	resolvedType := r.resolveTypeName(elemType, refNS)
	if resolvedType == "" {
		resolvedType = "string"
	}
	doc := rf.Doc
	if doc == "" {
		doc = r.allElementDocs[refKey]
	}
	f := parse.Field{
		GoName: parse.GoName(refName),
		XMLTag: buildXMLTag(refNS, refName, false),
		GoType: resolvedType,
		TypeNS: refNS,
		Omit:   rf.MinOccurs == "0" || rf.MinOccurs == "",
		Slice:  rf.MaxOccurs == "unbounded" || (rf.MaxOccurs != "" && rf.MaxOccurs != "1"),
		Doc:    doc,
	}
	if f.Slice {
		f.GoType = "[]" + strings.TrimPrefix(f.GoType, "[]")
	} else {
		f.GoType = "*" + strings.TrimPrefix(f.GoType, "*")
	}
	result := []parse.Field{f}
	// Expand substitutionGroup: add a parallel field for each concrete member.
	for _, member := range r.sortedSubstMembers(refKey) {
		memberElemType := r.allElements[member.NS+" "+member.Name]
		memberGoType := r.resolveTypeName(memberElemType, member.NS)
		if memberGoType == "" {
			memberGoType = "string"
		}
		var memberTypeNS string
		if !parse.IsBuiltinGoType(memberGoType) && memberElemType != "" {
			memberTypeNS, _ = r.resolveQName(memberElemType, member.NS)
		}
		if f.Slice {
			memberGoType = "[]" + strings.TrimPrefix(memberGoType, "[]")
		} else {
			memberGoType = "*" + strings.TrimPrefix(memberGoType, "*")
		}
		result = append(result, parse.Field{
			GoName: parse.GoName(member.Name),
			XMLTag: buildXMLTag(member.NS, member.Name, false),
			GoType: memberGoType,
			TypeNS: memberTypeNS,
			Omit:   !f.Slice,
			Slice:  f.Slice,
		})
	}
	return result
}

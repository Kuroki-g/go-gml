package resolve

import (
	"fmt"
	"os"
	"strings"

	"xsd2go-lite2/internal/parse"
)

// resolveAttrDecl converts an AttrDecl to zero or more Fields.
// prohibited attributes produce no Field (they are filtered here).
func (r *Resolver) resolveAttrDecl(ad parse.AttrDecl, schemaNS string) []parse.Field {
	if ad.Use == parse.UseProhibited {
		return nil
	}
	if ad.Ref != "" {
		return r.resolveAttrRef(ad.Ref, ad.Use, schemaNS)
	}

	goType := r.resolveTypeName(ad.TypeRef, schemaNS)
	if goType == "" {
		goType = "string"
	}

	isOmit := ad.Use != parse.UseRequired
	if isOmit {
		goType = "*" + strings.TrimPrefix(goType, "*")
	}

	return []parse.Field{{
		GoName: parse.GoName(ad.Name),
		XMLTag: buildXMLTag("", ad.Name, true),
		GoType: goType,
		IsAttr: true,
		Omit:   isOmit,
		Doc:    ad.Doc,
	}}
}

// resolveAttrRef resolves a global attribute ref QName to a Field.
func (r *Resolver) resolveAttrRef(ref, use, schemaNS string) []parse.Field {
	if use == parse.UseProhibited {
		return nil
	}
	refNS, refName := r.resolveQName(ref, schemaNS)
	refKey := refNS + " " + refName

	goType := "string"
	if ga, ok := r.allGlobalAttrs[refKey]; ok && ga.TypeRef != "" {
		if t := r.resolveTypeName(ga.TypeRef, refNS); t != "" {
			goType = t
		}
	}

	isOmit := use != parse.UseRequired
	if isOmit {
		goType = "*" + strings.TrimPrefix(goType, "*")
	}

	return []parse.Field{{
		GoName: parse.GoName(refName),
		XMLTag: buildXMLTag(refNS, refName, true),
		GoType: goType,
		IsAttr: true,
		Omit:   isOmit,
	}}
}

// resolveAttrGroupRef recursively expands an attributeGroup ref to Fields.
func (r *Resolver) resolveAttrGroupRef(ref, schemaNS string) []parse.Field {
	agNS, agName := r.resolveQName(ref, schemaNS)
	agKey := agNS + " " + agName
	ag, ok := r.allAttrGroups[agKey]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: attributeGroup not found: %s\n", ref)
		return nil
	}

	var fields []parse.Field
	for _, ad := range ag.Attrs {
		fields = append(fields, r.resolveAttrDecl(ad, agNS)...)
	}
	for _, nestedRef := range ag.NestedGroups {
		fields = append(fields, r.resolveAttrGroupRef(nestedRef, agNS)...)
	}
	return fields
}

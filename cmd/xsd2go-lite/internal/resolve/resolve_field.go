package resolve

import (
	"fmt"
	"os"
	"strings"

	"xsd2go-lite/internal/parse"
)

// resolveRawField converts a single RawField to zero or more Fields.
func (r *Resolver) resolveRawField(rf parse.RawField, schemaNS string, visiting map[string]bool) []parse.Field {
	switch {
	case rf.Ref != "" && strings.HasPrefix(rf.Ref, "__base_attrs__:"):
		return r.resolveBaseAttrsRef(rf, schemaNS, visiting)
	case rf.Ref != "" && strings.HasPrefix(rf.Ref, "__base__:"):
		return r.resolveBaseRef(rf, schemaNS, visiting)
	case rf.Ref != "":
		return r.resolveElemRef(rf, schemaNS)
	case rf.IsAttr && rf.AttrRef != "":
		return r.resolveAttrRef(rf, schemaNS)
	case rf.GroupRef != "":
		return r.resolveGroupRef(rf, schemaNS, visiting)
	case rf.IsChar:
		return r.resolveChardata(rf, schemaNS)
	default:
		return r.resolveLocalField(rf, schemaNS)
	}
}

// resolveBaseAttrsRef handles __base_attrs__: — inherit only attribute fields from base.
func (r *Resolver) resolveBaseAttrsRef(rf parse.RawField, schemaNS string, visiting map[string]bool) []parse.Field {
	baseRef := strings.TrimPrefix(rf.Ref, "__base_attrs__:")
	baseNS, baseName := r.resolveQName(baseRef, schemaNS)
	baseKey := baseNS + " " + baseName
	if baseCT, ok := r.allComplexTypes[baseKey]; ok {
		r.resolveComplexType(baseCT, visiting)
		var attrFields []parse.Field
		for _, f := range baseCT.Fields {
			if f.IsAttr {
				attrFields = append(attrFields, f)
			}
		}
		return attrFields
	}
	return nil
}

// resolveBaseRef handles __base__: — inherit all fields from base type.
func (r *Resolver) resolveBaseRef(rf parse.RawField, schemaNS string, visiting map[string]bool) []parse.Field {
	baseRef := strings.TrimPrefix(rf.Ref, "__base__:")
	baseNS, baseName := r.resolveQName(baseRef, schemaNS)
	baseKey := baseNS + " " + baseName
	if baseCT, ok := r.allComplexTypes[baseKey]; ok {
		r.resolveComplexType(baseCT, visiting)
		return baseCT.Fields
	}
	// If base is a simpleType, add a chardata field.
	if baseST, ok := r.allSimpleTypes[baseKey]; ok {
		goType := baseST.GoType
		if goType == "" {
			goType = "string"
		}
		return []parse.Field{{GoName: "Value", GoType: goType, IsChar: true, XMLTag: ",chardata"}}
	}
	// If it's an XSD built-in
	if goType := parse.XsBuiltinGoType(baseRef); goType != "" {
		return []parse.Field{{GoName: "Value", GoType: goType, IsChar: true, XMLTag: ",chardata"}}
	}
	fmt.Fprintf(os.Stderr, "warn: base type not found: %s\n", baseRef)
	return nil
}

// resolveAttrRef handles attribute ref.
func (r *Resolver) resolveAttrRef(rf parse.RawField, schemaNS string) []parse.Field {
	// Nested attributeGroup ref (encoded as "__group__:ref")
	if strings.HasPrefix(rf.AttrRef, "__group__:") {
		agRef := strings.TrimPrefix(rf.AttrRef, "__group__:")
		return r.resolveAttrGroupRef(agRef, schemaNS)
	}
	refNS, refName := r.resolveQName(rf.AttrRef, schemaNS)
	refKey := refNS + " " + refName
	goType := "string"
	if ga, ok := r.allAttrGroups[refKey]; ok {
		_ = ga // attribute group as attribute ref is unusual; treat as string
	}
	return []parse.Field{{
		GoName: parse.GoName(refName),
		XMLTag: buildXMLTag(refNS, refName, true),
		GoType: goType,
		IsAttr: true,
		Omit:   true,
	}}
}

// resolveGroupRef handles group ref.
func (r *Resolver) resolveGroupRef(rf parse.RawField, schemaNS string, visiting map[string]bool) []parse.Field {
	groupNS, groupName := r.resolveQName(rf.GroupRef, schemaNS)
	groupKey := groupNS + " " + groupName
	groupFields, ok := r.allGroups[groupKey]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: group ref not found: %s\n", rf.GroupRef)
		return nil
	}
	var result []parse.Field
	for _, gf := range groupFields {
		result = append(result, r.resolveRawField(gf, groupNS, visiting)...)
	}
	return result
}

// resolveChardata handles chardata (simpleContent) fields.
func (r *Resolver) resolveChardata(rf parse.RawField, schemaNS string) []parse.Field {
	goType := r.resolveTypeName(rf.TypeRef, schemaNS)
	// chardata for simpleContent is always string-like
	if goType == "" || goType == "string" || parse.IsStructType(goType) {
		goType = "string"
	}
	return []parse.Field{{
		GoName: "Value",
		XMLTag: ",chardata",
		GoType: goType,
		IsChar: true,
	}}
}

// resolveLocalField handles normal named element or attribute.
func (r *Resolver) resolveLocalField(rf parse.RawField, schemaNS string) []parse.Field {
	if rf.LocalName == "" || rf.LocalName == "__chardata__" {
		return nil
	}

	goType := r.resolveTypeName(rf.TypeRef, schemaNS)
	if goType == "" {
		goType = "string"
	}

	var typeNS string
	if !parse.IsBuiltinGoType(goType) && rf.TypeRef != "" {
		typeNS, _ = r.resolveQName(rf.TypeRef, schemaNS)
	}

	isSlice := rf.MaxOccurs == "unbounded" || (rf.MaxOccurs != "" && rf.MaxOccurs != "1" && rf.MaxOccurs != "0")
	isOmit := rf.IsAttr && rf.MinOccurs != "required"
	if !rf.IsAttr {
		isOmit = rf.MinOccurs == "0" || rf.MinOccurs == ""
	}

	xmlTag := buildXMLTag(schemaNS, rf.LocalName, rf.IsAttr)
	if rf.IsAttr {
		// Attributes don't use namespace URI in tag unless explicitly from another NS.
		xmlTag = buildXMLTag("", rf.LocalName, true)
	}

	f := parse.Field{
		GoName: parse.GoName(rf.LocalName),
		XMLTag: xmlTag,
		GoType: goType,
		TypeNS: typeNS,
		IsAttr: rf.IsAttr,
		Omit:   isOmit,
		Slice:  isSlice,
		Doc:    rf.Doc,
	}

	if isSlice && !rf.IsAttr {
		f.GoType = "[]" + strings.TrimPrefix(goType, "[]")
	} else if !isSlice && !rf.IsAttr && !parse.IsBuiltinGoType(goType) {
		f.GoType = "*" + strings.TrimPrefix(goType, "*")
	} else if rf.IsAttr && isOmit {
		// XSD optional attribute (use="optional" or use omitted) → pointer type
		// so encoding/xml can distinguish absent (nil) from present-but-empty ("").
		f.GoType = "*" + strings.TrimPrefix(f.GoType, "*")
	}

	return []parse.Field{f}
}

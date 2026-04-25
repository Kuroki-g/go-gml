package parse

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// Parse reads one XSD file and returns a Schema (not yet resolved).
func Parse(filename string) (*Schema, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", filename, err)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	var raw xsdSchema
	var nsMap map[string]string

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", filename, err)
		}
		if se, ok := tok.(xml.StartElement); ok {
			nsMap = make(map[string]string)
			for _, attr := range se.Attr {
				if attr.Name.Space == "xmlns" {
					nsMap[attr.Name.Local] = attr.Value
				} else if attr.Name.Space == "" && attr.Name.Local == "xmlns" {
					nsMap[""] = attr.Value
				}
			}
			if err := decoder.DecodeElement(&raw, &se); err != nil {
				return nil, fmt.Errorf("decode %s: %w", filename, err)
			}
			break
		}
	}

	s := &Schema{
		TargetNamespace: raw.TargetNamespace,
		Filename:        filename,
		NSMap:           nsMap,
		Elements:        make(map[string]Element),
		GlobalAttrs:     make(map[string]Attribute),
		AttributeGroups: make(map[string]AttrGroup),
		Groups:          make(map[string]Group),
	}

	for _, el := range raw.Elements {
		doc := ""
		if el.Annotation != nil {
			doc = strings.TrimSpace(el.Annotation.Documentation)
		}
		typeRef := el.Type
		if typeRef == "" {
			typeRef = "string"
		}
		s.Elements[el.Name] = Element{
			Name:              el.Name,
			TypeRef:           typeRef,
			SubstitutionGroup: el.SubstGroup,
			Abstract:          el.Abstract == "true",
			Doc:               doc,
		}
	}

	for _, a := range raw.Attributes {
		s.GlobalAttrs[a.Name] = Attribute{
			Name:    a.Name,
			TypeRef: a.Type,
			Default: a.Default,
		}
	}

	for _, ag := range raw.AttrGroups {
		s.AttributeGroups[ag.Name] = collectAttrGroupDecl(ag)
	}

	for _, g := range raw.Groups {
		s.Groups[g.Name] = Group{
			Name:    g.Name,
			Content: collectGroupContent(g),
		}
	}

	for _, st := range raw.SimpleTypes {
		s.SimpleTypes = append(s.SimpleTypes, convertSimpleType(st))
	}

	for _, ct := range raw.ComplexTypes {
		s.ComplexTypes = append(s.ComplexTypes, convertComplexType(ct, raw.TargetNamespace))
	}

	return s, nil
}

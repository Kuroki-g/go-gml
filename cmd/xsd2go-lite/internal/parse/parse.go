package parse

import (
	"encoding/xml"
	"fmt"
	"io"
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

	// We decode manually to capture namespace declarations.
	decoder := xml.NewDecoder(f)

	var raw xsdSchema
	var nsMap map[string]string

	// Find the root element and capture its NS declarations before Decode.
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
		Elements:        make(map[string]string),
		ElementDocs:     make(map[string]string),
		SubstGroups:     make(map[string]string),
		GlobalAttrs:     make(map[string]Attribute),
		AttributeGroups: make(map[string]AttrGroup),
		Groups:          make(map[string][]RawField),
	}

	// Global elements
	for _, el := range raw.Elements {
		t := el.Type
		if t == "" {
			t = "string"
		}
		s.Elements[el.Name] = t
		if el.SubstGroup != "" {
			s.SubstGroups[el.Name] = el.SubstGroup
		}
		if el.Annotation != nil {
			doc := strings.TrimSpace(el.Annotation.Documentation)
			if doc != "" {
				s.ElementDocs[el.Name] = doc
			}
		}
	}

	// Global attributes
	for _, a := range raw.Attributes {
		s.GlobalAttrs[a.Name] = Attribute{
			Name:    a.Name,
			Type:    a.Type,
			Default: a.Default,
		}
	}

	// attributeGroups
	for _, ag := range raw.AttrGroups {
		grp := AttrGroup{Name: ag.Name}
		grp.RawFields = collectAttrGroupFields(ag)
		s.AttributeGroups[ag.Name] = grp
	}

	// groups
	for _, g := range raw.Groups {
		fields := collectGroupFields(g)
		s.Groups[g.Name] = fields
	}

	// simpleTypes
	for _, st := range raw.SimpleTypes {
		s.SimpleTypes = append(s.SimpleTypes, convertSimpleType(st))
	}

	// complexTypes
	for _, ct := range raw.ComplexTypes {
		s.ComplexTypes = append(s.ComplexTypes, convertComplexType(ct, raw.TargetNamespace))
	}

	return s, nil
}

// XsBuiltinGoType maps XSD built-in type names to Go types.
func XsBuiltinGoType(xsType string) string {
	// Strip namespace prefix if present
	local := xsType
	if idx := strings.LastIndex(xsType, ":"); idx >= 0 {
		local = xsType[idx+1:]
	}
	switch local {
	case "string", "anyURI", "NCName", "QName", "NMTOKEN", "NMTOKENS",
		"Name", "language", "token", "normalizedString", "ID", "IDREF",
		"IDREFS", "ENTITY", "ENTITIES", "notation",
		"base64Binary", "hexBinary",
		"duration", "dateTime", "date", "time", "gYear", "gYearMonth",
		"gMonth", "gMonthDay", "gDay":
		return "string"
	case "boolean":
		return "bool"
	case "int":
		return "int32"
	case "long":
		return "int64"
	case "short":
		return "int16"
	case "byte":
		return "int8"
	case "unsignedLong":
		return "uint64"
	case "unsignedInt":
		return "uint32"
	case "unsignedShort":
		return "uint16"
	case "unsignedByte":
		return "uint8"
	case "nonNegativeInteger", "positiveInteger":
		return "uint"
	case "integer", "nonPositiveInteger", "negativeInteger":
		return "int"
	case "double":
		return "float64"
	// xs:decimal is an arbitrary-precision decimal (W3C XSD Part 2, §3.2.3), not a
	// binary floating-point type. The semantically correct Go representation would
	// require an external arbitrary-precision decimal library (math/big.Float is
	// binary, not decimal). Mapping to float64 is a known compromise: the precision
	// (~15 digits) is sufficient for all GML use cases (coordinates, arc-minutes,
	// arc-seconds), but the type meaning differs. Changing this would require adding
	// an external dependency, which is prohibited for parser modules.
	case "decimal":
		return "float64"
	case "float":
		return "float32"
	case "anyType", "anySimpleType":
		return "string"
	default:
		return "" // unknown → caller decides
	}
}

// decodeXML decodes an XSD file into v.
func decodeXML(r io.Reader, v any) error {
	return xml.NewDecoder(r).Decode(v)
}

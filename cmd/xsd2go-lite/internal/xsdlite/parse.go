package xsdlite

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// ---- XSD raw XML structures (for decoding only) ----

type xsdSchema struct {
	XMLName         xml.Name         `xml:"schema"`
	TargetNamespace string           `xml:"targetNamespace,attr"`
	Includes        []xsdInclude     `xml:"include"`
	Imports         []xsdImport      `xml:"import"`
	ComplexTypes    []xsdComplexType `xml:"complexType"`
	SimpleTypes     []xsdSimpleType  `xml:"simpleType"`
	Elements        []xsdElement     `xml:"element"`
	Attributes      []xsdAttribute   `xml:"attribute"`
	AttrGroups      []xsdAttrGroup   `xml:"attributeGroup"`
	Groups          []xsdGroup       `xml:"group"`
	NSDecls         []xml.Attr       // namespace declarations captured manually
}

type xsdInclude struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
}

type xsdImport struct {
	Namespace      string `xml:"namespace,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
}

type xsdComplexType struct {
	Name           string             `xml:"name,attr"`
	Abstract       string             `xml:"abstract,attr"`
	Mixed          string             `xml:"mixed,attr"`
	Annotation     *xsdAnnotation     `xml:"annotation"`
	SimpleContent  *xsdSimpleContent  `xml:"simpleContent"`
	ComplexContent *xsdComplexContent `xml:"complexContent"`
	Sequence       *xsdSequence       `xml:"sequence"`
	All            *xsdSequence       `xml:"all"`
	Choice         *xsdSequence       `xml:"choice"`
	AttrGroups     []xsdAttrGroupRef  `xml:"attributeGroup"`
	Attributes     []xsdAttribute     `xml:"attribute"`
}

type xsdSimpleContent struct {
	Extension   *xsdExtension   `xml:"extension"`
	Restriction *xsdRestriction `xml:"restriction"`
}

type xsdComplexContent struct {
	Extension   *xsdExtension   `xml:"extension"`
	Restriction *xsdRestriction `xml:"restriction"`
}

type xsdExtension struct {
	Base       string            `xml:"base,attr"`
	Sequence   *xsdSequence      `xml:"sequence"`
	All        *xsdSequence      `xml:"all"`
	Choice     *xsdSequence      `xml:"choice"`
	AttrGroups []xsdAttrGroupRef `xml:"attributeGroup"`
	Attributes []xsdAttribute    `xml:"attribute"`
}

type xsdRestriction struct {
	Base       string            `xml:"base,attr"`
	Sequence   *xsdSequence      `xml:"sequence"`
	AttrGroups []xsdAttrGroupRef `xml:"attributeGroup"`
	Attributes []xsdAttribute    `xml:"attribute"`
}

type xsdSequence struct {
	Elements   []xsdElement      `xml:"element"`
	Sequences  []xsdSequence     `xml:"sequence"`
	Choices    []xsdSequence     `xml:"choice"`
	Groups     []xsdGroupRef     `xml:"group"`
	AttrGroups []xsdAttrGroupRef `xml:"attributeGroup"`
	Attributes []xsdAttribute    `xml:"attribute"`
	MinOccurs  string            `xml:"minOccurs,attr"`
	MaxOccurs  string            `xml:"maxOccurs,attr"`
}

type xsdElement struct {
	Name        string          `xml:"name,attr"`
	Ref         string          `xml:"ref,attr"`
	Type        string          `xml:"type,attr"`
	MinOccurs   string          `xml:"minOccurs,attr"`
	MaxOccurs   string          `xml:"maxOccurs,attr"`
	Annotation  *xsdAnnotation  `xml:"annotation"`
	ComplexType *xsdComplexType `xml:"complexType"`
	SimpleType  *xsdSimpleType  `xml:"simpleType"`
	SubstGroup  string          `xml:"substitutionGroup,attr"`
	Abstract    string          `xml:"abstract,attr"`
}

type xsdAttribute struct {
	Name       string         `xml:"name,attr"`
	Ref        string         `xml:"ref,attr"`
	Type       string         `xml:"type,attr"`
	Use        string         `xml:"use,attr"`
	Default    string         `xml:"default,attr"`
	Annotation *xsdAnnotation `xml:"annotation"`
	Simple     *xsdSimpleType `xml:"simpleType"`
}

type xsdAttrGroup struct {
	Name       string            `xml:"name,attr"`
	Ref        string            `xml:"ref,attr"`
	Attributes []xsdAttribute    `xml:"attribute"`
	AttrGroups []xsdAttrGroupRef `xml:"attributeGroup"`
}

type xsdAttrGroupRef struct {
	Ref string `xml:"ref,attr"`
}

type xsdGroup struct {
	Name     string       `xml:"name,attr"`
	Ref      string       `xml:"ref,attr"`
	Sequence *xsdSequence `xml:"sequence"`
	Choice   *xsdSequence `xml:"choice"`
	All      *xsdSequence `xml:"all"`
}

type xsdGroupRef struct {
	Ref       string `xml:"ref,attr"`
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
}

type xsdAnnotation struct {
	Documentation string `xml:"documentation"`
}

type xsdSimpleType struct {
	Name        string        `xml:"name,attr"`
	Restriction *xsdSRestrict `xml:"restriction"`
	List        *xsdList      `xml:"list"`
	Union       *xsdUnion     `xml:"union"`
}

type xsdSRestrict struct {
	Base string `xml:"base,attr"`
}

type xsdList struct {
	ItemType string `xml:"itemType,attr"`
}

type xsdUnion struct {
	MemberTypes string `xml:"memberTypes,attr"`
}

// ---- Internal representation ----

// Schema is the fully-loaded (but not yet resolved) schema data.
type Schema struct {
	TargetNamespace string
	Filename        string
	NSMap           map[string]string // prefix → URI
	ComplexTypes    []ComplexType
	SimpleTypes     []SimpleType
	Elements        map[string]string    // name → type (global elements)
	ElementDocs     map[string]string    // name → xs:annotation/xs:documentation (global elements)
	SubstGroups     map[string]string    // concrete element name → abstract element head (QName)
	GlobalAttrs     map[string]Attribute // name → Attribute
	AttributeGroups map[string]AttrGroup
	Groups          map[string][]rawField
}

// rawField holds unresolved field info (before resolve phase).
type rawField struct {
	LocalName string // element/attribute local name (may be empty if ref)
	Ref       string // ref="ns:name" (for element refs)
	AttrRef   string // ref="ns:name" (for attribute refs)
	TypeRef   string // type="ns:name"
	IsAttr    bool
	IsChar    bool // simpleContent chardata
	MinOccurs string
	MaxOccurs string
	GroupRef  string // group ref="ns:name"
	Doc       string // xs:annotation/xs:documentation text
}

// ComplexType is the resolved complex type definition.
type ComplexType struct {
	Name       string
	Abstract   bool
	Mixed      bool
	Doc        string // xs:annotation/xs:documentation text
	RawFields  []rawField
	AttrGroups []string // list of attributeGroup ref names
	Source     string   // targetNamespace of the defining schema
	// After resolve:
	Fields []Field
}

// Field is a single resolved struct field.
type Field struct {
	GoName string
	XMLTag string
	GoType string
	TypeNS string // namespace URI of the referenced Go type (for omit filtering)
	IsAttr bool
	IsChar bool
	Omit   bool
	Slice  bool
	Doc    string // xs:annotation/xs:documentation text
}

// SimpleType is the resolved simple type definition.
type SimpleType struct {
	Name   string
	Kind   string // "list" | "restriction" | "union" | "alias"
	GoType string
}

// AttrGroup is an attribute group definition.
type AttrGroup struct {
	Name      string
	RawFields []rawField
}

// Attribute is a global attribute definition.
type Attribute struct {
	Name    string
	Type    string
	Default string
}

// parseSchema reads one XSD file and returns a Schema (not yet resolved).
func parseSchema(filename string) (*Schema, error) {
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
		Groups:          make(map[string][]rawField),
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

func convertSimpleType(st xsdSimpleType) SimpleType {
	r := SimpleType{Name: st.Name}
	switch {
	case st.List != nil:
		r.Kind = "list"
		r.GoType = "string"
	case st.Union != nil:
		r.Kind = "union"
		r.GoType = "string"
	case st.Restriction != nil:
		r.Kind = "restriction"
		r.GoType = xsBuiltinGoType(st.Restriction.Base)
	default:
		r.Kind = "alias"
		r.GoType = "string"
	}
	return r
}

func convertComplexType(ct xsdComplexType, ns string) ComplexType {
	result := ComplexType{
		Name:     ct.Name,
		Abstract: ct.Abstract == "true",
		Mixed:    ct.Mixed == "true",
		Source:   ns,
	}
	if ct.Annotation != nil {
		result.Doc = strings.TrimSpace(ct.Annotation.Documentation)
	}

	switch {
	case ct.SimpleContent != nil:
		if ext := ct.SimpleContent.Extension; ext != nil {
			// chardata field for the value
			result.RawFields = append(result.RawFields, rawField{
				LocalName: "__chardata__",
				TypeRef:   ext.Base,
				IsChar:    true,
			})
			result.RawFields = append(result.RawFields, collectExtensionFields(ext)...)
			for _, ag := range ext.AttrGroups {
				result.AttrGroups = append(result.AttrGroups, ag.Ref)
			}
		} else if res := ct.SimpleContent.Restriction; res != nil {
			result.RawFields = append(result.RawFields, rawField{
				LocalName: "__chardata__",
				TypeRef:   res.Base,
				IsChar:    true,
			})
			result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(res.AttrGroups)...)
			result.RawFields = append(result.RawFields, convertAttributes(res.Attributes)...)
		}

	case ct.ComplexContent != nil:
		if ext := ct.ComplexContent.Extension; ext != nil {
			// Base type placeholder (resolved later)
			result.RawFields = append(result.RawFields, rawField{
				Ref: "__base__:" + ext.Base,
			})
			result.RawFields = append(result.RawFields, collectExtensionFields(ext)...)
			result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(ext.AttrGroups)...)
		} else if res := ct.ComplexContent.Restriction; res != nil {
			// complexContent/restriction replaces the parent's element content model
			// but inherits its attribute declarations. Only the restriction's own
			// element fields are included; attribute fields are inherited from the base.
			if res.Sequence != nil {
				result.RawFields = append(result.RawFields, collectSequenceFields(res.Sequence)...)
			}
			result.RawFields = append(result.RawFields, convertAttributes(res.Attributes)...)
			result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(res.AttrGroups)...)
			// Inherit only attribute fields from the base type.
			result.RawFields = append(result.RawFields, rawField{
				Ref: "__base_attrs__:" + res.Base,
			})
		}

	default:
		// Direct sequence/all/choice
		if ct.Sequence != nil {
			result.RawFields = append(result.RawFields, collectSequenceFields(ct.Sequence)...)
		}
		if ct.All != nil {
			result.RawFields = append(result.RawFields, collectSequenceFields(ct.All)...)
		}
		if ct.Choice != nil {
			result.RawFields = append(result.RawFields, collectSequenceFields(ct.Choice)...)
		}
		result.AttrGroups = append(result.AttrGroups, collectAttrGroupRefs(ct.AttrGroups)...)
		result.RawFields = append(result.RawFields, convertAttributes(ct.Attributes)...)
	}

	return result
}

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

// xsBuiltinGoType maps XSD built-in type names to Go types.
func xsBuiltinGoType(xsType string) string {
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
	case "integer", "positiveInteger", "nonNegativeInteger", "nonPositiveInteger",
		"negativeInteger", "long", "int", "short", "byte",
		"unsignedLong", "unsignedInt", "unsignedShort", "unsignedByte":
		return "int"
	case "double", "decimal":
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

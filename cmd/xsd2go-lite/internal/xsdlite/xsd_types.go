package xsdlite

import "encoding/xml"

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

package resolve

import "xsd2go-lite/internal/parse"

// substMember holds a concrete element that is a member of a substitutionGroup.
type substMember struct {
	NS   string // target namespace URI of the concrete element
	Name string // local name of the concrete element
}

// Resolver holds all loaded schemas and resolves cross-schema references.
type Resolver struct {
	// All loaded schemas keyed by absolute filename.
	schemas map[string]*parse.Schema
	// All complexTypes across all schemas, keyed by "targetNS localName".
	allComplexTypes map[string]*parse.ComplexType
	// All simpleTypes across all schemas, keyed by "targetNS localName".
	allSimpleTypes map[string]*parse.SimpleType
	// All global elements across all schemas, keyed by "targetNS localName" → type string.
	allElements map[string]string
	// All global element docs, keyed by "targetNS localName" → documentation text.
	allElementDocs map[string]string
	// All attributeGroups across all schemas, keyed by "targetNS localName".
	allAttrGroups map[string]*parse.AttrGroup
	// All groups across all schemas, keyed by "targetNS localName".
	allGroups map[string][]parse.RawField
	// substHeads maps abstract element key ("ns localName") → concrete members.
	substHeads map[string][]substMember
	// nsMaps maps targetNamespace → {prefix → uri}, built from each loaded schema's
	// xmlns declarations. Used by resolveQName to deterministically resolve prefixes
	// against the schema that defined the type being resolved.
	nsMaps map[string]map[string]string
	// catalog maps namespace URI → local XSD file path. Used to resolve
	// external imports (HTTP schemaLocation) without modifying the original XSD.
	catalog map[string]string
}

// NewResolver creates a new Resolver.
func NewResolver() *Resolver {
	return &Resolver{
		schemas:         make(map[string]*parse.Schema),
		allComplexTypes: make(map[string]*parse.ComplexType),
		allSimpleTypes:  make(map[string]*parse.SimpleType),
		allElements:     make(map[string]string),
		allElementDocs:  make(map[string]string),
		allAttrGroups:   make(map[string]*parse.AttrGroup),
		allGroups:       make(map[string][]parse.RawField),
		substHeads:      make(map[string][]substMember),
		nsMaps:          make(map[string]map[string]string),
		catalog:         make(map[string]string),
	}
}

// AddCatalogEntry registers a namespace URI → local file path mapping.
// When an XSD import has an external (HTTP) schemaLocation for this namespace,
// the local path is used instead.
func (r *Resolver) AddCatalogEntry(namespaceURI, localPath string) {
	r.catalog[namespaceURI] = localPath
}

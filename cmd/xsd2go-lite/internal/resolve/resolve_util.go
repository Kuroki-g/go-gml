package resolve

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"xsd2go-lite/internal/parse"
)

// resolveAttrGroupRef resolves an attributeGroup ref to Fields.
func (r *Resolver) resolveAttrGroupRef(ref, schemaNS string) []parse.Field {
	agNS, agName := r.resolveQName(ref, schemaNS)
	agKey := agNS + " " + agName
	ag, ok := r.allAttrGroups[agKey]
	if !ok {
		fmt.Fprintf(os.Stderr, "warn: attributeGroup not found: %s\n", ref)
		return nil
	}

	var fields []parse.Field
	for _, rf := range ag.RawFields {
		// Nested group ref in attrGroup
		if rf.IsAttr && strings.HasPrefix(rf.AttrRef, "__group__:") {
			nestedRef := strings.TrimPrefix(rf.AttrRef, "__group__:")
			fields = append(fields, r.resolveAttrGroupRef(nestedRef, agNS)...)
			continue
		}
		f := r.resolveRawField(rf, agNS, nil)
		fields = append(fields, f...)
	}
	return fields
}

// resolveTypeName converts a QName type reference to a Go type string.
func (r *Resolver) resolveTypeName(typeRef, schemaNS string) string {
	if typeRef == "" || typeRef == "string" {
		return "string"
	}
	if typeRef == "__inline__" {
		return "string"
	}

	// Check XSD built-in first
	if goType := parse.XsBuiltinGoType(typeRef); goType != "" {
		return goType
	}

	ns, name := r.resolveQName(typeRef, schemaNS)
	key := ns + " " + name

	// simpleType?
	if st, ok := r.allSimpleTypes[key]; ok {
		// GoType can be "" when the simpleType is a restriction of a non-builtin
		// base (e.g. restriction base="someCustomType"). Fall back to string so
		// callers never receive an empty type string that would generate invalid Go.
		if st.GoType == "" {
			return "string"
		}
		return st.GoType
	}
	// complexType?
	if _, ok := r.allComplexTypes[key]; ok {
		return parse.GoTypeName(name)
	}

	// Fallback: use the local name as a struct type
	fmt.Fprintf(os.Stderr, "warn: type not found: %s (ns=%s)\n", typeRef, schemaNS)
	return "string"
}

// resolveQName resolves a possibly-prefixed name ("gml:Foo" or "Foo") to (ns, localName).
func (r *Resolver) resolveQName(qname, schemaNS string) (string, string) {
	idx := strings.Index(qname, ":")
	if idx < 0 {
		return schemaNS, qname
	}
	prefix := qname[:idx]
	local := qname[idx+1:]

	// First, try the namespace map for the schema that defines the type being
	// resolved (identified by schemaNS). This is deterministic and correct:
	// the prefix must be declared in the schema that uses it.
	if nsMap, ok := r.nsMaps[schemaNS]; ok {
		if uri, ok := nsMap[prefix]; ok {
			return uri, local
		}
	}
	// Fall back: search other schemas in deterministic (sorted) order.
	keys := make([]string, 0, len(r.nsMaps))
	for k := range r.nsMaps {
		keys = append(keys, k)
	}
	sortStrings(keys)
	for _, k := range keys {
		if k == schemaNS {
			continue // already tried above
		}
		if uri, ok := r.nsMaps[k][prefix]; ok {
			return uri, local
		}
	}
	// Unknown prefix — return as-is with original schemaNS.
	fmt.Fprintf(os.Stderr, "warn: unknown prefix %q in %s\n", prefix, qname)
	return schemaNS, local
}

// buildXMLTag constructs the xml struct tag value.
func buildXMLTag(ns, localName string, isAttr bool) string {
	var sb strings.Builder
	if ns != "" {
		sb.WriteString(ns)
		sb.WriteString(" ")
	}
	sb.WriteString(localName)
	if isAttr {
		sb.WriteString(",attr")
	}
	return sb.String()
}

func isHTTP(loc string) bool {
	return strings.HasPrefix(loc, "http://") || strings.HasPrefix(loc, "https://")
}

func sortStrings(s []string) {
	sort.Strings(s)
}

func decodeXML(r io.Reader, v any) error {
	return xml.NewDecoder(r).Decode(v)
}

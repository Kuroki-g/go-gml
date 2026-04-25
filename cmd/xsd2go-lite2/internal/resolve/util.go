package resolve

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"xsd2go-lite2/internal/parse"
)

// resolveTypeName converts a QName type reference to a Go type string.
func (r *Resolver) resolveTypeName(typeRef, schemaNS string) string {
	if typeRef == "" || typeRef == "string" {
		return "string"
	}
	if goType := parse.XsBuiltinGoType(typeRef); goType != "" {
		return goType
	}
	ns, name := r.resolveQName(typeRef, schemaNS)
	key := ns + " " + name
	if st, ok := r.allSimpleTypes[key]; ok {
		if st.GoType == "" {
			return "string"
		}
		return st.GoType
	}
	if _, ok := r.allComplexTypes[key]; ok {
		return parse.GoTypeName(name)
	}
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

	if nsMap, ok := r.nsMaps[schemaNS]; ok {
		if uri, ok := nsMap[prefix]; ok {
			return uri, local
		}
	}
	// Fall back: search other schemas in deterministic order.
	keys := make([]string, 0, len(r.nsMaps))
	for k := range r.nsMaps {
		keys = append(keys, k)
	}
	sortStrings(keys)
	for _, k := range keys {
		if k == schemaNS {
			continue
		}
		if uri, ok := r.nsMaps[k][prefix]; ok {
			return uri, local
		}
	}
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

// sortedSubstMembers returns all transitive substitution group members for a
// head element key, sorted deterministically by (NS, Name).
func (r *Resolver) sortedSubstMembers(headKey string) []substMember {
	var result []substMember
	r.collectSubstMembers(headKey, make(map[string]bool), &result)
	if len(result) == 0 {
		return nil
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].NS != result[j].NS {
			return result[i].NS < result[j].NS
		}
		return result[i].Name < result[j].Name
	})
	return result
}

func (r *Resolver) collectSubstMembers(headKey string, visited map[string]bool, out *[]substMember) {
	if visited[headKey] {
		return
	}
	visited[headKey] = true
	for _, m := range r.substHeads[headKey] {
		*out = append(*out, m)
		r.collectSubstMembers(m.NS+" "+m.Name, visited, out)
	}
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

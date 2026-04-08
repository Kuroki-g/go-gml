package resolve

import (
	"sort"
	"strings"

	"xsd2go-lite/internal/parse"
)

// ResolveAll resolves all loaded schemas: flattens inheritance, inlines attributeGroups.
func (r *Resolver) ResolveAll(targetNS string) []*parse.ComplexType {
	// Resolve each complex type's fields.
	for _, ct := range r.allComplexTypes {
		if len(ct.Fields) == 0 {
			r.resolveComplexType(ct, nil)
		}
	}

	// Return only types belonging to the target namespace.
	var result []*parse.ComplexType
	for key, ct := range r.allComplexTypes {
		if strings.HasPrefix(key, targetNS+" ") {
			result = append(result, ct)
		}
	}
	return result
}

// resolveComplexType builds the final Fields slice (depth-first, handles cycles).
func (r *Resolver) resolveComplexType(ct *parse.ComplexType, visiting map[string]bool) {
	if visiting == nil {
		visiting = make(map[string]bool)
	}
	key := ct.Source + " " + ct.Name
	if visiting[key] {
		return // cycle guard
	}
	visiting[key] = true
	defer delete(visiting, key)

	if len(ct.Fields) > 0 {
		return // already resolved
	}

	var fields []parse.Field

	for _, rf := range ct.RawFields {
		resolved := r.resolveRawField(rf, ct.Source, visiting)
		fields = append(fields, resolved...)
	}

	// Inline attributeGroups
	for _, agRef := range ct.AttrGroups {
		agFields := r.resolveAttrGroupRef(agRef, ct.Source)
		fields = append(fields, agFields...)
	}

	ct.Fields = deduplicateFields(fields)
}

// deduplicateFields removes fields with duplicate GoNames.
// When an attribute and an element conflict, the later-defined (own) field wins
// over the earlier (inherited) field regardless of which is attr or element:
//   - inherited element + own attribute → attribute wins (e.g. GML 3.1.1 srsName)
//   - inherited attribute + own element → element wins (e.g. GML 3.2.1 GridType axisLabels)
func deduplicateFields(fields []parse.Field) []parse.Field {
	type entry struct {
		idx    int
		isAttr bool
	}
	if len(fields) == 0 {
		return fields
	}
	seen := make(map[string]entry)
	result := make([]parse.Field, 0, len(fields))
	for _, f := range fields {
		e, exists := seen[f.GoName]
		if !exists {
			seen[f.GoName] = entry{len(result), f.IsAttr}
			result = append(result, f)
		} else if f.IsAttr != e.isAttr {
			// One is attr, the other is element: later-defined (own) wins.
			result[e.idx] = f
			seen[f.GoName] = entry{e.idx, f.IsAttr}
		}
		// else: same kind (both attr or both element) — first wins.
	}
	return result
}

// sortedSubstMembers returns all transitive substitution group members for a
// head element key, sorted deterministically by (NS, Name).
// Multi-level chains (e.g. AbstractGeometry → AbstractImplicitGeometry → Grid)
// are fully expanded via collectSubstMembers.
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

// collectSubstMembers recursively collects all transitive members of a
// substitution group head into out, using visited to prevent cycles.
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

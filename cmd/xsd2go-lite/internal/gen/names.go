package gen

import (
	"strings"

	"xsd2go-lite/internal/parse"
)

// qualifyType prefixes the non-builtin type name in goType with alias.
// Handles pointer (*T → *alias.T) and slice ([]T → []alias.T) prefixes.
func qualifyType(goType, alias string) string {
	prefix := ""
	base := goType
	if strings.HasPrefix(base, "[]") {
		prefix = "[]"
		base = base[2:]
	} else if strings.HasPrefix(base, "*") {
		prefix = "*"
		base = base[1:]
	}
	if parse.IsBuiltinGoType(base) || strings.Contains(base, ".") {
		return goType
	}
	return prefix + alias + "." + base
}

// deriveAlias derives a Go import alias from a package path.
// If the last segment is "generated", uses the preceding segment stripped of
// non-Go-identifier characters (hyphens, dots, etc.) plus "gen".
// Underscores are preserved as they are valid Go identifiers.
// Examples:
//
//	"github.com/foo/gml3_1_1/generated" → "gml3_1_1gen"
//	"github.com/foo/gml3_2_1/generated" → "gml3_2_1gen"
//	"github.com/foo/mypkg"              → "mypkg"
func deriveAlias(pkgPath string) string {
	parts := strings.Split(pkgPath, "/")
	// Strip characters that are not valid Go identifier chars (keep letters, digits, underscore).
	stripInvalid := func(s string) string {
		return strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
				return r
			}
			return -1
		}, s)
	}
	last := parts[len(parts)-1]
	if last == "generated" && len(parts) >= 2 {
		return stripInvalid(parts[len(parts)-2]) + "gen"
	}
	return stripInvalid(last)
}

// xmlTagNS extracts the namespace URI from an xml struct tag value.
// Tag format: "namespace localName" or "localName,attr" etc.
func xmlTagNS(tag string) string {
	if idx := strings.Index(tag, " "); idx >= 0 {
		return tag[:idx]
	}
	return ""
}

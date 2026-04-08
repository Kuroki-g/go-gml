package xsdlite

import (
	"strings"
	"unicode"
)

// goTypeName converts an XSD type name to a Go exported type name.
func goTypeName(name string) string {
	return goName(name)
}

// goName converts a name to a Go exported identifier.
func goName(name string) string {
	if name == "" {
		return "_"
	}

	// GML 3.1.1 uses "_Xxx" for abstract substitution group heads (e.g. "_Curve", "_Surface").
	// Prefix the result with "Abstract" so the concrete member field (e.g. "Curve *CurveType")
	// can coexist without being deduped away.
	leadingUnderscore := strings.HasPrefix(name, "_")

	// Replace hyphens and dots with underscores for processing.
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, ".", "_")

	// Split by underscore and capitalize each part.
	parts := strings.Split(name, "_")
	var sb strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		runes := []rune(p)
		runes[0] = unicode.ToUpper(runes[0])
		sb.WriteString(string(runes))
	}

	result := sb.String()
	if result == "" {
		return "_"
	}

	// If first char is a digit, prefix with underscore.
	if unicode.IsDigit([]rune(result)[0]) {
		result = "_" + result
	}

	// Avoid Go reserved words.
	result = avoidReserved(result)

	// Apply "Abstract" prefix for "_Xxx" abstract element names.
	if leadingUnderscore {
		result = "Abstract" + result
	}

	return result
}

func avoidReserved(name string) string {
	switch name {
	case "Break", "Case", "Chan", "Const", "Continue", "Default", "Defer",
		"Else", "Fallthrough", "For", "Func", "Go", "Goto", "If", "Import",
		"Interface", "Map", "Package", "Range", "Return", "Select", "Struct",
		"Switch", "Type", "Var":
		return name + "Field"
	}
	return name
}

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
	if isBuiltinGoType(base) || strings.Contains(base, ".") {
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

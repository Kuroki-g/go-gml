package parse

import (
	"strings"
	"unicode"
)

// GoTypeName converts an XSD type name to a Go exported type name.
func GoTypeName(name string) string {
	return GoName(name)
}

// GoName converts a name to a Go exported identifier.
func GoName(name string) string {
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

// IsBuiltinGoType reports whether t is a primitive Go type.
func IsBuiltinGoType(t string) bool {
	switch t {
	case "string", "bool", "int", "float64", "float32":
		return true
	}
	return false
}

// IsStructType reports whether t is a non-builtin, non-empty type (i.e. a struct/pointer type).
func IsStructType(t string) bool {
	return !IsBuiltinGoType(t) && t != ""
}

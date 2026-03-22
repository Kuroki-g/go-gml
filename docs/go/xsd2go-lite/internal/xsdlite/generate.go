package xsdlite

import (
	"fmt"
	"go/format"
	"sort"
	"strings"
	"unicode"
)

// Generate produces a Go source file string for the given types.
// withDoc controls whether XSD documentation is emitted as field comments.
func Generate(types []*ComplexType, pkgName string, skipAbstract bool, withDoc bool) (string, error) {
	// Sort for deterministic output.
	sort.Slice(types, func(i, j int) bool {
		return types[i].Name < types[j].Name
	})

	var sb strings.Builder

	sb.WriteString("package ")
	sb.WriteString(pkgName)
	sb.WriteString("\n\n")

	for _, ct := range types {
		if skipAbstract && ct.Abstract {
			continue
		}
		if ct.Name == "" {
			continue
		}
		writeStruct(&sb, ct, withDoc)
	}

	src := sb.String()
	formatted, err := format.Source([]byte(src))
	if err != nil {
		// Return unformatted with the error so caller can debug.
		return src, fmt.Errorf("go/format: %w\n--- unformatted source ---\n%s", err, src)
	}
	return string(formatted), nil
}

func writeStruct(sb *strings.Builder, ct *ComplexType, withDoc bool) {
	sb.WriteString("type ")
	sb.WriteString(goTypeName(ct.Name))
	sb.WriteString(" struct {\n")

	for _, f := range ct.Fields {
		if withDoc && f.Doc != "" {
			writeComment(sb, f.Doc)
		}
		sb.WriteString("\t")
		sb.WriteString(f.GoName)
		sb.WriteString(" ")
		sb.WriteString(f.GoType)
		sb.WriteString(" `xml:\"")
		sb.WriteString(f.XMLTag)
		tag := buildTagSuffix(f)
		sb.WriteString(tag)
		sb.WriteString("\"`\n")
	}

	sb.WriteString("}\n\n")
}

// writeComment emits a doc string as Go line comments inside a struct body.
func writeComment(sb *strings.Builder, doc string) {
	for _, line := range strings.Split(doc, "\n") {
		sb.WriteString("\t// ")
		sb.WriteString(strings.TrimSpace(line))
		sb.WriteString("\n")
	}
}

// buildTagSuffix returns the options part of the xml tag after the name/ns.
func buildTagSuffix(f Field) string {
	// The XMLTag already contains ",attr" or ",chardata" if applicable.
	// We just need to append ",omitempty" if needed.
	if f.Omit && !f.IsChar {
		// Don't add omitempty to chardata or slices.
		if !f.Slice {
			return ",omitempty"
		}
	}
	return ""
}

// goTypeName converts an XSD type name to a Go exported type name.
func goTypeName(name string) string {
	return goName(name)
}

// goName converts a name to a Go exported identifier.
func goName(name string) string {
	if name == "" {
		return "_"
	}

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

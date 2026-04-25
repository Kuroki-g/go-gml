package gen

import (
	"strings"

	"xsd2go-lite2/internal/parse"
)

// qualifyType prefixes the non-builtin type name in goType with alias.
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
func deriveAlias(pkgPath string) string {
	parts := strings.Split(pkgPath, "/")
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

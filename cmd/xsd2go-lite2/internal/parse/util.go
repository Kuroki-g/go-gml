package parse

import "strings"

// XsBuiltinGoType maps XSD built-in type names to Go types.
func XsBuiltinGoType(xsType string) string {
	local := xsType
	if idx := strings.LastIndex(xsType, ":"); idx >= 0 {
		local = xsType[idx+1:]
	}
	switch local {
	case "string", "anyURI", "NCName", "QName", "NMTOKEN", "NMTOKENS",
		"Name", "language", "token", "normalizedString", "ID", "IDREF",
		"IDREFS", "ENTITY", "ENTITIES", "notation",
		"base64Binary", "hexBinary",
		"duration", "dateTime", "date", "time", "gYear", "gYearMonth",
		"gMonth", "gMonthDay", "gDay":
		return "string"
	case "boolean":
		return "bool"
	case "int":
		return "int32"
	case "long":
		return "int64"
	case "short":
		return "int16"
	case "byte":
		return "int8"
	case "unsignedLong":
		return "uint64"
	case "unsignedInt":
		return "uint32"
	case "unsignedShort":
		return "uint16"
	case "unsignedByte":
		return "uint8"
	case "nonNegativeInteger", "positiveInteger":
		return "uint"
	case "integer", "nonPositiveInteger", "negativeInteger":
		return "int"
	case "double":
		return "float64"
	// xs:decimal: arbitrary-precision decimal; mapped to float64 as a known compromise.
	// See v1 parse.go for rationale.
	case "decimal":
		return "float64"
	case "float":
		return "float32"
	case "anyType", "anySimpleType":
		return "string"
	default:
		return ""
	}
}

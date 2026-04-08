package parse

import "testing"

// ---- GoName ----

func TestGoName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", "_"},
		{"pointType", "PointType"},
		{"point_type", "PointType"},
		{"point-type", "PointType"},
		{"PolygonType", "PolygonType"},
		{"multiSurface", "MultiSurface"},
		{"multi_surface_type", "MultiSurfaceType"},
		// digit prefix
		{"3DGeometry", "_3DGeometry"},
		// reserved word
		{"type", "TypeField"},
		{"map", "MapField"},
		{"range", "RangeField"},
		// dot separator
		{"gml.geometry", "GmlGeometry"},
		// already exported
		{"AbstractGeometryType", "AbstractGeometryType"},
		// single char
		{"x", "X"},
		// all underscores → empty parts → "_"
		{"_", "_"},
		// GML 3.1.1 abstract element names: "_Xxx" → "AbstractXxx"
		{"_Curve", "AbstractCurve"},
		{"_Surface", "AbstractSurface"},
		{"_Ring", "AbstractRing"},
		{"_CurveSegment", "AbstractCurveSegment"},
		{"_Geometry", "AbstractGeometry"},
	}

	for _, tt := range tests {
		got := GoName(tt.in)
		if got != tt.want {
			t.Errorf("GoName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// ---- avoidReserved ----

func TestAvoidReserved(t *testing.T) {
	reserved := []string{
		"Break", "Case", "Chan", "Const", "Continue", "Default", "Defer",
		"Else", "Fallthrough", "For", "Func", "Go", "Goto", "If", "Import",
		"Interface", "Map", "Package", "Range", "Return", "Select", "Struct",
		"Switch", "Type", "Var",
	}
	for _, r := range reserved {
		got := avoidReserved(r)
		if got != r+"Field" {
			t.Errorf("avoidReserved(%q) = %q, want %q", r, got, r+"Field")
		}
	}
	// Non-reserved should pass through unchanged.
	for _, ok := range []string{"Point", "Line", "Geometry", "Hello"} {
		if got := avoidReserved(ok); got != ok {
			t.Errorf("avoidReserved(%q) = %q, want %q (unchanged)", ok, got, ok)
		}
	}
}

// ---- IsBuiltinGoType ----

func TestIsBuiltinGoType(t *testing.T) {
	for _, bt := range []string{"string", "bool", "int", "float64", "float32"} {
		if !IsBuiltinGoType(bt) {
			t.Errorf("IsBuiltinGoType(%q) = false, want true", bt)
		}
	}
	for _, ct := range []string{"PointType", "[]string", "*Foo", "", "int64", "uint"} {
		if IsBuiltinGoType(ct) {
			t.Errorf("IsBuiltinGoType(%q) = true, want false", ct)
		}
	}
}

// ---- IsStructType ----

func TestIsStructType(t *testing.T) {
	if IsStructType("") {
		t.Error("empty string should not be struct type")
	}
	if IsStructType("string") {
		t.Error("string is not a struct type")
	}
	if !IsStructType("PointType") {
		t.Error("PointType should be a struct type")
	}
	if !IsStructType("*DirectPositionType") {
		t.Error("pointer to struct should be a struct type")
	}
}

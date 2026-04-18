package core

import "testing"

func pointEq(a, b Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPointFromFlat(t *testing.T) {
	tests := []struct {
		name    string
		coords  []float64
		dim     uint
		want    Point
		wantErr bool
	}{
		{"2D", []float64{139.7, 35.6}, 2, Point{139.7, 35.6}, false},
		{"3D keeps Z", []float64{139.7, 35.6, 10.5}, 3, Point{139.7, 35.6, 10.5}, false},
		{"dim 0 infer", []float64{139.7, 35.6}, 0, Point{139.7, 35.6}, false},
		{"too short", []float64{139.7}, 2, Point{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PointFromFlat(tt.coords, tt.dim)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !pointEq(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineStringFromFlat(t *testing.T) {
	coords := []float64{139.7, 35.6, 139.8, 35.7, 139.9, 35.8}
	got, err := LineStringFromFlat(coords, 2)
	if err != nil {
		t.Fatal(err)
	}
	want := LineString{{139.7, 35.6}, {139.8, 35.7}, {139.9, 35.8}}
	if len(got) != len(want) {
		t.Fatalf("len=%d want %d", len(got), len(want))
	}
	for i := range got {
		if !pointEq(got[i], want[i]) {
			t.Errorf("[%d] got %v want %v", i, got[i], want[i])
		}
	}
}

func TestRingFromFlat(t *testing.T) {
	coords := []float64{139.7, 35.6, 139.8, 35.6, 139.8, 35.7, 139.7, 35.6}
	got, err := RingFromFlat(coords, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("len=%d want 4", len(got))
	}
	if !pointEq(got[0], got[3]) {
		t.Error("first and last point should be equal for a closed ring")
	}
}

func TestPointFromPosString(t *testing.T) {
	d2, d3 := uint(2), uint(3)
	tests := []struct {
		input string
		dim   *uint
		want  Point
	}{
		{"139.691667 35.689722", &d2, Point{139.691667, 35.689722}},
		{"139.7 35.6 10.5", &d3, Point{139.7, 35.6, 10.5}},
		{"139.7 35.6 10.5", nil, Point{139.7, 35.6, 10.5}}, // nil → infer from data
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := PointFromPosString(tt.input, tt.dim)
			if err != nil {
				t.Fatal(err)
			}
			if !pointEq(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestLineStringFromPosListString(t *testing.T) {
	d := uint(2)
	got, err := LineStringFromPosListString("139.7 35.6 139.8 35.7", &d)
	if err != nil {
		t.Fatal(err)
	}
	want := LineString{{139.7, 35.6}, {139.8, 35.7}}
	if len(got) != len(want) {
		t.Fatalf("len=%d want %d", len(got), len(want))
	}
	for i := range got {
		if !pointEq(got[i], want[i]) {
			t.Errorf("[%d] %v want %v", i, got[i], want[i])
		}
	}
}

func TestRingFromPosListString(t *testing.T) {
	d := uint(2)
	got, err := RingFromPosListString("139.7 35.6 139.8 35.6 139.8 35.7 139.7 35.6", &d)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("len=%d want 4", len(got))
	}
}

func TestRingFromCoordinatesString(t *testing.T) {
	got, err := RingFromCoordinatesString("139.7,35.6 139.8,35.6 139.8,35.7 139.7,35.6", "", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("len=%d want 4", len(got))
	}
}

func TestEffectiveDim(t *testing.T) {
	p := func(n uint) *uint { return &n }
	check := func(dim *uint, nValues int, want uint) {
		t.Helper()
		got, err := effectiveDim(dim, nValues)
		if err != nil {
			t.Fatalf("effectiveDim(%v, %d) unexpected error: %v", dim, nValues, err)
		}
		if got != want {
			t.Errorf("effectiveDim(%v, %d) = %d, want %d", dim, nValues, got, want)
		}
	}

	check(p(2), 4, 2) // explicit 2D
	check(p(3), 9, 3) // explicit 3D
	check(p(1), 0, 1) // positiveInteger: 1D is valid per XSD
	check(p(4), 0, 4) // positiveInteger: 4D is valid per XSD
	check(nil, 4, 2)  // omitted, even → 2D
	check(nil, 9, 3)  // omitted, odd → 3D (PLATEAU pattern)
	check(nil, 27, 3) // omitted, odd → 3D
	check(nil, 12, 2) // omitted, even (ambiguous) → 2D fallback
}

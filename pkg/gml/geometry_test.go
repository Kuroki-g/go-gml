package gml

import "testing"

func TestPointFromFlat(t *testing.T) {
	tests := []struct {
		name    string
		coords  []float64
		dim     int
		want    Point
		wantErr bool
	}{
		{"2D", []float64{139.7, 35.6}, 2, Point{139.7, 35.6}, false},
		{"3D drops Z", []float64{139.7, 35.6, 10.5}, 3, Point{139.7, 35.6}, false},
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
			if got != tt.want {
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
		if got[i] != want[i] {
			t.Errorf("[%d] got %v want %v", i, got[i], want[i])
		}
	}
}

func TestRingFromFlat(t *testing.T) {
	// closed ring: 4 points (first == last)
	coords := []float64{139.7, 35.6, 139.8, 35.6, 139.8, 35.7, 139.7, 35.6}
	got, err := RingFromFlat(coords, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("len=%d want 4", len(got))
	}
	if got[0] != got[3] {
		t.Error("first and last point should be equal for a closed ring")
	}
}

func TestPointFromPosString(t *testing.T) {
	tests := []struct {
		input string
		dim   int
		want  Point
	}{
		{"139.691667 35.689722", 2, Point{139.691667, 35.689722}},
		{"139.7 35.6 10.5", 3, Point{139.7, 35.6}},
		{"139.7 35.6 10.5", 0, Point{139.7, 35.6}}, // infer 3D
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := PointFromPosString(tt.input, tt.dim)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestLineStringFromPosListString(t *testing.T) {
	got, err := LineStringFromPosListString("139.7 35.6 139.8 35.7", 2)
	if err != nil {
		t.Fatal(err)
	}
	want := LineString{{139.7, 35.6}, {139.8, 35.7}}
	if len(got) != len(want) {
		t.Fatalf("len=%d want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("[%d] %v want %v", i, got[i], want[i])
		}
	}
}

func TestRingFromPosListString(t *testing.T) {
	got, err := RingFromPosListString("139.7 35.6 139.8 35.6 139.8 35.7 139.7 35.6", 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("len=%d want 4", len(got))
	}
}

func TestRingFromCoordinatesString(t *testing.T) {
	// deprecated gml:coordinates format
	got, err := RingFromCoordinatesString("139.7,35.6 139.8,35.6 139.8,35.7 139.7,35.6", "", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("len=%d want 4", len(got))
	}
}

func TestEffectiveDim(t *testing.T) {
	if effectiveDim(2, 6) != 2 {
		t.Error("explicit dim should be returned as-is")
	}
	if effectiveDim(0, 6) != 3 {
		t.Error("total=6 should infer dim=3")
	}
	if effectiveDim(0, 4) != 2 {
		t.Error("total=4 should infer dim=2")
	}
	if effectiveDim(0, 0) != 2 {
		t.Error("total=0 should default to dim=2")
	}
}

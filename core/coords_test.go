package core

import (
	"testing"
)

func TestParsePosList(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []float64
		wantErr bool
	}{
		{"2D two points", "139.7 35.6 139.8 35.7", []float64{139.7, 35.6, 139.8, 35.7}, false},
		{"single pos", "139.691667 35.689722", []float64{139.691667, 35.689722}, false},
		{"3D point", "139.7 35.6 10.5", []float64{139.7, 35.6, 10.5}, false},
		{"extra whitespace", "  139.7  35.6  \n  139.8  35.7  ", []float64{139.7, 35.6, 139.8, 35.7}, false},
		{"negative coordinates", "-73.9857 40.7484", []float64{-73.9857, 40.7484}, false},
		{"empty string", "", []float64{}, false},
		{"invalid token", "139.7 abc 35.6", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePosList(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("len=%d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("[%d] got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestParseCoordinates(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		cs, ts, decimal string
		want            []float64
		wantDim         uint
		wantErr         bool
	}{
		{"2D default separators", "139.7,35.6 139.8,35.7", "", "", "", []float64{139.7, 35.6, 139.8, 35.7}, 2, false},
		{"3D default separators", "139.7,35.6,10.5 139.8,35.7,11.0", "", "", "", []float64{139.7, 35.6, 10.5, 139.8, 35.7, 11.0}, 3, false},
		{"custom separators", "139.7;35.6|139.8;35.7", ";", "|", "", []float64{139.7, 35.6, 139.8, 35.7}, 2, false},
		{"custom decimal", "139,7;35,6|139,8;35,7", ";", "|", ",", []float64{139.7, 35.6, 139.8, 35.7}, 2, false},
		{"leading trailing whitespace", "\n  139.7,35.6 139.8,35.7  \n", "", "", "", []float64{139.7, 35.6, 139.8, 35.7}, 2, false},
		{"empty string", "", "", "", "", nil, 0, true},
		{"single value tuple", "139.7 35.6", ",", " ", "", nil, 0, true},
		{"invalid coordinate", "139.7,bad 139.8,35.7", "", "", "", nil, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, dim, err := ParseCoordinates(tt.input, tt.cs, tt.ts, tt.decimal)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dim != tt.wantDim {
				t.Errorf("dim=%d, want %d", dim, tt.wantDim)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("len=%d, want %d: %v", len(got), len(tt.want), got)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("[%d] got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestToPoints(t *testing.T) {
	pointEq := func(a, b Point) bool {
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
	tests := []struct {
		name    string
		coords  []float64
		dim     uint
		want    []Point
		wantErr bool
	}{
		{"2D two points", []float64{139.7, 35.6, 139.8, 35.7}, 2, []Point{{139.7, 35.6}, {139.8, 35.7}}, false},
		{"3D keeps Z", []float64{139.7, 35.6, 10.5, 139.8, 35.7, 11.0}, 3, []Point{{139.7, 35.6, 10.5}, {139.8, 35.7, 11.0}}, false},
		{"dim 0 defaults to 2", []float64{139.7, 35.6}, 0, []Point{{139.7, 35.6}}, false},
		{"empty coords", []float64{}, 2, nil, false},
		{"nil coords", nil, 2, nil, false},
		{"odd count for 2D", []float64{139.7, 35.6, 139.8}, 2, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToPoints(tt.coords, tt.dim)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("len=%d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if !pointEq(got[i], tt.want[i]) {
					t.Errorf("[%d] got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

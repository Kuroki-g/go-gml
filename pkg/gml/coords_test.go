package gml

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
		{
			name:  "2D two points",
			input: "139.7 35.6 139.8 35.7",
			want:  []float64{139.7, 35.6, 139.8, 35.7},
		},
		{
			name:  "single pos",
			input: "139.691667 35.689722",
			want:  []float64{139.691667, 35.689722},
		},
		{
			name:  "3D point",
			input: "139.7 35.6 10.5",
			want:  []float64{139.7, 35.6, 10.5},
		},
		{
			name:  "extra whitespace",
			input: "  139.7  35.6  \n  139.8  35.7  ",
			want:  []float64{139.7, 35.6, 139.8, 35.7},
		},
		{
			name:  "negative coordinates",
			input: "-73.9857 40.7484",
			want:  []float64{-73.9857, 40.7484},
		},
		{
			name:  "empty string",
			input: "",
			want:  []float64{},
		},
		{
			name:    "invalid token",
			input:   "139.7 abc 35.6",
			wantErr: true,
		},
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
		name    string
		input   string
		cs, ts  string
		want    []float64
		wantErr bool
	}{
		{
			name:  "default separators",
			input: "139.7,35.6 139.8,35.7",
			cs:    "",
			ts:    "",
			want:  []float64{139.7, 35.6, 139.8, 35.7},
		},
		{
			name:  "3D default separators",
			input: "139.7,35.6,10.5 139.8,35.7,11.0",
			cs:    "",
			ts:    "",
			want:  []float64{139.7, 35.6, 10.5, 139.8, 35.7, 11.0},
		},
		{
			name:  "custom separators",
			input: "139.7;35.6|139.8;35.7",
			cs:    ";",
			ts:    "|",
			want:  []float64{139.7, 35.6, 139.8, 35.7},
		},
		{
			name:  "leading trailing whitespace",
			input: "\n  139.7,35.6 139.8,35.7  \n",
			cs:    "",
			ts:    "",
			want:  []float64{139.7, 35.6, 139.8, 35.7},
		},
		{
			name:  "empty string",
			input: "",
			cs:    "",
			ts:    "",
			want:  nil,
		},
		{
			name:    "invalid coordinate",
			input:   "139.7,bad 139.8,35.7",
			cs:      "",
			ts:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCoordinates(tt.input, tt.cs, tt.ts)
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
	tests := []struct {
		name    string
		coords  []float64
		dim     int
		want    [][2]float64
		wantErr bool
	}{
		{
			name:   "2D two points",
			coords: []float64{139.7, 35.6, 139.8, 35.7},
			dim:    2,
			want:   [][2]float64{{139.7, 35.6}, {139.8, 35.7}},
		},
		{
			name:   "3D drops Z",
			coords: []float64{139.7, 35.6, 10.5, 139.8, 35.7, 11.0},
			dim:    3,
			want:   [][2]float64{{139.7, 35.6}, {139.8, 35.7}},
		},
		{
			name:   "dim 0 defaults to 2",
			coords: []float64{139.7, 35.6},
			dim:    0,
			want:   [][2]float64{{139.7, 35.6}},
		},
		{
			name:   "empty coords",
			coords: []float64{},
			dim:    2,
			want:   nil,
		},
		{
			name:   "nil coords",
			coords: nil,
			dim:    2,
			want:   nil,
		},
		{
			name:    "odd count for 2D",
			coords:  []float64{139.7, 35.6, 139.8},
			dim:     2,
			wantErr: true,
		},
		{
			name:    "count not multiple of 3",
			coords:  []float64{139.7, 35.6, 10.5, 139.8},
			dim:     3,
			wantErr: true,
		},
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
				if got[i] != tt.want[i] {
					t.Errorf("[%d] got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

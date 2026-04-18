package core

import (
	"testing"
)

func ptr(s string) *string { return &s }

func TestEPSGFromSRSName(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"EPSG:4326", 4326},
		{"EPSG:6668", 6668},
		{"epsg:3857", 3857},
		{"urn:ogc:def:crs:EPSG::4326", 4326},
		{"urn:ogc:def:crs:EPSG::6668", 6668},
		{"urn:ogc:def:crs:EPSG:6.18:4326", 4326},
		{"urn:ogc:def:crs:EPSG:6.18.3:4326", 4326},
		{"http://www.opengis.net/def/crs/EPSG/0/4326", 4326},
		{"http://www.opengis.net/def/crs/EPSG/0/6668", 6668},
		{"", 0},
		{"CRS84", 0},
		{"urn:ogc:def:crs:OGC:1.3:CRS84", 0},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := EPSGFromSRSName(tt.input)
			if got != tt.want {
				t.Errorf("EPSGFromSRSName(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func u(v uint) *uint { return &v }

func TestDimFromSRSName(t *testing.T) {
	tests := []struct {
		name  string
		input *string
		want  *uint
	}{
		// nil input
		{"nil", nil, nil},
		// EPSG URI forms (via epsgDimTable)
		{"EPSG:4326", ptr("EPSG:4326"), u(2)},
		{"EPSG:6668", ptr("EPSG:6668"), u(2)},
		{"urn EPSG 4326", ptr("urn:ogc:def:crs:EPSG::4326"), u(2)},
		{"http EPSG 6668", ptr("http://www.opengis.net/def/crs/EPSG/0/6668"), u(2)},
		// non-EPSG srsName strings (via nonEPSGDimTable)
		{"JGD2011 / (B, L)", ptr("JGD2011 / (B, L)"), u(2)},
		{"JGD2000 / (B, L)", ptr("JGD2000 / (B, L)"), u(2)},
		{"JGD2011/(B,L)", ptr("JGD2011/(B,L)"), u(2)},
		{"TD / (B, L)", ptr("TD / (B, L)"), u(2)},
		{"fguuid:jgd2011.bl", ptr("fguuid:jgd2011.bl"), u(2)},
		// unknown
		{"unknown", ptr("CRS84"), nil},
		{"empty", ptr(""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DimFromSRSName(tt.input)
			if tt.want == nil {
				if got != nil {
					t.Errorf("DimFromSRSName(%v) = %d, want nil", tt.input, *got)
				}
			} else {
				if got == nil {
					t.Errorf("DimFromSRSName(%q) = nil, want %d", *tt.input, *tt.want)
				} else if *got != *tt.want {
					t.Errorf("DimFromSRSName(%q) = %d, want %d", *tt.input, *got, *tt.want)
				}
			}
		})
	}
}

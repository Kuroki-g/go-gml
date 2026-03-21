package gml

import "testing"

func TestEPSGFromSRSName(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		// simple prefix
		{"EPSG:4326", 4326},
		{"EPSG:6668", 6668},
		{"epsg:3857", 3857},

		// URN double-colon (empty version)
		{"urn:ogc:def:crs:EPSG::4326", 4326},
		{"urn:ogc:def:crs:EPSG::6668", 6668},

		// URN with version
		{"urn:ogc:def:crs:EPSG:6.18:4326", 4326},
		{"urn:ogc:def:crs:EPSG:6.18.3:4326", 4326},

		// HTTP URL
		{"http://www.opengis.net/def/crs/EPSG/0/4326", 4326},
		{"http://www.opengis.net/def/crs/EPSG/0/6668", 6668},

		// empty / unknown
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

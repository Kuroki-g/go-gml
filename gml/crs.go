package gml

import "github.com/Kuroki-g/go-gml/core"

// EPSGFromSRSName extracts the EPSG code from a GML srsName attribute.
// Returns 0 if the code cannot be determined.
//
// Supported formats:
//   - "EPSG:4326"
//   - "urn:ogc:def:crs:EPSG::4326"
//   - "urn:ogc:def:crs:EPSG:6.18:4326"
//   - "http://www.opengis.net/def/crs/EPSG/0/4326"
func EPSGFromSRSName(srsName string) int {
	return core.EPSGFromSRSName(srsName)
}

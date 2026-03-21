package gml

import (
	"strconv"
	"strings"
)

// EPSGFromSRSName extracts the EPSG code from a GML srsName attribute.
// Returns 0 if the code cannot be determined.
//
// Supported formats:
//   - "EPSG:4326"
//   - "urn:ogc:def:crs:EPSG::4326"
//   - "urn:ogc:def:crs:EPSG:6.18:4326"  (version-aware URN)
//   - "http://www.opengis.net/def/crs/EPSG/0/4326"
func EPSGFromSRSName(srsName string) int {
	if srsName == "" {
		return 0
	}

	// "EPSG:4326"
	if upper := strings.ToUpper(srsName); strings.HasPrefix(upper, "EPSG:") {
		return parseTrailingInt(srsName[5:])
	}

	// "urn:ogc:def:crs:EPSG:...:4326"
	if strings.HasPrefix(strings.ToLower(srsName), "urn:ogc:def:crs:epsg:") {
		parts := strings.Split(srsName, ":")
		// parts: ["urn","ogc","def","crs","EPSG","<version>","<code>"]
		if len(parts) >= 7 {
			return parseTrailingInt(parts[6])
		}
		// "urn:ogc:def:crs:EPSG::4326" → parts[5]=="" parts[6]="4326" handled above
		// but also "urn:ogc:def:crs:EPSG::4326" splits into 7 parts with parts[5]=""
		return 0
	}

	// "http://www.opengis.net/def/crs/EPSG/0/4326"
	if strings.Contains(srsName, "/EPSG/") || strings.Contains(srsName, "/epsg/") {
		idx := strings.LastIndex(srsName, "/")
		if idx >= 0 && idx < len(srsName)-1 {
			return parseTrailingInt(srsName[idx+1:])
		}
	}

	return 0
}

func parseTrailingInt(s string) int {
	s = strings.TrimSpace(s)
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 {
		return 0
	}
	return v
}

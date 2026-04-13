package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml2_1_2/generated"
)

// resolver caches GML 2.1.2 geometry elements by their gid attribute value,
// enabling xlink:href references in Multi* members to be resolved.
type resolver struct {
	pointByID      map[string]core.Point
	lineStringByID map[string]core.LineString
	polygonByID    map[string]core.Polygon
}

func newResolver() *resolver {
	return &resolver{
		pointByID:      make(map[string]core.Point),
		lineStringByID: make(map[string]core.LineString),
		polygonByID:    make(map[string]core.Polygon),
	}
}

// stripHash strips a leading "#" from an xlink:href value.
func stripHash(href string) string {
	return strings.TrimPrefix(href, "#")
}

// extractGID returns the value of the gid attribute (no namespace) from a StartElement, or "".
// GML 2.1.2 AbstractGeometryType defines gid as a plain (non-namespace-qualified) attribute.
// Do NOT use extractGMLID here — that function matches gml:id (GML 3.x only).
func extractGID(se xml.StartElement) string {
	for _, a := range se.Attr {
		if a.Name.Local == "gid" && a.Name.Space == "" {
			return a.Value
		}
	}
	return ""
}

// preScanGeometries scans a GML 2.1.2 document for Point, LineString, and Polygon
// elements that carry a gid attribute, decoding and caching them in the resolver.
// This enables forward xlink:href references in Multi* members to be resolved.
func preScanGeometries(dec *xml.Decoder, res *resolver) error {
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if !isGMLNS(se.Name.Space) {
			continue
		}
		gid := extractGID(se)
		if gid == "" {
			continue
		}
		switch se.Name.Local {
		case gmlPoint:
			var x gen.PointType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Point %q: %w", gid, err)
			}
			if pt, err := pointFromXML(&x); err == nil {
				res.pointByID[gid] = pt
			}
		case gmlLineString:
			var x gen.LineStringType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("LineString %q: %w", gid, err)
			}
			if ls, err := lineStringFromXML(&x); err == nil {
				res.lineStringByID[gid] = ls
			}
		case gmlPolygon:
			var x gen.PolygonType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Polygon %q: %w", gid, err)
			}
			if poly, err := polygonFromXML(&x); err == nil {
				res.polygonByID[gid] = poly
			}
		}
	}
	return nil
}

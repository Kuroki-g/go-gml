package gml

import (
	"encoding/xml"
	"fmt"
	"io"

	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// preScanGeometries scans the GML document for geometry elements with gml:id,
// decoding and caching them in the resolver. This populates the resolver cache
// before the main parse pass, enabling forward xlink:href references to be resolved.
//
// Elements without gml:id are skipped (descended naturally via the token loop).
// DecodeElement errors are returned immediately. Geometry conversion errors (e.g.
// polygonFromXML, lineStringFromXML) are silently ignored; the main pass will
// report them if the element is encountered again.
func preScanGeometries(dec *xml.Decoder, resolver *curveResolver) error {
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return nil
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
		id := extractGMLID(se)
		if id == "" {
			// No gml:id: descend into children naturally via the token loop.
			continue
		}
		switch se.Name.Local {
		case gmlLineString:
			var x v3_2_1.LineStringType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("LineString %q: %w", id, err)
			}
			if ls, err := lineStringFromXML(&x); err == nil {
				resolver.lineStringByID[id] = ls
			}
		case gmlCompositeCurve:
			var x v3_2_1.CompositeCurveType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("CompositeCurve %q: %w", id, err)
			}
			if ls, err := lineStringFromCompositeCurveType(&x, 0, resolver); err == nil {
				resolver.lineStringByID[id] = ls
			}
		case gmlPolygon:
			var x v3_2_1.PolygonType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Polygon %q: %w", id, err)
			}
			if poly, err := polygonFromXML(&x); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlSurface:
			var x v3_2_1.SurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Surface %q: %w", id, err)
			}
			if poly, err := polygonFromSurface(&x, resolver); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlCompositeSurface:
			var x v3_2_1.CompositeSurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("CompositeSurface %q: %w", id, err)
			}
			if poly, err := polygonFromCompositeSurface(&x, resolver); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlOrientableSurface:
			var x v3_2_1.OrientableSurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("OrientableSurface %q: %w", id, err)
			}
			if x.BaseSurface != nil {
				if poly, err := polygonFromSurfaceProperty(x.BaseSurface, derefDim(x.SrsDimension), resolver); err == nil {
					resolver.polygonByID[id] = poly
				}
			}
		case gmlCurve:
			var x v3_2_1.CurveType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Curve %q: %w", id, err)
			}
			resolver.curves[id] = &x
		case gmlOrientableCurve:
			var x v3_2_1.OrientableCurveType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("OrientableCurve %q: %w", id, err)
			}
			resolver.orientable[id] = &x
		}
		// Elements with gml:id not matched above: descend into children naturally.
	}
}

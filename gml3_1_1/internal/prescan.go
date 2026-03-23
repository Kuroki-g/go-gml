package internal

import (
	"encoding/xml"
	"fmt"
	"io"

	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// preScanGeometries scans the GML document for geometry elements with gml:id,
// decoding and caching them in the resolver. This populates the resolver cache
// before the main parse pass, enabling forward xlink:href references to be resolved.
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
			continue
		}
		switch se.Name.Local {
		case gmlLineString:
			var x gen.LineStringType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("LineString %q: %w", id, err)
			}
			if ls, err := lineStringFromXML(&x); err == nil {
				resolver.lineStringByID[id] = ls
			}
		case gmlCompositeCurve:
			var x gen.CompositeCurveType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("CompositeCurve %q: %w", id, err)
			}
			if ls, err := lineStringFromCompositeCurveType(&x, 0, resolver); err == nil {
				resolver.lineStringByID[id] = ls
			}
		case gmlPolygon:
			var x gen.PolygonType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Polygon %q: %w", id, err)
			}
			if poly, err := polygonFromXML(&x); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlSurface:
			var x gen.SurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Surface %q: %w", id, err)
			}
			if poly, err := polygonFromSurface(&x, resolver); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlCompositeSurface:
			var x gen.CompositeSurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("CompositeSurface %q: %w", id, err)
			}
			if poly, err := polygonFromCompositeSurface(&x, resolver); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlOrientableSurface:
			var x gen.OrientableSurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("OrientableSurface %q: %w", id, err)
			}
			if x.BaseSurface != nil {
				if poly, err := polygonFromSurfaceProperty(x.BaseSurface, derefDim(x.SrsDimension), resolver); err == nil {
					resolver.polygonByID[id] = poly
				}
			}
		case gmlCurve:
			var x gen.CurveType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Curve %q: %w", id, err)
			}
			resolver.curves[id] = &x
		case gmlOrientableCurve:
			var x gen.OrientableCurveType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("OrientableCurve %q: %w", id, err)
			}
			resolver.orientable[id] = &x
		case gmlGrid:
			var x gen.GridType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Grid %q: %w", id, err)
			}
			if gb, err := gridBoundsFromGridLimits(x.Limits); err == nil {
				resolver.gridByID[id] = gb
			}
		case gmlRectifiedGrid:
			var x gen.RectifiedGridType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("RectifiedGrid %q: %w", id, err)
			}
			if gb, err := gridBoundsFromGridLimits(x.Limits); err == nil {
				resolver.gridByID[id] = gb
			}
		}
	}
}

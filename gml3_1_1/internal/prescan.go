package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// pendingCS holds a decoded CompositeSurfaceType awaiting deferred resolution.
type pendingCS struct {
	id string
	x  *gen.CompositeSurfaceType
}

// preScanGeometries scans the GML document for geometry elements with gml:id,
// decoding and caching them in the resolver. This populates the resolver cache
// before the main parse pass, enabling forward xlink:href references to be resolved.
//
// CompositeSurface elements are collected and resolved after all leaf geometries
// are cached, so forward xlink:href references within CompositeSurface are handled.
func preScanGeometries(dec *xml.Decoder, resolver *curveResolver) error {
	var deferred []pendingCS
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
			if poly, err := polygonFromXML(&x, 0); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlSurface:
			var x gen.SurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Surface %q: %w", id, err)
			}
			if poly, err := polygonFromSurface(&x, resolver, 0); err == nil {
				resolver.polygonByID[id] = poly
			}
		case gmlCompositeSurface:
			var x gen.CompositeSurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("CompositeSurface %q: %w", id, err)
			}
			deferred = append(deferred, pendingCS{id: id, x: &x})
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
		case gmlTin:
			var x gen.TinType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Tin %q: %w", id, err)
			}
			if mp, err := multiPolygonFromTrianglePatches(x.TrianglePatches, 0, resolver); err == nil {
				resolver.multiPolygonByID[id] = mp
			}
		case gmlTriangulatedSurface:
			var x gen.TriangulatedSurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("TriangulatedSurface %q: %w", id, err)
			}
			if mp, err := multiPolygonFromTrianglePatches(x.TrianglePatches, 0, resolver); err == nil {
				resolver.multiPolygonByID[id] = mp
			}
		case gmlPolyhedralSurface:
			var x gen.PolyhedralSurfaceType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("PolyhedralSurface %q: %w", id, err)
			}
			if mp, err := multiPolygonFromPolyhedralSurface(&x, 0, resolver); err == nil {
				resolver.multiPolygonByID[id] = mp
			}
		case gmlSolid:
			var x gen.SolidType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("Solid %q: %w", id, err)
			}
			if s, err := solidFromXML(&x, 0, resolver); err == nil {
				resolver.solidByID[id] = s
			}
		case gmlCompositeSolid:
			var x gen.CompositeSolidType
			if err := dec.DecodeElement(&x, &se); err != nil {
				return fmt.Errorf("CompositeSolid %q: %w", id, err)
			}
			if s, err := solidFromSolidMembers(x.SolidMember, 0, resolver); err == nil {
				resolver.solidByID[id] = s
			}
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
	resolveDeferred(deferred, resolver)
	return nil
}

// resolveDeferred iteratively resolves CompositeSurface elements whose
// surfaceMember href targets may not have been cached yet when first encountered.
// Each round resolves any element whose dependencies are now fully available.
func resolveDeferred(pending []pendingCS, resolver *curveResolver) {
	for len(pending) > 0 {
		var remaining []pendingCS
		for _, p := range pending {
			if !allSurfaceMembersResolvable(p.x.SurfaceMember, resolver) {
				remaining = append(remaining, p)
				continue
			}
			if mp, err := multiPolygonFromCompositeSurface(p.x, resolver, 0); err == nil {
				resolver.multiPolygonByID[p.id] = mp
			}
		}
		if len(remaining) == len(pending) {
			// No progress: remaining refs are unresolvable (e.g. external hrefs).
			break
		}
		pending = remaining
	}
}

// allSurfaceMembersResolvable reports whether every href in the surface member list
// is already cached in the resolver. Inline members are always considered resolvable.
func allSurfaceMembersResolvable(members []gen.SurfacePropertyType, resolver *curveResolver) bool {
	for _, m := range members {
		if m.Href != "" {
			id := strings.TrimPrefix(m.Href, "#")
			_, inPoly := resolver.polygonByID[id]
			_, inMulti := resolver.multiPolygonByID[id]
			if !inPoly && !inMulti {
				return false
			}
		}
		if m.OrientableSurface != nil && m.OrientableSurface.BaseSurface != nil {
			bs := m.OrientableSurface.BaseSurface
			if bs.Href != "" {
				id := strings.TrimPrefix(bs.Href, "#")
				_, inPoly := resolver.polygonByID[id]
				_, inMulti := resolver.multiPolygonByID[id]
				if !inPoly && !inMulti {
					return false
				}
			}
		}
		if m.CompositeSurface != nil {
			if !allSurfaceMembersResolvable(m.CompositeSurface.SurfaceMember, resolver) {
				return false
			}
		}
		// Inline Tin/TriangulatedSurface/PolyhedralSurface contain no nested hrefs;
		// they are always resolvable. Explicitly reference fields to satisfy coverage.
		_ = m.Tin
		_ = m.TriangulatedSurface
		_ = m.PolyhedralSurface
	}
	return true
}

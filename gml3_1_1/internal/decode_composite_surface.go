package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// handleCompositeSurface decodes a gml:CompositeSurface, caches the result by gml:id, and returns a MultiPolygon.
func (r *Reader) handleCompositeSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	var x gen.CompositeSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: CompositeSurface: %w", err)
	}
	mp, err := multiPolygonFromCompositeSurface(&x, r.resolver, r.globalDim)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		r.resolver.multiPolygonByID[id] = mp
	}
	return core.Geometry{Value: mp, SRSName: x.SrsName}, nil
}

// handleOrientableSurface decodes a gml:OrientableSurface and returns a Polygon.
func (r *Reader) handleOrientableSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	var x gen.OrientableSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: OrientableSurface: %w", err)
	}
	if x.BaseSurface == nil {
		return core.Geometry{Value: core.Polygon(nil), SRSName: x.SrsName}, nil
	}
	poly, err := polygonFromSurfaceProperty(x.BaseSurface, preferDim(derefDim(x.SrsDimension), r.globalDim), r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		r.resolver.polygonByID[id] = poly
	}
	return core.Geometry{Value: poly, SRSName: x.SrsName}, nil
}

// multiPolygonFromCompositeSurface returns one Polygon per surfaceMember.
// Nested CompositeSurface members are flattened into the result.
func multiPolygonFromCompositeSurface(x *gen.CompositeSurfaceType, resolver *curveResolver, fallbackDim int) (core.MultiPolygon, error) {
	dim := preferDim(derefDim(x.SrsDimension), fallbackDim)
	var result core.MultiPolygon
	for i, m := range x.SurfaceMember {
		if m.CompositeSurface != nil {
			nested, err := multiPolygonFromCompositeSurface(m.CompositeSurface, resolver, fallbackDim)
			if err != nil {
				return nil, fmt.Errorf("gml: CompositeSurface surfaceMember[%d]: %w", i, err)
			}
			result = append(result, nested...)
			continue
		}
		if m.Href != "" {
			id := strings.TrimPrefix(m.Href, "#")
			if mp, ok := resolver.multiPolygonByID[id]; ok {
				result = append(result, mp...)
				continue
			}
		}
		poly, err := polygonFromSurfaceProperty(&m, dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: CompositeSurface surfaceMember[%d]: %w", i, err)
		}
		if poly != nil {
			result = append(result, poly)
		}
	}
	return result, nil
}

// multiPolygonFromSurfaceProperty converts a SurfacePropertyType to a MultiPolygon.
// Used for gml:Solid exterior/interior where the surface may be a CompositeSurface.
func multiPolygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim int, resolver *curveResolver) (core.MultiPolygon, error) {
	if m.CompositeSurface != nil {
		return multiPolygonFromCompositeSurface(m.CompositeSurface, resolver, inheritDim)
	}
	if m.PolyhedralSurface != nil {
		ps := m.PolyhedralSurface
		dim := preferDim(derefDim(ps.SrsDimension), inheritDim)
		var result core.MultiPolygon
		if ps.PolygonPatches != nil {
			for i := range ps.PolygonPatches.PolygonPatch {
				poly, err := polygonFromPatch(&ps.PolygonPatches.PolygonPatch[i], dim, resolver)
				if err != nil {
					return nil, fmt.Errorf("gml: PolyhedralSurface polygonPatches[%d]: %w", i, err)
				}
				result = append(result, poly)
			}
		}
		return result, nil
	}
	if m.TriangulatedSurface != nil {
		ts := m.TriangulatedSurface
		return multiPolygonFromTrianglePatches(ts.TrianglePatches, preferDim(derefDim(ts.SrsDimension), inheritDim), resolver)
	}
	if m.Tin != nil {
		t := m.Tin
		return multiPolygonFromTrianglePatches(t.TrianglePatches, preferDim(derefDim(t.SrsDimension), inheritDim), resolver)
	}
	if m.Href != "" {
		id := strings.TrimPrefix(m.Href, "#")
		if mp, ok := resolver.multiPolygonByID[id]; ok {
			return mp, nil
		}
	}
	poly, err := polygonFromSurfaceProperty(m, inheritDim, resolver)
	if err != nil {
		return nil, err
	}
	if poly == nil {
		return nil, nil
	}
	return core.MultiPolygon{poly}, nil
}

// polygonFromSurfaceProperty converts a SurfacePropertyType to a Polygon.
// For CompositeSurface members, use multiPolygonFromSurfaceProperty instead.
func polygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim int, resolver *curveResolver) (core.Polygon, error) {
	if m.Polygon != nil {
		return polygonFromXML(m.Polygon, inheritDim)
	}
	if m.Surface != nil {
		return polygonFromSurface(m.Surface, resolver, inheritDim)
	}
	if m.OrientableSurface != nil {
		os := m.OrientableSurface
		if os.BaseSurface != nil {
			return polygonFromSurfaceProperty(os.BaseSurface, preferDim(inheritDim, derefDim(os.SrsDimension)), resolver)
		}
	}
	if m.PolyhedralSurface != nil {
		ps := m.PolyhedralSurface
		if ps.PolygonPatches == nil || len(ps.PolygonPatches.PolygonPatch) == 0 {
			return core.Polygon(nil), nil
		}
		dim := preferDim(derefDim(ps.SrsDimension), inheritDim)
		return polygonFromPatch(&ps.PolygonPatches.PolygonPatch[0], dim, resolver)
	}
	if m.TriangulatedSurface != nil {
		ts := m.TriangulatedSurface
		if ts.TrianglePatches == nil || len(ts.TrianglePatches.Triangle) == 0 {
			return core.Polygon(nil), nil
		}
		return polygonFromTriangle(&ts.TrianglePatches.Triangle[0], preferDim(derefDim(ts.SrsDimension), inheritDim), resolver)
	}
	if m.Tin != nil {
		t := m.Tin
		if t.TrianglePatches == nil || len(t.TrianglePatches.Triangle) == 0 {
			return core.Polygon(nil), nil
		}
		return polygonFromTriangle(&t.TrianglePatches.Triangle[0], preferDim(derefDim(t.SrsDimension), inheritDim), resolver)
	}
	if m.Href != "" {
		id := strings.TrimPrefix(m.Href, "#")
		if poly, ok := resolver.polygonByID[id]; ok {
			return poly, nil
		}
	}
	return core.Polygon(nil), nil
}

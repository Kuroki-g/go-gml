package internal

import (
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// multiPolygonFromSurfaceProperty converts a SurfacePropertyType to a MultiPolygon.
// CompositeSurface/Shell members produce multiple Polygons; all other surface types produce one.
func multiPolygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim int, resolver *curveResolver) (core.MultiPolygon, error) {
	if m.CompositeSurface != nil {
		return multiPolygonFromCompositeSurface(m.CompositeSurface, resolver, inheritDim)
	}
	if m.Shell != nil {
		return multiPolygonFromShell(m.Shell, preferDim(derefDim(m.Shell.SrsDimension), inheritDim), resolver)
	}
	if m.PolyhedralSurface != nil {
		ps := m.PolyhedralSurface
		dim := preferDim(derefDim(ps.SrsDimension), inheritDim)
		var result core.MultiPolygon
		if ps.Patches != nil {
			for i := range ps.Patches.PolygonPatch {
				poly, err := polygonFromPatch(&ps.Patches.PolygonPatch[i], dim, resolver)
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
		return multiPolygonFromPatchArray(ts.Patches, preferDim(derefDim(ts.SrsDimension), inheritDim), resolver)
	}
	if m.Tin != nil {
		t := m.Tin
		return multiPolygonFromPatchArray(t.TrianglePatches, preferDim(derefDim(t.SrsDimension), inheritDim), resolver)
	}
	if m.Polygon != nil {
		poly, err := polygonFromXML(m.Polygon, inheritDim)
		if err != nil {
			return nil, err
		}
		return core.MultiPolygon{poly}, nil
	}
	if m.Surface != nil {
		poly, err := polygonFromSurface(m.Surface, resolver, inheritDim)
		if err != nil {
			return nil, err
		}
		if poly == nil {
			return nil, nil
		}
		return core.MultiPolygon{poly}, nil
	}
	if m.OrientableSurface != nil {
		os := m.OrientableSurface
		if os.BaseSurface != nil {
			return multiPolygonFromSurfaceProperty(os.BaseSurface, preferDim(inheritDim, derefDim(os.SrsDimension)), resolver)
		}
		return nil, nil
	}
	if m.Href != "" {
		id := strings.TrimPrefix(m.Href, "#")
		if mp, ok := resolver.multiPolygonByID[id]; ok {
			return mp, nil
		}
		if poly, ok := resolver.polygonByID[id]; ok {
			return core.MultiPolygon{poly}, nil
		}
	}
	// xlink metadata attributes — not used for geometry.
	_ = m.NilReason
	_ = m.RemoteSchema
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	_ = m.Owns
	return nil, nil
}

// polygonFromSurfaceProperty converts a SurfacePropertyType to a Polygon.
// For CompositeSurface/Shell members, use multiPolygonFromSurfaceProperty instead.
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
	if m.CompositeSurface != nil {
		mp, err := multiPolygonFromCompositeSurface(m.CompositeSurface, resolver, inheritDim)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.Shell != nil {
		mp, err := multiPolygonFromShell(m.Shell, preferDim(derefDim(m.Shell.SrsDimension), inheritDim), resolver)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.PolyhedralSurface != nil {
		return polygonFromSurface(m.PolyhedralSurface, resolver, inheritDim)
	}
	if m.TriangulatedSurface != nil {
		ts := m.TriangulatedSurface
		if ts.Patches == nil || len(ts.Patches.Triangle) == 0 {
			return polygonFromSurface(ts, resolver, inheritDim)
		}
		return polygonFromTriangle(&ts.Patches.Triangle[0], preferDim(derefDim(ts.SrsDimension), inheritDim), resolver)
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
	// xlink metadata attributes — not used for geometry.
	_ = m.NilReason
	_ = m.RemoteSchema
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	_ = m.Owns
	return core.Polygon(nil), nil
}

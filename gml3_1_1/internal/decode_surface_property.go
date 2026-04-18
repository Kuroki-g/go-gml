package internal

import (
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// multiPolygonFromSurfaceProperty converts a SurfacePropertyType to a MultiPolygon.
// Used for gml:Solid exterior/interior where the surface may be a CompositeSurface.
func multiPolygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim *uint, resolver *curveResolver) (core.MultiPolygon, error) {
	if m.CompositeSurface != nil {
		return multiPolygonFromCompositeSurface(m.CompositeSurface, resolver, inheritDim)
	}
	if m.PolyhedralSurface != nil {
		ps := m.PolyhedralSurface
		dim := preferDim(ps.SrsDimension, inheritDim)
		return multiPolygonFromPolyhedralSurface(ps, dim, resolver)
	}
	if m.TriangulatedSurface != nil {
		ts := m.TriangulatedSurface
		return multiPolygonFromTrianglePatchArrayProperty(ts.TrianglePatches, preferDim(ts.SrsDimension, inheritDim), resolver)
	}
	if m.Tin != nil {
		t := m.Tin
		return multiPolygonFromTrianglePatchArrayProperty(t.TrianglePatches, preferDim(t.SrsDimension, inheritDim), resolver)
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
			return multiPolygonFromSurfaceProperty(os.BaseSurface, preferDim(os.SrsDimension, inheritDim), resolver)
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
		return nil, fmt.Errorf("gml: unresolved xlink:href %q", m.Href)
	}
	// xlink metadata attributes — not used for geometry.
	_ = m.RemoteSchema
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	return nil, nil
}

// polygonFromSurfaceProperty converts a SurfacePropertyType to a Polygon.
// For CompositeSurface members, use multiPolygonFromSurfaceProperty instead.
func polygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim *uint, resolver *curveResolver) (core.Polygon, error) {
	if m.Polygon != nil {
		return polygonFromXML(m.Polygon, inheritDim)
	}
	if m.Surface != nil {
		return polygonFromSurface(m.Surface, resolver, inheritDim)
	}
	if m.OrientableSurface != nil {
		os := m.OrientableSurface
		if os.BaseSurface != nil {
			return polygonFromSurfaceProperty(os.BaseSurface, preferDim(os.SrsDimension, inheritDim), resolver)
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
	if m.PolyhedralSurface != nil {
		ps := m.PolyhedralSurface
		dim := preferDim(ps.SrsDimension, inheritDim)
		mp, err := multiPolygonFromPolyhedralSurface(ps, dim, resolver)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.TriangulatedSurface != nil {
		ts := m.TriangulatedSurface
		mp, err := multiPolygonFromTrianglePatchArrayProperty(ts.TrianglePatches, preferDim(ts.SrsDimension, inheritDim), resolver)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.Tin != nil {
		t := m.Tin
		mp, err := multiPolygonFromTrianglePatchArrayProperty(t.TrianglePatches, preferDim(t.SrsDimension, inheritDim), resolver)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.Href != "" {
		id := strings.TrimPrefix(m.Href, "#")
		if poly, ok := resolver.polygonByID[id]; ok {
			return poly, nil
		}
		return core.Polygon(nil), fmt.Errorf("gml: unresolved xlink:href %q", m.Href)
	}
	// xlink metadata attributes — not used for geometry.
	_ = m.RemoteSchema
	_ = m.TypeField
	_ = m.Role
	_ = m.Arcrole
	_ = m.Title
	_ = m.Show
	_ = m.Actuate
	return core.Polygon(nil), nil
}

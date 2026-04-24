package internal

import (
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// multiPolygonFromSurfaceProperty converts a SurfacePropertyType to a MultiPolygon.
// CompositeSurface/Shell members produce multiple Polygons; all other surface types produce one.
func multiPolygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim *uint, inheritSrsName *string, resolver *curveResolver) (core.MultiPolygon, error) {
	if m.NilReason != nil {
		return nil, nil
	}
	if m.CompositeSurface != nil {
		return multiPolygonFromCompositeSurface(m.CompositeSurface, resolver, inheritDim, inheritSrsName)
	}
	if m.Shell != nil {
		return multiPolygonFromShell(m.Shell, preferDim(m.Shell.SrsDimension, inheritDim), inheritSrsName, resolver)
	}
	if m.PolyhedralSurface != nil {
		ps := m.PolyhedralSurface
		dim := preferDim(ps.SrsDimension, inheritDim)
		return multiPolygonFromSurfacePatchArrayProperty(ps.Patches, dim, preferSrsName(ps.SrsName, inheritSrsName), resolver)
	}
	if m.TriangulatedSurface != nil {
		ts := m.TriangulatedSurface
		return multiPolygonFromSurfacePatchArrayProperty(ts.Patches, preferDim(ts.SrsDimension, inheritDim), preferSrsName(ts.SrsName, inheritSrsName), resolver)
	}
	if m.Tin != nil {
		t := m.Tin
		return multiPolygonFromSurfacePatchArrayProperty(t.TrianglePatches, preferDim(t.SrsDimension, inheritDim), preferSrsName(t.SrsName, inheritSrsName), resolver)
	}
	if m.Polygon != nil {
		poly, err := polygonFromXML(m.Polygon, inheritDim, inheritSrsName)
		if err != nil {
			return nil, err
		}
		return core.MultiPolygon{poly}, nil
	}
	if m.Surface != nil {
		s := m.Surface
		if s.Patches == nil {
			return nil, nil
		}
		dim := preferDim(s.SrsDimension, inheritDim)
		return multiPolygonFromSurfacePatchArrayProperty(s.Patches, dim, preferSrsName(s.SrsName, inheritSrsName), resolver)
	}
	if m.OrientableSurface != nil {
		os := m.OrientableSurface
		if os.BaseSurface != nil {
			return multiPolygonFromSurfaceProperty(os.BaseSurface, preferDim(os.SrsDimension, inheritDim), preferSrsName(os.SrsName, inheritSrsName), resolver)
		}
		return nil, nil
	}
	if m.Href != nil {
		id := strings.TrimPrefix(*m.Href, "#")
		if mp, ok := resolver.multiPolygonByID[id]; ok {
			return mp, nil
		}
		if poly, ok := resolver.polygonByID[id]; ok {
			return core.MultiPolygon{poly}, nil
		}
		return nil, fmt.Errorf("gml: unresolved xlink:href %q", *m.Href)
	}
	// xlink metadata attributes — not used for geometry.
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
func polygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim *uint, inheritSrsName *string, resolver *curveResolver) (core.Polygon, error) {
	if m.NilReason != nil {
		return core.Polygon(nil), nil
	}
	if m.Polygon != nil {
		return polygonFromXML(m.Polygon, inheritDim, inheritSrsName)
	}
	if m.Surface != nil {
		return polygonFromSurface(m.Surface, resolver, inheritDim, inheritSrsName)
	}
	if m.OrientableSurface != nil {
		os := m.OrientableSurface
		if os.BaseSurface != nil {
			return polygonFromSurfaceProperty(os.BaseSurface, preferDim(os.SrsDimension, inheritDim), preferSrsName(os.SrsName, inheritSrsName), resolver)
		}
	}
	if m.CompositeSurface != nil {
		mp, err := multiPolygonFromCompositeSurface(m.CompositeSurface, resolver, inheritDim, inheritSrsName)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.Shell != nil {
		mp, err := multiPolygonFromShell(m.Shell, preferDim(m.Shell.SrsDimension, inheritDim), inheritSrsName, resolver)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.PolyhedralSurface != nil {
		return polygonFromSurface(m.PolyhedralSurface, resolver, inheritDim, inheritSrsName)
	}
	if m.TriangulatedSurface != nil {
		ts := m.TriangulatedSurface
		mp, err := multiPolygonFromSurfacePatchArrayProperty(ts.Patches, preferDim(ts.SrsDimension, inheritDim), preferSrsName(ts.SrsName, inheritSrsName), resolver)
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
		mp, err := multiPolygonFromSurfacePatchArrayProperty(t.TrianglePatches, preferDim(t.SrsDimension, inheritDim), preferSrsName(t.SrsName, inheritSrsName), resolver)
		if err != nil {
			return nil, err
		}
		if len(mp) == 0 {
			return core.Polygon(nil), nil
		}
		return mp[0], nil
	}
	if m.Href != nil {
		id := strings.TrimPrefix(*m.Href, "#")
		if poly, ok := resolver.polygonByID[id]; ok {
			return poly, nil
		}
		return core.Polygon(nil), fmt.Errorf("gml: unresolved xlink:href %q", *m.Href)
	}
	// xlink metadata attributes — not used for geometry.
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

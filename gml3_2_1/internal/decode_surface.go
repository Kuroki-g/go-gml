package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// handleSurface decodes a gml:Surface, caches the resulting geometry by gml:id, and returns it.
func (r *Reader) handleSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	g, err := decodeSurfaceElement(dec, se, r.resolver, r.globalDim)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		switch v := g.Value.(type) {
		case core.Polygon:
			r.resolver.polygonByID[id] = v
		case core.MultiPolygon:
			r.resolver.multiPolygonByID[id] = v
		}
	}
	return g, nil
}

func decodeSurfaceElement(dec *xml.Decoder, se xml.StartElement, resolver *curveResolver, fallbackDim uint) (core.Geometry, error) {
	var x gen.SurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Surface: %w", err)
	}
	if x.Patches != nil && (len(x.Patches.PolygonPatch)+len(x.Patches.Rectangle)) > 1 {
		dim := preferDim(x.SrsDimension, fallbackDim)
		mp, err := multiPolygonFromSurfacePatchArrayProperty(x.Patches, dim, resolver)
		if err != nil {
			return core.Geometry{}, err
		}
		return core.Geometry{Value: mp, SRSName: x.SrsName}, nil
	}
	poly, err := polygonFromSurface(&x, resolver, fallbackDim)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func polygonFromSurface(x *gen.SurfaceType, resolver *curveResolver, fallbackDim uint) (core.Polygon, error) {
	if x.Patches == nil {
		return core.Polygon{}, nil
	}
	dim := preferDim(x.SrsDimension, fallbackDim)
	return polygonFromSurfacePatchArrayProperty(x.Patches, dim, resolver)
}

func polygonFromSurfacePatchArrayProperty(sp *gen.SurfacePatchArrayPropertyType, dim uint, resolver *curveResolver) (core.Polygon, error) {
	if len(sp.PolygonPatch) > 0 {
		return polygonFromPatch(&sp.PolygonPatch[0], dim, resolver)
	}
	if len(sp.Rectangle) > 0 {
		return polygonFromRectangle(&sp.Rectangle[0], dim, resolver)
	}
	_ = sp.Triangle // SF-2: 未実装 (docs/issues/sf2-curves.md)
	_ = sp.Cone
	_ = sp.Cylinder
	_ = sp.Sphere
	return core.Polygon{}, nil
}

func polygonFromPatch(patch *gen.PolygonPatchType, inheritDim uint, resolver *curveResolver) (core.Polygon, error) {
	var rings []core.Ring
	if patch.Exterior != nil {
		r, err := ringFromAbstractRingProperty(patch.Exterior, inheritDim, "exterior", resolver)
		if err != nil {
			return nil, err
		}
		if r != nil {
			rings = append(rings, r)
		}
	}
	for i, ir := range patch.Interior {
		r, err := ringFromAbstractRingProperty(&ir, inheritDim, fmt.Sprintf("interior[%d]", i), resolver)
		if err != nil {
			return nil, err
		}
		if r != nil {
			rings = append(rings, r)
		}
	}
	return core.Polygon(rings), nil
}

func ringFromAbstractRingProperty(prop *gen.AbstractRingPropertyType, inheritDim uint, label string, resolver *curveResolver) (core.Ring, error) {
	if prop.LinearRing != nil {
		lr := prop.LinearRing
		dim := preferDim(lr.SrsDimension, inheritDim)
		if lr.PosList == nil {
			return nil, nil
		}
		r, err := core.RingFromPosListString(lr.PosList.Value, preferDim(lr.PosList.SrsDimension, dim))
		if err != nil {
			return nil, fmt.Errorf("gml: PolygonPatch %s LinearRing: %w", label, err)
		}
		return r, nil
	}
	if prop.Ring != nil {
		r, err := ringFromRingType(prop.Ring, inheritDim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: PolygonPatch %s Ring: %w", label, err)
		}
		return r, nil
	}
	return nil, nil
}

func ringFromRingType(ring *gen.RingType, inheritDim uint, resolver *curveResolver) (core.Ring, error) {
	var pts core.Ring
	dim := preferDim(ring.SrsDimension, inheritDim)
	for i := range ring.CurveMember {
		ls, err := lineStringFromCurveProperty(&ring.CurveMember[i], dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("curveMember[%d]: %w", i, err)
		}
		if ls == nil {
			return nil, fmt.Errorf("curveMember[%d]: unresolved curve reference (href=%q)", i, ring.CurveMember[i].Href)
		}
		if len(pts) > 0 && len(ls) > 0 {
			pts = append(pts, ls[1:]...)
		} else {
			pts = append(pts, ls...)
		}
	}
	return pts, nil
}

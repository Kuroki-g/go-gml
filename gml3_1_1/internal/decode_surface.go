package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// handleSurface decodes a gml:Surface, caches the resulting Polygon by gml:id, and returns it.
func (r *Reader) handleSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	g, err := decodeSurfaceElement(dec, se, r.resolver, r.globalDim)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		if poly, ok := g.Value.(core.Polygon); ok {
			r.resolver.polygonByID[id] = poly
		}
	}
	return g, err
}

func decodeSurfaceElement(dec *xml.Decoder, se xml.StartElement, resolver *curveResolver, fallbackDim int) (core.Geometry, error) {
	var x gen.SurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Surface: %w", err)
	}
	poly, err := polygonFromSurface(&x, resolver, fallbackDim)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func polygonFromSurface(x *gen.SurfaceType, resolver *curveResolver, fallbackDim int) (core.Polygon, error) {
	if x.Patches == nil {
		return core.Polygon{}, nil
	}
	dim := preferDim(derefDim(x.SrsDimension), fallbackDim)
	return polygonFromSurfacePatchArrayProperty(x.Patches, dim, resolver)
}

func polygonFromSurfacePatchArrayProperty(sp *gen.SurfacePatchArrayPropertyType, dim int, resolver *curveResolver) (core.Polygon, error) {
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

func polygonFromRectangle(rect *gen.RectangleType, inheritDim int, resolver *curveResolver) (core.Polygon, error) {
	prop := rect.Exterior
	if prop == nil {
		prop = rect.OuterBoundaryIs
	}
	if prop == nil {
		return core.Polygon(nil), nil
	}
	r, err := ringFromAbstractRingProperty(prop, inheritDim, "exterior", resolver)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return core.Polygon(nil), nil
	}
	return core.Polygon{r}, nil
}

func polygonFromPatch(patch *gen.PolygonPatchType, inheritDim int, resolver *curveResolver) (core.Polygon, error) {
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

func ringFromAbstractRingProperty(prop *gen.AbstractRingPropertyType, inheritDim int, label string, resolver *curveResolver) (core.Ring, error) {
	if prop.LinearRing != nil {
		lr := prop.LinearRing
		dim := preferDim(inheritDim, derefDim(lr.SrsDimension))
		if lr.PosList == nil {
			return nil, nil
		}
		r, err := core.RingFromPosListString(lr.PosList.Value, preferDim(dim, derefDim(lr.PosList.SrsDimension)))
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

func ringFromRingType(ring *gen.RingType, inheritDim int, resolver *curveResolver) (core.Ring, error) {
	var pts core.Ring
	dim := preferDim(inheritDim, derefDim(ring.SrsDimension))
	for i, cm := range ring.CurveMember {
		curve := cm.Curve
		if curve == nil && cm.Href != "" {
			curve = resolver.resolve(strings.TrimPrefix(cm.Href, "#"))
		}
		if curve == nil {
			continue
		}
		ls, err := lineStringFromCurve(curve, dim)
		if err != nil {
			return nil, fmt.Errorf("curveMember[%d]: %w", i, err)
		}
		if len(pts) > 0 && len(ls) > 0 {
			pts = append(pts, ls[1:]...)
		} else {
			pts = append(pts, ls...)
		}
	}
	return pts, nil
}

package gml

import (
	"encoding/xml"
	"fmt"
	"strings"

	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// handleSurface decodes a gml:Surface, caches the resulting Polygon by gml:id, and returns it.
func (r *Reader) handleSurface(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	id := extractGMLID(se)
	g, err := decodeSurfaceElement(dec, se, r.resolver)
	if err != nil {
		return Geometry{}, err
	}
	if id != "" {
		if poly, ok := g.Value.(Polygon); ok {
			r.resolver.polygonByID[id] = poly
		}
	}
	return g, err
}

// decodeSurfaceElement handles gml:Surface > gml:patches > gml:PolygonPatch.
// resolver is used to resolve xlink:href references in gml:Ring/curveMember.
func decodeSurfaceElement(dec *xml.Decoder, se xml.StartElement, resolver *curveResolver) (Geometry, error) {
	var x v3_2_1.SurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Surface: %w", err)
	}
	poly, err := polygonFromSurface(&x, resolver)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func polygonFromSurface(x *v3_2_1.SurfaceType, resolver *curveResolver) (Polygon, error) {
	if x.Patches == nil || len(x.Patches.PolygonPatch) == 0 {
		return Polygon{}, nil
	}
	return polygonFromPatch(&x.Patches.PolygonPatch[0], derefDim(x.SrsDimension), resolver)
}

func polygonFromPatch(patch *v3_2_1.PolygonPatchType, inheritDim int, resolver *curveResolver) (Polygon, error) {
	var rings []Ring
	if patch.Exterior != nil {
		r, err := ringFromAbstractRingProp(patch.Exterior, inheritDim, "exterior", resolver)
		if err != nil {
			return nil, err
		}
		if r != nil {
			rings = append(rings, r)
		}
	}
	for i, ir := range patch.Interior {
		r, err := ringFromAbstractRingProp(&ir, inheritDim, fmt.Sprintf("interior[%d]", i), resolver)
		if err != nil {
			return nil, err
		}
		if r != nil {
			rings = append(rings, r)
		}
	}
	return Polygon(rings), nil
}

// ringFromAbstractRingProp extracts a Ring from an AbstractRingPropertyType,
// supporting both gml:LinearRing (posList) and gml:Ring (curveMember > Curve).
func ringFromAbstractRingProp(prop *v3_2_1.AbstractRingPropertyType, inheritDim int, label string, resolver *curveResolver) (Ring, error) {
	if prop.LinearRing != nil {
		lr := prop.LinearRing
		dim := preferDim(inheritDim, derefDim(lr.SrsDimension))
		if lr.PosList == nil {
			return nil, nil
		}
		r, err := RingFromPosListString(lr.PosList.Value, preferDim(dim, derefDim(lr.PosList.SrsDimension)))
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

// ringFromRingType builds a Ring by concatenating curveMember Curves.
// Inline Curves are used directly; xlink:href references are resolved via resolver
// (supports OrientableCurve → Curve indirection used in old-format N03 files).
func ringFromRingType(ring *v3_2_1.RingType, inheritDim int, resolver *curveResolver) (Ring, error) {
	var pts Ring
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
		// Avoid duplicating shared endpoints between consecutive segments.
		if len(pts) > 0 && len(ls) > 0 {
			pts = append(pts, ls[1:]...)
		} else {
			pts = append(pts, ls...)
		}
	}
	return pts, nil
}

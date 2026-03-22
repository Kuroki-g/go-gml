package gml

import (
	"encoding/xml"
	"fmt"

	v3 "github.com/Kuroki-g/go-gml/pkg/gml/v3"
)

// decodeSurfaceElement handles gml:Surface > gml:patches > gml:PolygonPatch.
// Returns a Polygon built from the first PolygonPatch's exterior/interior rings.
// Supports both gml:LinearRing (posList) and gml:Ring with inline curveMember Curves.
func decodeSurfaceElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3.SurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Surface: %w", err)
	}
	poly, err := polygonFromSurface(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func polygonFromSurface(x *v3.SurfaceType) (Polygon, error) {
	if x.Patches == nil || len(x.Patches.PolygonPatch) == 0 {
		return Polygon{}, nil
	}
	return polygonFromPatch(&x.Patches.PolygonPatch[0], derefDim(x.SrsDimension))
}

func polygonFromPatch(patch *v3.PolygonPatchType, inheritDim int) (Polygon, error) {
	var rings []Ring
	if patch.Exterior != nil {
		r, err := ringFromAbstractRingProp(patch.Exterior, inheritDim, "exterior")
		if err != nil {
			return nil, err
		}
		if r != nil {
			rings = append(rings, r)
		}
	}
	for i, ir := range patch.Interior {
		r, err := ringFromAbstractRingProp(&ir, inheritDim, fmt.Sprintf("interior[%d]", i))
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
func ringFromAbstractRingProp(prop *v3.AbstractRingPropertyType, inheritDim int, label string) (Ring, error) {
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
		r, err := ringFromRingType(prop.Ring, inheritDim)
		if err != nil {
			return nil, fmt.Errorf("gml: PolygonPatch %s Ring: %w", label, err)
		}
		return r, nil
	}
	return nil, nil
}

// ringFromRingType builds a Ring by concatenating all inline curveMember Curves.
func ringFromRingType(ring *v3.RingType, inheritDim int) (Ring, error) {
	var pts Ring
	dim := preferDim(inheritDim, derefDim(ring.SrsDimension))
	for i, cm := range ring.CurveMember {
		if cm.Curve == nil {
			continue
		}
		ls, err := lineStringFromCurve(cm.Curve, dim)
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

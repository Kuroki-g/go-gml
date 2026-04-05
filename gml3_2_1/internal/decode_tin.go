package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

func (r *Reader) handleTriangulatedSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	var x gen.SurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: TriangulatedSurface: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	mp, err := multiPolygonFromPatchArray(x.Patches, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		r.resolver.multiPolygonByID[id] = mp
	}
	return core.Geometry{Value: mp, SRSName: x.SrsName}, nil
}

func (r *Reader) handleTin(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	var x gen.TinType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: Tin: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	mp, err := multiPolygonFromPatchArray(x.TrianglePatches, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		r.resolver.multiPolygonByID[id] = mp
	}
	return core.Geometry{Value: mp, SRSName: x.SrsName}, nil
}

// multiPolygonFromPatchArray converts a SurfacePatchArrayPropertyType to a MultiPolygon,
// handling PolygonPatch, Rectangle, and Triangle patch types.
func multiPolygonFromPatchArray(sp *gen.SurfacePatchArrayPropertyType, dim int, resolver *curveResolver) (core.MultiPolygon, error) {
	if sp == nil {
		return nil, nil
	}
	// SF-2 curved patch types — not yet implemented (see docs/issues/sf2-curves.md).
	_ = sp.Cone
	_ = sp.Cylinder
	_ = sp.Sphere
	var result core.MultiPolygon
	for i := range sp.Triangle {
		poly, err := polygonFromTriangle(&sp.Triangle[i], dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: Triangle[%d]: %w", i, err)
		}
		if poly != nil {
			result = append(result, poly)
		}
	}
	for i := range sp.PolygonPatch {
		poly, err := polygonFromPatch(&sp.PolygonPatch[i], dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: PolygonPatch[%d]: %w", i, err)
		}
		if poly != nil {
			result = append(result, poly)
		}
	}
	for i := range sp.Rectangle {
		poly, err := polygonFromRectangle(&sp.Rectangle[i], dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: Rectangle[%d]: %w", i, err)
		}
		if poly != nil {
			result = append(result, poly)
		}
	}
	return result, nil
}

func polygonFromRectangle(rect *gen.RectangleType, inheritDim int, resolver *curveResolver) (core.Polygon, error) {
	if rect.Exterior == nil {
		return core.Polygon(nil), nil
	}
	r, err := ringFromAbstractRingProp(rect.Exterior, inheritDim, "exterior", resolver)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return core.Polygon(nil), nil
	}
	return core.Polygon{r}, nil
}

func polygonFromTriangle(t *gen.TriangleType, dim int, resolver *curveResolver) (core.Polygon, error) {
	if t.Exterior == nil {
		return core.Polygon(nil), nil
	}
	ring, err := ringFromAbstractRingProp(t.Exterior, dim, "exterior", resolver)
	if err != nil {
		return nil, fmt.Errorf("gml: Triangle exterior: %w", err)
	}
	if ring == nil {
		return core.Polygon(nil), nil
	}
	return core.Polygon{ring}, nil
}

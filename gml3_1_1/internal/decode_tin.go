package internal

import (
	"encoding/xml"
	"fmt"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

func (r *Reader) handleTriangulatedSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	var x gen.TriangulatedSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: TriangulatedSurface: %w", err)
	}
	dim := preferDim(derefDim(x.SrsDimension), r.globalDim)
	mp, err := multiPolygonFromTrianglePatches(x.TrianglePatches, dim, r.resolver)
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
	mp, err := multiPolygonFromTrianglePatches(x.TrianglePatches, dim, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		r.resolver.multiPolygonByID[id] = mp
	}
	return core.Geometry{Value: mp, SRSName: x.SrsName}, nil
}

func multiPolygonFromTrianglePatches(tp *gen.TrianglePatchArrayPropertyType, dim int, resolver *curveResolver) (core.MultiPolygon, error) {
	if tp == nil {
		return nil, nil
	}
	var result core.MultiPolygon
	for i := range tp.Triangle {
		poly, err := polygonFromTriangle(&tp.Triangle[i], dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: trianglePatches[%d]: %w", i, err)
		}
		if poly != nil {
			result = append(result, poly)
		}
	}
	return result, nil
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

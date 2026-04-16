package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
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
	poly, err := polygonFromSurfaceProperty(x.BaseSurface, preferDim(x.SrsDimension, r.globalDim), r.resolver)
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
func multiPolygonFromCompositeSurface(x *gen.CompositeSurfaceType, resolver *curveResolver, fallbackDim uint) (core.MultiPolygon, error) {
	dim := preferDim(x.SrsDimension, fallbackDim)
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

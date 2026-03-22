package gml

import (
	"encoding/xml"
	"fmt"
	"strings"

	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// handleCompositeSurface decodes a gml:CompositeSurface, caches the result by gml:id, and returns a Polygon.
// The rings from all surface members are collected into a single Polygon:
// the first member's exterior ring becomes rings[0]; subsequent rings follow.
func (r *Reader) handleCompositeSurface(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	id := extractGMLID(se)
	var x v3_2_1.CompositeSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: CompositeSurface: %w", err)
	}
	poly, err := polygonFromCompositeSurface(&x, r.resolver)
	if err != nil {
		return Geometry{}, err
	}
	if id != "" {
		r.resolver.polygonByID[id] = poly
	}
	return Geometry{Value: poly, SRSName: x.SrsName}, nil
}

// handleOrientableSurface decodes a gml:OrientableSurface, caches the result by gml:id,
// and returns a Polygon by following its baseSurface property.
func (r *Reader) handleOrientableSurface(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	id := extractGMLID(se)
	var x v3_2_1.OrientableSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: OrientableSurface: %w", err)
	}
	if x.BaseSurface == nil {
		return Geometry{Value: Polygon(nil), SRSName: x.SrsName}, nil
	}
	poly, err := polygonFromSurfaceProperty(x.BaseSurface, derefDim(x.SrsDimension), r.resolver)
	if err != nil {
		return Geometry{}, err
	}
	if id != "" {
		r.resolver.polygonByID[id] = poly
	}
	return Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func polygonFromCompositeSurface(x *v3_2_1.CompositeSurfaceType, resolver *curveResolver) (Polygon, error) {
	var allRings []Ring
	dim := derefDim(x.SrsDimension)
	for i, m := range x.SurfaceMember {
		poly, err := polygonFromSurfaceProperty(&m, dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: CompositeSurface surfaceMember[%d]: %w", i, err)
		}
		allRings = append(allRings, poly...)
	}
	return Polygon(allRings), nil
}

// polygonFromSurfaceProperty converts a SurfacePropertyType to a Polygon.
// Supports inline Polygon, Surface, OrientableSurface, CompositeSurface,
// and xlink:href references resolved via the resolver's polygonByID cache.
func polygonFromSurfaceProperty(m *v3_2_1.SurfacePropertyType, inheritDim int, resolver *curveResolver) (Polygon, error) {
	if m.Polygon != nil {
		return polygonFromXML(m.Polygon)
	}
	if m.Surface != nil {
		return polygonFromSurface(m.Surface, resolver)
	}
	if m.OrientableSurface != nil {
		os := m.OrientableSurface
		if os.BaseSurface != nil {
			return polygonFromSurfaceProperty(os.BaseSurface, preferDim(inheritDim, derefDim(os.SrsDimension)), resolver)
		}
	}
	if m.CompositeSurface != nil {
		return polygonFromCompositeSurface(m.CompositeSurface, resolver)
	}
	if m.Href != "" {
		id := strings.TrimPrefix(m.Href, "#")
		if poly, ok := resolver.polygonByID[id]; ok {
			return poly, nil
		}
	}
	return Polygon(nil), nil
}

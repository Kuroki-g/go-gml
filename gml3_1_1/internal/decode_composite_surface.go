package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
)

// handleCompositeSurface decodes a gml:CompositeSurface, caches the result by gml:id, and returns a Polygon.
func (r *Reader) handleCompositeSurface(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	id := extractGMLID(se)
	var x gen.CompositeSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: CompositeSurface: %w", err)
	}
	poly, err := polygonFromCompositeSurface(&x, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		r.resolver.polygonByID[id] = poly
	}
	return core.Geometry{Value: poly, SRSName: x.SrsName}, nil
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
	poly, err := polygonFromSurfaceProperty(x.BaseSurface, derefDim(x.SrsDimension), r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	if id != "" {
		r.resolver.polygonByID[id] = poly
	}
	return core.Geometry{Value: poly, SRSName: x.SrsName}, nil
}

func polygonFromCompositeSurface(x *gen.CompositeSurfaceType, resolver *curveResolver) (core.Polygon, error) {
	var allRings []core.Ring
	dim := derefDim(x.SrsDimension)
	for i, m := range x.SurfaceMember {
		poly, err := polygonFromSurfaceProperty(&m, dim, resolver)
		if err != nil {
			return nil, fmt.Errorf("gml: CompositeSurface surfaceMember[%d]: %w", i, err)
		}
		allRings = append(allRings, poly...)
	}
	return core.Polygon(allRings), nil
}

// polygonFromSurfaceProperty converts a SurfacePropertyType to a Polygon.
func polygonFromSurfaceProperty(m *gen.SurfacePropertyType, inheritDim int, resolver *curveResolver) (core.Polygon, error) {
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
	return core.Polygon(nil), nil
}

package gml

import (
	"encoding/xml"
	"fmt"
	"strings"

	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// curveResolver resolves gml:id references to geometry elements.
// Handles Curve/OrientableCurve for curve xlink:href resolution, and
// Polygon for surface xlink:href resolution in CompositeSurface/OrientableSurface.
type curveResolver struct {
	curves         map[string]*v3_2_1.CurveType
	orientable     map[string]*v3_2_1.OrientableCurveType
	polygonByID    map[string]Polygon
	lineStringByID map[string]LineString
	gridByID       map[string]*gridBounds
}

func newCurveResolver() *curveResolver {
	return &curveResolver{
		curves:         make(map[string]*v3_2_1.CurveType),
		orientable:     make(map[string]*v3_2_1.OrientableCurveType),
		polygonByID:    make(map[string]Polygon),
		lineStringByID: make(map[string]LineString),
		gridByID:       make(map[string]*gridBounds),
	}
}

// resolve returns the CurveType for the given id (without leading '#').
// Resolution order:
//  1. Direct Curve lookup by id.
//  2. OrientableCurve lookup by id → follow baseCurve to Curve.
//
// Unresolvable references return nil; the caller emits an empty ring.
// Per GML 3.2.1 XSD, xlink:href="#foo" must match an existing gml:id="foo"
// exactly (gml:id is xs:ID; xlink:href is anyURI with fragment semantics).
func (cr *curveResolver) resolve(id string) *v3_2_1.CurveType {
	if c, ok := cr.curves[id]; ok {
		return c
	}
	if oc, ok := cr.orientable[id]; ok && oc.BaseCurve != nil {
		if oc.BaseCurve.Curve != nil {
			return oc.BaseCurve.Curve
		}
		if oc.BaseCurve.Href != "" {
			return cr.curves[strings.TrimPrefix(oc.BaseCurve.Href, "#")]
		}
	}
	return nil
}

func decodeCurveType(dec *xml.Decoder, se xml.StartElement) (*v3_2_1.CurveType, error) {
	var x v3_2_1.CurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return nil, fmt.Errorf("gml: Curve: %w", err)
	}
	return &x, nil
}

// handleCurve decodes a gml:Curve, caches it for xlink:href resolution, and returns a LineString.
func (r *Reader) handleCurve(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	x, err := decodeCurveType(dec, se)
	if err != nil {
		return Geometry{}, err
	}
	if x.Id != "" {
		r.resolver.curves[x.Id] = x
	}
	ls, err := lineStringFromCurve(x, 0)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: ls, SRSName: x.SrsName}, nil
}

// cacheOrientableCurve decodes a gml:OrientableCurve and caches it by gml:id.
// These elements are not emitted as geometry; they serve as xlink:href targets
// that redirect to the underlying Curve via baseCurve.
func (r *Reader) cacheOrientableCurve(dec *xml.Decoder, se xml.StartElement) error {
	var x v3_2_1.OrientableCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return fmt.Errorf("gml: OrientableCurve: %w", err)
	}
	if x.Id != "" {
		r.resolver.orientable[x.Id] = &x
	}
	return nil
}

// lineStringFromCurve converts a gml:Curve to a LineString.
// inheritDim is used when the Curve element itself has no srsDimension attribute.
func lineStringFromCurve(x *v3_2_1.CurveType, inheritDim int) (LineString, error) {
	if x.Segments == nil || len(x.Segments.LineStringSegment) == 0 {
		return nil, fmt.Errorf("gml: Curve has no LineStringSegment")
	}
	var result LineString
	for i, seg := range x.Segments.LineStringSegment {
		var ls LineString
		var err error
		if seg.PosList != nil {
			dim := preferDim(preferDim(inheritDim, derefDim(x.SrsDimension)), derefDim(seg.PosList.SrsDimension))
			ls, err = LineStringFromPosListString(seg.PosList.Value, dim)
		} else if seg.Coordinates != nil {
			var coords []float64
			coords, err = ParseCoordinates(seg.Coordinates.Value, derefStrOr(seg.Coordinates.Cs, ","), derefStrOr(seg.Coordinates.Ts, " "))
			if err == nil {
				ls, err = LineStringFromFlat(coords, 2)
			}
		} else {
			return nil, fmt.Errorf("gml: LineStringSegment[%d] has no coordinate data", i)
		}
		if err != nil {
			return nil, err
		}
		if len(result) > 0 && len(ls) > 0 {
			result = append(result, ls[1:]...)
		} else {
			result = append(result, ls...)
		}
	}
	return result, nil
}

package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// curveResolver resolves gml:id references to geometry elements.
type curveResolver struct {
	curves         map[string]*gen.CurveType
	orientable     map[string]*gen.OrientableCurveType
	polygonByID    map[string]core.Polygon
	lineStringByID map[string]core.LineString
	gridByID       map[string]*gridBounds
}

func newCurveResolver() *curveResolver {
	return &curveResolver{
		curves:         make(map[string]*gen.CurveType),
		orientable:     make(map[string]*gen.OrientableCurveType),
		polygonByID:    make(map[string]core.Polygon),
		lineStringByID: make(map[string]core.LineString),
		gridByID:       make(map[string]*gridBounds),
	}
}

// resolve returns the CurveType for the given id (without leading '#').
func (cr *curveResolver) resolve(id string) *gen.CurveType {
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

func decodeCurveType(dec *xml.Decoder, se xml.StartElement) (*gen.CurveType, error) {
	var x gen.CurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return nil, fmt.Errorf("gml: Curve: %w", err)
	}
	return &x, nil
}

// handleCurve decodes a gml:Curve, caches it for xlink:href resolution, and returns a LineString.
func (r *Reader) handleCurve(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	x, err := decodeCurveType(dec, se)
	if err != nil {
		return core.Geometry{}, err
	}
	if x.Id != "" {
		r.resolver.curves[x.Id] = x
	}
	ls, err := lineStringFromCurve(x, 0)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: ls, SRSName: x.SrsName}, nil
}

// cacheOrientableCurve decodes a gml:OrientableCurve and caches it by gml:id.
func (r *Reader) cacheOrientableCurve(dec *xml.Decoder, se xml.StartElement) error {
	var x gen.OrientableCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return fmt.Errorf("gml: OrientableCurve: %w", err)
	}
	if x.Id != "" {
		r.resolver.orientable[x.Id] = &x
	}
	return nil
}

// lineStringFromCurve converts a gml:Curve to a LineString.
func lineStringFromCurve(x *gen.CurveType, inheritDim int) (core.LineString, error) {
	if x.Segments == nil || len(x.Segments.LineStringSegment) == 0 {
		return nil, fmt.Errorf("gml: Curve has no LineStringSegment")
	}
	var result core.LineString
	for i, seg := range x.Segments.LineStringSegment {
		var ls core.LineString
		var err error
		if seg.PosList != nil {
			dim := preferDim(preferDim(inheritDim, derefDim(x.SrsDimension)), derefDim(seg.PosList.SrsDimension))
			ls, err = LineStringFromPosListString(seg.PosList.Value, dim)
		} else if seg.Coordinates != nil {
			var coords []float64
			coords, err = core.ParseCoordinates(seg.Coordinates.Value, derefStrOr(seg.Coordinates.Cs, ","), derefStrOr(seg.Coordinates.Ts, " "))
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

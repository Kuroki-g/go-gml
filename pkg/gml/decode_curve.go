package gml

import (
	"encoding/xml"
	"fmt"

	v3 "github.com/Kuroki-g/go-gml/pkg/gml/v3"
)

// decodeCurveElement handles gml:Curve > gml:segments > gml:LineStringSegment.
// Returns a LineString built from all LineStringSegment children.
func decodeCurveElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3.CurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Curve: %w", err)
	}
	ls, err := lineStringFromCurve(&x, 0)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: ls, SRSName: x.SrsName}, nil
}

// lineStringFromCurve converts a gml:Curve to a LineString.
// inheritDim is used when the Curve element itself has no srsDimension attribute.
func lineStringFromCurve(x *v3.CurveType, inheritDim int) (LineString, error) {
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

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
	return Geometry{Value: ls, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

// lineStringFromCurve converts a gml:Curve to a LineString.
// inheritDim is used when the Curve element itself has no srsDimension attribute.
func lineStringFromCurve(x *v3.CurveType, inheritDim int) (LineString, error) {
	if x.Segments == nil || x.Segments.LineStringSegment == nil {
		return nil, fmt.Errorf("gml: Curve has no LineStringSegment")
	}
	seg := x.Segments.LineStringSegment
	if seg.PosList != nil {
		dim := preferDim(preferDim(inheritDim, x.SrsDimension), seg.PosList.SrsDimension)
		return LineStringFromPosListString(seg.PosList.Value, dim)
	}
	if seg.Coordinates != nil {
		coords, err := ParseCoordinates(seg.Coordinates.Value, seg.Coordinates.Cs, seg.Coordinates.Ts)
		if err != nil {
			return nil, err
		}
		return LineStringFromFlat(coords, 2)
	}
	return nil, fmt.Errorf("gml: LineStringSegment has no coordinate data")
}

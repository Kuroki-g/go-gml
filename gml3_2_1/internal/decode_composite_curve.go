package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// handleCompositeCurve decodes a gml:CompositeCurve and returns a LineString.
func (r *Reader) handleCompositeCurve(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	var x gen.CompositeCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: CompositeCurve: %w", err)
	}
	ls, err := lineStringFromCompositeCurveType(&x, nil, nil, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: ls, SRSName: x.SrsName}, nil
}

func lineStringFromCompositeCurveType(x *gen.CompositeCurveType, inheritDim *uint, inheritSrsName *string, resolver *curveResolver) (core.LineString, error) {
	var result core.LineString
	dim := preferDim(x.SrsDimension, inheritDim)
	srsName := preferSrsName(x.SrsName, inheritSrsName)
	for i, cm := range x.CurveMember {
		ls, err := lineStringFromCurveProperty(&cm, dim, srsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("curveMember[%d]: %w", i, err)
		}
		if len(result) > 0 && len(ls) > 0 {
			result = append(result, ls[1:]...)
		} else {
			result = append(result, ls...)
		}
	}
	return result, nil
}

// lineStringFromCurveProperty converts a single CurvePropertyType to a LineString.
func lineStringFromCurveProperty(cm *gen.CurvePropertyType, inheritDim *uint, inheritSrsName *string, resolver *curveResolver) (core.LineString, error) {
	if cm.NilReason != nil {
		return nil, nil
	}
	if cm.Curve != nil {
		return lineStringFromCurve(cm.Curve, inheritDim, inheritSrsName)
	}
	if cm.LineString != nil {
		return lineStringFromXML(cm.LineString, inheritDim, inheritSrsName)
	}
	if cm.CompositeCurve != nil {
		return lineStringFromCompositeCurveType(cm.CompositeCurve, inheritDim, inheritSrsName, resolver)
	}
	if cm.OrientableCurve != nil {
		oc := cm.OrientableCurve
		if oc.BaseCurve != nil {
			if oc.BaseCurve.Curve != nil {
				return lineStringFromCurve(oc.BaseCurve.Curve, inheritDim, inheritSrsName)
			}
			if oc.BaseCurve.Href != "" {
				if c := resolver.resolve(strings.TrimPrefix(oc.BaseCurve.Href, "#")); c != nil {
					return lineStringFromCurve(c, inheritDim, inheritSrsName)
				}
			}
		}
	}
	if cm.LinearRing != nil {
		return lineStringFromLinearRingType(cm.LinearRing, inheritDim, inheritSrsName)
	}
	if cm.Ring != nil {
		return lineStringFromRingType(cm.Ring, inheritDim, inheritSrsName, resolver)
	}
	if cm.Href != "" {
		id := strings.TrimPrefix(cm.Href, "#")
		if c := resolver.resolve(id); c != nil {
			return lineStringFromCurve(c, inheritDim, inheritSrsName)
		}
		if ls, ok := resolver.lineStringByID[id]; ok {
			return ls, nil
		}
	}
	// xlink metadata attributes — not used for curve resolution.
	_ = cm.RemoteSchema
	_ = cm.TypeField
	_ = cm.Role
	_ = cm.Arcrole
	_ = cm.Title
	_ = cm.Show
	_ = cm.Actuate
	_ = cm.Owns
	return nil, nil
}

// lineStringFromLinearRingType extracts coordinates from a LinearRingType as a LineString.
func lineStringFromLinearRingType(x *gen.LinearRingType, inheritDim *uint, inheritSrsName *string) (core.LineString, error) {
	dim := preferDim(x.SrsDimension, inheritDim)
	srsName := preferSrsName(x.SrsName, inheritSrsName)
	if x.PosList != nil {
		return core.LineStringFromPosListString(x.PosList.Value, preferDim(x.PosList.SrsDimension, dim), preferSrsName(x.PosList.SrsName, srsName))
	}
	if len(x.Pos) > 0 {
		var flat []float64
		for _, p := range x.Pos {
			vals, err := core.ParsePosList(p.Value)
			if err != nil {
				return nil, err
			}
			flat = append(flat, vals...)
		}
		d := uint(len(strings.Fields(x.Pos[0].Value)))
		if d == 0 {
			return nil, fmt.Errorf("gml: pos has no values")
		}
		return core.LineStringFromFlat(flat, d)
	}
	if x.Coordinates != nil {
		return core.LineStringFromCoordinatesString(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "), derefStrOr(x.Coordinates.Decimal, "."))
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

// lineStringFromRingType concatenates curveMember segments of a RingType into a LineString.
func lineStringFromRingType(x *gen.RingType, inheritDim *uint, inheritSrsName *string, resolver *curveResolver) (core.LineString, error) {
	var result core.LineString
	dim := preferDim(x.SrsDimension, inheritDim)
	srsName := preferSrsName(x.SrsName, inheritSrsName)
	for i := range x.CurveMember {
		ls, err := lineStringFromCurveProperty(&x.CurveMember[i], dim, srsName, resolver)
		if err != nil {
			return nil, fmt.Errorf("curveMember[%d]: %w", i, err)
		}
		if len(result) > 0 && len(ls) > 0 {
			result = append(result, ls[1:]...)
		} else {
			result = append(result, ls...)
		}
	}
	return result, nil
}

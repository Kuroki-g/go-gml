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
	ls, err := lineStringFromCompositeCurveType(&x, 0, r.resolver)
	if err != nil {
		return core.Geometry{}, err
	}
	return core.Geometry{Value: ls, SRSName: x.SrsName}, nil
}

func lineStringFromCompositeCurveType(x *gen.CompositeCurveType, inheritDim uint, resolver *curveResolver) (core.LineString, error) {
	var result core.LineString
	dim := preferDim(inheritDim, derefDim(x.SrsDimension))
	for i, cm := range x.CurveMember {
		ls, err := lineStringFromCurveProperty(&cm, dim, resolver)
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
func lineStringFromCurveProperty(cm *gen.CurvePropertyType, inheritDim uint, resolver *curveResolver) (core.LineString, error) {
	if cm.NilReason != nil {
		return nil, nil
	}
	if cm.Curve != nil {
		return lineStringFromCurve(cm.Curve, inheritDim)
	}
	if cm.LineString != nil {
		return lineStringFromXML(cm.LineString)
	}
	if cm.CompositeCurve != nil {
		return lineStringFromCompositeCurveType(cm.CompositeCurve, inheritDim, resolver)
	}
	if cm.OrientableCurve != nil {
		oc := cm.OrientableCurve
		if oc.BaseCurve != nil {
			if oc.BaseCurve.Curve != nil {
				return lineStringFromCurve(oc.BaseCurve.Curve, inheritDim)
			}
			if oc.BaseCurve.Href != "" {
				if c := resolver.resolve(strings.TrimPrefix(oc.BaseCurve.Href, "#")); c != nil {
					return lineStringFromCurve(c, inheritDim)
				}
			}
		}
	}
	if cm.LinearRing != nil {
		return lineStringFromLinearRingType(cm.LinearRing, inheritDim)
	}
	if cm.Ring != nil {
		return lineStringFromRingType(cm.Ring, inheritDim, resolver)
	}
	if cm.Href != "" {
		id := strings.TrimPrefix(cm.Href, "#")
		if c := resolver.resolve(id); c != nil {
			return lineStringFromCurve(c, inheritDim)
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
func lineStringFromLinearRingType(x *gen.LinearRingType, inheritDim uint) (core.LineString, error) {
	dim := preferDim(inheritDim, derefDim(x.SrsDimension))
	if x.PosList != nil {
		return core.LineStringFromPosListString(x.PosList.Value, preferDim(dim, derefDim(x.PosList.SrsDimension)))
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
		d := preferDim(dim, derefDim(x.Pos[0].SrsDimension))
		if d == 0 {
			d = uint(len(strings.Fields(x.Pos[0].Value)))
			if d < 2 {
				d = 2
			}
		}
		return core.LineStringFromFlat(flat, d)
	}
	if x.Coordinates != nil {
		coords, err := core.ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return nil, err
		}
		return core.LineStringFromFlat(coords, 2)
	}
	return nil, fmt.Errorf("gml: LinearRing has no coordinate data")
}

// lineStringFromRingType concatenates curveMember segments of a RingType into a LineString.
func lineStringFromRingType(x *gen.RingType, inheritDim uint, resolver *curveResolver) (core.LineString, error) {
	var result core.LineString
	dim := preferDim(inheritDim, derefDim(x.SrsDimension))
	for i := range x.CurveMember {
		ls, err := lineStringFromCurveProperty(&x.CurveMember[i], dim, resolver)
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

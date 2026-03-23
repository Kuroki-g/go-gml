package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_1_1/generated"
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

func lineStringFromCompositeCurveType(x *gen.CompositeCurveType, inheritDim int, resolver *curveResolver) (core.LineString, error) {
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
func lineStringFromCurveProperty(cm *gen.CurvePropertyType, inheritDim int, resolver *curveResolver) (core.LineString, error) {
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
	if cm.Href != "" {
		id := strings.TrimPrefix(cm.Href, "#")
		if c := resolver.resolve(id); c != nil {
			return lineStringFromCurve(c, inheritDim)
		}
		if ls, ok := resolver.lineStringByID[id]; ok {
			return ls, nil
		}
	}
	return nil, nil
}

package gml

import (
	"encoding/xml"
	"fmt"
	"strings"

	v3 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// handleCompositeCurve decodes a gml:CompositeCurve and returns a LineString.
func (r *Reader) handleCompositeCurve(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3.CompositeCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: CompositeCurve: %w", err)
	}
	ls, err := lineStringFromCompositeCurveType(&x, 0, r.resolver)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: ls, SRSName: x.SrsName}, nil
}

func lineStringFromCompositeCurveType(x *v3.CompositeCurveType, inheritDim int, resolver *curveResolver) (LineString, error) {
	var result LineString
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
// Handles Curve, LineString, OrientableCurve, CompositeCurve, and xlink:href references.
func lineStringFromCurveProperty(cm *v3.CurvePropertyType, inheritDim int, resolver *curveResolver) (LineString, error) {
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
		if c := resolver.resolve(strings.TrimPrefix(cm.Href, "#")); c != nil {
			return lineStringFromCurve(c, inheritDim)
		}
	}
	return nil, nil
}

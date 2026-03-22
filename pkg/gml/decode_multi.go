package gml

import (
	"encoding/xml"
	"fmt"

	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// ---- multi-point ----

func decodeMultiPointElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3_2_1.MultiPointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: MultiPoint: %w", err)
	}
	var pts MultiPoint
	for _, m := range x.PointMember {
		if m.Point == nil {
			continue
		}
		if m.Point.SrsDimension == nil {
			m.Point.SrsDimension = x.SrsDimension
		}
		pt, err := pointFromXML(m.Point)
		if err != nil {
			return Geometry{}, err
		}
		pts = append(pts, pt)
	}
	return Geometry{Value: pts, SRSName: x.SrsName}, nil
}

// ---- multi-curve / multi-linestring ----

func decodeMultiCurveElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		return decodeMultiLineStringV2(dec, se)
	}
	var x v3_2_1.MultiCurveType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	var lines MultiLineString
	for _, m := range x.CurveMember {
		if m.LineString == nil {
			continue
		}
		if m.LineString.SrsDimension == nil {
			m.LineString.SrsDimension = x.SrsDimension
		}
		ls, err := lineStringFromXML(m.LineString)
		if err != nil {
			return Geometry{}, err
		}
		lines = append(lines, ls)
	}
	return Geometry{Value: lines, SRSName: x.SrsName}, nil
}

// ---- multi-surface / multi-polygon ----

func decodeMultiSurfaceElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		return decodeMultiPolygonV2(dec, se)
	}
	var x v3_2_1.MultiSurfaceType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	var polys MultiPolygon
	for _, m := range x.SurfaceMember {
		if m.Polygon == nil {
			continue
		}
		if m.Polygon.SrsDimension == nil {
			m.Polygon.SrsDimension = x.SrsDimension
		}
		p, err := polygonFromXML(m.Polygon)
		if err != nil {
			return Geometry{}, err
		}
		polys = append(polys, p)
	}
	return Geometry{Value: polys, SRSName: x.SrsName}, nil
}

// ---- envelope ----

func decodeEnvelopeElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v3_2_1.EnvelopeType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Envelope: %w", err)
	}
	b, err := boundFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: b, SRSName: x.SrsName}, nil
}

func boundFromXML(x *v3_2_1.EnvelopeType) (Bound, error) {
	dim := derefDim(x.SrsDimension)
	if x.LowerCorner != nil && x.UpperCorner != nil {
		d := preferDim(dim, derefDim(x.LowerCorner.SrsDimension))
		lo, err := PointFromPosString(x.LowerCorner.Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope lowerCorner: %w", err)
		}
		hi, err := PointFromPosString(x.UpperCorner.Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope upperCorner: %w", err)
		}
		return Bound{Min: lo, Max: hi}, nil
	}
	if len(x.Pos) >= 2 {
		d := preferDim(dim, derefDim(x.Pos[0].SrsDimension))
		lo, err := PointFromPosString(x.Pos[0].Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope pos[0]: %w", err)
		}
		hi, err := PointFromPosString(x.Pos[1].Value, d)
		if err != nil {
			return Bound{}, fmt.Errorf("gml: Envelope pos[1]: %w", err)
		}
		return Bound{Min: lo, Max: hi}, nil
	}
	return Bound{}, fmt.Errorf("gml: Envelope has no coordinate data")
}

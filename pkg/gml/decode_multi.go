package gml

import (
	"encoding/xml"
	"fmt"
)

// ---- multi-point ----

type xmlPointMember struct {
	Point *xmlPoint `xml:"Point"`
}

type xmlMultiPoint struct {
	SrsName      string           `xml:"srsName,attr,omitempty"`
	SrsDimension int              `xml:"srsDimension,attr,omitempty"`
	PointMember  []xmlPointMember `xml:"pointMember"`
}

func decodeMultiPointElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x xmlMultiPoint
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: MultiPoint: %w", err)
	}
	var pts MultiPoint
	for _, m := range x.PointMember {
		if m.Point == nil {
			continue
		}
		if m.Point.SrsDimension == 0 {
			m.Point.SrsDimension = x.SrsDimension
		}
		pt, err := pointFromXML(m.Point)
		if err != nil {
			return Geometry{}, err
		}
		pts = append(pts, pt)
	}
	return Geometry{Value: pts, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

// ---- multi-curve / multi-linestring ----

type xmlCurveMember struct {
	LineString *xmlLineString `xml:"LineString"`
}

type xmlMultiCurve struct {
	SrsName          string           `xml:"srsName,attr,omitempty"`
	SrsDimension     int              `xml:"srsDimension,attr,omitempty"`
	CurveMember      []xmlCurveMember `xml:"curveMember"`      // GML 3.x
	LineStringMember []xmlCurveMember `xml:"lineStringMember"` // GML 2.x
}

func decodeMultiCurveElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x xmlMultiCurve
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	var lines MultiLineString
	for _, members := range [2][]xmlCurveMember{x.CurveMember, x.LineStringMember} {
		for _, m := range members {
			if m.LineString == nil {
				continue
			}
			if m.LineString.SrsDimension == 0 {
				m.LineString.SrsDimension = x.SrsDimension
			}
			ls, err := lineStringFromXML(m.LineString)
			if err != nil {
				return Geometry{}, err
			}
			lines = append(lines, ls)
		}
	}
	return Geometry{Value: lines, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

// ---- multi-surface / multi-polygon ----

type xmlSurfaceMember struct {
	Polygon *xmlPolygon `xml:"Polygon"`
}

type xmlMultiSurface struct {
	SrsName       string             `xml:"srsName,attr,omitempty"`
	SrsDimension  int                `xml:"srsDimension,attr,omitempty"`
	SurfaceMember []xmlSurfaceMember `xml:"surfaceMember"` // GML 3.x
	PolygonMember []xmlSurfaceMember `xml:"polygonMember"` // GML 2.x
}

func decodeMultiSurfaceElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x xmlMultiSurface
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: %s: %w", se.Name.Local, err)
	}
	var polys MultiPolygon
	for _, members := range [2][]xmlSurfaceMember{x.SurfaceMember, x.PolygonMember} {
		for _, m := range members {
			if m.Polygon == nil {
				continue
			}
			if m.Polygon.SrsDimension == 0 {
				m.Polygon.SrsDimension = x.SrsDimension
			}
			p, err := polygonFromXML(m.Polygon)
			if err != nil {
				return Geometry{}, err
			}
			polys = append(polys, p)
		}
	}
	return Geometry{Value: polys, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

// ---- envelope ----

type xmlEnvelope struct {
	SrsName      string         `xml:"srsName,attr,omitempty"`
	SrsDimension int            `xml:"srsDimension,attr,omitempty"`
	LowerCorner  *xmlDirectPos  `xml:"lowerCorner"`
	UpperCorner  *xmlDirectPos  `xml:"upperCorner"`
	Pos          []xmlDirectPos `xml:"pos"`
}

func decodeEnvelopeElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x xmlEnvelope
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Envelope: %w", err)
	}
	b, err := boundFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: b, EPSG: EPSGFromSRSName(x.SrsName)}, nil
}

func boundFromXML(x *xmlEnvelope) (Bound, error) {
	dim := x.SrsDimension
	if x.LowerCorner != nil && x.UpperCorner != nil {
		d := preferDim(dim, x.LowerCorner.SrsDimension)
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
		d := preferDim(dim, x.Pos[0].SrsDimension)
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

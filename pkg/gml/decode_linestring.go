package gml

import (
	"encoding/xml"
	"fmt"
	"strings"

	v3_1_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_1_1"
	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

func decodeLineStringElement(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	if se.Name.Space == gmlNS2 {
		var x v3_1_1.LineStringType
		if err := dec.DecodeElement(&x, &se); err != nil {
			return Geometry{}, fmt.Errorf("gml: LineString(ns2): %w", err)
		}
		ls, err := lineStringFromV31(&x)
		if err != nil {
			return Geometry{}, err
		}
		return Geometry{Value: ls, SRSName: x.SrsName}, nil
	}
	var x v3_2_1.LineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: LineString: %w", err)
	}
	ls, err := lineStringFromXML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: ls, SRSName: x.SrsName}, nil
}

func lineStringFromXML(x *v3_2_1.LineStringType) (LineString, error) {
	dim := derefDim(x.SrsDimension)
	if x.PosList != nil {
		return LineStringFromPosListString(x.PosList.Value, preferDim(dim, derefDim(x.PosList.SrsDimension)))
	}
	if len(x.Pos) > 0 {
		var flat []float64
		for _, p := range x.Pos {
			vals, err := ParsePosList(p.Value)
			if err != nil {
				return nil, err
			}
			flat = append(flat, vals...)
		}
		d := preferDim(dim, derefDim(x.Pos[0].SrsDimension))
		if d == 0 {
			d = len(strings.Fields(x.Pos[0].Value))
			if d < 2 {
				d = 2
			}
		}
		return LineStringFromFlat(flat, d)
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return nil, err
		}
		return LineStringFromFlat(coords, 2)
	}
	return nil, fmt.Errorf("gml: LineString has no coordinate data")
}

// lineStringFromV31 decodes a LineString from a v3_1_1.LineStringType, covering
// GML 3.1.1 primary fields and GML 2.x deprecated fields.
//
// Priority:
//  1. gml:posList     — GML 3.1.1 primary (compact list)
//  2. gml:pos (list)  — GML 3.1.1 alternative (individual positions)
//  3. gml:coordinates — GML 2.x / GML 3.1.1 deprecated
//  4. gml:coord (list) — GML 2.x deprecated (removed in GML 3.0)
func lineStringFromV31(x *v3_1_1.LineStringType) (LineString, error) {
	dim := derefDim(x.SrsDimension)
	if x.PosList != nil {
		return LineStringFromPosListString(x.PosList.Value, preferDim(dim, derefDim(x.PosList.SrsDimension)))
	}
	if len(x.Pos) > 0 {
		return lineStringFromPosSlice(x.Pos, dim)
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return nil, err
		}
		return LineStringFromFlat(coords, 2)
	}
	if len(x.Coord) > 0 {
		return lineStringFromCoordSlice(x.Coord)
	}
	return nil, fmt.Errorf("gml: LineString(ns2) has no coordinate data")
}

// lineStringFromPosSlice builds a LineString from a gml:pos slice (GML 3.1.1).
func lineStringFromPosSlice(pos []v3_1_1.DirectPositionType, inheritDim int) (LineString, error) {
	flat, d, err := flatFromPosSlice(pos, inheritDim)
	if err != nil {
		return nil, err
	}
	return LineStringFromFlat(flat, d)
}

// flatFromPosSlice converts a gml:pos slice to a flat float64 slice.
// Dimension is resolved as: inheritDim → pos[0].srsDimension → token count.
func flatFromPosSlice(pos []v3_1_1.DirectPositionType, inheritDim int) ([]float64, int, error) {
	var flat []float64
	for _, p := range pos {
		vals, err := ParsePosList(p.Value)
		if err != nil {
			return nil, 0, err
		}
		flat = append(flat, vals...)
	}
	d := preferDim(inheritDim, derefDim(pos[0].SrsDimension))
	if d == 0 {
		d = len(strings.Fields(pos[0].Value))
		if d < 2 {
			d = 2
		}
	}
	return flat, d, nil
}

// lineStringFromCoordSlice builds a LineString from a gml:coord slice (GML 2.x).
func lineStringFromCoordSlice(coords []v3_1_1.CoordType) (LineString, error) {
	return LineStringFromFlat(coordSliceToFlat(coords), 2)
}

// coordSliceToFlat converts a gml:coord slice (X/Y) to a flat float64 slice.
func coordSliceToFlat(coords []v3_1_1.CoordType) []float64 {
	flat := make([]float64, 0, len(coords)*2)
	for _, c := range coords {
		flat = append(flat, c.X, c.Y)
	}
	return flat
}

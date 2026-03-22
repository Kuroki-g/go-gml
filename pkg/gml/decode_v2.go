package gml

// decode_v2.go — namespace "http://www.opengis.net/gml" 向けデコード
//
// # Namespace の共有と GML バージョン判別不能問題
//
// GML 2.1.2 と GML 3.1.1 はどちらも同一の namespace URI
// "http://www.opengis.net/gml" (gmlNS2) を使用する。
// XML の要素レベルではバージョンを区別できないため、このファイルの
// decode 関数は両バージョンを単一パスで処理する。
//
// # なぜ v3_1_1 struct を使うか
//
// v2_1_2.PointType 等の GML 2.x struct は座標を gml:coordinates / gml:coord
// でしか保持できず、GML 3.1.1 が primary とする gml:pos / gml:posList に
// 対応するフィールドを持たない。
//
// v3_1_1 struct は GML 3.1.1 XSD (geometryBasic0d1d.xsd, geometryBasic2d.xsd)
// から生成されており、deprecated として残された GML 2.x 互換フィールドも含む:
//
//	GML 3.1.1 primary    → GML 2.x / deprecated fallback
//	Pos / PosList        → Coordinates / Coord
//	Exterior / Interior  → OuterBoundaryIs / InnerBoundaryIs
//
// したがって v3_1_1 struct は gmlNS2 で受信し得るすべてのエンコーディングを
// 網羅するスーパーセットであり、GML 2.x ファイルも GML 3.1.1 ファイルも
// 同一コードパスで正しくデコードできる。

import (
	"encoding/xml"
	"fmt"
	"strings"

	v31 "github.com/Kuroki-g/go-gml/pkg/gml/v3_1_1"
)

func decodePointV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v31.PointType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: Point(ns2): %w", err)
	}
	pt, err := pointFromNS2XML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: pt, SRSName: x.SrsName}, nil
}

// pointFromNS2XML は GML 3.1.1 / GML 2.x 共通の座標 fallback 順で Point を構築する。
//
// 優先順位 (GML 3.1.1 XSD PointType の choice 定義に従う):
//  1. gml:pos        — GML 3.1.1 primary
//  2. gml:coordinates — GML 2.x / GML 3.1.1 deprecated
//  3. gml:coord      — GML 2.x deprecated (GML 3.0 で廃止)
func pointFromNS2XML(x *v31.PointType) (Point, error) {
	dim := preferDim(derefDim(x.SrsDimension), 0)
	if x.Pos != nil {
		return PointFromPosString(x.Pos.Value, preferDim(dim, derefDim(x.Pos.SrsDimension)))
	}
	if x.Coordinates != nil {
		coords, err := ParseCoordinates(x.Coordinates.Value, derefStrOr(x.Coordinates.Cs, ","), derefStrOr(x.Coordinates.Ts, " "))
		if err != nil {
			return Point{}, err
		}
		return PointFromFlat(coords, 2)
	}
	if x.Coord != nil {
		return Point{x.Coord.X, x.Coord.Y}, nil
	}
	return Point{}, fmt.Errorf("gml: Point(ns2) has no coordinate data")
}

func decodeLineStringV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v31.LineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: LineString(ns2): %w", err)
	}
	ls, err := lineStringFromNS2XML(&x)
	if err != nil {
		return Geometry{}, err
	}
	return Geometry{Value: ls, SRSName: x.SrsName}, nil
}

// lineStringFromNS2XML は GML 3.1.1 / GML 2.x 共通の座標 fallback 順で LineString を構築する。
//
// 優先順位 (GML 3.1.1 XSD LineStringType の choice 定義に従う):
//  1. gml:posList     — GML 3.1.1 primary (compact list)
//  2. gml:pos (list)  — GML 3.1.1 alternative (individual positions)
//  3. gml:coordinates — GML 2.x / GML 3.1.1 deprecated
//  4. gml:coord (list) — GML 2.x deprecated (GML 3.0 で廃止)
func lineStringFromNS2XML(x *v31.LineStringType) (LineString, error) {
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

func decodeMultiLineStringV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v31.MultiLineStringType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: MultiLineString(ns2): %w", err)
	}
	var lines MultiLineString
	for _, m := range x.LineStringMember {
		if m.LineString == nil {
			continue
		}
		if m.LineString.SrsDimension == nil {
			m.LineString.SrsDimension = x.SrsDimension
		}
		ls, err := lineStringFromNS2XML(m.LineString)
		if err != nil {
			return Geometry{}, err
		}
		lines = append(lines, ls)
	}
	return Geometry{Value: lines, SRSName: x.SrsName}, nil
}

func decodeMultiPolygonV2(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	var x v31.MultiPolygonType
	if err := dec.DecodeElement(&x, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: MultiPolygon(ns2): %w", err)
	}
	var polys MultiPolygon
	for _, m := range x.PolygonMember {
		if m.Polygon == nil {
			continue
		}
		if m.Polygon.SrsDimension == nil {
			m.Polygon.SrsDimension = x.SrsDimension
		}
		poly, err := polygonFromNS2XML(m.Polygon)
		if err != nil {
			return Geometry{}, err
		}
		polys = append(polys, poly)
	}
	return Geometry{Value: polys, SRSName: x.SrsName}, nil
}

// polygonFromNS2XML は GML 3.1.1 / GML 2.x 共通の境界 fallback 順で Polygon を構築する。
//
// 優先順位 (GML 3.1.1 XSD PolygonType の定義に従う):
//  1. gml:exterior / gml:interior  — GML 3.1.1 primary
//  2. gml:outerBoundaryIs / gml:innerBoundaryIs — GML 2.x / GML 3.0 deprecated
func polygonFromNS2XML(x *v31.PolygonType) (Polygon, error) {
	dim := derefDim(x.SrsDimension)
	var rings []Ring

	// exterior ring: GML 3.1.1 → GML 2.x fallback
	extRing := x.Exterior
	if extRing == nil {
		extRing = x.OuterBoundaryIs // GML 2.x: outerBoundaryIs
	}
	if extRing != nil && extRing.LinearRing != nil {
		r, err := ringFromNS2LinearRing(extRing.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon(ns2) exterior: %w", err)
		}
		rings = append(rings, r)
	}

	// interior rings: GML 3.1.1 → GML 2.x fallback
	intRings := x.Interior
	if len(intRings) == 0 {
		intRings = x.InnerBoundaryIs // GML 2.x: innerBoundaryIs
	}
	for i, ir := range intRings {
		if ir.LinearRing == nil {
			continue
		}
		r, err := ringFromNS2LinearRing(ir.LinearRing, dim)
		if err != nil {
			return nil, fmt.Errorf("gml: Polygon(ns2) interior[%d]: %w", i, err)
		}
		rings = append(rings, r)
	}
	return Polygon(rings), nil
}

// ringFromNS2LinearRing は GML 3.1.1 / GML 2.x 共通の座標 fallback 順で Ring を構築する。
//
// 優先順位 (GML 3.1.1 XSD LinearRingType の choice 定義に従う):
//  1. gml:posList     — GML 3.1.1 primary (compact list)
//  2. gml:pos (list)  — GML 3.1.1 alternative (individual positions)
//  3. gml:coordinates — GML 2.x / GML 3.1.1 deprecated
//  4. gml:coord (list) — GML 2.x deprecated (GML 3.0 で廃止)
func ringFromNS2LinearRing(lr *v31.LinearRingType, inheritDim int) (Ring, error) {
	if lr == nil {
		return nil, fmt.Errorf("gml: nil LinearRing(ns2)")
	}
	dim := preferDim(inheritDim, derefDim(lr.SrsDimension))
	if lr.PosList != nil {
		return RingFromPosListString(lr.PosList.Value, preferDim(dim, derefDim(lr.PosList.SrsDimension)))
	}
	if len(lr.Pos) > 0 {
		flat, d, err := flatFromPosSlice(lr.Pos, dim)
		if err != nil {
			return nil, err
		}
		return RingFromFlat(flat, d)
	}
	if lr.Coordinates != nil {
		return RingFromCoordinatesString(lr.Coordinates.Value, derefStrOr(lr.Coordinates.Cs, ","), derefStrOr(lr.Coordinates.Ts, " "))
	}
	if len(lr.Coord) > 0 {
		flat := coordSliceToFlat(lr.Coord)
		return RingFromFlat(flat, 2)
	}
	return nil, fmt.Errorf("gml: LinearRing(ns2) has no coordinate data")
}

// ---- 内部ヘルパー ----

// lineStringFromPosSlice は gml:pos 列から LineString を構築する (GML 3.1.1)。
func lineStringFromPosSlice(pos []v31.DirectPositionType, inheritDim int) (LineString, error) {
	flat, d, err := flatFromPosSlice(pos, inheritDim)
	if err != nil {
		return nil, err
	}
	return LineStringFromFlat(flat, d)
}

// flatFromPosSlice は gml:pos 列を flat float64 スライスに変換する。
// 次元数は inheritDim → pos[0].srsDimension → トークン数の順で決定する。
func flatFromPosSlice(pos []v31.DirectPositionType, inheritDim int) ([]float64, int, error) {
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

// lineStringFromCoordSlice は gml:coord 列から LineString を構築する (GML 2.x)。
func lineStringFromCoordSlice(coords []v31.CoordType) (LineString, error) {
	return LineStringFromFlat(coordSliceToFlat(coords), 2)
}

// coordSliceToFlat は gml:coord 列 (X/Y) を flat float64 スライスに変換する。
// Z が非ゼロの場合は 3D として扱う (GML 2.x coord の慣習)。
func coordSliceToFlat(coords []v31.CoordType) []float64 {
	flat := make([]float64, 0, len(coords)*2)
	for _, c := range coords {
		flat = append(flat, c.X, c.Y)
	}
	return flat
}

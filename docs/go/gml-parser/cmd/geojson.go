package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/Kuroki-g/go-gml/gml"
)

// ---------- GeoJSON types ----------

type featureCollection struct {
	Type     string    `json:"type"`
	Features []feature `json:"features"`
}

type feature struct {
	Type       string         `json:"type"`
	Geometry   geojsonGeom    `json:"geometry"`
	Properties map[string]any `json:"properties"`
}

type geojsonGeom struct {
	Type        string `json:"type"`
	Coordinates any    `json:"coordinates"`
}

// ---------- converters ----------

func toGeoJSON(g gml.Geometry, swap bool) (geojsonGeom, error) {
	switch v := g.Value.(type) {
	case gml.Point:
		return geojsonGeom{Type: "Point", Coordinates: pt(v, swap)}, nil
	case gml.LineString:
		return geojsonGeom{Type: "LineString", Coordinates: lineCoords(v, swap)}, nil
	case gml.Polygon:
		return geojsonGeom{Type: "Polygon", Coordinates: polyCoords(v, swap)}, nil
	case gml.MultiPoint:
		coords := make([][]float64, len(v))
		for i, p := range v {
			coords[i] = pt(p, swap)
		}
		return geojsonGeom{Type: "MultiPoint", Coordinates: coords}, nil
	case gml.MultiLineString:
		coords := make([][][]float64, len(v))
		for i, ls := range v {
			coords[i] = lineCoords(ls, swap)
		}
		return geojsonGeom{Type: "MultiLineString", Coordinates: coords}, nil
	case gml.MultiPolygon:
		coords := make([][][][]float64, len(v))
		for i, poly := range v {
			coords[i] = polyCoords(poly, swap)
		}
		return geojsonGeom{Type: "MultiPolygon", Coordinates: coords}, nil
	case gml.Bound:
		// Envelope → closed bounding-box Polygon
		min, max := pt(v.Min, swap), pt(v.Max, swap)
		ring := [][]float64{min, {max[0], min[1]}, max, {min[0], max[1]}, min}
		return geojsonGeom{Type: "Polygon", Coordinates: [][][]float64{ring}}, nil
	default:
		return geojsonGeom{}, fmt.Errorf("gml-parser: unknown geometry type %T", g.Value)
	}
}

func geomTypeName(v interface{}) string {
	switch v.(type) {
	case gml.Point:
		return "Point"
	case gml.LineString:
		return "LineString"
	case gml.Polygon:
		return "Polygon"
	case gml.MultiPoint:
		return "MultiPoint"
	case gml.MultiLineString:
		return "MultiLineString"
	case gml.MultiPolygon:
		return "MultiPolygon"
	case gml.Bound:
		return "Bound"
	case gml.GridCoverage:
		return "GridCoverage"
	default:
		return fmt.Sprintf("unknown(%T)", v)
	}
}

// ---------- coordinate helpers ----------

func pt(p gml.Point, swap bool) []float64 {
	if swap {
		return []float64{p[1], p[0]}
	}
	return []float64{p[0], p[1]}
}

func lineCoords(ls gml.LineString, swap bool) [][]float64 {
	c := make([][]float64, len(ls))
	for i, p := range ls {
		c[i] = pt(p, swap)
	}
	return c
}

func polyCoords(poly gml.Polygon, swap bool) [][][]float64 {
	c := make([][][]float64, len(poly))
	for i, ring := range poly {
		c[i] = make([][]float64, len(ring))
		for j, p := range ring {
			c[i][j] = pt(p, swap)
		}
	}
	return c
}

// ---------- IO ----------

// openInput returns an io.ReadSeeker for the given file path.
// If path is empty, stdin is read entirely into memory (stdin is not seekable).
func openInput(path string) (io.ReadSeeker, func(), error) {
	if path == "" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, nil, err
		}
		return bytes.NewReader(b), func() {}, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return f, func() { f.Close() }, nil
}

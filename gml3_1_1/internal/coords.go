package internal

import (
	"fmt"
	"strconv"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
)

// ParsePosList parses a GML posList or pos value.
// Both use the same format: whitespace-separated flat sequence of numbers.
func ParsePosList(s string) ([]float64, error) {
	fields := strings.Fields(s)
	result := make([]float64, 0, len(fields))
	for _, f := range fields {
		v, err := strconv.ParseFloat(f, 64)
		if err != nil {
			return nil, fmt.Errorf("gml: invalid coordinate %q: %w", f, err)
		}
		result = append(result, v)
	}
	return result, nil
}

// ParseCoordinates parses the deprecated gml:coordinates format.
// cs is the coordinate separator (default ","), ts is the tuple separator (default " ").
func ParseCoordinates(s, cs, ts string) ([]float64, error) {
	if cs == "" {
		cs = ","
	}
	if ts == "" {
		ts = " "
	}
	s = strings.TrimSpace(s)
	var result []float64
	for _, tuple := range strings.Split(s, ts) {
		tuple = strings.TrimSpace(tuple)
		if tuple == "" {
			continue
		}
		for _, coord := range strings.Split(tuple, cs) {
			coord = strings.TrimSpace(coord)
			if coord == "" {
				continue
			}
			v, err := strconv.ParseFloat(coord, 64)
			if err != nil {
				return nil, fmt.Errorf("gml: invalid coordinate %q: %w", coord, err)
			}
			result = append(result, v)
		}
	}
	return result, nil
}

// ToPoints converts a flat coordinate slice into a slice of Points.
// dim specifies the coordinate dimension (2 or 3). If dim <= 0 it defaults to 2.
func ToPoints(coords []float64, dim int) ([]core.Point, error) {
	if dim <= 0 {
		dim = 2
	}
	if dim < 2 {
		return nil, fmt.Errorf("gml: dimension must be >= 2, got %d", dim)
	}
	if len(coords) == 0 {
		return nil, nil
	}
	if len(coords)%dim != 0 {
		return nil, fmt.Errorf("gml: coordinate count %d is not a multiple of dimension %d", len(coords), dim)
	}
	pts := make([]core.Point, len(coords)/dim)
	for i := range pts {
		pts[i] = append(core.Point(nil), coords[i*dim:(i+1)*dim]...)
	}
	return pts, nil
}

package gml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	v3_2_1 "github.com/Kuroki-g/go-gml/pkg/gml/v3_2_1"
)

// gridBounds stores the index range of a gml:Grid or gml:RectifiedGrid element.
// Low and High are derived from gml:GridEnvelope/gml:low and gml:high,
// using the length of the index list as the axis count (dimension attribute is unreliable).
type gridBounds struct {
	low  []int
	high []int
}

// handleDomainSet decodes a gml:domainSet element and stores the resulting
// gridBounds in r.pendingGrid for consumption by handleRangeSet.
// Handles both inline Grid/RectifiedGrid and xlink:href references.
func (r *Reader) handleDomainSet(dec *xml.Decoder, se xml.StartElement) error {
	var ds v3_2_1.DomainSetType
	if err := dec.DecodeElement(&ds, &se); err != nil {
		return fmt.Errorf("gml: domainSet: %w", err)
	}
	if ds.Href != "" {
		id := strings.TrimPrefix(ds.Href, "#")
		r.pendingGrid = r.resolver.gridByID[id]
		return nil
	}
	if ds.Grid != nil {
		gb, err := gridBoundsFromGridLimits(ds.Grid.Limits)
		if err == nil {
			r.pendingGrid = gb
		}
		return nil
	}
	if ds.RectifiedGrid != nil {
		gb, err := gridBoundsFromGridLimits(ds.RectifiedGrid.Limits)
		if err == nil {
			r.pendingGrid = gb
		}
		return nil
	}
	return nil
}

// handleRangeSet decodes a gml:rangeSet element and returns a GridCoverage.
// r.pendingGrid must be set by a prior handleDomainSet call.
func (r *Reader) handleRangeSet(dec *xml.Decoder, se xml.StartElement) (Geometry, error) {
	grid := r.pendingGrid
	r.pendingGrid = nil
	var rs v3_2_1.RangeSetType
	if err := dec.DecodeElement(&rs, &se); err != nil {
		return Geometry{}, fmt.Errorf("gml: rangeSet: %w", err)
	}
	var tuples string
	if rs.DataBlock != nil && rs.DataBlock.TupleList != nil {
		tuples = strings.TrimSpace(rs.DataBlock.TupleList.Value)
	}
	return Geometry{Value: GridCoverage{
		Low:    grid.low,
		High:   grid.high,
		Tuples: tuples,
	}}, nil
}

// gridBoundsFromGridLimits extracts index bounds from a gml:GridLimitsType.
func gridBoundsFromGridLimits(limits *v3_2_1.GridLimitsType) (*gridBounds, error) {
	if limits == nil || limits.GridEnvelope == nil {
		return nil, fmt.Errorf("Grid has no limits")
	}
	env := limits.GridEnvelope
	low, err := parseIntList(env.Low)
	if err != nil {
		return nil, fmt.Errorf("Grid low: %w", err)
	}
	high, err := parseIntList(env.High)
	if err != nil {
		return nil, fmt.Errorf("Grid high: %w", err)
	}
	return &gridBounds{low: low, high: high}, nil
}

// parseIntList parses a space-separated list of integers (e.g. "0 0" or "9 9").
func parseIntList(s string) ([]int, error) {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return nil, fmt.Errorf("empty int list")
	}
	result := make([]int, len(fields))
	for i, f := range fields {
		v, err := strconv.Atoi(f)
		if err != nil {
			return nil, fmt.Errorf("invalid int %q: %w", f, err)
		}
		result[i] = v
	}
	return result, nil
}

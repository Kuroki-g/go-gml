package internal

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	core "github.com/Kuroki-g/go-gml/core"
	gen "github.com/Kuroki-g/go-gml/gml3_2_1/generated"
)

// gridBounds stores the index range of a gml:Grid or gml:RectifiedGrid element.
type gridBounds struct {
	low  []int
	high []int
}

// handleDomainSet decodes a gml:domainSet element and stores the resulting
// gridBounds in r.pendingGrid for consumption by handleRangeSet.
func (r *Reader) handleDomainSet(dec *xml.Decoder, se xml.StartElement) error {
	var ds gen.DomainSetType
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
func (r *Reader) handleRangeSet(dec *xml.Decoder, se xml.StartElement) (core.Geometry, error) {
	grid := r.pendingGrid
	r.pendingGrid = nil
	var rs gen.RangeSetType
	if err := dec.DecodeElement(&rs, &se); err != nil {
		return core.Geometry{}, fmt.Errorf("gml: rangeSet: %w", err)
	}
	var tuples string
	if rs.DataBlock != nil && rs.DataBlock.TupleList != nil {
		tuples = strings.TrimSpace(rs.DataBlock.TupleList.Value)
	}
	return core.Geometry{Value: core.GridCoverage{
		Low:    grid.low,
		High:   grid.high,
		Tuples: tuples,
	}}, nil
}

// gridBoundsFromGridLimits extracts index bounds from a gml:GridLimitsType.
func gridBoundsFromGridLimits(limits *gen.GridLimitsType) (*gridBounds, error) {
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

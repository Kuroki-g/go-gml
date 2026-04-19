package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var swap bool

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a GML file to GeoJSON",
	Long: `Read GML geometry elements and write a GeoJSON FeatureCollection to stdout.
Non-geometry XML is silently skipped.

Coordinate axis order:
  GML with srsName in urn:ogc:def:crs:EPSG:: form (e.g. 国土数値情報) stores
  coordinates as lat/lon. GeoJSON requires lon/lat. Use --swap in that case.`,
	Example: `  gml-parser convert -i data.gml
  gml-parser convert -i data.gml --swap > out.geojson
  cat data.gml | gml-parser convert --swap`,
	RunE: func(cmd *cobra.Command, args []string) error {
		r, closer, err := openInput(inFile)
		if err != nil {
			return err
		}
		defer closer()

		reader, err := newGMLReader(r, gmlVersion, nil)
		if err != nil {
			return err
		}
		fc := &featureCollection{Type: "FeatureCollection", Features: []feature{}}
		skipped := map[string]int{}

		for {
			g, err := reader.Next()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return err
			}

			geom, err := toGeoJSON(g, swap)
			if err != nil {
				skipped[geomTypeName(g.Value)]++
				continue
			}

			fc.Features = append(fc.Features, feature{
				Type:       "Feature",
				Geometry:   geom,
				Properties: map[string]any{},
			})
		}

		for typeName, count := range skipped {
			fmt.Fprintf(os.Stderr, "warning: %d %s element(s) skipped (not representable in GeoJSON)\n", count, typeName)
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(fc)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().BoolVarP(&swap, "swap", "s", false, "swap coordinate axes (lat/lon → lon/lat)")
}

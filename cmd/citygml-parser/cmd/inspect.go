package cmd

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Kuroki-g/go-gml/citygml2_0"
	"github.com/Kuroki-g/go-gml/gml3_1_1"
	"github.com/spf13/cobra"
)

type inspectStats struct {
	File           string         `json:"file"`
	TotalBuildings int            `json:"total_buildings"`
	LodCounts      map[string]int `json:"lod_counts"`
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Show building statistics for a CityGML 2.0 file",
	Long:  `Stream through a CityGML 2.0 file and report the count of buildings and LoD geometry types.`,
	Example: `  citygml-parser inspect -i data.gml
  cat data.gml | citygml-parser inspect`,
	RunE: func(cmd *cobra.Command, args []string) error {
		r, closer, err := openInput(inFile)
		if err != nil {
			return err
		}
		defer closer()

		gmlDec := gml3_1_1.NewReader(r)
		reader := citygml2_0.NewReader(r, gmlDec)

		stats := &inspectStats{
			File:      inFile,
			LodCounts: map[string]int{},
		}

		for {
			b, err := reader.Next()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return err
			}

			stats.TotalBuildings++
			if b.Lod0FootPrint != nil {
				stats.LodCounts["lod0FootPrint"]++
			}
			if b.Lod0RoofEdge != nil {
				stats.LodCounts["lod0RoofEdge"]++
			}
			if b.Lod1Solid != nil {
				stats.LodCounts["lod1Solid"]++
			}
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(stats)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}

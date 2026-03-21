package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type inspectStats struct {
	File          string         `json:"file"`
	TotalFeatures int            `json:"total_features"`
	GeometryTypes map[string]int `json:"geometry_types"`
	EPSGCodes     map[string]int `json:"epsg_codes"`
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Show geometry statistics for a GML file",
	Long:  `Stream through a GML file and report the count of each geometry type and EPSG code seen.`,
	Example: `  gml-parser inspect -i data.gml
  cat data.gml | gml-parser inspect`,
	RunE: func(cmd *cobra.Command, args []string) error {
		r, closer, err := openInput(inFile)
		if err != nil {
			return err
		}
		defer closer()

		stats := &inspectStats{
			File:          inFile,
			GeometryTypes: map[string]int{},
			EPSGCodes:     map[string]int{},
		}

		reader := newGMLReader(r)
		for {
			g, err := reader.Next()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return err
			}

			stats.TotalFeatures++
			stats.GeometryTypes[geomTypeName(g.Value)]++

			epsgKey := fmt.Sprintf("EPSG:%d", g.EPSG)
			if g.EPSG == 0 {
				epsgKey = "unknown"
			}
			stats.EPSGCodes[epsgKey]++
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(stats)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}

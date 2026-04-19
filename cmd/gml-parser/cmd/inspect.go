package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Kuroki-g/go-gml/gml"
	"github.com/spf13/cobra"
)

type inspectStats struct {
	File            string         `json:"file"`
	TotalFeatures   int            `json:"total_features"`
	GeometryTypes   map[string]int `json:"geometry_types"`
	EPSGCodes       map[string]int `json:"epsg_codes"`
	SRSNames        map[string]int `json:"srs_names"`
	UnknownElements map[string]int `json:"unknown_elements,omitempty"`
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
			File:            inFile,
			GeometryTypes:   map[string]int{},
			EPSGCodes:       map[string]int{},
			SRSNames:        map[string]int{},
			UnknownElements: map[string]int{},
		}

		reader, err := newGMLReader(r, gmlVersion, func(name string) {
			stats.UnknownElements[name]++
		})
		if err != nil {
			return err
		}
		var parseErr error
		for {
			g, err := reader.Next()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				parseErr = err
				break
			}

			stats.TotalFeatures++
			stats.GeometryTypes[geomTypeName(g.Value)]++

			srsName := ""
			if g.SRSName != nil {
				srsName = *g.SRSName
			}
			epsg := gml.EPSGFromSRSName(srsName)
			epsgKey := fmt.Sprintf("EPSG:%d", epsg)
			if epsg == 0 {
				epsgKey = "unknown"
			}
			stats.EPSGCodes[epsgKey]++

			srsKey := srsName
			if srsKey == "" {
				srsKey = "(none)"
			}
			stats.SRSNames[srsKey]++
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(stats); err != nil {
			return err
		}
		return parseErr
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}

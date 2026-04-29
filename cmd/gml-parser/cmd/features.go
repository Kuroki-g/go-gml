package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Kuroki-g/go-gml/gml"
	"github.com/spf13/cobra"
)

type featuresStats struct {
	File          string         `json:"file"`
	TotalFeatures int            `json:"total_features"`
	ElementTypes  map[string]int `json:"element_types"`
	WithID        int            `json:"with_id"`
	WithBoundedBy int            `json:"with_bounded_by"`
}

var featuresCmd = &cobra.Command{
	Use:   "features",
	Short: "Show feature statistics for a GML 3.2.1 file",
	Long:  `Stream through a GML 3.2.1 file and report counts per feature element type, IDs, and boundedBy.`,
	Example: `  gml-parser features -i data.gml
  cat data.gml | gml-parser features`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gmlVersion != "3.2.1" {
			return fmt.Errorf("features command only supports GML 3.2.1 (got %s)", gmlVersion)
		}
		r, closer, err := openInput(inFile)
		if err != nil {
			return err
		}
		defer closer()

		reader := gml.NewReader321(r)
		reader.SetCharsetReader(charsetReader)

		stats := &featuresStats{
			File:         inFile,
			ElementTypes: map[string]int{},
		}

		for f, err := range reader.Features() {
			if err != nil {
				return err
			}
			stats.TotalFeatures++

			key := f.ElementName.Local
			if f.ElementName.Space != "" {
				key = f.ElementName.Space + ":" + f.ElementName.Local
			}
			stats.ElementTypes[key]++

			if f.ID != nil {
				stats.WithID++
			}
			if f.BoundedBy != nil {
				stats.WithBoundedBy++
			}
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(stats)
	},
}

func init() {
	rootCmd.AddCommand(featuresCmd)
}

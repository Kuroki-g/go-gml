package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Persistent flags shared by all subcommands.
var inFile string
var gmlVersion string

var rootCmd = &cobra.Command{
	Use:   "gml-parser",
	Short: "Parse GML files using the go-gml SDK",
}

// Execute is called by main.main().
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inFile, "in", "i", "", "input GML file (default: stdin)")
	rootCmd.PersistentFlags().StringVar(&gmlVersion, "gml-version", "3.2.1", "GML version to parse (3.2.1, 3.1.1, 2.1.2)")
}

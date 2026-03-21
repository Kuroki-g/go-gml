package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// inFile is a persistent flag shared by all subcommands.
var inFile string

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
}

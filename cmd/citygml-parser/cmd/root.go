package cmd

import (
	"bytes"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var inFile string

var rootCmd = &cobra.Command{
	Use:   "citygml-parser",
	Short: "Parse CityGML 2.0 files using the go-gml SDK",
}

// Execute is called by main.main().
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inFile, "in", "i", "", "input CityGML file (default: stdin)")
}

// openInput returns an io.ReadSeeker for the given file path.
// If path is empty, stdin is read entirely into memory (stdin is not seekable).
func openInput(path string) (io.ReadSeeker, func(), error) {
	if path == "" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, nil, err
		}
		return bytes.NewReader(b), func() {}, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return f, func() { f.Close() }, nil
}

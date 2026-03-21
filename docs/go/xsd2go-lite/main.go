package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	namespace    string
	pkgName      string
	outputFile   string
	skipAbstract bool
)

var rootCmd = &cobra.Command{
	Use:   "xsd2go-lite <input.xsd>",
	Short: "Generate Go structs (Unmarshal-only) from XSD schema",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(args[0])
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Target namespace URI to output (required)")
	rootCmd.Flags().StringVarP(&pkgName, "package", "p", "gml", "Go package name")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (default: stdout)")
	rootCmd.Flags().BoolVar(&skipAbstract, "skip-abstract", false, "Skip abstract types")
	_ = rootCmd.MarkFlagRequired("namespace")
}

func main() {
	Execute()
}

func run(inputXSD string) error {
	resolver := newResolver()
	if _, err := resolver.Load(inputXSD); err != nil {
		return fmt.Errorf("load schema: %w", err)
	}

	types := resolver.ResolveAll(namespace)
	if len(types) == 0 {
		fmt.Fprintf(os.Stderr, "warn: no types found for namespace %q\n", namespace)
	}

	src, err := Generate(types, pkgName, skipAbstract)
	if err != nil {
		// Print unformatted source for debugging, but still report error.
		fmt.Fprint(os.Stderr, src)
		return fmt.Errorf("generate: %w", err)
	}

	var out io.Writer = os.Stdout
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("create output: %w", err)
		}
		defer f.Close()
		out = f
	}

	_, err = fmt.Fprint(out, src)
	return err
}

// decodeXML is a helper used by resolve.go to decode an XSD file.
func decodeXML(r io.Reader, v any) error {
	return xml.NewDecoder(r).Decode(v)
}

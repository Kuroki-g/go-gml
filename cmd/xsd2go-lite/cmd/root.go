package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"xsd2go-lite/internal/gen"
	"xsd2go-lite/internal/resolve"
)

// Version is the tool version string, injected at build time via
// -ldflags "-X xsd2go-lite/cmd.Version=$(git describe --tags --always)".
var Version = "dev"

var (
	namespace    string
	pkgName      string
	outputFile   string
	skipAbstract bool
	withDoc      bool
	catalogPairs []string
	mapNSPairs   []string
)

var rootCmd = &cobra.Command{
	Use:   "xsd2go-lite <input.xsd>",
	Short: "Generate Go structs (Unmarshal-only) from XSD schema",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(args[0])
	},
}

// Execute is called by main.main().
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
	rootCmd.Flags().BoolVar(&withDoc, "with-doc", false, "Include XSD documentation as field comments")
	rootCmd.Flags().StringArrayVar(&catalogPairs, "catalog", nil, "Namespace-to-local-path mapping: ns=path (repeatable)")
	rootCmd.Flags().StringArrayVar(&mapNSPairs, "map-namespace", nil, "Map namespace to external Go package: ns=pkgPath (repeatable)")
	_ = rootCmd.MarkFlagRequired("namespace")
}

func run(inputXSD string) error {
	resolver := resolve.NewResolver()
	for _, pair := range catalogPairs {
		ns, path, ok := strings.Cut(pair, "=")
		if !ok {
			return fmt.Errorf("--catalog: invalid format %q, expected ns=path", pair)
		}
		resolver.AddCatalogEntry(ns, path)
	}
	if _, err := resolver.Load(inputXSD); err != nil {
		return fmt.Errorf("load schema: %w", err)
	}

	types := resolver.ResolveAll(namespace)
	if len(types) == 0 {
		fmt.Fprintf(os.Stderr, "warn: no types found for namespace %q\n", namespace)
	}

	mapNS := make(map[string]string, len(mapNSPairs))
	for _, pair := range mapNSPairs {
		ns, pkgPath, ok := strings.Cut(pair, "=")
		if !ok {
			return fmt.Errorf("--map-namespace: invalid format %q, expected ns=pkgPath", pair)
		}
		mapNS[ns] = pkgPath
	}

	src, err := gen.Generate(types, pkgName, skipAbstract, withDoc, mapNS, inputXSD, Version)
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

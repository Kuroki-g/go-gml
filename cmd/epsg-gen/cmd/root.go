package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	outputPath string
	force      bool
)

var rootCmd = &cobra.Command{
	Use:   "epsg-gen <code> [code...]",
	Short: "Fetch EPSG CRS dimension info and update epsg_dim_table.go",
	Args:  cobra.MinimumNArgs(1),
	RunE:  run,
}

func run(_ *cobra.Command, args []string) error {
	codes := make([]int, 0, len(args))
	for _, a := range args {
		code, err := strconv.Atoi(a)
		if err != nil {
			return fmt.Errorf("invalid EPSG code %q: %w", a, err)
		}
		codes = append(codes, code)
	}

	existing, err := readExisting(outputPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		existing = make(map[int]crsEntry)
	}

	for _, code := range codes {
		if _, ok := existing[code]; ok && !force {
			fmt.Fprintf(os.Stderr, "skip EPSG %d: already exists (use --force to overwrite)\n", code)
			continue
		}
		entry, err := fetchEntry(code)
		if err != nil {
			return fmt.Errorf("fetch EPSG %d: %w", code, err)
		}
		existing[code] = entry
		fmt.Printf("added EPSG %d: dim=%d (%s)\n", code, entry.Dim, entry.Name)
	}

	return writeTable(outputPath, existing)
}

// Execute is the entry point called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "core/epsg_dim_table.go", "output file path")
	rootCmd.Flags().BoolVar(&force, "force", false, "overwrite existing entries")
}

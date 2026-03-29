package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "check-coverage [module-dir...]",
	Short: "Detect unhandled PropertyType fields in GML parser modules",
	Long: `Parses generated/*.go and internal/*.go to find *PropertyType struct fields
that are never accessed in dispatch functions. Exits non-zero if any are found.

If no module directories are given, defaults to: gml2_1_2 gml3_1_1 gml3_2_1 citygml2_0`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(args)
	},
}

func init() {
	rootCmd.SilenceUsage = true
}

// Execute is called by main.main().
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(dirs []string) error {
	if len(dirs) == 0 {
		dirs = defaultModules()
	}

	var all []Finding
	for _, d := range dirs {
		found, err := checkModule(d)
		if err != nil {
			return fmt.Errorf("%s: %w", d, err)
		}
		all = append(all, found...)
	}

	sort.Slice(all, func(i, j int) bool {
		if all[i].Module != all[j].Module {
			return all[i].Module < all[j].Module
		}
		if all[i].Struct != all[j].Struct {
			return all[i].Struct < all[j].Struct
		}
		return all[i].Field < all[j].Field
	})

	if len(all) == 0 {
		fmt.Println("OK: no unhandled fields found")
		return nil
	}

	for _, f := range all {
		kind := "elem"
		if f.IsAttr {
			kind = "attr"
		}
		fmt.Printf("[UNHANDLED] %s / %s / %s.%s (%s:%s)\n",
			f.Module, f.Func, f.Struct, f.Field, kind, f.XMLName)
	}
	return fmt.Errorf("%d unhandled field(s) found", len(all))
}

func defaultModules() []string {
	candidates := []string{"gml2_1_2", "gml3_1_1", "gml3_2_1", "citygml2_0"}
	var result []string
	for _, m := range candidates {
		if _, err := os.Stat(m); err == nil {
			result = append(result, m)
		}
	}
	return result
}

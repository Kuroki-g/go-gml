package cmd

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of gml-parser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version())
	},
}

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(devel)"
	}

	// Tagged release: clean semver with no pseudo-version suffix.
	v := info.Main.Version
	if v != "" && v != "(devel)" && !strings.Contains(v, "-0.") {
		return v
	}

	// Fall back to VCS commit hash from build settings.
	var rev string
	var dirty bool
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			if len(s.Value) >= 7 {
				rev = s.Value[:7]
			} else {
				rev = s.Value
			}
		case "vcs.modified":
			dirty = s.Value == "true"
		}
	}
	if rev == "" {
		return "(devel)"
	}
	if dirty {
		return rev + "-dirty"
	}
	return rev
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

func buildVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}
	var rev, modified string
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			if len(s.Value) > 7 {
				rev = s.Value[:7]
			} else {
				rev = s.Value
			}
		case "vcs.modified":
			if s.Value == "true" {
				modified = "-dirty"
			}
		}
	}
	if rev == "" {
		return "dev"
	}
	return rev + modified
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of citygml-parser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

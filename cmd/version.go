package cmd

import (
	"fmt"

	"github.com/morikuni/aec"
	"github.com/openbrisk/brisk-cli/version"
	"github.com/spf13/cobra"
)

var (
	shortVersion bool
)

func init() {
	versionCmd.Flags().BoolVarP(&shortVersion, "short", "s", false, "just print the version")

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version [--short]",
	Short: "Prints the version number of the OpenBrisk CLI",
	Long:  "The version command returns the current version information for the OpenBrisk CLI.",
	Example: `  brisk version
  brisk version --short`,
	Run: runVersionCommand,
}

func runVersionCommand(cmd *cobra.Command, args []string) {
	if shortVersion {
		fmt.Printf("%s\n", version.Version[1:len(version.Version)])
	} else {
		printFiglet()
		fmt.Printf("Version: %s\n", version.Version[1:len(version.Version)])
		fmt.Printf("GitHash: %s\n", version.GitCommitHash)
	}
}

func printFiglet() {
	coloredFiglet := aec.GreenF.Apply(figlet)
	fmt.Printf(coloredFiglet)
}

const figlet = `  ___                   ____       _     _
 / _ \ _ __   ___ _ __ | __ ) _ __(_)___| | __
| | | | '_ \ / _ \ '_ \|  _ \| '__| / __| |/ /
| |_| | |_) |  __/ | | | |_) | |  | \__ \   <
 \___/| .__/ \___|_| |_|____/|_|  |_|___/_|\_\
      |_|

`

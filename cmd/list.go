package cmd

import "github.com/spf13/cobra"

var (
	listVerbose bool
)

func init() {
	listCmd.Flags().BoolVarP(&listVerbose, "verbose", "v", false, "Verbose output for function list")

	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     `list [--gateway GATEWAY_URL] [--verbose]`,
	Aliases: []string{"ls"},
	Short:   "Lists OpenBrisk functions",
	Long:    "Lists OpenBrisk functions",
	Example: `  brisk list
  brisk list --verbose`,
	RunE: runListCommand,
}

func runListCommand(cmd *cobra.Command, args []string) error {

	return nil

}

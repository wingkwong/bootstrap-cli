package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wingkwong/bootstrap-cli/internal/ui"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:     "launch",
	Short:   "launches text user interface view",
	Long:    "launches the TUI (text user interface) view for the application",
	Aliases: []string{"tui"},
	Run: func(cmd *cobra.Command, args []string) {

		ui.Execute()
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}

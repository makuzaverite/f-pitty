package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "fitty",
	Short: "Interract with filesystem on all oses with the same commands",
	Long: `Fitty is Cross Platform fileSystem utility cli
	Build with Productivity for productivity`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute initialized cli
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

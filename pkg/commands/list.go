package commands

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "ls",
	Short: "List all files in a directory",
	Long:  "Listing all files in the directory",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCommand.AddCommand(listCmd)
}

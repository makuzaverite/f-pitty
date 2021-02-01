package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addFile = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "touch",
	Short: "add file",
	Long:  "add new file in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		// newFile()
	},
}

func init() {
	rootCommand.AddCommand(addFile)
}

func newFile() {

}

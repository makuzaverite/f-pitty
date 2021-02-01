package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addFile = &cobra.Command{
	Use:   "touch",
	Short: "Add file",
	Args:  cobra.MaximumNArgs(25),
	Long:  "Add new file in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		newFile(args)
	},
}

func init() {
	rootCommand.AddCommand(addFile)
}

func newFile(fNames []string) {
	currDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Failed to create files")
	}

	fmt.Print(currDir)

	for _, fName := range fNames {
		fmt.Println(fName)
	}
}

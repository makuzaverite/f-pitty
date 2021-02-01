package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "ls",
	Short: "List all files in a directory",
	Long:  "Listing all files in the directory",
	Run: func(cmd *cobra.Command, args []string) {
		handleListingFiles()
	},
}

// listCmd.Flags().IntVarP(1,"list","")
func init() {
	rootCommand.AddCommand(listCmd)
}

func handleListingFiles() {
	currDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Failed to open this directory")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(currDir)

	for _, file := range files {
		fmt.Print(file.Name() + "  ")
	}
}

func isDir(file os.FileInfo) bool {
	return file.IsDir()
}

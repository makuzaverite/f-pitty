package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

func init() {
	rootCommand.AddCommand(listCmd)
}

func handleListingFiles() {
	currDir, err := os.Getwd()
	filesCount := 0
	dirsCount := 0

	if err != nil {
		fmt.Println("Failed to open this directory")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(currDir)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")
	for _, file := range files {
		if !checkDotFile(file) {
			if file.IsDir() {
				dirsCount++
			} else {
				filesCount++
			}
			fmt.Print(file.Name() + "  ")
		}
	}

	var fileName, dirName string

	if filesCount > 1 {
		fileName = "files"
	} else {
		fileName = "file"
	}

	if dirsCount > 1 {
		dirName = "directories"
	} else {
		dirName = "directory"
	}

	fmt.Printf("\n\n%d %s, %d %s\n", dirsCount, dirName, filesCount, fileName)
}

func checkDotFile(file os.FileInfo) bool {

	index := strings.Index(file.Name(), ".")

	if index == 0 {
		return true
	}

	return false
}

package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/makuzaverite/fitty/pkg/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "ls",
	Short: "List all files in a directory",
	Long:  "Listing all files in the directory",
	Run: func(cmd *cobra.Command, args []string) {

		listMore := utils.GetBool("more", cmd)

		if listMore {
			handleListingMore()
		} else {
			handleListingFiles()
		}

	},
}

func init() {

	listCmd.Flags().Bool("more", false, "Get more info about the list command")

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

		if file.IsDir() {
			dirsCount++
		} else {
			filesCount++
		}
		color.Cyan(file.Name() + "  ")

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

func handleListingMore() {
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

		info, err := os.Stat(file.Name())

		if err != nil {
			fmt.Println("info error")
		}

		kilos := float64(info.Size() / 1000)

		fmt.Print(kilos, " K ", "\t")

		fmt.Print(info.Mode(), "\t")

		// fmt.Print(inf)

		fmt.Println(file.Name())

		if file.IsDir() {
			dirsCount++
		} else {
			filesCount++
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

	currentDirectoryInfo, errorInOpenCurrDir := os.Stat(currDir)

	if errorInOpenCurrDir != nil {
		fmt.Println("Error in the process")
	}

	fmt.Printf("\n\n%d %s, %d %s, total size %d K\n", dirsCount, dirName, filesCount, fileName, currentDirectoryInfo.Size())
}

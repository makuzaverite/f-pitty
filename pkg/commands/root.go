package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/makuzaverite/fitty/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var rootCommand = &cobra.Command{
	Use:   "fitty <command> <subcommand>",
	Args:  cobra.ExactArgs(0),
	Short: "Interract with file system on all oses with the same commands",
	Long: `

    ____    _    __    __
   / __/   (_)  / /_  / /_   __  __
  / /_    / /  / __/ / __/  / / / /
 / __/   / /  / /_  / /_   / /_/ /
/_/     /_/   \__/  \__/   \__, /
                          /____/



Fitty the cross platform file system utility
Version: 0.0.1
	`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Example:       "$ fitty ls",
	Run: func(cmd *cobra.Command, args []string) {

		versionFlag := utils.GetBool("version", cmd)

		if versionFlag {
			version := "0.0.1"

			fmt.Printf("\n")
			fmt.Println("Fitty version ", version)
			fmt.Println("Check Current release https://github.com/makuzaverite/fitty/releases")

		} else {
			err := cmd.Help()

			if err != nil {
				fmt.Println("Failed to display help ", err, 1)
			}
			os.Exit(0)
		}
	},
}

// Execute initialized cli
func Execute() {

	cmd := rootCommand

	if err := doc.GenMarkdownTree(cmd, "./doc"); err != nil {
		log.Fatal("Doc gen failed")
	}

	rootCommand.Flags().Bool("version", false, "Get the current version of fitty")

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package commands

import (
	"fmt"
	"os"

	"github.com/makuzaverite/fitty/pkg/utils"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "fitty",
	Short: "Interract with filesystem on all oses with the same commands",
	Long: `


    ____    _    __    __
   / __/   (_)  / /_  / /_   __  __
  / /_    / /  / __/ / __/  / / / /
 / __/   / /  / /_  / /_   / /_/ /
/_/     /_/   \__/  \__/   \__, /
                          /____/






	Fitty the cross platform filesystem utility

	Version: 0.0.1
	`,
	Run: func(cmd *cobra.Command, args []string) {

		versionFlag := utils.GetBool("version", cmd)

		if versionFlag {
			version := "v0.0.1"

			fmt.Println("fitty version ", version)
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
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

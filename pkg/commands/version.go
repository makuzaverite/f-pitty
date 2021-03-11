package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

//VersionCMD cmd to check the version of the app
func VersionCMD(in string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Version check",
		Long:  "Display the current version of fitty you are running",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), in)
			fmt.Println("The current version is 0.0.0")
			return nil
		},
	}
}

func init() {
	//RootCommand.AddCommand(...VersionCMD().Commands())
}

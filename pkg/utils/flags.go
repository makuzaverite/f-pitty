package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

//GetBool a flag with type of boolean
func GetBool(name string, cmd *cobra.Command) bool {
	v, err := cmd.Flags().GetBool(name)
	if err != nil {
		fmt.Printf("Failed to get %v flag ", name)
	}
	return v
}

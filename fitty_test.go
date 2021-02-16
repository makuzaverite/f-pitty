package main

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func emptyRun(*cobra.Command, []string) {}

func checkStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func TestVersionFlag(t *testing.T) {
	rootCmd := &cobra.Command{Version: "0.0.1", Run: emptyRun}

	_, err := rootCmd.ExecuteC()

	if err != nil {
		t.Errorf("Un expected error: %v", err)
	}

	// output.Run(rootCmd)

}

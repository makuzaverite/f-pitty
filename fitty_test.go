package main

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// func Test_printError(t *testing.T) {
// 	cmd := &cobra.Command{}

// 	type args struct {
// 		err   error
// 		cmd   *cobra.Command
// 		debug bool
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantOut string
// 	}{
// 		{
// 			name: "Cobra flag error",
// 			args: args{
// 				err:   errors.New("The app field"),
// 				cmd:   nil,
// 				debug: false,
// 			},
// 			wantOut: "The app field",
// 		},
// 		{
// 			name: "unknown cobra command error",
// 			args: args{
// 				err:   errors.New("unknown command"),
// 				cmd:   cmd,
// 				debug: false,
// 			},
// 			wantOut: "unknown command",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			out := &bytes.Buffer{}
// 			//fmt.Println(out, tt.args.err)
// 			if goOut := out.String(); goOut != tt.wantOut {
// 				t.Errorf("Print error %q, wanted %q ", goOut, tt.wantOut)
// 			}
// 		})
// 	}
// }

// func TestCommand(t *testing.T) {
// 	var rootCmdArgs []string

// 	rootCmd := &cobra.Command{
// 		Use:  "root",
// 		Args: cobra.ExactArgs(1),
// 		Run:  func(_ *Command, args []string) { rootCmdArgs = args },
// 	}
// }

func emptyRun(*cobra.Command, []string) {}

func checkStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func TestVersionFlag(t *testing.T) {
	rootCmd := &cobra.Command{Version: "0.0.1", Run: emptyRun}

	output, err := rootCmd.ExecuteC()

	if err != nil {
		t.Errorf("Un expected error: %v", err)
	}

	// fmt.Println(output.Run())

	// checkStringContains(t,output,"fitty version 0.0.1")
}

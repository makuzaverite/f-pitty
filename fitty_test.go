package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func Test_printError(t *testing.T) {
	cmd := &cobra.Command{}

	type args struct {
		err   error
		cmd   *cobra.Command
		debug bool
	}

	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "Cobra flag error",
			args: args{
				err:   errors.New("The app field"),
				cmd:   nil,
				debug: false,
			},
			wantOut: "The app field",
		},
		{
			name: "unknow cobra command error",
			args: args{
				err:   errors.New("unkown command"),
				cmd:   cmd,
				debug: false,
			},
			wantOut: "unkwon command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			fmt.Println(out, tt.args.err)
			if goOut := out.String(); goOut != tt.wantOut {
				t.Errorf("Print error %q, wanted %q ", goOut, tt.wantOut)
			}
		})
	}
}

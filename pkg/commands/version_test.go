package commands

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_Version(t *testing.T) {
	cmd := VersionCMD("version")
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	// if we have string args
	//	cmd.SetArgs([]string{"hi-via-args"})

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "version" {
		t.Fatalf("Expected \"%s\" got \"%s\"", "hi", string(out))
	}

}

package main

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func testinit(t *testing.T) {

	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var (
		expectExitCode = 0
		expectMsg      = "Usage:"

		actualMsg      string
		actualExitCode int
		myExit         = func(code int) {
			actualExitCode = code
		}
	)

	// Assign the mock
	osExit = myExit
	// Run main() to capture STDOUT message and its exit-status-code
	actualMsg = capturer.CaptureOutput(func() {
		main()
	})

	// Assertion
	assert.Equal(t, actualExitCode, expectExitCode, "Unexpected exit code.")
	assert.Contains(t, actualMsg, expectMsg, "Should contain help message")
}

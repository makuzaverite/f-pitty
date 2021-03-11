package main

import (
	"fmt"
	"testing"
)

func testinit(t *testing.T) {

	fmt.Println("Hello World!")

	// oldOsExit := osExit
	// defer func() { osExit = oldOsExit }()

	// var (
	// 	expectExitCode = 0
	// 	expectMsg      = "Usage:"

	// 	actualMsg      string
	// 	actualExitCode int
	// 	myExit         = func(code int) {
	// 		actualExitCode = code
	// 	}
	// )

	// // Assign the mock
	// osExit = myExit
	// // Run main() to capture STDOUT message and its exit-status-code
	// actualMsg = capturer.CaptureOutput(func() {
	// 	main()
	// })

	// // Assertion
	// assert.Equal(t, actualExitCode, expectExitCode, "Unexpected exit code.")
	// assert.Contains(t, actualMsg, expectMsg, "Should contain help message")
	fmt.Println("Hello World!")
}

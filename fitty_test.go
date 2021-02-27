package main

import "testing"

func TestInit(t *testing.T) {
	ans := "Hello World!"

	if ans != Check() {
		t.Errorf("Expected to return %s", ans)
	}

}

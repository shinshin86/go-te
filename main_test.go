package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestValidateFileName(t *testing.T) {
	if os.Getenv("BE_ERROR") == "1" {
		validateFileName("foo_test.go")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestValidateFileName")
	cmd.Env = append(os.Environ(), "BE_ERROR=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with error %v, want exit status 1", err)
}

// This file acts as a test runner.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	. "github.com/shinshin86/go-te/te"
)

type TeSummary struct {
	testFiles      []string
	fileCnt        int
	fileSuccessCnt int
	fileFailCnt    int
	exitCode       int
}

func runnerExit(t *TeSummary) {
	report(t)

	fmt.Println("=== Tiny Expect: Finish ===")

	os.Exit(t.exitCode)
}

func report(t *TeSummary) {
	fmt.Println("=== Tiny Expect: Report ===")

	msg := "Test Files: "
	if t.fileFailCnt > 0 {
		msg = msg + strconv.Itoa(t.fileFailCnt) + " failed "
		t.exitCode = 1
	}
	msg = msg + strconv.Itoa(t.fileSuccessCnt) + " passed, " + strconv.Itoa(t.fileCnt) + " total"

	if t.fileFailCnt > 0 {
		DisplayFailMessage(msg, 0)
	} else {
		DisplaySuccessMessage(msg, 0)
	}
}

func listFiles(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	var testDir = flag.String("d", "test", "Specify test directory")
	flag.Parse()

	fmt.Println("=== Tiny Expect: Start ===")

	t := &TeSummary{
		fileCnt:        0,
		fileSuccessCnt: 0,
		fileFailCnt:    0,
		exitCode:       0,
	}

	files, err := listFiles(*testDir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, file := range files {
		// TODO: Ensure that only test files are retrieved.
		if filepath.Ext(file) != ".go" {
			continue
		}

		t.testFiles = append(t.testFiles, file)
		t.fileCnt++
	}

	testFilesMsg := "Found " + strconv.Itoa(t.fileCnt) + " test files"

	for _, f := range t.testFiles {
		DisplayTestFile(f, 0)

		out, err := exec.Command("go", "run", f).Output()
		if err != nil {
			t.fileFailCnt++
		} else {
			t.fileSuccessCnt++
		}

		fmt.Println(string(out))
	}

	DisplaySuccessMessage(testFilesMsg, 0)

	runnerExit(t)
}

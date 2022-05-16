// This file acts as a test runner.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

type config struct {
	TestMatch []string `json:testMath`
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

func listGoFiles(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(info.Name()) == ".go" {
				files = append(files, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func readConfig() *config {
	var cfg config

	b, err := ioutil.ReadFile("te.config.json")
	if err == nil {
		json.Unmarshal(b, &cfg)
	}

	return &cfg
}

func main() {
	var testDir = flag.String("d", "test", "Specify test directory")

	flag.Parse()

	cfg := readConfig()

	fmt.Println("=== Tiny Expect: Start ===")

	t := &TeSummary{
		fileCnt:        0,
		fileSuccessCnt: 0,
		fileFailCnt:    0,
		exitCode:       0,
	}

	if cfg.TestMatch != nil {
		files, err := filepath.Glob(cfg.TestMatch[0])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, file := range files {
			t.testFiles = append(t.testFiles, file)
			t.fileCnt++
		}
	} else {
		files, err := listGoFiles(*testDir)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, file := range files {
			t.testFiles = append(t.testFiles, file)
			t.fileCnt++
		}
	}

	testFilesMsg := "Found " + strconv.Itoa(t.fileCnt) + " test files"

	for _, f := range t.testFiles {
		DisplayTestFile(f, 0)

		out, err := exec.Command("go", "run", f).CombinedOutput()
		if err != nil {
			t.fileFailCnt++
			fmt.Println(err)
		} else {
			t.fileSuccessCnt++
		}

		fmt.Println(string(out))
	}

	DisplaySuccessMessage(testFilesMsg, 0)

	runnerExit(t)
}

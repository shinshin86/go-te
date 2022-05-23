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
	TestMatch []string
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

func readConfig() *config {
	var cfg config
	cfgDefaultPath := "te.config.json"

	var cfgPath = flag.String("c", cfgDefaultPath, "Specify config file path")

	flag.Parse()

	b, err := ioutil.ReadFile(*cfgPath)
	if err != nil && *cfgPath != cfgDefaultPath {
		fmt.Println("=== Tiny Expect: read config file error ===")
		fmt.Println(err)
	} else {
		json.Unmarshal(b, &cfg)
	}

	return &cfg
}

func main() {
	var testDir = flag.String("d", "test", "Specify test directory")
	var matchFiles = flag.String("m", "*.go", "Specify match files")

	cfg := readConfig()

	fmt.Println("=== Tiny Expect: Start ===")

	t := &TeSummary{
		fileCnt:        0,
		fileSuccessCnt: 0,
		fileFailCnt:    0,
		exitCode:       0,
	}

	if cfg.TestMatch != nil {
		var files []string

		for _, tm := range cfg.TestMatch {
			f, err := filepath.Glob(tm)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			files = append(files, f...)
		}

		for _, file := range files {
			t.testFiles = append(t.testFiles, file)
			t.fileCnt++
		}
	} else {
		files, err := filepath.Glob(filepath.Join(*testDir, *matchFiles))

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

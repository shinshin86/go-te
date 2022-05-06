package te

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Te struct {
	testCnt        int
	testSuccessCnt int
	testFailCnt    int
	exitCode       int
}

func Init() *Te {
	return &Te{
		exitCode:       0,
		testCnt:        0,
		testSuccessCnt: 0,
		testFailCnt:    0,
	}
}

func TeExit(t *Te) {
	report(t)
	os.Exit(t.exitCode)
}

func report(t *Te) {
	msg := "Test Case: "
	if t.testFailCnt > 0 {
		msg = msg + strconv.Itoa(t.testFailCnt) + " failed "
	}
	msg = msg + strconv.Itoa(t.testSuccessCnt) + " passed, " + strconv.Itoa(t.testCnt) + " total"

	if t.testFailCnt > 0 {
		DisplayFailMessage(msg, 0)
	} else {
		DisplaySuccessMessage(msg, 0)
	}

}

func DisplayTestFile(filename string, space int) {
	s := strings.Repeat(" ", space)
	text := s + "File: " + filename
	fmt.Printf("\x1b[36m%s\x1b[0m\n", text)
}

func DisplayTestTitle(title string, space int) {
	s := strings.Repeat(" ", space)
	fmt.Println(s + title)
}

func DisplaySuccessMessage(msg string, space int) {
	s := strings.Repeat(" ", space)
	fmt.Printf("\x1b[32m%s\x1b[0m\n", s+msg)
}

func DisplayFailMessage(msg string, space int) {
	s := strings.Repeat(" ", space)
	fmt.Printf("\x1b[31m%s\x1b[0m\n", s+msg)
}

func Describe(name string, fn func()) {
	DisplayTestTitle(name, 0)
	fn()
}

func It(t *Te, name string, fn func(t *Te)) {
	t.testCnt++
	title := strconv.Itoa(t.testCnt) + ": " + name

	DisplayTestTitle(title, 2)
	fn(t)
}

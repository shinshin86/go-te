package te

import (
	"os"
	"reflect"
	"strconv"
)

type Te struct {
	testCnt          int
	testSuccessCnt   int
	testFailCnt      int
	exitCode         int
	CurrentTestName  string
	CurrentTestValue interface{}
	CurrentTestType  reflect.Kind
}

func Init() *Te {
	return &Te{
		exitCode:       0,
		testCnt:        0,
		testSuccessCnt: 0,
		testFailCnt:    0,
	}
}

func (t *Te) Exit() {
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

func (t *Te) Describe(name string, fn func()) {
	DisplayTestTitle(name, 0)
	fn()
}

func (t *Te) It(name string, fn func()) {
	t.test(name, fn)
}

func (t *Te) Test(name string, fn func()) {
	t.test(name, fn)
}

func (t *Te) test(name string, fn func()) {
	t.testCnt++
	t.CurrentTestName = name
	title := strconv.Itoa(t.testCnt) + ": " + name

	DisplayTestTitle(title, 2)
	fn()
}

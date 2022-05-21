package te

import (
	"fmt"
	"reflect"
)

type CompareType interface {
	bool | int | float64 | string | rune
}

type CompareArrayType interface {
	[]uint8 | []string
}

func (t *Te) Expect(i interface{}) *Te {
	v := reflect.ValueOf(i)
	t.CurrentTestType = v.Kind()
	t.CurrentTestValue = i
	return t
}

func (t *Te) notExpectedTypeMsg() {
	msg := fmt.Sprintf("Error: Not expected type. (value: %v, type: %T)", t.CurrentTestValue, t.CurrentTestValue)
	DisplayFailMessage(msg, 4)
}

func notExpectedTypeMsg(i interface{}) {
	msg := fmt.Sprintf("Error: Not expected type. (value: %v, type: %T)", i, i)
	DisplayFailMessage(msg, 4)
}

func (t *Te) NotToBe(i interface{}) {
	switch t.CurrentTestType {
	case reflect.Bool:
		actual, ok := t.CurrentTestValue.(bool)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(bool)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertNotEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Int:
		actual, ok := t.CurrentTestValue.(int)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(int)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertNotEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Float64:
		actual, ok := t.CurrentTestValue.(float64)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(float64)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertNotEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.String:
		actual, ok := t.CurrentTestValue.(string)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(string)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertNotEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Int32: // rune
		actual, ok := t.CurrentTestValue.(rune)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(rune)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertNotEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Slice:
		switch i.(type) {
		case []uint8:
			actual, ok := t.CurrentTestValue.([]uint8)
			if ok == false {
				t.notExpectedTypeMsg()
				t.exitCode = 1
			}

			expect, ok := i.([]uint8)
			if ok == false {
				notExpectedTypeMsg(i)
				t.exitCode = 1
			}

			if assertNotDeepEqual(actual, expect, t.CurrentTestName) {
				t.testSuccessCnt++
			} else {
				t.testFailCnt++
			}
		}
	case reflect.Array:
		actual := reflect.ValueOf(t.CurrentTestValue)
		expect := reflect.ValueOf(i)

		if actual.Len() != expect.Len() {
			msg := fmt.Sprintf("Error: The number of arrays to be compared does not match.\n")
			DisplayFailMessage(msg, 4)
			t.exitCode = 1
		}

		itemLen := actual.Len()
		equalCount := 0

		for idx := 0; idx < actual.Len(); idx++ {
			actualItem := actual.Index(idx).Interface()
			expectItem := expect.Index(idx).Interface()

			if actualItem == expectItem {
				equalCount++
			}
		}

		if itemLen == equalCount {
			msg := fmt.Sprintf("Failed!: " + t.CurrentTestName + "\n")
			msg = msg + fmt.Sprintf("    Actual: %v, Not expected: %v", actual, expect)
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		} else {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		}
	case reflect.Pointer:
		actual := reflect.ValueOf(t.CurrentTestValue)
		expect := reflect.ValueOf(i)

		if !reflect.DeepEqual(actual, expect) {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		} else {
			msg := fmt.Sprintf("Failed!: " + t.CurrentTestName + "\n")
			msg = msg + fmt.Sprintf("    Actual: %s, Expected: %s", actual, expect)
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		}
	default:
		msg := fmt.Sprintf("ERROR: Not found invalid type. %s\n", t.CurrentTestType)
		DisplayFailMessage(msg, 4)
		t.exitCode = 1
	}
}

func (t *Te) ToBe(i interface{}) {
	switch t.CurrentTestType {
	case reflect.Bool:
		actual, ok := t.CurrentTestValue.(bool)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(bool)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Int:
		actual, ok := t.CurrentTestValue.(int)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(int)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Float64:
		actual, ok := t.CurrentTestValue.(float64)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(float64)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.String:
		actual, ok := t.CurrentTestValue.(string)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(string)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Int32: // rune
		actual, ok := t.CurrentTestValue.(rune)
		if ok == false {
			t.notExpectedTypeMsg()
			t.exitCode = 1
		}

		expect, ok := i.(rune)
		if ok == false {
			notExpectedTypeMsg(i)
			t.exitCode = 1
		}

		if assertEqual(actual, expect, t.CurrentTestName) {
			t.testSuccessCnt++
		} else {
			t.testFailCnt++
		}
	case reflect.Slice:
		switch i.(type) {
		case []uint8:
			actual, ok := t.CurrentTestValue.([]uint8)
			if ok == false {
				t.notExpectedTypeMsg()
				t.exitCode = 1
			}

			expect, ok := i.([]uint8)
			if ok == false {
				notExpectedTypeMsg(i)
				t.exitCode = 1
			}

			if assertDeepEqual(actual, expect, t.CurrentTestName) {
				t.testSuccessCnt++
			} else {
				t.testFailCnt++
			}
		}
	case reflect.Array:
		actual := reflect.ValueOf(t.CurrentTestValue)
		expect := reflect.ValueOf(i)

		if actual.Len() != expect.Len() {
			msg := fmt.Sprintf("Error: The number of arrays to be compared does not match.\n")
			DisplayFailMessage(msg, 4)
			t.exitCode = 1
		}

		for idx := 0; idx < actual.Len(); idx++ {
			actualItem := actual.Index(idx).Interface()
			expectItem := expect.Index(idx).Interface()

			if !reflect.DeepEqual(actualItem, expectItem) {
				msg := fmt.Sprintf("Failed!: " + t.CurrentTestName + "\n")
				msg = msg + fmt.Sprintf("    Actual: %v, Expected: %v", actual, expect)
				DisplayFailMessage(msg, 4)
				t.testFailCnt++
			}
		}

		DisplaySuccessMessage("Succeeded", 4)
		t.testSuccessCnt++
	case reflect.Pointer:
		actual := reflect.ValueOf(t.CurrentTestValue)
		expect := reflect.ValueOf(i)

		if reflect.DeepEqual(actual, expect) {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		} else {
			msg := fmt.Sprintf("Failed!: " + t.CurrentTestName + "\n")
			msg = msg + fmt.Sprintf("    Actual: %s, Expected: %s", actual, expect)
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		}
	default:
		msg := fmt.Sprintf("ERROR: Not found invalid type. %s\n", t.CurrentTestType)
		DisplayFailMessage(msg, 4)
		t.exitCode = 1
	}
}

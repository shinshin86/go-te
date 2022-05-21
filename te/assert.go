package te

import (
	"fmt"
	"reflect"
)

func assertEqual[T CompareType](actual, expect T, testName string) bool {
	if actual == expect {
		DisplaySuccessMessage("Succeeded", 4)
		return true
	} else {
		msg := fmt.Sprintf("Failed!: " + testName + "\n")
		msg = msg + fmt.Sprintf("    Actual: %v, Expected: %v", actual, expect)
		DisplayFailMessage(msg, 4)
		return false
	}
}

func assertDeepEqual[T CompareArrayType](actual, expect T, testName string) bool {
	if reflect.DeepEqual(actual, expect) {
		DisplaySuccessMessage("Succeeded", 4)
		return true
	} else {
		msg := fmt.Sprintf("Failed!: " + testName + "\n")
		msg = msg + fmt.Sprintf("    Actual: %v, Expected: %v", actual, expect)
		DisplayFailMessage(msg, 4)
		return false
	}
}

func assertNotDeepEqual[T CompareArrayType](actual, expect T, testName string) bool {
	if !reflect.DeepEqual(actual, expect) {
		DisplaySuccessMessage("Succeeded", 4)
		return true
	} else {
		msg := fmt.Sprintf("Failed!: " + testName + "\n")
		msg = msg + fmt.Sprintf("    Actual: %v, Expected: %v", actual, expect)
		DisplayFailMessage(msg, 4)
		return false
	}
}

func assertNotEqual[T CompareType](actual, expect T, testName string) bool {
	if actual != expect {
		DisplaySuccessMessage("Succeeded", 4)
		return true
	} else {
		msg := fmt.Sprintf("Failed!: " + testName + "\n")
		msg = msg + fmt.Sprintf("    Actual: %v, Not expected: %v", actual, expect)
		DisplayFailMessage(msg, 4)
		return false
	}
}

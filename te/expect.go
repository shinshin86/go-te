package te

import (
	"fmt"
	"reflect"
	"strconv"
)

func (t *Te) Expect(i interface{}) *Te {
	v := reflect.ValueOf(i)
	t.ActualType = v.Kind()
	t.Actual = i
	return t
}

func (t *Te) ToBe(i interface{}) {
	switch t.ActualType {
	case reflect.Bool:
		actual, ok := t.Actual.(bool)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", t.Actual, t.Actual)
			t.exitCode = 1
		}

		expect, ok := i.(bool)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", i, i)
			t.exitCode = 1
		}

		if actual == expect {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		} else {
			msg := fmt.Sprintf("Failed! Actual: %s, Expected: %s", strconv.FormatBool(actual), strconv.FormatBool(expect))
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		}
	case reflect.Int:
		actual, ok := t.Actual.(int)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", t.Actual, t.Actual)
			t.exitCode = 1
		}

		expect, ok := i.(int)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", i, i)
			t.exitCode = 1
		}

		if actual == expect {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		} else {
			msg := fmt.Sprintf("Failed! Actual: %d, Expected: %d", actual, expect)
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		}
	case reflect.Float64:
		actual, ok := t.Actual.(float64)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", t.Actual, t.Actual)
			t.exitCode = 1
		}

		expect, ok := i.(float64)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", i, i)
			t.exitCode = 1
		}

		if actual == expect {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		} else {
			msg := fmt.Sprintf("Failed! Actual: %f, Expected: %f", actual, expect)
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		}
	case reflect.String:
		actual, ok := t.Actual.(string)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", t.Actual, t.Actual)
			t.exitCode = 1
		}

		expect, ok := i.(string)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", i, i)
			t.exitCode = 1
		}

		if actual == expect {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		} else {
			msg := fmt.Sprintf("Failed! Actual: %s, Expected: %s", actual, expect)
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		}
	case reflect.Int32: // rune
		actual, ok := t.Actual.(rune)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", t.Actual, t.Actual)
			t.exitCode = 1
		}

		expect, ok := i.(rune)
		if ok == false {
			fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", i, i)
			t.exitCode = 1
		}

		if actual == expect {
			DisplaySuccessMessage("Succeeded", 4)
			t.testSuccessCnt++
		} else {
			msg := fmt.Sprintf("Failed! Actual: %d, Expected: %d", actual, expect)
			DisplayFailMessage(msg, 4)
			t.testFailCnt++
		}
	case reflect.Slice:
		switch i.(type) {
		case []uint8:
			actual, ok := t.Actual.([]uint8)
			if ok == false {
				fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", t.Actual, t.Actual)
				t.exitCode = 1
			}

			expect, ok := i.([]uint8)
			if ok == false {
				fmt.Printf("Error: Not expected type. (value: %v, type: %T)\n", i, i)
				t.exitCode = 1
			}

			if reflect.DeepEqual(actual, expect) {
				DisplaySuccessMessage("Succeeded", 4)
				t.testSuccessCnt++
			} else {
				msg := fmt.Sprintf("Failed! Actual: %s, Expected: %s", actual, expect)
				DisplayFailMessage(msg, 4)
				t.testFailCnt++
			}
		}

	default:
		fmt.Printf("ERROR: Not found invalid type. %s\n", t.ActualType)
		t.exitCode = 1
	}
}

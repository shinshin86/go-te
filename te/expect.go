package te

import (
	"fmt"
	"strconv"
)

func Expect(i interface{}) func(*Te, interface{}) {
	switch i.(type) {
	case bool:
		actual := i.(bool)

		return func(t *Te, i2 interface{}) {
			expect, ok := i2.(bool)
			if ok == false {
				fmt.Println("Type error")
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
		}
	case int:
		actual := i.(int)

		return func(t *Te, i2 interface{}) {
			expect, ok := i2.(int)
			if ok == false {
				fmt.Println("Type error")
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
		}
	case float64:
		actual := i.(float64)

		return func(t *Te, i2 interface{}) {
			expect, ok := i2.(float64)
			if ok == false {
				fmt.Println("Type error")
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
		}
	case string:
		actual := i.(string)

		return func(t *Te, i2 interface{}) {
			expect, ok := i2.(string)
			if ok == false {
				fmt.Println("Type error")
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
		}
	default:
		return func(t *Te, i2 interface{}) {
			fmt.Println("ERROR")
			t.exitCode = 1
		}
	}
}

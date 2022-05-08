package main

import (
	. "github.com/shinshin86/go-te/te"
)

func add(x, y int) int {
	return x + y
}

func main() {
	t := Init()

	Describe("add function test(It)", func() {
		It(t, "Expected to be output sum of the two arguments", func(t *Te) {
			t.Expect(add(2, 3)).ToBe(5)
		})
	})

	Describe("add function test(Test)", func() {
		Test(t, "Expected to be output sum of the two arguments", func(t *Te) {
			t.Expect(add(2, 3)).ToBe(5)
		})
	})

	TeExit(t)
}

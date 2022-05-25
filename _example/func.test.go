package main

import (
	. "github.com/shinshin86/go-te/te"
)

func add(x, y int) int {
	return x + y
}

func main() {
	t := Init()

	t.Describe("add function test(It)", func() {
		t.It("Expected to be output sum of the two arguments", func() {
			t.Expect(add(2, 3)).ToBe(5)
		})
	})

	t.Describe("add function test(Test)", func() {
		t.Test("Expected to be output sum of the two arguments", func() {
			t.Expect(add(2, 3)).ToBe(5)
		})
	})

	t.Exit()
}

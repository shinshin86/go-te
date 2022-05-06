package main

import (
	. "github.com/shinshin86/go-te/te"
)

func add(x, y int) int {
	return x + y
}

func main() {
	t := Init()

	Describe("Minimal sample 2", func() {
		It(t, "Expected to be output sum of the two arguments", func(t *Te) {
			toBe := Expect(add(2, 3))
			toBe(t, 5)
		})
	})

	TeExit(t)
}

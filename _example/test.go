package main

import (
	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	Describe("Minimal sample", func() {
		It(t, "Expect b is true", func(t *Te) {
			b := true
			toBe := Expect(b)
			toBe(t, true)
		})

		It(t, "Expect i is 1", func(t *Te) {
			i := 1
			toBe := Expect(i)
			toBe(t, 1)
		})

		It(t, "Expect s is helloworld", func(t *Te) {
			s := "helloworld"
			toBe := Expect(s)
			toBe(t, "helloworld")
		})
	})

	TeExit(t)
}

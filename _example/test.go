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

		It(t, "Expect f is 1.2", func(t *Te) {
			f := 1.2
			toBe := Expect(f)
			toBe(t, 1.2)
		})

		It(t, "Expect f is 1.23", func(t *Te) {
			f := 1.23
			toBe := Expect(f)
			toBe(t, 1.23)
		})

		It(t, "Expect f is 1.234", func(t *Te) {
			f := 1.234
			toBe := Expect(f)
			toBe(t, 1.234)
		})

		It(t, "Expect f is 1.2345", func(t *Te) {
			f := 1.2345
			toBe := Expect(f)
			toBe(t, 1.2345)
		})

		It(t, "Expect f is 1.23456", func(t *Te) {
			f := 1.23456
			toBe := Expect(f)
			toBe(t, 1.23456)
		})

		It(t, "Expect s is helloworld", func(t *Te) {
			s := "helloworld"
			toBe := Expect(s)
			toBe(t, "helloworld")
		})

		It(t, "Expect r is a", func(t *Te) {
			r := 'a'
			toBe := Expect(r)
			toBe(t, 'a')
		})
	})

	TeExit(t)
}

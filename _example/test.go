package main

import (
	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	Describe("Minimal sample", func() {
		It(t, "Expect b is true", func(t *Te) {
			b := true
			t.Expect(b).ToBe(true)
		})

		It(t, "Expect i is 1", func(t *Te) {
			i := 1
			t.Expect(i).ToBe(1)
		})

		It(t, "Expect f is 1.2", func(t *Te) {
			f := 1.2
			t.Expect(f).ToBe(1.2)
		})

		It(t, "Expect f is 1.23", func(t *Te) {
			f := 1.23
			t.Expect(f).ToBe(1.23)
		})

		It(t, "Expect f is 1.234", func(t *Te) {
			f := 1.234
			t.Expect(f).ToBe(1.234)
		})

		It(t, "Expect f is 1.2345", func(t *Te) {
			f := 1.2345
			t.Expect(f).ToBe(1.2345)
		})

		It(t, "Expect f is 1.23456", func(t *Te) {
			f := 1.23456
			t.Expect(f).ToBe(1.23456)
		})

		It(t, "Expect s is helloworld", func(t *Te) {
			s := "helloworld"
			t.Expect(s).ToBe("helloworld")
		})

		It(t, "Expect r is a", func(t *Te) {
			r := 'a'
			t.Expect(r).ToBe('a')
		})

		It(t, "Expect b is abc", func(t *Te) {
			b := []byte("abc")
			t.Expect(b).ToBe([]byte("abc"))
		})

		It(t, "Expect b is abc([]byte{97, 98, 99}", func(t *Te) {
			b := []byte("abc")
			t.Expect(b).ToBe([]byte{97, 98, 99})
		})
	})

	TeExit(t)
}

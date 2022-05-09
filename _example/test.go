package main

import (
	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	t.Describe("Minimal sample", func() {
		t.It("Expect b is true", func() {
			b := true
			t.Expect(b).ToBe(true)
		})

		t.It("Expect i is 1", func() {
			i := 1
			t.Expect(i).ToBe(1)
		})

		t.It("Expect f is 1.2", func() {
			f := 1.2
			t.Expect(f).ToBe(1.2)
		})

		t.It("Expect f is 1.23", func() {
			f := 1.23
			t.Expect(f).ToBe(1.23)
		})

		t.It("Expect f is 1.234", func() {
			f := 1.234
			t.Expect(f).ToBe(1.234)
		})

		t.It("Expect f is 1.2345", func() {
			f := 1.2345
			t.Expect(f).ToBe(1.2345)
		})

		t.It("Expect f is 1.23456", func() {
			f := 1.23456
			t.Expect(f).ToBe(1.23456)
		})

		t.It("Expect s is helloworld", func() {
			s := "helloworld"
			t.Expect(s).ToBe("helloworld")
		})

		t.It("Expect r is a", func() {
			r := 'a'
			t.Expect(r).ToBe('a')
		})

		t.It("Expect b is abc", func() {
			b := []byte("abc")
			t.Expect(b).ToBe([]byte("abc"))
		})

		t.It("Expect b is abc([]byte{97, 98, 99}", func() {
			b := []byte("abc")
			t.Expect(b).ToBe([]byte{97, 98, 99})
		})
	})

	t.Exit()
}

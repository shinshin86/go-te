package main

import (
	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	t.Describe("Minimal sample (sub2 dir)", func() {
		t.It("Expect b is true", func() {
			b := true
			t.Expect(b).ToBe(true)
		})
	})

	t.Exit()
}

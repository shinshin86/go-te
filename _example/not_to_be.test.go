package main

import (
	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	t.Describe("Expoect not to be test", func() {
		t.It("Expect b is not true", func() {
			b := true
			t.Expect(b).NotToBe(false)
		})
	})

	t.Exit()
}

package main

import (
	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	t.Describe("Expect to be sample", func() {
		t.It("Expect b is true", func() {
			b := true
			t.Expect(b).ToBe(true)
		})

		t.It("Expect i is 1", func() {
			i := 1
			t.Expect(i).ToBe(1)
		})

		// The name Test is also available. The functionality is the same.
		t.Test("Expect s is helloworld", func() {
			s := "helloworld"
			t.Expect(s).ToBe("helloworld")
		})
	})

	t.Describe("Expoect not to be test", func() {
		t.It("Expect b is not true", func() {
			b := true
			t.Expect(b).NotToBe(false)
		})

		t.It("Expect i is not 1", func() {
			i := 1
			t.Expect(i).NotToBe(2)
		})
	})

	t.Exit()
}

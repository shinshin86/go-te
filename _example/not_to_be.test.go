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

		t.It("Expect i is not 1", func() {
			i := 1
			t.Expect(i).NotToBe(2)
		})

		t.It("Expect f is not 1.2", func() {
			f := 1.2
			t.Expect(f).NotToBe(1.3)
		})

		t.It("Expect f is not 1.23", func() {
			f := 1.23
			t.Expect(f).NotToBe(1.24)
		})

		t.It("Expect f is not 1.234", func() {
			f := 1.234
			t.Expect(f).NotToBe(1.235)
		})

		t.It("Expect f is not 1.2345", func() {
			f := 1.2345
			t.Expect(f).NotToBe(1.2346)
		})

		t.It("Expect f is not 1.23456", func() {
			f := 1.23456
			t.Expect(f).NotToBe(1.23457)
		})

		t.It("Expect s is not helloworld", func() {
			s := "helloworld"
			t.Expect(s).NotToBe("helloworldd")
		})

		t.It("Expect r is not a", func() {
			r := 'a'
			t.Expect(r).NotToBe('b')
		})

		t.It("Expect b is not abc", func() {
			b := []byte("abc")
			t.Expect(b).NotToBe([]byte("abd"))
		})

		t.It("Expect b is not abc([]byte{97, 98, 99}", func() {
			b := []byte("abc")
			t.Expect(b).NotToBe([]byte{97, 98, 100})
		})

		t.It("Expect a is not [1 2]", func() {
			var a [2]int
			a[0] = 1
			a[1] = 2
			t.Expect(a).NotToBe([2]int{1, 3})
		})

		t.It("Expect a is not [Hello World]", func() {
			var a [2]string
			a[0] = "Hello"
			a[1] = "World"
			t.Expect(a).NotToBe([2]string{"Hello", "Worldd"})
		})

		t.It("Expect p is not ", func() {
			i := 123
			i2 := 123
			p := &i
			p2 := &i2
			t.Expect(p).NotToBe(p2)
		})
	})

	t.Exit()
}

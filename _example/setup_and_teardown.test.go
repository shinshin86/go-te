package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	cwd, _ := os.Getwd()

	filename := "test.txt"
	testtxt := "Setup and Teardown test"

	num := 1

	t.BeforeAll(func() {
		fp, err := os.Create(filepath.Join(cwd, "_example", filename))
		if err != nil {
			fmt.Println(err)
			return
		}

		defer fp.Close()

		fp.WriteString(testtxt)
	})

	t.AfterAll(func() {
		if err := os.Remove(filepath.Join(cwd, "_example", filename)); err != nil {
			fmt.Println(err)
		}
	})

	t.BeforeEach(func() {
		num++
		// fmt.Printf("before each: %d\n", num)
	})

	t.AfterEach(func() {
		num++
		// fmt.Printf("after each: %d\n", num)
	})

	t.Describe("Setup and Teardown test", func() {
		t.It("Expected a file strings are the same", func() {
			f, _ := os.Open(filepath.Join(cwd, "_example", filename))

			defer f.Close()

			b, _ := ioutil.ReadAll(f)

			t.Expect(string(b)).ToBe(testtxt)
		})

		t.It("Expected the variable num to be incremented a specified number of times", func() {
			// 1 -> 2(BeforeEach) -> 3(AfterEach) -> 4(BeforeEach: this function)
			t.Expect(num).ToBe(4)
		})
	})

	t.Exit()
}

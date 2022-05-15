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

	t.Describe("Setup and Teardown test", func() {
		t.It("Expected a file strings are the same", func() {
			f, _ := os.Open(filepath.Join(cwd, "_example", filename))

			defer f.Close()

			b, _ := ioutil.ReadAll(f)

			t.Expect(string(b)).ToBe(testtxt)
		})
	})

	t.Exit()
}

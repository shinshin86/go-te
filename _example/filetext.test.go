package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/shinshin86/go-te/te"
)

func main() {
	t := Init()

	t.Describe("File text comparison test", func() {
		t.It("Expected to be output same of the two file", func() {
			cwd, _ := os.Getwd()

			f, _ := os.Open(filepath.Join(cwd, "_example", "testdata.txt"))
			f2, _ := os.Open(filepath.Join(cwd, "_example", "testdata.txt"))

			defer f.Close()
			defer f2.Close()

			b, _ := ioutil.ReadAll(f)
			b2, _ := ioutil.ReadAll(f2)

			t.Expect(b).ToBe(b2)
		})
	})

	t.Exit()
}

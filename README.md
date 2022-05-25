# (WIP) go-te 💨
[![CI](https://github.com/shinshin86/go-te/actions/workflows/ci.yml/badge.svg)](https://github.com/shinshin86/go-te/actions/workflows/ci.yml)

te(Tiny Expect) is a tiny test library for go.

## Note
te is not recommended for use, although I have using it myself in very limited circumstances.  
(This one was created with a bit of an idea in mind).

It also does not have all the features needed for testing, so please use it for fun only.

## Install
Please install it to perform the test with the command `go-te`.

```sh
go install github.com/shinshin86/go-te@latest
```

It must then be added to the project to be used.

```sh
go get github.com/shinshin86/go-te
```


## Usage

This is a simple example.

```go
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
```

## Setup and Teardown

Often, there are cases where you want to do specific processing before and after the test is run. te also provides helper functions to assist in these cases.

Note that these helper functions must be defined prior to test execution.

The functions provided are as follows.

* BeforeAll
* AfterAll
* BeforeEach
* AfterEach

An example of use can be found here.

```go
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
	})

	t.AfterEach(func() {
		num++
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
```

## Command line options

```sh
go-te --help

# -c string
#   	Specify config file path (default "te.config.json")
# -d string
#   	Specify test directory (default "test")
# -m string
#   	Specify match files (default "*.go")
```

## Config file
You can create a `te.config.json` file and put runtime test options there.

```json
{
    "testMatch": ["_example/*.go", "_example/sub/*.go"]
}
```

 - `testMatch` - You can specify the path to the test file.

## About Test File Names
There are no specific rules about files for testing in `go-te`.
However, as a specification of go, files with `_test.go` such as `foo_test.go` cannot be used.

## Examples

You can also try the example code for this project.  
Execute the following command.

```sh
# install
go install github.com/shinshin86/go-te@latest

# project clone
git clone https://github.com/shinshin86/go-te.git
cd go-te

# run
go-te -d _example
```

## Running Tests
Tests are performed on code written under the `_example` directory.

```sh
make test
```

## License
[MIT](https://github.com/shinshin86/go-te/blob/main/LICENSE)

## Author
[Yuki Shindo](https://shinshin86.com/en)
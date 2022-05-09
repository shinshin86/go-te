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

	t.Describe("Minimal sample", func() {
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

	t.Exit()
}
```

### Example

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
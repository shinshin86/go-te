package te

import (
	"fmt"
	"strings"
)

func DisplayTestFile(filename string, space int) {
	s := strings.Repeat(" ", space)
	text := s + "File: " + filename
	fmt.Printf("\x1b[36m%s\x1b[0m\n", text)
}

func DisplayTestTitle(title string, space int) {
	s := strings.Repeat(" ", space)
	fmt.Println(s + title)
}

func DisplaySuccessMessage(msg string, space int) {
	s := strings.Repeat(" ", space)
	fmt.Printf("\x1b[32m%s\x1b[0m\n", s+msg)
}

func DisplayFailMessage(msg string, space int) {
	s := strings.Repeat(" ", space)
	fmt.Printf("\x1b[31m%s\x1b[0m\n", s+msg)
}

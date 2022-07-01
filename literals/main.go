package main

import "fmt"

// Create string literals
var (
	s1 = "Interpreted literal strings use a double-quote\nand need to escape special characters\n\n"
	s2 = `Raw string literal uses back-quote
and do not need to escape special characters

`
)

const (
	c1 = "Interpreted literal constant\n"
	c2 = `Raw string literal constant
`
)

func main() {
	fmt.Print(s1)
	fmt.Print(s2)
	fmt.Print(c1)
	fmt.Print(c2)
}

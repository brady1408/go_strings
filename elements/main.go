package main

import (
	"fmt"
)

func main() {
	s := "Hello Nav!"
	// fmt.Println(s[0])
	// fmt.Println(string(s[0]))
	// fmt.Println(s[:5])

	// fmt.Print("\n\n")

	// a string has to be set fully it you cannot set its elements
	// s[0] = byte('W')
	// s = fmt.Sprintf("O%s", s[2:])
	// fmt.Println(s)

	// fmt.Print("\n\n")

	// If a string has an underlying type of []byte does this mean we can range a string?
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}

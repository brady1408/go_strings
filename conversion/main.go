package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func main() {
	s := "Nav!"
	bs := []byte(s)      // string -> []byte
	s = string(bs)       // []byte -> string
	rs := []rune(s)      // string -> []rune
	s = string(rs)       // []rune -> string
	rs = bytes.Runes(bs) // []byte -> []rune
	bs = Runes2Bytes(rs) // []rune -> []byte
	s = string(72)       // int	   -> string

	stringToByteSlice(s)
}

func stringToByteSlice(s string) {
	greetingPerson := []byte("Hello ")
	greetingPerson = append(greetingPerson, s...)
	fmt.Println(string(greetingPerson))
}

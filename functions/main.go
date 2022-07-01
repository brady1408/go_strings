package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Hello Nav!"
	fmt.Println(len(s)) // Built in len()

	// strings package
	greeting := strings.Split(s, " ")[0]
	fmt.Println(greeting, "Navaroni!")

	fmt.Print("\n\n")

	// concat with strings.Join
	people := []string{"Nav", "Navster", "Navaroni and Cheese"}
	fmt.Println("strings.join")
	for _, v := range people {
		fmt.Println(strings.Join([]string{greeting, " ", v, "!"}, ""))
	}

	fmt.Print("\n\n")

	// concat with strings.Builder
	fmt.Println("strings.Builder")
	builder := &strings.Builder{}
	for _, v := range people {
		fmt.Fprintf(builder, "%s %s!\n", greeting, v)
	}
	fmt.Println(builder.String())

	// strconv package
	// convert from a int to a string
	fmt.Println(strconv.Itoa(42))
	i, _ := strconv.Atoi("42")
	fmt.Println(i)
	i64, _ := strconv.ParseInt("42", 10, 64)
	fmt.Println(i64)
}

package main

import (
	"bytes"
	"fmt"
)

func main() {
	greeting := "Hello"
	people := []string{"Nav", "Navster", "Navaronie and Cheese"}

	fmt.Println("Using go + operator")
	for _, v := range people {
		fmt.Println(greeting + " " + v + "!")
	}

	fmt.Print("\n\n")

	fmt.Println("using the += operator")
	var greetPeople string
	for _, v := range people {
		greetPerson := greeting + " " + v + "!\n"
		greetPeople += greetPerson
	}
	fmt.Println(greetPeople)

	fmt.Print("\n\n")

	fmt.Println("using fmt package")
	greetPeople = ""
	for _, v := range people {
		greetPerson := fmt.Sprintf("%s %s!\n", greeting, v)
		greetPeople += greetPerson
	}
	fmt.Println(greetPeople)

	fmt.Println("using bytes package")
	b := bytes.Buffer{}
	for _, v := range people {
		b.Write([]byte(greeting))
		b.Write([]byte(" "))
		b.Write([]byte(v))
		b.Write([]byte("!\n"))
	}
	fmt.Println(b.String())
}

package main

import "fmt"

func main() {
	s := "Thundercats"
	fmt.Printf("string s: %s equals Thundercats? %t\n", s, s == "Thundercats")
	fmt.Printf("string s: %s not equal Thundercats? %t\n", s, s != "Thundercats")
	fmt.Printf("string s: %s > thundercats? %t\n", s, s > "thundercats")
	fmt.Printf("string s: %s < thundercats? %t\n", s, s < "thundercats")
	fmt.Printf("string s: %s >= thundercats? %t\n", s, s >= "thundercats")
	fmt.Printf("string s: %s <= thundercats? %t\n", s, s <= "thundercats")
}

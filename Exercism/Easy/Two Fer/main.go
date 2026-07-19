package main

import (
	"fmt"
)

func ShareWith(name string) string {
	if len(name) != 0 {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}

func main() {
	fmt.Println(ShareWith("Alice"))
	fmt.Println(ShareWith("Bohdan"))
	fmt.Println(ShareWith(""))
	fmt.Println(ShareWith("Zaphod"))
	fmt.Println(ShareWith("Rain"))
}

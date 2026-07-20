package main

/*
	The user first enter a poisitive integer n, then n words one after the other. The program outputs "jeste" if
	the user entered the word "programiranje" exactly two times, else it outputs "nije".
*/

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Enter a positive integer n:")
	fmt.Scan(&n)

	var w string
	counter := 0
	fmt.Printf("Enter %d words:\n", n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&w)
		if w == "programiranje" {
			counter++
		}
	}

	if counter == 2 {
		fmt.Println("jeste")
	} else {
		fmt.Println("nije")
	}
}

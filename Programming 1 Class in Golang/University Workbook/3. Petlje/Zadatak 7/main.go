package main

/*
	Write a programme where the user enters words, each on a new line. The end of the input is marked by a blank line.
	The program has to output the written words on the same line, separated by spaces.
*/

import (
	"fmt"
)

func main() {
	var w string
	fmt.Println("Enter words. The end is marked by a blank line.")
	words := []string{}
	for {
		_, err := fmt.Scanln(&w)
		if err != nil {
			break
		}

		words = append(words, w)
	}

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s ", words[i])
	}
	fmt.Println("")
}

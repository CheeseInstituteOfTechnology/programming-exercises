package main

/*
	The user enters whole numbers one after the other until they enter -1, to which the loop ends. The program has
	to calculate and output the sum of the last digits of the entered numbers. So 11, 1022, and 102103 would be 1 + 2 + 3.
*/

import (
	"fmt"
)

func main() {
	var n int
	sum := 0

	fmt.Println("Enter numbers. To end the program enter -1.")
	for true {
		fmt.Scan(&n)
		if n == -1 || n < 0 {
			break
		}
		sum += n % 10
	}

	fmt.Println(sum)
}

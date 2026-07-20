package main

/*
	The user enters positive integer numbers one after the other. The end of the input is when the user enters -1.
	The program outputs the sum of the last digits for one digit numbers, and the second last digits of numbers that
	contain more than one digit in them.
*/

import (
	"fmt"
)

func main() {
	fmt.Println("Enter numbers. To end, input -1.")
	var n int
	sum := 0
	for true {
		fmt.Scan(&n)
		if n == -1 || n < 0 {
			break
		}

		if n >= 1 && n < 10 {
			sum += n
		} else {
			sum += (n / 10) % 10
		}
	}

	fmt.Println(sum)
}

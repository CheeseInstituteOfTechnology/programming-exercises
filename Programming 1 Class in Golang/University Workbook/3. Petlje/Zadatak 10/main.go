package main

/*
	The user enters a whole digit number. The program outputs the difference of the product and quotient of the largest
	and smallest digit in the number. You can assume that every entered number does not have the value 0.
*/

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	var largest int64 = 0
	var smallest int64 = 10
	for n != 0 {
		digit := n % 10
		if digit > int64(largest) {
			largest = digit
		}
		if digit < smallest {
			smallest = digit
		}

		n = int64(n / 10)
	}

	product := largest * smallest
	quotient := largest / smallest
	fmt.Println(product - quotient)
}

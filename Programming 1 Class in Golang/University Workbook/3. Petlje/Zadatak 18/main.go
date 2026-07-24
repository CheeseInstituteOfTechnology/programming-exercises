package main

/*
	Write a program where the user has to enter an integer number n. The program outputs the product of the odd digits
	that are less than 6.
*/

import (
	"fmt"
)

func main() {
	var n int64

	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	var p int64 = 1
	for n != 0 {
		d := n % 10
		if d%2 != 0 && d < 6 {
			p *= d
		}
		n = int64(n / 10)
	}
	fmt.Println(p)
}

package main

/*
	The function f is defined via segments like so:
	f(x) = {
		-1 when x <= -2
		1 / (3 + (2 / x)) when 0 < x <= 100
		200 when x >= 105
		0 else.
	}
	Write a program that prints out the value of f for an entered value of x.
*/

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("Enter the value of x:")
	fmt.Scan(&x)

	if x <= -2 {
		fmt.Println(1)
	} else if x > 0 && x <= 100 {
		result := 1.0 / (3.0 + (2.0 / float64(x)))
		fmt.Println(result)
	} else if x >= 105 {
		fmt.Println(200)
	} else {
		fmt.Println(0)
	}
}

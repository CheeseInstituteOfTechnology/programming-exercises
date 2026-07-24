package main

/*
	The user enters natural numbers m and n, one after the other. The program outputs a matrix with m rows and n columns
	in the way that's represented in the next two examples.
	Input:
	4
	7
	Output:
	1 1 1 1 1 1 1
	1 2 1 2 1 2 1
	1 2 3 1 2 3 1
	1 2 3 4 1 2 3

	Input:
	6
	4
	1 1 1 1
	1 2 1 2
	1 2 3 1
	1 2 3 4
	1 2 3 4
	1 2 3 4
*/

import (
	"fmt"
)

func main() {
	var m, n int

	fmt.Println("Enter the number m:")
	fmt.Scan(&m)
	fmt.Println("Enter the number n:")
	fmt.Scan(&n)

	for i := 1; i <= m; i++ {
		output := 1
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", output)
			output++
			if output > i {
				output = 1
			}
		}
		fmt.Println("")
	}
}

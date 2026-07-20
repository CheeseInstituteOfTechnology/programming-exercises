package main

/*
	The user enters a positive integer n, then n integer numbers. The program then has to output lines, whose length
	is determined by the entered n numbers. The first line is represented by -, the second by *, the third by -, and
	so on.
*/

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Enter the number n:")
	fmt.Scan(&n)

	var num int
	nums := []int{}
	fmt.Printf("Enter %d numbers:\n", n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	indicator := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < nums[i]; j++ {
			if indicator%2 != 0 {
				fmt.Print("-")
			} else {
				fmt.Print("*")
			}
		}
		indicator++
		fmt.Println("")
	}
}

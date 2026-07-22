package main

/*
	The user enters positive numbers, one after the other. The end of the programme is marked when the user enters -1.
	if the entered number is even, then the programme outputs its negative value, else it outputs the original number.
*/

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter numbers. To end, input -1.")
	nums := []int{}
	for {
		fmt.Scan(&n)
		if n == -1 || n < 1 {
			break
		}

		if n%2 == 0 {
			nums = append(nums, -n)
		} else {
			nums = append(nums, n)
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

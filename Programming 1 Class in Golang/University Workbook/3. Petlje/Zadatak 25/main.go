package main

/*
	The user enters natural numbers one after the other. The end of the input is marked by a blank line. The program
	outputs the product of the second and the second-to-last entered number. It can be assumed that the user will enter
	at least two numbers.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Println("Enter numbers. The end is marked by a blank line.")

	nums := []uint64{}
	for {
		_, err := fmt.Scanln(&n)

		if err != nil {
			break
		}

		if num, err := strconv.ParseUint(n, 10, 64); err == nil {
			nums = append(nums, num)
		}
	}

	fmt.Println(nums[1] * nums[len(nums)-2])
}

package main

/*
	The programme that asks the user to enter whole numbers, each in a new line. The end of the input is marked by a
	blank line. The programme outputs the negated value of the numbers. The numbers are also outputted on the same line,
	with a space in between them.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Println("Enter numbers. The end is marked by a blank line.")
	nums := []int64{}
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}

		if num, err := strconv.ParseInt(n, 0, 64); err == nil {
			nums = append(nums, -num)
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println("")
}

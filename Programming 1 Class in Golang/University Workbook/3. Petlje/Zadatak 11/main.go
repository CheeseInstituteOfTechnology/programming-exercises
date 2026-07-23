package main

/*
	The user enters a big number n. The program outputs the longest interval where all the digits are greater than four.
	Example:
	928632978158
	All intervals where all the digits are greater than four:
	9, 86, 978, 58. The longest interval is 978, therefore the output is 3.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64
	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	intervals := []string{}
	var interval string = ""
	for n != 0 {
		digit := n % 10
		if digit > 4 {
			interval += strconv.FormatInt(digit, 10)
		} else {
			if len(interval) != 0 {
				fmt.Println(interval)
				intervals = append(intervals, interval)
			}
			interval = ""
		}
		n = int64(n / 10)
	}

	if len(interval) != 0 {
		fmt.Println(interval)
		intervals = append(intervals, interval)
	}

	largest := intervals[0]
	for i := 0; i < len(intervals); i++ {
		if len(largest) < len(intervals[i]) {
			largest = intervals[i]
		}
	}

	fmt.Println(len(largest))
}

package main

/*
	While writing larger numbers, it is possible to improve their clearity, so that every trhee digits we separate them
	by a dot. Write a program that requires the user to enter different values, each in a new line, and the program
	outputs those values, but so that the digits are separated with a dot like how it was previously explained.
	It can be assumed that thye user will always enter positive integer numbers, and the end of the input is marked by
	a -1.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64

	fmt.Println("Input numbers. The end is marked by -1.")
	nums := []string{}
	for {
		fmt.Scan(&n)
		if n == -1 || n < 0 {
			break
		}

		if n <= 99 {
			nums = append(nums, strconv.FormatInt(n, 10))
		} else {
			nCopy := n
			c := 0
			str := ""
			for nCopy != 0 {
				d := nCopy % 10
				if c == 3 {
					str += "."
					c = 0
				}
				str += strconv.FormatInt(d, 10)
				c++
				nCopy = int64(nCopy / 10)
			}

			strR := ""
			for i := len(str) - 1; i >= 0; i-- {
				strR += string(str[i])
			}
			nums = append(nums, strR)
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

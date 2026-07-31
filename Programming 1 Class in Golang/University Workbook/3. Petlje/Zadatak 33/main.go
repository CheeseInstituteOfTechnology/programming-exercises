package main

/*
	Write a program that requires the user to input the number n (that's not greater than 10). The program outputs the multiplication
	table of n x n numbers. The table must be formatted as well.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n uint

	fmt.Println("n:")
	fmt.Scan(&n)

	if n > 10 {
		fmt.Println("The number must be from 1-10!")
	} else {
		square := n * n
		length := len(strconv.FormatUint(uint64(square), 10))
		for i := uint(0); i <= n; i++ {
			if i == 0 {
				output := ""
				for j := 0; j < length; j++ {
					output += " "
				}
				fmt.Printf("%s", output)
			} else {
				output := ""
				currentLength := len(strconv.FormatUint(uint64(i), 10))
				minLength := length - currentLength
				for j := 0; j < minLength; j++ {
					output += " "
				}
				output += strconv.FormatUint(uint64(i), 10)
				fmt.Printf("%s", output)
			}
		}
		fmt.Println("")

		nums := uint(1)
		for i := uint(0); i < n; i++ {
			for j := uint(0); j <= n; j++ {
				if j == 0 {
					output := ""
					currentLength := len(strconv.FormatUint(uint64(nums), 10))
					minLength := length - currentLength
					for k := 0; k < minLength; k++ {
						output += " "
					}
					output += strconv.FormatUint(uint64(nums), 10)
					fmt.Printf("%s", output)
				} else {
					output := ""
					currentLength := len(strconv.FormatUint(uint64(j*nums), 10))
					minLength := length - currentLength
					for k := 0; k < minLength; k++ {
						output += " "
					}
					output += strconv.FormatUint(uint64(j*nums), 10)
					fmt.Printf("%s", output)
				}
			}
			fmt.Println("")
			nums++
		}
	}
}

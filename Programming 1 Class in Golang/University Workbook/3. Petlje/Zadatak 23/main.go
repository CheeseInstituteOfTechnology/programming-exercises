package main

/*
	The program asks the user to enter the numbers n and m. The program after that outputs all numbers from 1 to n, but
	so that it outputs m numbers in one row. The numbers are separated by a blank line, but during the output of one
	digit numbers the program adds one blank space in front of the number so they can all be aligned perfectly.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m int

	fmt.Println("n:")
	fmt.Scan(&n)
	fmt.Println("m:")
	fmt.Scan(&m)

	c := 1
	l := strconv.FormatInt(int64(n), 10)
	for i := 1; i <= n; i++ {
		str := strconv.FormatInt(int64(i), 10)
		n := ""
		if len(l) > len(str) {
			loop := len(l) - len(str)
			for range loop {
				n += " "
			}
		}
		n += str
		fmt.Printf("%s ", n)
		if c == m {
			fmt.Println("")
			c = 0
		}
		c++
	}
}

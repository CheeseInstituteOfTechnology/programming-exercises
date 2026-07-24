package main

/*
	Write a program that asks the user to enter an integer number. The program checks if the digits of the entered
	number are in a monotone decreasing order from the left to the right.
	In other words, the program checks if an >= an+1 is true, where a represents one digit of a number, while the n
	is the digit's position. If the digits of the entered number are monotonously decreasing, the program outputs "yes",
	otherwise "no".
*/

import (
	"fmt"
)

func main() {
	var n int64

	fmt.Println("Please input a number:")
	fmt.Scan(&n)

	digits := []int64{}
	for n != 0 {
		d := n % 10
		digits = append(digits, d)
		n = int64(n / 10)
	}

	order := true
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			fmt.Println("No.")
			order = false
			break
		}
	}

	if order {
		fmt.Println("Yes.")
	}
}

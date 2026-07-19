package main

/*
	Write a program that asks the user to enter a three digit number. The program checks if the entered number is a
	three digit number, and if it is, it calculates and prints out the sum of its digits. Otherwise, print out the
	message "Invalid number!"
*/

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Enter a three digit number:")
	fmt.Scan(&n)

	if n >= 100 && n <= 999 {
		n1 := int(n / 100)
		n2 := int((n / 10) % 10)
		n3 := n % 10
		fmt.Println(n1 + n2 + n3)
	} else {
		fmt.Println("Invalid number!")
	}
}

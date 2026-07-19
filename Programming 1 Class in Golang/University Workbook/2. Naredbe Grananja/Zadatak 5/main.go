package main

/*
	Write a program that asks the user to enter a five digit number. The program checks if the enter number is five
	digits long, and if it is, it calculates and prints out the difference between its largest and smallest digit.
	Otherwise, it prints out "Wrong number!"
*/

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Enter a five digit number:")
	fmt.Scan(&n)

	if n >= 10000 && n <= 99999 {
		numCopy := n
		digitsSlice := []int{}
		for numCopy != 0 {
			digit := numCopy % 10
			digitsSlice = append(digitsSlice, digit)
			numCopy /= 10
		}
		lowest, highest := digitsSlice[0], 0
		for i := 0; i < len(digitsSlice); i++ {
			if lowest > digitsSlice[i] {
				lowest = digitsSlice[i]
			}

			if highest < digitsSlice[i] {
				highest = digitsSlice[i]
			}
		}
		fmt.Println(highest - lowest)
	} else {
		fmt.Println("Wrong number!")
	}
}

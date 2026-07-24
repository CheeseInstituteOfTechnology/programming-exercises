package main

/*
	Write a program that asks the user to input the number n. The program draws out a rectangle with the dimensions of
	n rows and 2n + 1 columns, in a format that's shown in the examples below:

	n = 2:
	xxxxx
	x0x0x

	n = 5:
	xxxxxxxxxxx
	x0x0x0x0x0x
	xxxxxxxxxxx
	x0x0x0x0x0x
	xxxxxxxxxxx
*/

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Enter the number n:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		iter := 1
		for j := 1; j <= 2*n+1; j++ {
			if i%2 != 0 {
				fmt.Print("x")
			}
			if i%2 == 0 {
				if iter == 1 {
					fmt.Print("x")
				} else {
					fmt.Print("0")
					iter = 0
				}
				iter++
			}
		}
		fmt.Println("")
	}
}

package main

/*
	Write a program that outputs all three digit numbers that meet this condition:
	abc = a^3 + b^3 + c^3
	where a, b, and c represent the digits of the thre digit number. The program outputs numbers from the lowest to
	highest in the same row, with space in between.
*/

import (
	"fmt"
)

func main() {
	for i := 100; i <= 999; i++ {
		s := int(i / 100)
		d := int((i / 10) % 10)
		j := i % 10
		if i == (s*s*s)+(d*d*d)+(j*j*j) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}

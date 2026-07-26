package main

/*
	For some number n we say that it's perfect if it is equal to the sum of its positive divisors that are less than n.
	For example, 28 is a perfect number: its divisors are 1, 2, 4, 7, and 14, and 1 + 2 + 4 + 7 + 14 = 28.
	Write a program that requires the user to enter integers a and b, and then outputs all perfect numbers in the range
	of a and b, one after the other.
*/

import (
	"fmt"
)

func main() {
	var a, b int64

	fmt.Println("Enter range of a:")
	fmt.Scan(&a)
	fmt.Println("Enter range of b:")
	fmt.Scan(&b)

	perfect := []int64{}
	for i := a; i <= b; i++ {
		s := int64(0)
		for j := int64(1); j < i; j++ {
			if i%j == 0 {
				s += j
			}
		}
		if s == i {
			perfect = append(perfect, i)
		}
	}

	for i := 0; i < len(perfect); i++ {
		fmt.Println(perfect[i])
	}
}

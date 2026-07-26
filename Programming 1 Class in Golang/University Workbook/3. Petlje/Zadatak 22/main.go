package main

/*
	Write a program that finds all two-digit numbers whose three digit square has the last two digits as the original
	two digit number. That is, if A, B, C represent the digits, then the two digit number is in the format of AB,
	it's necessary to calculate AB * AB so that AB * AB = CAB.

	(Example:)
	number = 25
	25 * 25 = 625
	does 625 have the last two digits as 25? yes, so we output it to the console.
*/

import (
	"fmt"
)

func main() {
	for i := 10; i <= 99; i++ {
		s := int64(i * i)
		if s >= 100 && s <= 999 {
			d1 := int64(s % 10)
			d2 := int64((s/10)%10) * 10
			if d1+d2 == int64(i) {
				fmt.Printf("%d\n%d\n\n", i, s)
			}
		}
	}
}

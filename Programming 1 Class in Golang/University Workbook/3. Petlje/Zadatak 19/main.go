package main

/*
	For any integer number 0 <= n <= 10000 that's not divisible by 2 or 5, it's possible to find a product with a
	different number where each digit is 1. The program has to find how many digits is in the lowest product n.
	The program takes the number n from the user, and outputs the amount of digits in the lowest product n where each
	digit is 1. For an example, if the user inputs 3, the output will be 3, as if 3 is multiplied by 37, the result is 111,
	and because the number has three ones, the program outputs 3.
*/

import (
	"fmt"
)

func main() {
	var n int64

	fmt.Println("Input number n:")
	fmt.Scan(&n)

	var i int64 = 1
	for {
		p := n * i
		pCopy := p
		hasOne := true
		c := 0
		for pCopy != 0 {
			d := pCopy % 10
			if d != 1 {
				hasOne = false
				break
			}
			c++
			pCopy = int64(pCopy / 10)
		}
		if hasOne {
			fmt.Println(c)
			break
		}
		i++
	}
}

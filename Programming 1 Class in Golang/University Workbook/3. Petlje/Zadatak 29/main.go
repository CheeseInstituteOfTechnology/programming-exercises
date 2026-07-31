package main

/*
	The user enters two numbers, one after the other, where the first one is longer (has more digits). The program outputs the position
	of the second number that's inside the first number. For numbers 123456 and 234 the answer is 1 because the numbers 234 start
	appearing on the second position of the numbner 123456 (the indexing starts at 0, so the 2nd position has the index of 1). If the
	smaller number does not appear inside the bigger on then the program outputs -1. So for numbers 987654 and 9887 the output is -1.
	[IMPORTANT! While doing the exercise it's important to treat the values as integers, and never as strings.]
*/

import (
	"fmt"
)

func main() {
	var n1, n2 uint64

	fmt.Println("n1:")
	fmt.Scan(&n1)
	fmt.Println("n2:")
	fmt.Scan(&n2)

	smallSlice := []uint64{}
	for n2 != 0 {
		d := n2 % 10
		smallSlice = append(smallSlice, d)
		n2 = uint64(n2 / 10)
	}

	i := uint64(0)
	found := false
	n1C := n1
	c := 1
	for n1 != 0 {
		d := n1 % 10
		if d == smallSlice[i] {
			if i == uint64(len(smallSlice)-1) {
				found = true
				break
			}
			i++
		} else {
			i = 0
		}
		c++
		n1 = uint64(n1 / 10)
	}

	n1Slice := []uint64{}
	for n1C != 0 {
		d := n1C % 10
		n1Slice = append(n1Slice, d)
		n1C = uint64(n1C / 10)
	}

	if found {
		fmt.Println(len(n1Slice) - c)
	} else {
		fmt.Println(-1)
	}
}

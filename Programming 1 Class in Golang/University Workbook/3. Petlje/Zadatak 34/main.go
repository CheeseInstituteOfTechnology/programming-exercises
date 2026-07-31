package main

/*
	With the use of the next two rules it's possible to generate a sequence of numbers:
	n = n / 2 (when n is even)
	n = 3n + 1 (when n is odd)
	Using these rules and starting with 13 we get the next sequence:
	13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
	This sequence contains 10 elements. Every sequence ends when we get to number 1. The program has to find and output a number
	in the range from 1 to 1000 that, when it's used like the beginning of the sequence, generates a sequence with the largest number
	of elements.
*/

import (
	"fmt"
)

func main() {
	largest := uint64(0)
	currentOutput := uint64(0)
	for i := uint64(1); i <= 10000; i++ {
		current := i
		steps := uint64(1)
		for {
			if current%2 == 0 {
				current = uint64(current / 2)
			} else {
				current = uint64(3*current + 1)
			}
			steps++
			if current == 1 {
				break
			}
		}
		if steps > largest {
			largest = steps
			currentOutput = i
		}
	}

	fmt.Println(currentOutput, largest)
}

package main

/*
	The professor on the whiteboard wrote a natural number. Students one by one go in front of the whiteboard and instead of that number
	they write the sum of the squares of the previous written number, all until the number one appears on the whiteboard.
	If on the whiteboard the number 654 is written, the next number will be 77, and then 98, etc.
	The user enters a natural number n. The program outputs how many numbers will be written on the whiteboard. If the number one will
	never appear on the whiteboard, the program outputs -1.
*/

import (
	"fmt"
)

func main() {
	var n uint64

	fmt.Println("n:")
	fmt.Scan(&n)

	current := n
	i := 1
	for {
		currentC := current
		square := uint64(0)
		for currentC != 0 {
			d := currentC % 10
			square += d * d
			currentC = uint64(currentC / 10)
		}
		current = square
		if current == 1 {
			break
		}
		if current == 4 {
			i = -1
			break
		}
		i++
	}

	fmt.Println(i)
}

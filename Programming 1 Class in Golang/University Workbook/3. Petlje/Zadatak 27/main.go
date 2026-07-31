package main

/*
	On three horizontal strips made out of fields are placed three grasshoppers. Every grasshopper can be placed on any part of the
	strip of its field. The user enters three pairs of numbers. The first pair represent the starting position of that grasshopper,
	and the second pair is the number of fields that the grasshopper jumps over. Also, there is the last input that's required for the
	user to enter; and that is the field that represents the finish for all grasshoppers.

	The grasshoppers jump one by one, starting with the first one and they all are trying to get to the end of the field. The game ends
	as soon as one of the grasshoppers get to the finish line (either goes over it, or directly on the finish line)

	In the example below the first grasshopper is placed on field 3 and jumps over 2 fields, the second grasshopper is placed on field 1
	and jumps over 4 fields, and the third grasshopper is placed on field 6 and jumps over one field.

	The program has to draw the situation of the end of the race. The fields are marked by '-', the grasshoppers by '*', and the finish
	line by '|'. If the grasshopper is standing on the finish line, the mark for the finish line ('|') is replaced with '*'. The length
	of the fields are determined by the winning grasshopper's position. The fields begin with the number 1, not 0.
*/

import (
	"fmt"
)

func main() {
	var g1, g2, g3 uint
	var j1, j2, j3 uint
	var f uint

	fmt.Println("Starting position of g1:")
	fmt.Scan(&g1)
	fmt.Println("Number of jumps of g1:")
	fmt.Scan(&j1)
	fmt.Println("Starting position of g2:")
	fmt.Scan(&g2)
	fmt.Println("Number of jumps of g2:")
	fmt.Scan(&j2)
	fmt.Println("Starting position of g3:")
	fmt.Scan(&g3)
	fmt.Println("Number of jumps of g3:")
	fmt.Scan(&j3)
	fmt.Println("Position of the finish line:")
	fmt.Scan(&f)

	w := uint(0)
	for {
		g1 += j1
		if g1 >= f {
			w = g1
			break
		}
		g2 += j2
		if g2 >= f {
			w = g2
			break
		}
		g3 += j3
		if g3 >= f {
			w = g3
			break
		}
	}

	for i := uint(1); i <= w; i++ {
		switch i {
		case f:
			fmt.Print("|")
		case g1:
			fmt.Print("*")
		default:
			fmt.Print("-")
		}
	}
	fmt.Println("")

	for i := uint(1); i <= w; i++ {
		switch i {
		case f:
			fmt.Print("|")
		case g2:
			fmt.Print("*")
		default:
			fmt.Print("-")
		}
	}
	fmt.Println("")

	for i := uint(1); i <= w; i++ {
		switch i {
		case f:
			fmt.Print("|")
		case g3:
			fmt.Print("*")
		default:
			fmt.Print("-")
		}
	}
	fmt.Println("")
}

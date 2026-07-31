package main

/*
	The program draws out stairs on the basis of two natural numbers that the user enters, one after the other. The first number
	represents the height of the stairs, while the second number represents the height at which they finish. The stairs are drawn with
	the use of the character '*'. The columns are seperated by a blank line.

	Example:
	7
	3
	Output:
	*
	* *
	* * *
	* * * *
	* * * * *
	* * * * *
	* * * * *
	Note: I have never been more confused by dotted lines ever. The height is 7, which is true, but the stairs end at height 4, not
	3. So I assume it's treated like endingHeight + 1. Also, the width of the rows are confusing, too, as it caps out at 5. 7 - 3 is 4,
	not 5. So I assume it's treated like maxHeight - endingHeight + 1? This is all the translated exercise says. Do I love the Bosnian
	education system.
*/

import (
	"fmt"
)

func main() {
	var n1, n2 uint

	fmt.Println("n1:")
	fmt.Scan(&n1)
	fmt.Println("n2:")
	fmt.Scan(&n2)

	w := n1 - n2 + 1
	o := 1
	for i := uint(0); i < n1; i++ {
		s := "*"
		oLocal := 0
		for j := uint(0); j < w; j++ {
			if i <= n2 {
				if oLocal >= o {
					s = " "
				}
			}
			fmt.Printf("%s ", s)
			oLocal++
		}
		o++
		fmt.Println("")
	}
}

package main

/*
	The user enters two integer numbers a and b. The program outputs every number in the range from a to b (including
	both values) that are divisible by 3 or those that contain at least one digit that's divisible by 3. Numbers
	are outputted from the smallest to the largest on one line, with space in between them.
	[Note: 0 is divisible by 3.]
*/

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Println("Enter a:")
	fmt.Scan(&a)
	fmt.Println("Enter b:")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		isDivisible := false
		if i%3 == 0 {
			fmt.Printf("%d ", i)
			isDivisible = true
		}

		if !isDivisible {
			copyNumber := i
			for copyNumber != 0 {
				digit := copyNumber % 10
				if digit%3 == 0 || digit == 0 {
					fmt.Printf("%d ", i)
					break
				}
				copyNumber = int(copyNumber / 10)
			}
		}
	}

	fmt.Println("")
}

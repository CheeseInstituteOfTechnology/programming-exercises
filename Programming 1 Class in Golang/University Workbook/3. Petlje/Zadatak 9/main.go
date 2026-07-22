package main

/*
	Write a programme where the user enters integer numbers, each on a new line. The end is marked by a blank line.
	The programme outputs "Appears" if the number 2 is present in the entered values. The programme outputs "Doesn't appear!
	if the number 2 is not present in the entered values.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Println("Enter numbers. The end is marked by a blank line.")
	twoAppears := false
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}

		if num, err := strconv.ParseInt(n, 0, 64); err == nil {
			if num == 2 {
				twoAppears = true
			}
		}
	}

	if twoAppears {
		fmt.Println("Appears!")
	} else {
		fmt.Println("Doesnt appear!")
	}
}

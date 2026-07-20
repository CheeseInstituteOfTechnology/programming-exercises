package main

/*
	The user enters real numbers one after the other. The end of the input is indicated by a blank line. The numbers
	belong to Rain and Link. The first number belongs to Rain, the second one to Link, the third to Rain, the fourth to
	Link, etc.

	The program calculates and prints out the difference of sums of Rain and Link's numbers. For example, if the
	sum of Rain's numbers is 19.14, and Link's 11.48, the difference is 7.66.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var n string

	fmt.Println("Enter numbers below. To end the input enter a blank line.")

	rainsSum, linksSum := 0.0, 0.0
	indicator := 1
	for true {
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}

		if n == " " {
			break
		}

		if num, err := strconv.ParseFloat(n, 64); err == nil {
			if indicator%2 != 0 {
				rainsSum += num
			} else {
				linksSum += num
			}
			indicator++
		}
	}

	fmt.Printf("Rain's sum: %f\nLink's sum: %f\n", rainsSum, linksSum)
	fmt.Println("The difference:")
	if rainsSum >= linksSum {
		fmt.Println(rainsSum - linksSum)
	} else {
		fmt.Println(linksSum - rainsSum)
	}
}

package main

/*
	Write a program that asks the user to enter the x and y coordinates of the centers of two circles in a
	two-dimensional plane. The coordinates should be entered on separate lines. The user also enters the radius of the
	first circle. The radius of the first circle is equal to the diameter of the second circle. Check whether the
	entered circles intersect and print one of the following three messages accordingly: "The circles intersect,"
	"The circles touch," or "The circles have no common points."
*/

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, R1 float64

	fmt.Println("Coordinates for (x1, y1)")
	fmt.Scan(&x1)
	fmt.Scan(&y1)

	fmt.Println("Coordinates for (x2, y2)")
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	fmt.Println("Radius of the first circle")
	fmt.Scan(&R1)

	R2 := R1 / 2.0

	d := math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)

	sumR := R1 + R2
	differenceR := R1 + R2

	sumSquare := sumR * sumR
	differenceSquare := differenceR * differenceR
	if d == sumSquare || d == differenceSquare {
		fmt.Println("The circles touch.")
	} else if d > differenceR && d < sumSquare {
		fmt.Println("The circles intersect.")
	} else {
		fmt.Println("The circles have no common points.")
	}
}

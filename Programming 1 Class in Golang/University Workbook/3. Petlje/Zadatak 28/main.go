package main

/*
	The user enters five positive integers one after the other: r, k, x1, y1, and d. The value of r represents the number of rows, and
	k the number of colums of a matrix in which a right-angled isosceles triangle is being drawn. The values x1 and y1 represent
	the coordinates of the upper-left point of the triangle, and the value d represents the length of the triangle's sides.

	The triangle is drawn so that all points in the matrix that belong to the triangle are marked by 'X', and all other points by '-'.
	In between the characters in the same row there's white space in between them.
*/

import (
	"fmt"
)

func main() {
	var r, k, x1, y1, d uint

	fmt.Println("rows:")
	fmt.Scan(&r)
	fmt.Println("cols:")
	fmt.Scan(&k)
	fmt.Println("x1:")
	fmt.Scan(&x1)
	fmt.Println("y1:")
	fmt.Scan(&y1)
	fmt.Println("d:")
	fmt.Scan(&d)

	for i := uint(0); i < r; i++ {
		found := false
		foundIncrease := false
		c := uint(1)
		for j := uint(0); j < k; j++ {
			if j == x1 && i == y1 {
				found = true
				foundIncrease = true
			}
			if foundIncrease && c <= d {
				fmt.Print(" X ")
				c++
			} else {
				fmt.Print(" - ")
			}
		}
		if found {
			x1++
			y1++
			d--
		}
		fmt.Println("")
	}
}

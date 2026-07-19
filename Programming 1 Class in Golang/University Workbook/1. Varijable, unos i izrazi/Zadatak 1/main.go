package main

import (
	"fmt"
	"math"
)

/*
	Write a program that asks the user to input the next four values: alfa, n, k, and h. Program calculates and
	prints out the following equation:
	((2*k^(2/3)) / (sin(alfa)*sqrt(h + n))) * cos(alfa)
*/

func main() {
	var alfa, n, k, h float64
	fmt.Println("Please input the next four values.")
	fmt.Println("")
	fmt.Println("Alfa: ")
	fmt.Scan(&alfa)
	fmt.Println("n: ")
	fmt.Scan(&n)
	fmt.Println("k: ")
	fmt.Scan(&k)
	fmt.Println("h: ")
	fmt.Scan(&h)

	toRadians := alfa * (math.Pi / 180.0)
	result := ((2 * math.Pow(k, 2.0/3.0)) / (math.Sin(toRadians) * math.Sqrt(h+n))) * math.Cos(toRadians)
	fmt.Println(result)
}

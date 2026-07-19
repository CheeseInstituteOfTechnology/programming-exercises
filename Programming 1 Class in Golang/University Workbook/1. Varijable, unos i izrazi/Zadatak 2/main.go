package main

/*
	In geometry, the entered circle of a triangle is the largest circle that the triangle has. The circle touches
	all three sides of the triangle. The radius of the netered circle r is calculated like so:

	r = c * ((sin(alfa / 2) * sin(beta / 2)) / cos(gamma / 2))

	Write a program that asks the user to enter the values of alfa, beta, and the length of side c, and then calculates
	the value of the radius r.
*/

import (
	"fmt"
	"math"
)

func main() {
	var alfa, beta, c float64
	fmt.Println("Enter the value of the next three variables:")
	fmt.Println("")
	fmt.Println("Alfa:")
	fmt.Scan(&alfa)
	fmt.Println("Beta:")
	fmt.Scan(&beta)
	fmt.Println("Side c: ")
	fmt.Scan(&c)

	gamma := 180 - alfa - beta

	radiansAlfa := alfa * (math.Pi / 180.0)
	radiansBeta := beta * (math.Pi / 180.0)
	radiansGamma := gamma * (math.Pi / 180.0)

	result := c * ((math.Sin(radiansAlfa/2.0) * math.Sin(radiansBeta/2.0)) / math.Cos(radiansGamma/2.0))
	fmt.Println(result)
}

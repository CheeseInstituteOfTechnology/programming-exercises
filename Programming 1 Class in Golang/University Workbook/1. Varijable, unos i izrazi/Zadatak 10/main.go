package main

/*
	An ultrasonic flow meter measures fluid velocity using ultrasound to determine flow volume. Using the upstream transit
	time (tu) the distance between the upstream and downstream transducers (L), the inclination angle (theta), and the
	average fluid velocity (Va), it is possible to calculate the downstream transit time (td) using the formula:

	td = 1 / ((2Vacos(theta) / L) + (1 / tu))

	Write a program that asks the user to enter the values of Va, theta, L, and tu, and calculates and prints out the
	value of td.
*/

import (
	"fmt"
	"math"
)

func main() {
	var Va, theta, L, tu float64

	fmt.Println("Va:")
	fmt.Scan(&Va)

	fmt.Println("Theta:")
	fmt.Scan(&theta)

	fmt.Println("K:")
	fmt.Scan(&L)

	fmt.Println("tu:")
	fmt.Scan(&tu)

	radiansTheta := theta * (math.Pi / 180.0)

	bottomHalf := ((2.0 * Va * math.Cos(radiansTheta)) / L) + (1.0 / tu)
	td := 1 / bottomHalf

	fmt.Println(td)
}

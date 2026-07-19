package main

/*
	When simulating projectile motion, it is necessary to calculate the object's initial velocity. If the object's
	instantaneous position (x, y) and the launch angle alfa are given, the initial velocity is calculated using the
	formula:
	v0 = sqrt((x^2g) / (xsin(2alfa) - 2ycos^2(alfa)))
	where g is the speed of Earth's gravity and is equal to 9.81.
	Write a program that asks the user to enter the values of x, y, and alfa, and then calculate the initial velocity
	using the formula above.
*/

import (
	"fmt"
	"math"
)

func main() {
	var x, y, alfa float64

	fmt.Println("Enter the value of x:")
	fmt.Scan(&x)

	fmt.Println("Enter the value of y:")
	fmt.Scan(&y)

	fmt.Println("Enter the value of alfa:")
	fmt.Scan(&alfa)

	radiansAlfa := alfa * (math.Pi / 180.0)
	g := 9.81

	v0 := math.Sqrt((x * x * g) / (x*math.Sin(2*radiansAlfa) - 2*y*math.Cos(radiansAlfa)*math.Cos(radiansAlfa)))
	fmt.Println(v0)
}

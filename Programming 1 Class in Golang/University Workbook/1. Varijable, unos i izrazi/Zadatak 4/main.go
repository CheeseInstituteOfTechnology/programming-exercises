package main

/*
	Write a program that takes the value of a temperature in Celcius, and converts it to Fahrenheit and prints the output.
	The formula for converting from Celcius to Fahrenheit is:

	F = C * (9 / 5) + 32
*/

import (
	"fmt"
)

func main() {
	var C float64

	fmt.Println("Enter the temperature in Celcius:")
	fmt.Scan(&C)

	F := C*(9.0/5.0) + 32
	fmt.Println(F)
}

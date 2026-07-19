package main

import (
	"fmt"
)

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return 2 * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

func main() {
	fmt.Println(RemainingOvenTime(30))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(3, 20))
}

package main

import (
	"fmt"
)

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	individualCars := carsCount % 10
	return uint(groupsOfTen*95000 + individualCars*10000)
}

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println(CalculateCost(37))
	fmt.Println(CalculateCost(21))
}

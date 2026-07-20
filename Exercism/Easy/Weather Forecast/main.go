// Package weather provides tools to tell the forecast of a city.
package main

var (
	// CurrentCondition is a string variable that tells the current weather condition of a place.
	CurrentCondition string
	// CurrentLocation is a string variable that tells the current location of a place to report the weather on.
	CurrentLocation string
)

// Forecast returns a string that describes the current weather condition of the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

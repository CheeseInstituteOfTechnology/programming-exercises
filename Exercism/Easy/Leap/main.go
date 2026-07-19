package main

import (
	"fmt"
)

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(IsLeapYear(1997))
	fmt.Println(IsLeapYear(1900))
	fmt.Println(IsLeapYear(2000))
}

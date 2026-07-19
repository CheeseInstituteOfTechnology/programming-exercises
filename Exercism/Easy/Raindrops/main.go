package main

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	result := ""
	isDivisible := false
	if number%3 == 0 {
		result += "Pling"
		isDivisible = true
	}
	if number%5 == 0 {
		result += "Plang"
		isDivisible = true
	}
	if number%7 == 0 {
		result += "Plong"
		isDivisible = true
	}
	if !isDivisible {
		result += strconv.FormatInt(int64(number), 10)
	}

	return result
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(Convert(i))
	}
}

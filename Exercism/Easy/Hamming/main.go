package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strings aren't equal length")
	}

	totalDistance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			totalDistance++
		}
	}

	return totalDistance, nil
}

func main() {
	fmt.Println(Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))
}

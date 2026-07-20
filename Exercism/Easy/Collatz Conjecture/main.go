package main

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Number must be positive!")
	}

	steps := 0
	for {
		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = int(n / 2)
		} else {
			n = n*3 + 1
		}
		steps++
	}

	return steps, nil
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(CollatzConjecture(i))
	}
}

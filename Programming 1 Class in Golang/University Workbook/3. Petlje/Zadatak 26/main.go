package main

/*
	The user enters the natural number n, then n natural numbers. The first number belongs to Rain, the next two belong
	to Link, then the next three belong to Rain, the next four belong to Link, and so on and so forth until all n numbers
	are given out to them. On the first line output the sum of all numbers that Rain got, then on the second line output
	the sum of all numbers that Link got.
*/

import (
	"fmt"
)

func main() {
	var n uint64

	fmt.Println("n:")
	fmt.Scan(&n)

	fmt.Printf("Enter %d numbers:\n", n)
	var n1 uint64
	nums := []uint64{}
	for i := uint64(1); i <= n; i++ {
		fmt.Scan(&n1)
		nums = append(nums, n1)
	}

	rainsSum := uint64(0)
	linksSum := uint64(0)

	o := 1
	c := 1
	iter := 0
	var order bool
	for i := 0; i < len(nums); i++ {
		if o%2 != 0 {
			order = true
		} else {
			order = false
		}

		if order {
			rainsSum += nums[i]
			iter++
			if iter == c {
				o = 2
				iter = 0
				c++
			}
		} else {
			linksSum += nums[i]
			iter++
			if iter == c {
				o = 1
				iter = 0
				c++

			}
		}
	}

	fmt.Printf("%d\n%d\n", rainsSum, linksSum)
}

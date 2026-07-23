package main

/*
	We call a natural number "fine" if it contains two digits back-to-back whose sum is equal to four.
	The user enters a natural number n. The programme outputs the nth fine natural number.
*/

import (
	"fmt"
)

func main() {
	var n int64

	fmt.Println("Enter a number n:")
	fmt.Scan(&n)

	fine := []int64{}
	var i int64 = 10

	for {
		copyI := i
		isFine := false

		for copyI >= 10 {
			d1 := copyI % 10
			copyI = int64(copyI / 10)
			d2 := copyI % 10
			if d1 != 0 && d2 != 0 && d1+d2 == 4 {
				isFine = true
				break
			}
		}

		if isFine {
			fine = append(fine, i)
			if int64(len(fine)) == n {
				fmt.Println(fine[n-1])
				break
			}
		}

		i++
	}
}

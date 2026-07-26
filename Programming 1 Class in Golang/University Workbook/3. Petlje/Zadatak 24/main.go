package main

/*
	We call a natural number "weird" if its digits only consist of 2 and 3. The user enters a natural number n. The program
	outputs the nth weird number.
*/

import (
	"fmt"
)

func main() {
	var n uint64

	fmt.Println("n:")
	fmt.Scan(&n)

	i := uint64(0)
	weird := []uint64{}
	for {
		i++
		copyI := i
		has := true
		for copyI != 0 {
			d := copyI % 10
			if d != 2 && d != 3 {
				has = false
				break
			}
			copyI = uint64(copyI / 10)
		}

		if has {
			weird = append(weird, i)
			if uint64(len(weird)) == n {
				fmt.Println(weird[n-1])
				break
			}
		}
	}

}

package main

/*
	Write a program that asks the user to input the total amount of points they gathered on the subject "Programming".
	The program prints out the given grade for that subject, based on the user's entered point. The grading system is
	like so:
	95 - 100: 10
	85 - 94: 9
	75 - 84: 8
	65 - 74: 7
	55 - 64: 6
	If the user enters points less than the ones in the table, they didn't pass the subject, and the grade is a 5.
*/

import (
	"fmt"
)

func main() {
	var points int
	fmt.Println("Enter the amount of points you got in the subject 'Programing':")
	fmt.Scan(&points)

	fmt.Print("The grade: ")
	if points >= 95 {
		fmt.Println(10)
	} else if points >= 85 {
		fmt.Println(9)
	} else if points >= 75 {
		fmt.Println(8)
	} else if points >= 65 {
		fmt.Println(7)
	} else if points >= 55 {
		fmt.Println(6)
	} else {
		fmt.Println(5)
	}
}

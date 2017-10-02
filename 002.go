package main

import (
	"fmt"
)

func two() {
	var termOne, termTwo, sum, temp int
	termOne = 1
	termTwo = 2
	sum = 2
	for true {
		temp = termOne + termTwo
		if temp > 4000000 {
			break
		}

		if temp%2 == 0 {
			sum = temp + sum
		}

		termOne = termTwo
		termTwo = temp
	}
	fmt.Println(sum)
}

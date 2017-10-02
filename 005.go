package main

import (
	"fmt"
)

func five() {
	var result int
	result = 0
	for i := 1; i <= 100000000; i++ {
		for n := 1; n <= 20; n++ {
			if i%n != 0 {
				break
			}
			result = i
		}
		if result == i {
			break
		}
	}
	fmt.Println(result)
}

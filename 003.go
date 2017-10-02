package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	i := 2
	for ; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func three() {
	n := 600851475143
	i := int(math.Sqrt(float64(n)))
	for ; i > 1; i-- {
		if n%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}

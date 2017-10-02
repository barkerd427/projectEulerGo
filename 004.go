package main

import (
	"fmt"
)

func four() {
	var first, second, product, reversed, largestPalindrome int
	second = 999
	largestPalindrome = 0

	for ; second >= 100; second-- {
		first = 999
		for ; first >= 100; first-- {
			product = first * second
			reversed = 0
			n := product
			for ; n > 0; n = n / 10 {
				reversed = 10*reversed + n%10
			}
			if product == reversed && product > largestPalindrome {
				largestPalindrome = product
				break
			}
		}
	}
	fmt.Println(largestPalindrome)
}

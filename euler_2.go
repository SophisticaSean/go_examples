// https://projecteuler.net/problem=2

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
	"fmt"
	"time"
)

func main() {
	total := 2
	current := 1
	next := 2
	fib := 0
	for fib < 4000000 {
		fib = current + next
		time.Sleep(20 * time.Millisecond)
		current = next
		next = fib
		if fib%2 == 0 {
			total += fib
			fmt.Println(fib)
		}
	}
	fmt.Println(total)
}

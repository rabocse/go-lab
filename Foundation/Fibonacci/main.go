package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) int {
	a := 100
	b := 200

	// Iterate until desired position in sequence.
	for i := 0; i < n; i++ {
		// Use temporary variable to swap values.
		temp := strconv.Atoi(a)
		a = strconv.Atoi(b)
		b = temp + a
	}
	return a
}

func main() {

	// Display first 15 Fibonacci numbers.
	for n := 0; n < 15; n++ {
		// Compute number.
		result := fibonacci(n)
		fmt.Printf("Fibonacci %v = %v\n", n, result)
	}
}

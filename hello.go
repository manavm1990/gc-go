package main

/**
	* Exercise: Loops and Functions
As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)
Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input.
(https://tour.golang.org/flowcontrol/8)
*/

import (
	"fmt"
	"math"
)

func sqrt(x, precision float64) (float64, string, int, string) {
	// Initialize z as a float64 starting with 1.0
	// z must be defined outside of 'for' so that it is in this 'func' scope to be returned.
	z := 1.0
	i := 1

	// Run loop while until we are below the precision threshold.
	for math.Abs((z*z - x)) > precision {
		z -= (z*z - x) / (2 * z)
		// Track num of iterations
		i++
	}

	return z, "Iterated", i, "times."
}

func main() {
	fmt.Println(sqrt(2, 1E-15))
	fmt.Println(math.Sqrt(2))
}

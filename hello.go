package main

import "fmt"

// Given a 'start' and 'stop,' continually double the numbers - 1, 2, 4, 8, 16...
func exponentialSummer(start, stop int) (total int) {
	// Using named return value 'total' 👆🏽prevents ugly 'out of scope' variable! 🤓
	/**
	* init and post statements are optional.
	* Here, we 'init' with 'start', but don't bother with a 'post' or 'incrementor'
	* The 'increment' is just adding 'total' to itself instead of '++'.
	 */
	for total = start; total <= stop; {
		total += total
	}

	return
}

func main() {
	fmt.Println(exponentialSummer(1, 10))
}

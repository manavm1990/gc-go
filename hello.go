package main

import "fmt"

func exponentialSummerFrom0(stop int) (total int) {
	// If just starting from 0 and incrementing by 1s, we can omit both 'init' and 'post'.
	// Behold! Ur beloved 'while' loop.
	for total <= stop {
		total++
	}

	return
}

func main() {
	fmt.Println(exponentialSummerFrom0(10))
}

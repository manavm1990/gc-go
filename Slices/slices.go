package main

import "fmt"

func main() {
	// Slices provide a dynamic 'subset' of an underlying, inflexible Array
	a := [6]int{1, 2, 4, 5, 6, 7}

	// Starts at index 1 👆🏽UP TO (not including) index 3.
	s := a[1:3]

	fmt.Println(a, s)
}

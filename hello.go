package main

import "fmt"

// Structs and Pointers
func main() {
	type coordinates struct {
		x int
		y int
	}

	// Cannot do pointers with 'const'! 🙅🏽‍♂️
	const y = 236

	x := 10
	pointerToX := &x

	test := coordinates{1, 2}
	pointerToTest := &test

	test.x = 5

	fmt.Println(test)
	fmt.Println(test.x)

	fmt.Println((*pointerToTest).x)

	// Syntactic sugar 👆🏽
	fmt.Println(pointerToTest.x)

	// Current 'state' of 'struct' for pointer to 'custom' type
	fmt.Printf("pointerToTest has a value of: %v and a type of: %T", pointerToTest, pointerToTest)

	// A memory address for pointer to 'built-in' type
	fmt.Println(pointerToX)
}

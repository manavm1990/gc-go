package main

import "fmt"

// Structs
func main() {
	type coordinates struct {
		x int
		y int
	}

	test := coordinates{1, 2}
	fmt.Println(test)
	fmt.Println(test.x)
}

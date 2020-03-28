package main

import "fmt"

func main() {
	// Stacked 'defer's
	defer fmt.Println("Goodbye!")
	defer fmt.Println("Goodbye2!")
	defer fmt.Println("Goodby3!")

	fmt.Println("Hello")
}

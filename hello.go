package main

import "fmt"

// Pointers
func main() {
	i := 3
	fmt.Println(i)

	// 'j' is of type 'int pointer' - a memory address
	j := &i
	fmt.Printf("j is of type: %T, and points to a value of: %v\n", j, *j)
	fmt.Printf("The actual value of j is: %v\n", j)

	// As 'i' changes, 'j' keeps up to date b/c it is 'pointing.'
	i++
	fmt.Println(i, *j)

	fmt.Printf("j is of type: %T, and points to a value of: %v\n", j, *j)
	fmt.Printf("The actual value of j is: %v\n", j)
}

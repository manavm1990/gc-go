package main

import "fmt"

func main() {
	// Declare length or Array and what type of data upfront
	// length is part of the static type and cannot be changed
	var a [1]string

	a[0] = "mutate"

	b := [1]string{"shortcut!"}

	fmt.Println(a, b)
}

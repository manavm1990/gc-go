package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func lastNameFirst(fname, lname string) (string, string) {
	return lname, fname
}

func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x

	// Implicity returns 'x' and 'y' as named returns 👆🏽.
	return
}

func main() {
	fmt.Println(add(2, 3))

	fmt.Println(lastNameFirst("Manav", "Misra"))

	fmt.Println(split(17))
}

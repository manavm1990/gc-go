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
	num1, num2 := 18, 22
	fname, lname := "Manav", "Misra"

	fmt.Println(add(num1, num2))

	fmt.Println(lastNameFirst(fname, lname))

	fmt.Println(split(num1))
}

package main

import (
	"fmt"
	"math/cmplx"
)

// Globals wrapped in 'var' block like 'import' 👆🏽
var (
	toBe bool = false

	// Globals allow code pollution! 👎🏽
	pollution string

	// (https://stackoverflow.com/a/60866548/1653236)
	maxValuePositiveInt uint64 = 1<<64 - 1

	// Imaginary numbers - advanced math stuff 😕
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func lastNameFirst(fname, lname string) (string, string) {
	return lname, fname
}

func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x

	// Implicitly returns 'x' and 'y' as named returns 👆🏽.
	return
}

func main() {
	num1, num2 := 18, 22
	fname, lname := "Manav", "Misra"

	fmt.Println(add(num1, num2), toBe, maxValuePositiveInt)

	fmt.Println(lastNameFirst(fname, lname))

	fmt.Println(split(num1))
}

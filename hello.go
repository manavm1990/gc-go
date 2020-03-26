package main

import "fmt"

func main() {
	i := 30
	fmt.Printf("%T\n", i)

	//
	f := float64(i)

	fmt.Printf("%T, %v\n", f, f)

	a := f

	fmt.Printf("%T, %v, %T, %v\n", f, f, a, a)
}

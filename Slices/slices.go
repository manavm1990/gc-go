package main

import "fmt"

func main() {
	// Slices provide a dynamic 'subset' of an underlying, inflexible Array
	a := [6]int{1, 2, 4, 5, 6, 7}

	// 'starting index : ending index' - doesn't include element at ending index.
	s := a[1:4]

	/**
	* The length of a slice is the number of elements it contains.
	* The capacity is the length of the underlying Array starting with lower bound index.
	 */
	fmt.Printf("s cap is: %v. s len is : %v\n", cap(s), len(s))
	fmt.Printf("s looks like: %v\n", s)

	// Mutate the Array via its slice
	// All slices referencing elements are updated b/c they are just proxies for the underlying Array subsets.
	// The index here must exist on the slice - if length is only 1, attempting to access index 1...🙅🏽‍♂️
	s[1] = 333

	// Be sure to stay in bounds of underlying Array.
	s = a[:6]
	fmt.Printf("Re-sliced s. Cap is: %v. Len is : %v\n", cap(s), len(s))

	// Default lower bound is '0'
	anotherSlice := a[:4]
	fmt.Printf("another s cap is: %v. s len is : %v\n", cap(anotherSlice), len(anotherSlice))
	fmt.Printf("anotherSlice looks like: %v\n", anotherSlice)

	// Array updating itself
	a[0] = 999
	fmt.Printf("anotherSlice looks like: %v\n", anotherSlice)
}

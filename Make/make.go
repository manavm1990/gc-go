package main

import "fmt"

func printSliceDeets(name string, l int, c int) {
	fmt.Printf("%v has length: %v and capacity of: %v\n", name, l, c)
}

func main() {
	/**
	* Create a slice with make w/o bothering to 'fill' an Array first.
	* make(type, length, optional capacity)
	 */
	mySlice := make([]int, 3, 10)

	mySlicedSlice := mySlice[:2]

	printSliceDeets("mySlice", len(mySlice), cap(mySlice))
	printSliceDeets("mySlicedSlice", len(mySlicedSlice), cap(mySlicedSlice))

}

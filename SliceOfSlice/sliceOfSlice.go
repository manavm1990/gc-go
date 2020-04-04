package main

import "fmt"

func main() {
	/**
	* [] - designates a slice (no length specified, so not an array)
	* []string - designates that 👆🏽slice will hold string slices.
	* {} - create an underlying array of string slices that have an underlying array of 'string underscores.'
	 */
	ticTacToe := [][]string{[]string{"1", "2", "3"}, []string{"4", "5", "6"}, []string{"7", "8", "9"}}

	// fmt.Printf("ticTacToe has length: %v and capacity of: %v\n", len(ticTacToe), cap(ticTacToe))
	fmt.Println(ticTacToe[0][2])
}

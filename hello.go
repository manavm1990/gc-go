package main

import "fmt"

// Given a 'start' and 'stop,' sum all of the numbers in between, inclusively
func numberSummer(start, stop int) (total int) {
	// Using named return value 'total' 👆🏽prevents ugly 'out of scope' variable! 🤓
	for i := start; i <= stop; i++ {
		total += i
	}

	return
}

func main() {
	fmt.Println(numberSummer(1, 10))
}

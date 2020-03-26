package main

import "fmt"

func main() {
	// ⚠️ Without type specified, defaults to 'int' and overflows.
	const x = 1 << 100

	fmt.Println(x)
}

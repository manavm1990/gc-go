package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	n, _ := fmt.Println("hello world!")
	fmt.Println(n, " bytes written at: ", time.Now())

	fmt.Println("And now, here's something random: ", rand.Intn(10))

	fmt.Println("How about some pi?", math.Pi)
}

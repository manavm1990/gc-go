package main

import (
	"fmt"
	"time"
)

func getGreeting() string {
	// Get current time
	t := time.Now()

	// Check the 'hour' to create greeting
	if t.Hour() < 12 {
		return "Mornin'"
	}
	if t.Hour() < 17 {
		return "Good afternoon"
	}
	return "Good evening!"
}

func main() {
	fmt.Println(getGreeting())
}

package main

import "fmt"

func card(age int) string {
	if age < 21 {
		return "No kids allowed!"
	}
	return "🙆🏽‍♂️, Let's party! 👯‍♂️"
}

func main() {
	fmt.Println(card(10))
}

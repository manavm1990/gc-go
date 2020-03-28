package main

import "fmt"

func card(age int) string {
	// We can optionally initialize a 'scoped' variable with our 'if' 😑
	if msg := "There's a 2 drink 🍹 minimum 2nite!"; age >= 21 {
		return msg + " 🙆🏽‍♂️, Let's party! 👯‍♂️"
	} else { // 'scoped' 'if' variable 👆🏽is available in 'else.'
		return "No kids allowed! But, tell your parents to come by and LTK: " + msg
	}

}

func main() {
	fmt.Println(card(10))
}

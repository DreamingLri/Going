package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(100) + 1
	var input int

	fmt.Println("Game Start!")
	for {
		fmt.Println("Please enter a number:")
		fmt.Scanln(&input)
		if input > num {
			fmt.Println("You are too high!")
		} else if input < num {
			fmt.Println("You are too low!")
		} else {
			fmt.Println("You are win!")
			break
		}
	}
	fmt.Println("Game End!")
}

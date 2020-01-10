package main

import (
	"fmt"
)

func xmove() {
	goodInput := false
	for goodInput == false {
		fmt.Println("Your move pick 1-9")
		fmt.Scanln(&x)
		switch {
		case x == 1 && grid[0] != "[X]" && grid[0] != "[O]":
			grid[0] = "[X]"
			goodInput = true

		case x == 2 && grid[1] != "[X]" && grid[1] != "[O]":
			grid[1] = "[X]"
			goodInput = true

		case x == 3 && grid[2] != "[X]\n" && grid[2] != "[O]\n":
			grid[2] = "[X]\n"
			goodInput = true

		case x == 4 && grid[3] != "[X]" && grid[3] != "[O]":
			grid[3] = "[X]"
			goodInput = true

		case x == 5 && grid[4] != "[X]" && grid[4] != "[O]":
			grid[4] = "[X]"
			goodInput = true

		case x == 6 && grid[5] != "[X]\n" && grid[5] != "[O]\n":
			grid[5] = "[X]\n"
			goodInput = true

		case x == 7 && grid[6] != "[X]" && grid[6] != "[O]":
			grid[6] = "[X]"
			goodInput = true

		case x == 8 && grid[7] != "[X]" && grid[7] != "[O]":
			grid[7] = "[X]"
			goodInput = true

		case x == 9 && grid[8] != "[X]\n" && grid[8] != "[O]\n":
			grid[8] = "[X]\n"
			goodInput = true

		default:
			fmt.Println("Error: The Number you entered has been chosen or you didn't enter a number.")

		}
	}
}

package main

import (
	"math/rand"
)

func aiOmove() {
	checkRepeat := true
	for checkRepeat == true {

		n = rand.Intn(9)
		if moves == 8 {
			break
		}

		switch {

		case n == 1 && grid[0] != "[X]" && grid[0] != "[O]":
			grid[0] = "[O]"
			checkRepeat = false
			moves++
		case n == 2 && grid[1] != "[X]" && grid[1] != "[O]":
			grid[1] = "[O]"
			checkRepeat = false
			moves++
		case n == 3 && grid[2] != "[X]\n" && grid[2] != "[O]\n":
			grid[2] = "[O]\n"
			checkRepeat = false
			moves++
		case n == 4 && grid[3] != "[X]" && grid[3] != "[O]":
			grid[3] = "[O]"
			checkRepeat = false
			moves++
		case n == 5 && grid[4] != "[X]" && grid[4] != "[O]":
			grid[4] = "[O]"
			checkRepeat = false
			moves++
		case n == 6 && grid[5] != "[X]\n" && grid[5] != "[O]\n":
			grid[5] = "[O]\n"
			checkRepeat = false
		case n == 7 && grid[6] != "[X]" && grid[6] != "[O]":
			grid[6] = "[O]"
			checkRepeat = false
			moves++
		case n == 8 && grid[7] != "[X]" && grid[7] != "[O]":
			grid[7] = "[O]"
			checkRepeat = false
			moves++
		case n == 9 && grid[8] != "[X]\n" && grid[8] != "[O]\n":
			grid[8] = "[O]\n"
			checkRepeat = false
			moves++
		default:
			checkRepeat = true

		}
	}
}

package main

func checkWin() {
	switch {
	// X Win Conditions
	case grid[0] == "[X]" && grid[1] == "[X]" && grid[2] == "[X]\n":
		wonChat = "X Wins!"
		win = true
	case grid[0] == "[X]" && grid[4] == "[X]" && grid[8] == "[X]\n":
		wonChat = "X Wins!"
		win = true
	case grid[0] == "[X]" && grid[3] == "[X]" && grid[6] == "[X]":
		wonChat = "X Wins!"
		win = true
	case grid[1] == "[X]" && grid[4] == "[X]" && grid[7] == "[X]":
		wonChat = "X Wins!"
		win = true
	case grid[2] == "[X]\n" && grid[4] == "[X]" && grid[6] == "[X]":
		wonChat = "X Wins!"
		win = true
	case grid[2] == "[X]\n" && grid[5] == "[X]\n" && grid[8] == "[X]\n":
		wonChat = "X Wins!"
		win = true
	case grid[3] == "[X]" && grid[4] == "[X]" && grid[5] == "[X]\n":
		wonChat = "X Wins!"
		win = true
	case grid[6] == "[X]" && grid[7] == "[X]" && grid[8] == "[X]\n":
		wonChat = "X Wins!"
		win = true

		// O win conditions
	case grid[0] == "[0]" && grid[1] == "[0]" && grid[2] == "[0]\n":
		wonChat = "O Wins!"
		win = true
	case grid[0] == "[0]" && grid[4] == "[0]" && grid[8] == "[0]\n":
		wonChat = "O Wins!"
		win = true
	case grid[0] == "[O]" && grid[3] == "[O]" && grid[6] == "[O]":
		wonChat = "O Wins!"
		win = true
	case grid[1] == "[O]" && grid[4] == "[O]" && grid[7] == "[O]":
		wonChat = "O Wins!"
		win = true
	case grid[2] == "[O]\n" && grid[4] == "[O]" && grid[6] == "[O]":
		wonChat = "O Wins!"
		win = true
	case grid[2] == "[O]\n" && grid[5] == "[O]\n" && grid[8] == "[O]\n":
		wonChat = "O Wins!"
		win = true
	case grid[3] == "[O]" && grid[4] == "[O]" && grid[5] == "[O]\n":
		wonChat = "O Wins!"
		win = true
	case grid[6] == "[O]" && grid[7] == "[O]" && grid[8] == "[O]\n":
		wonChat = "O Wins!"
		win = true
	default:
		win = false
	}

}

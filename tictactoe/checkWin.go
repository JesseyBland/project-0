package main

func checkWin() {
	switch {
	// X Win Conditions
	case cells[0].slogic == "X" && cells[1].slogic == "X" && cells[2].slogic == "X":
		wonChat = "X Wins!"
		win = true
	case cells[0].slogic == "X" && cells[3].slogic == "X" && cells[6].slogic == "X":
		wonChat = "X Wins!"
		win = true
	case cells[0].slogic == "X" && cells[4].slogic == "X" && cells[8].slogic == "X":
		wonChat = "X Wins!"
		win = true
	case cells[1].slogic == "X" && cells[4].slogic == "X" && cells[7].slogic == "X":
		wonChat = "X Wins!"
		win = true
	case cells[2].slogic == "X" && cells[4].slogic == "X" && cells[6].slogic == "X":
		wonChat = "X Wins!"
		win = true
	case cells[2].slogic == "X" && cells[5].slogic == "X" && cells[8].slogic == "X":
		wonChat = "X Wins!"
		win = true
	case cells[3].slogic == "X" && cells[4].slogic == "X" && cells[5].slogic == "X":
		wonChat = "X Wins!"
		win = true
	case cells[6].slogic == "X" && cells[7].slogic == "X" && cells[8].slogic == "X":
		wonChat = "X Wins!"
		win = true

	// O Win Conditions
	case cells[0].slogic == "O" && cells[1].slogic == "O" && cells[2].slogic == "O":
		wonChat = "O Wins!"
		win = true
	case cells[0].slogic == "O" && cells[3].slogic == "O" && cells[6].slogic == "O":
		wonChat = "O Wins!"
		win = true
	case cells[0].slogic == "O" && cells[4].slogic == "O" && cells[8].slogic == "O":
		wonChat = "O Wins!"
		win = true
	case cells[1].slogic == "O" && cells[4].slogic == "O" && cells[7].slogic == "O":
		wonChat = "O Wins!"
		win = true
	case cells[2].slogic == "O" && cells[4].slogic == "O" && cells[6].slogic == "O":
		wonChat = "O Wins!"
		win = true
	case cells[2].slogic == "O" && cells[5].slogic == "O" && cells[8].slogic == "O":
		wonChat = "O Wins!"
		win = true
	case cells[3].slogic == "O" && cells[4].slogic == "O" && cells[5].slogic == "O":
		wonChat = "O Wins!"
		win = true
	case cells[6].slogic == "O" && cells[7].slogic == "O" && cells[8].slogic == "O":
		wonChat = "O Wins!"
		win = true
		//No Winner Stalemate changes the win variable but displays No winner chat.
	case moves == 9:
		win = true
	default:
		win = false
	}

}

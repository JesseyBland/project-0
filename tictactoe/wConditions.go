package main

func wConditions() {
	switch {

	// X Win Conditions
	case cells[0].slogic == "X" && cells[1].slogic == "X" && cells[2].slogic == "X" && win[0] == false:
		wonChat = "X Wins!"

		win[0] = true

	case cells[0].slogic == "X" && cells[3].slogic == "X" && cells[6].slogic == "X" && win[1] == false:
		wonChat = "X Wins!"

		win[1] = true

	case cells[0].slogic == "X" && cells[4].slogic == "X" && cells[8].slogic == "X" && win[2] == false:
		wonChat = "X Wins!"

		win[2] = true

	case cells[1].slogic == "X" && cells[4].slogic == "X" && cells[7].slogic == "X" && win[3] == false:
		wonChat = "X Wins!"

		win[3] = true

	case cells[2].slogic == "X" && cells[4].slogic == "X" && cells[6].slogic == "X" && win[4] == false:
		wonChat = "X Wins!"

		win[4] = true

	case cells[2].slogic == "X" && cells[5].slogic == "X" && cells[8].slogic == "X" && win[5] == false:
		wonChat = "X Wins!"

		win[5] = true

	case cells[3].slogic == "X" && cells[4].slogic == "X" && cells[5].slogic == "X" && win[6] == false:
		wonChat = "X Wins!"

		win[6] = true

	case cells[6].slogic == "X" && cells[7].slogic == "X" && cells[8].slogic == "X" && win[7] == false:
		wonChat = "X Wins!"
		win[7] = true
	// O Win Conditions

	case cells[0].slogic == "O" && cells[1].slogic == "O" && cells[2].slogic == "O" && win[0] == false:
		wonChat = "O Wins!"

		win[0] = true

	case cells[0].slogic == "O" && cells[3].slogic == "O" && cells[6].slogic == "O" && win[1] == false:
		wonChat = "O Wins!"

		win[1] = true

	case cells[0].slogic == "O" && cells[4].slogic == "O" && cells[8].slogic == "O" && win[2] == false:
		wonChat = "O Wins!"

		win[2] = true

	case cells[1].slogic == "O" && cells[4].slogic == "O" && cells[7].slogic == "O" && win[3] == false:
		wonChat = "O Wins!"

		win[3] = true

	case cells[2].slogic == "O" && cells[4].slogic == "O" && cells[6].slogic == "O" && win[4] == false:
		wonChat = "O Wins!"

		win[4] = true

	case cells[2].slogic == "O" && cells[5].slogic == "O" && cells[8].slogic == "O" && win[5] == false:
		wonChat = "O Wins!"

		win[5] = true

	case cells[3].slogic == "O" && cells[4].slogic == "O" && cells[5].slogic == "O" && win[6] == false:
		wonChat = "O Wins!"

		win[6] = true

	case cells[6].slogic == "O" && cells[7].slogic == "O" && cells[8].slogic == "O" && win[7] == false:
		wonChat = "O Wins!"
		win[7] = true

	case moves == 9:
		win[8] = true
	default:
		break
	}

}

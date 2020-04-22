// Package gamewin is a package that contains the games winning condiditons and functions
//that track win.
package main

//Win is an array that houses the win states.
var Win [17]bool

//WonChat is a string array which outputs a Win or Lose message
var WonChat string

//Moves is a int that tracks number of X and O moves.
var Moves int

func wConditions() {

	switch {

	// X Win Conditions
	case board[0].Slogic == "X" && board[1].Slogic == "X" && board[2].Slogic == "X" && Win[0] == false:
		WonChat = "X Wins!"

		Win[0] = true

	case board[0].Slogic == "X" && board[3].Slogic == "X" && board[6].Slogic == "X" && Win[1] == false:
		WonChat = "X Wins!"

		Win[1] = true

	case board[0].Slogic == "X" && board[4].Slogic == "X" && board[8].Slogic == "X" && Win[2] == false:
		WonChat = "X Wins!"

		Win[2] = true

	case board[1].Slogic == "X" && board[4].Slogic == "X" && board[7].Slogic == "X" && Win[3] == false:
		WonChat = "X Wins!"

		Win[3] = true

	case board[2].Slogic == "X" && board[4].Slogic == "X" && board[6].Slogic == "X" && Win[4] == false:
		WonChat = "X Wins!"

		Win[4] = true

	case board[2].Slogic == "X" && board[5].Slogic == "X" && board[8].Slogic == "X" && Win[5] == false:
		WonChat = "X Wins!"

		Win[5] = true

	case board[3].Slogic == "X" && board[4].Slogic == "X" && board[5].Slogic == "X" && Win[6] == false:
		WonChat = "X Wins!"

		Win[6] = true

	case board[6].Slogic == "X" && board[7].Slogic == "X" && board[8].Slogic == "X" && Win[7] == false:
		WonChat = "X Wins!"

		Win[7] = true

	// O Win Conditions

	case board[0].Slogic == "O" && board[1].Slogic == "O" && board[2].Slogic == "O" && Win[0] == false:
		WonChat = "O Wins!"

		Win[9] = true

	case board[0].Slogic == "O" && board[3].Slogic == "O" && board[6].Slogic == "O" && Win[1] == false:
		WonChat = "O Wins!"

		Win[10] = true

	case board[0].Slogic == "O" && board[4].Slogic == "O" && board[8].Slogic == "O" && Win[2] == false:
		WonChat = "O Wins!"

		Win[11] = true

	case board[1].Slogic == "O" && board[4].Slogic == "O" && board[7].Slogic == "O" && Win[3] == false:
		WonChat = "O Wins!"

		Win[12] = true

	case board[2].Slogic == "O" && board[4].Slogic == "O" && board[6].Slogic == "O" && Win[4] == false:
		WonChat = "O Wins!"

		Win[13] = true

	case board[2].Slogic == "O" && board[5].Slogic == "O" && board[8].Slogic == "O" && Win[5] == false:
		WonChat = "O Wins!"

		Win[14] = true

	case board[3].Slogic == "O" && board[4].Slogic == "O" && board[5].Slogic == "O" && Win[6] == false:
		WonChat = "O Wins!"

		Win[15] = true

	case board[6].Slogic == "O" && board[7].Slogic == "O" && board[8].Slogic == "O" && Win[7] == false:
		WonChat = "O Wins!"
		Win[16] = true
	case Moves == 9 && WonChat != "X Wins!":
		WonChat = "No winners"
		Win[8] = true

	default:
		break
	}

}

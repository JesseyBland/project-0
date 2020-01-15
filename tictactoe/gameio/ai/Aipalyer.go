package ai

import (
	"fmt"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
	"github.com/JesseyBland/project-0/tictactoe/gameio"
	"github.com/JesseyBland/project-0/tictactoe/gamewin"
)

//Aiplayer runs the game with ai controlling O.
func Aiplayer() {
	for gamewin.CheckWin() == false {
		gameboard.PrintBoard()
		gameio.Xmove()
		gamewin.CheckWin()
		if gamewin.CheckWin() == false {
			aiOmove()
			gamewin.CheckWin()
		}
	}
	gameboard.PrintBoard()
	fmt.Println(gamewin.WonChat)
}

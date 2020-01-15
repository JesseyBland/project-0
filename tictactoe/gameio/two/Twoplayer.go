// Package two houses the Two player option when -tp flag is used. X and O are user
//inputs.
package two

import (
	"fmt"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
	"github.com/JesseyBland/project-0/tictactoe/gameio"
	"github.com/JesseyBland/project-0/tictactoe/gamewin"
)

// Twoplayer this is the twopalyer routine
func Twoplayer() {
	for gamewin.CheckWin() == false {
		gameboard.PrintBoard()
		gameio.Xmove()
		gamewin.CheckWin()

		if gamewin.CheckWin() == false {
			gameboard.PrintBoard()
			twomove()
			gamewin.CheckWin()
		}
	}
	gameboard.PrintBoard()
	fmt.Println(gamewin.WonChat)
}

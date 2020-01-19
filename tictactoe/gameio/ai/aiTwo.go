package ai

import (
	"fmt"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
	"github.com/JesseyBland/project-0/tictactoe/gamewin"
)

//Aitwo runs the game with ai controlling both X and O.
func Aitwo() {
Reset:
	again := ""
	gamewin.CheckWin()
	for gamewin.CheckWin() == false {
		fmt.Print(gameboard.PrintBoard())
		aiXmove()
		gamewin.CheckWin()
		if gamewin.CheckWin() == false {
			fmt.Print(gameboard.PrintBoard())
			AiOmove()
			gamewin.CheckWin()
		}
	}
	fmt.Print(gameboard.PrintBoard())
	fmt.Println("**************")
	fmt.Println(gamewin.WonChat)
	fmt.Println("Play Again? y/n")

	fmt.Scan(&again)
	if again == "y" || again == "Y" || again == "Yes" || again == "yes" {

		gamewin.ResetWin()
		goto Reset

	}
}

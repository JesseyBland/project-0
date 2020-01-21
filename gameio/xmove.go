// Package gameio contains all the files regaurding player input and ai output
package gameio

import (
	"fmt"
	"os"
	"strings"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

// Xmove is the X turn in the game.
func Xmove() {
	goodInput := false
	x := 0
	for goodInput == false {
		user := os.Getenv("USER")

		fmt.Printf("%v's move pick 1-9\n", strings.Title(user))
		fmt.Scan(&x)

		for i := 0; i < len(gameboard.Board); i++ {
			if x == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "X"
				gameboard.Board[i].Fill = true
				gamewin.Moves++
				goodInput = true
				break
			}
		}

		if goodInput == false {
			fmt.Println("Error: The Number you entered has been chosen or you didn't enter a number.")
			fmt.Print(gameboard.PrintBoard())
		}
	}
}

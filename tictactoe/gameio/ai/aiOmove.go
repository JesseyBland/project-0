// Package ai contains the ai move inputs
package ai

import (
	"math/rand"
	"time"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
	"github.com/JesseyBland/project-0/tictactoe/gamewin"
)

func aiOmove() {
	checkRepeat := true
	for checkRepeat == true {
		rand.Seed(time.Now().UnixNano())

		n := rand.Intn(10)
		if gamewin.Moves == 8 {
			break
		}
		for i := 0; i < len(gameboard.Board); i++ {
			if n == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "O"
				gameboard.Board[i].Fill = true
				gamewin.Moves++
				checkRepeat = false
			}
		}
	}
}

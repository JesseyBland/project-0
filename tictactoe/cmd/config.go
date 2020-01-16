// Package main is the Command Directory housing the main function and configuration variables.
package main

import (
	"flag"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
)

var av, tp bool

func init() {
	gameboard.LoadCells("|", "|")
	flag.BoolVar(&tp, "tp", false, "Two Player Enabled")
	flag.BoolVar(&av, "av", false, "Ai vs Ai Enable")
	flag.Parse()

}

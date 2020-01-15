// Package main is the Command Directory housing the main function and configuration variables.
package main

import (
	"flag"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
)

var tp, tuse bool

func init() {
	gameboard.LoadCells("(", ")")

	flag.BoolVar(&tp, "tp", false, "Two Player Enabled")
	flag.Parse()

}

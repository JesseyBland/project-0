package main

import (
	"flag"
	"math/rand"
	"time"
)

var win [9]bool
var wonChat string
var x, gridSize, n, moves int
var tp, tuse bool

func init() {
	loadCells("(", ")")
	wonChat = "N0 WINNER!"
	moves = 0
	rand.Seed(time.Now().UnixNano())

	flag.BoolVar(&tp, "tp", false, "Two Player Enabled")
	flag.Parse()

}

package main

import (
	"flag"
	"math/rand"
	"time"
)

var grid [9]string
var wonChat string
var x, gridSize, n, moves int
var win, tp bool

func init() {
	grid = [9]string{"[1]", "[2]", "[3]\n", "[4]", "[5]", "[6]\n", "[7]", "[8]", "[9]\n"}
	win = false
	wonChat = "N0 WINNER!"
	rand.Seed(time.Now().UnixNano())
	flag.BoolVar(&tp, "tp", false, "Two Player Enabled")
	flag.Parse()
}

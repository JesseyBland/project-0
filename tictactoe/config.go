package main

import (
	"math/rand"
	"time"
)

var grid [9]string
var wonChat string
var x, gridSize, n int
var win bool

func init() {
	grid = [9]string{"[1]", "[2]", "[3]\n", "[4]", "[5]", "[6]\n", "[7]", "[8]", "[9]\n"}
	win = false
	rand.Seed(time.Now().UnixNano())
}

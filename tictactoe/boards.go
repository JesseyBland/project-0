//Creates a square Board filled with numbers 1 - 9 each number surrounded by a string.
//looking at cell 1 of 9: {1} : This is a struct with elements that define the left side { the right side } and
// the number 1, plus a bool variable saying if it is filled or empty.
package main

import (
	"fmt"
	"strconv"
)

var cells = [100]board{}
var scale = 3
var boardsize = scale * scale
var row = 0

type board struct {
	slogic string
	lGrid  string
	rGrid  string
	fill   bool
}

func (p board) String() string {
	return fmt.Sprintf("%v%v%v", p.lGrid, p.slogic, p.rGrid)
}
func loadCells(lGrid, rGrid string) {

	for i := 0; i < boardsize; i++ {
		offset := i + 1
		slogic := strconv.Itoa(offset)
		fill := false
		cells[i] = board{slogic, lGrid, rGrid, fill}
	}
}

// Prints my board.
func printCells() {
	for i := 1; i < boardsize+1; i++ {

		row++
		if row == scale {
			fmt.Printf("%v\n", cells[i-1])
			row = 0
		} else {
			fmt.Printf("%v", cells[i-1])
		}

	}

}

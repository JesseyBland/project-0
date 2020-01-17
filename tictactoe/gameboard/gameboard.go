// Package gameboard Creates a square Board filled with numbers 1 - 9 each number surrounded by a string.
//looking at cell 1 of 9: {1} : This is a struct with elements that define the left side { the right side } and
// the number 1, plus a bool variable saying if it is filled or empty.
package gameboard

import (
	"fmt"
	"strconv"
)

//Board is a of structured cells.
var Board = make([]Cells, boardsize)

//Scale is the size of the row of the Board.
var scale = 3

//Boardsize is how many cells.
var boardsize = scale * scale

//row exsists to add a counter so that when it reaches scale it makes a new line for the next row.
var row = 0

//Cells the Cell structure of my board
type Cells struct {
	Slogic string
	LGrid  string
	RGrid  string
	Fill   bool
}

func (p Cells) String() string {
	return fmt.Sprintf("%v%v%v", p.LGrid, p.Slogic, p.RGrid)
}

//LoadCells fills the Cells structs with the board values.
func LoadCells(LGrid, RGrid string) {

	for i := 0; i < boardsize; i++ {
		offset := i + 1
		Slogic := strconv.Itoa(offset)
		Fill := false
		Board[i] = Cells{Slogic, LGrid, RGrid, Fill}
	}
}

//PrintBoard loops my board and prints the resulting cells.
func PrintBoard() string {
	var sboard string
	for i := 1; i < boardsize+1; i++ {

		row++
		if row == scale {
			sboard += fmt.Sprintf("%v\n", Board[i-1])
			row = 0
		} else {
			sboard += fmt.Sprintf("%v", Board[i-1])
		}

	}
	return sboard
}

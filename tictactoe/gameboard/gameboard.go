// Package gameboard Creates a square Board filled with numbers 1 - 9 each number surrounded by a string.
//looking at cell 1 of 9: {1} : This is a struct with elements that define the left side { the right side } and
// the number 1, plus a bool variable saying if it is filled or empty.
package gameboard

import (
	"fmt"
	"strconv"
)

//Board is an Array of structured cells.
var Board = [9]Cells{}
var scale = 3
var boardsize = scale * scale
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
func PrintBoard() {
	for i := 1; i < boardsize+1; i++ {

		row++
		if row == scale {
			fmt.Printf("%v\n", Board[i-1])
			row = 0
		} else {
			fmt.Printf("%v", Board[i-1])
		}

	}

}

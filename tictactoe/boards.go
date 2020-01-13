//Creates a 9x9 Board filled with numbers 1 - 9 each number surrounded by a string.
//looking at cell 1 of 9: {1} : This is a struct with elements that define the left side { the right side } and
// the number 1, plus a bool variable saying if it is filled or empty.
package main

import (
	"fmt"
	"strconv"
)

var cells = [9]board{}

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

	//i, v := range cells
	for i := 0; i < len(cells); i++ {
		offset := i + 1
		slogic := strconv.Itoa(offset)
		fill := false
		cells[i] = board{slogic, lGrid, rGrid, fill}
	}
}
func printCells() {
	for i := 1; i < len(cells)+1; i++ {
		if i%3 == 0 {
			fmt.Printf("%v\n", cells[i-1])
		} else {
			fmt.Printf("%v", cells[i-1])
		}
	}

}

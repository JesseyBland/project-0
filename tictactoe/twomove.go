package main

import (
	"fmt"
)

func twomove() {
	goodInput := false
	x := 0
	for goodInput == false {

		fmt.Println("O's move pick 1-9")
		fmt.Scanln(&x)

		for i := 0; i < len(cells); i++ {
			if x == (i+1) && cells[i].fill == false {
				cells[i].slogic = "O"
				cells[i].fill = true
				goodInput = true
				moves++
			}
		}
		if goodInput == false {
			fmt.Println("Error: The Number you entered has been chosen or you didn't enter a number.")
			printCells()
		}
	}
}

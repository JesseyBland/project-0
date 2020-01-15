package main

import (
	"fmt"
)

func main() {

	if tp == true {
		for checkWin() == false {
			printCells()
			xmove()
			checkWin()

			if checkWin() == false {
				printCells()
				twomove()
				checkWin()
			}
		}
		printCells()
		fmt.Println(wonChat)

	} else {
		for checkWin() == false {

			printCells()
			xmove()
			checkWin()
			if checkWin() == false {
				aiOmove()
				checkWin()
			}

		}
		printCells()
		fmt.Println(wonChat)
	}
}

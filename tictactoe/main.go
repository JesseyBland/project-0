package main

import (
	"fmt"
)

func main() {
	if tp == true {
		for win == false {
			printCells()
			xmove()
			checkWin()

			if !win {
				printCells()
				twomove()
				checkWin()
			}
		}
		printCells()
		fmt.Println(wonChat)

	} else {
		for win == false {

			printCells()
			xmove()
			checkWin()
			if !win {
				aiOmove()
				checkWin()
			}

		}
		printCells()
		fmt.Println(wonChat)
	}
}

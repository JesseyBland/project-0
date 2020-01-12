package main

import "fmt"

func main() {
	if tp == true {
		for win == false {
			makeGrid()
			xmove()
			checkWin()

			if !win {
				makeGrid()
				twomove()
				checkWin()
			}
		}
		makeGrid()
		fmt.Println(wonChat)

	} else {
		for win == false {

			makeGrid()
			xmove()
			checkWin()
			if !win {
				aiOmove()
				checkWin()
			}

		}
		makeGrid()
		fmt.Println(wonChat)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func xmove() {
	goodInput := false
	for goodInput == false {
		user := os.Getenv("USER")

		fmt.Printf("%v's move pick 1-9\n", strings.Title(user))
		fmt.Scan(&x)

		for i := 0; i < len(cells); i++ {
			if x == (i+1) && cells[i].fill == false {
				cells[i].slogic = "X"
				cells[i].fill = true
				moves++
				goodInput = true
				break
			}
		}

		if goodInput == false {
			fmt.Println("Error: The Number you entered has been chosen or you didn't enter a number.")
			printCells()
		}
	}
}

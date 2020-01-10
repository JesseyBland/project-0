package main

import "fmt"

func makeGrid() { //Prints the 9x9 grid
	for i := 0; i < 9; i++ {
		fmt.Print(grid[i])
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var grid [9]string
var x int

func makeGrid() { //Prints the 9x9 grid
	for i := 0; i < 9; i++ {
		fmt.Print(grid[i])
	}
}

func init() {
	grid = [9]string{"[1]", "[2]", "[3]\n", "[4]", "[5]", "[6]\n", "[7]", "[8]", "[9]\n"}

}

func main() {

	for i := 0; i < 5; i++ {
		makeGrid()
		fmt.Println("Your move pick 1-9")
		fmt.Scan(&x)
		switch {
		case x == 1 && grid[0] != "[X]" && grid[0] != "[O]":
			grid[0] = "[X]"
		case x == 2 && grid[1] != "[X]" && grid[1] != "[O]":
			grid[1] = "[X]"
		case x == 3 && grid[2] != "[X]\n" && grid[2] != "[O]\n":
			grid[2] = "[X]\n"
		case x == 4 && grid[3] != "[X]" && grid[3] != "[O]":
			grid[3] = "[X]"
		case x == 5 && grid[4] != "[X]" && grid[4] != "[O]":
			grid[4] = "[X]"
		case x == 6 && grid[5] != "[X]\n" && grid[5] != "[O]\n":
			grid[5] = "[X]\n"
		case x == 7 && grid[6] != "[X]" && grid[6] != "[O]":
			grid[6] = "[X]"
		case x == 8 && grid[7] != "[X]" && grid[7] != "[O]":
			grid[7] = "[X]"
		case x == 9 && grid[8] != "[X]\n" && grid[8] != "[O]\n":
			grid[8] = "[X]\n"
		default:
			fmt.Println("Error: The Number you entered has been chosen or you didn't enter a number.")
			i--
			break

		}
		//while loop that holds the ai logic
		var y int = 0
		for y < 1 {
			// the ai is a random number generator that places its O on grid
			var t int64
			t = int64(time.Now().Second())
			rand.Seed(t)
			n := rand.Intn(10)
			//This checks to see if the game is finished 5 x moves only 4 circle moves
			if i > 3 {
				makeGrid()
				fmt.Println("You Finished!")
				y++
			} else {
				//Random number is filtered to 1-9 spots on grid depending
				//on if the grid is used already it loops back and trys a new random number
				switch {

				case n == 1 && grid[0] != "[X]" && grid[0] != "[O]":
					grid[0] = "[O]"
					y++
				case n == 2 && grid[1] != "[X]" && grid[1] != "[O]":
					grid[1] = "[O]"
					y++
				case n == 3 && grid[2] != "[X]\n" && grid[2] != "[O]\n":
					grid[2] = "[O]\n"
					y++
				case n == 4 && grid[3] != "[X]" && grid[3] != "[O]":
					grid[3] = "[O]"
					y++
				case n == 5 && grid[4] != "[X]" && grid[4] != "[O]":
					grid[4] = "[O]"
					y++
				case n == 6 && grid[5] != "[X]\n" && grid[5] != "[O]\n":
					grid[5] = "[O]\n"
					y++
				case n == 7 && grid[6] != "[X]" && grid[6] != "[O]":
					grid[6] = "[O]"
					y++
				case n == 8 && grid[7] != "[X]" && grid[7] != "[O]":
					grid[7] = "[O]"
					y++
				case n == 9 && grid[8] != "[X]\n" && grid[8] != "[O]\n":
					grid[8] = "[O]\n"
					y++

				}
			}

		}

	}

}

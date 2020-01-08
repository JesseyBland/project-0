package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//the tic tax toe grid array
	map1 := [9]string{"[1]", "[2]", "[3]", "[4]", "[5]", "[6]", "[7]", "[8]", "[9]"}
	var x int
	//Nested for loop first loop asks for x placement the second loop is the ai move
	for i := 0; i < 5; i++ {
		fmt.Println(map1[0], map1[1], map1[2])
		fmt.Println(map1[3], map1[4], map1[5])
		fmt.Println(map1[6], map1[7], map1[8])
		fmt.Println("Your move pick 1-9")
		fmt.Scan(&x)
		switch {
		case x == 1 && map1[0] != "[X]" && map1[0] != "[O]":
			map1[0] = "[X]"
		case x == 2 && map1[1] != "[X]" && map1[1] != "[O]":
			map1[1] = "[X]"
		case x == 3 && map1[2] != "[X]" && map1[2] != "[O]":
			map1[2] = "[X]"
		case x == 4 && map1[3] != "[X]" && map1[3] != "[O]":
			map1[3] = "[X]"
		case x == 5 && map1[4] != "[X]" && map1[4] != "[O]":
			map1[4] = "[X]"
		case x == 6 && map1[5] != "[X]" && map1[5] != "[O]":
			map1[5] = "[X]"
		case x == 7 && map1[6] != "[X]" && map1[6] != "[O]":
			map1[6] = "[X]"
		case x == 8 && map1[7] != "[X]" && map1[7] != "[O]":
			map1[7] = "[X]"
		case x == 9 && map1[8] != "[X]" && map1[8] != "[O]":
			map1[8] = "[X]"
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
				fmt.Println(map1[0], map1[1], map1[2])
				fmt.Println(map1[3], map1[4], map1[5])
				fmt.Println(map1[6], map1[7], map1[8])
				fmt.Println("You Finished!")
				y++
			} else {
				//Random number is filtered to 1-9 spots on grid depending
				//on if the grid is used already it loops back and trys a new random number
				switch {

				case n == 1 && map1[0] != "[X]" && map1[0] != "[O]":
					map1[0] = "[O]"
					y++
				case n == 2 && map1[1] != "[X]" && map1[1] != "[O]":
					map1[1] = "[O]"
					y++
				case n == 3 && map1[2] != "[X]" && map1[2] != "[O]":
					map1[2] = "[O]"
					y++
				case n == 4 && map1[3] != "[X]" && map1[3] != "[O]":
					map1[3] = "[O]"
					y++
				case n == 5 && map1[4] != "[X]" && map1[4] != "[O]":
					map1[4] = "[O]"
					y++
				case n == 6 && map1[5] != "[X]" && map1[5] != "[O]":
					map1[5] = "[O]"
					y++
				case n == 7 && map1[6] != "[X]" && map1[6] != "[O]":
					map1[6] = "[O]"
					y++
				case n == 8 && map1[7] != "[X]" && map1[7] != "[O]":
					map1[7] = "[O]"
					y++
				case n == 9 && map1[8] != "[X]" && map1[8] != "[O]":
					map1[8] = "[O]"
					y++

				}
			}

		}

	}

}

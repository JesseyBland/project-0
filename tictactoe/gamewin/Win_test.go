package gamewin

import (
	"fmt"
	"testing"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
)

//TestXwins tests all X Cases of win conditions in my CheckWin function

func TestXwins(t *testing.T) {
	for i := range gameboard.Board {
		gameboard.Board[i].Slogic = "X"
	}
	for range Win {
		CheckWin()
	}
	for i := 0; i < len(Win)-1; i++ {

		if Win[i] == false {
			t.Errorf("Error: The case %v X win condition statement is not working", i)
		} else {
			t.Logf("The case %v X win condition Passed!", i)
		}
	}
}

//TestOwins tests all O Cases of win conditions in my CheckWin and wCondtions function
func TestOwins(t *testing.T) {
	for i := range gameboard.Board {
		gameboard.Board[i].Slogic = "O"
	}
	for range Win {
		CheckWin()
	}
	for i := 0; i < len(Win)-1; i++ {

		if Win[i] == false {
			t.Errorf("Error: The case %v O win condition statement is not working", i)
		} else {
			t.Logf("The case %v O win condition Passed!", i)
		}
	}
}
func TestDraw(t *testing.T) {
	//Board DRAW
	drawBoard := [9]string{"O", "X", "O", "X", "O", "X", "X", "O", "X"}
	for i := range drawBoard {
		gameboard.Board[i].Slogic = drawBoard[i]
		Moves++
	}
	CheckWin()
	if Win[8] == true {
		t.Logf("Draw : PASSED")
	} else {
		t.Error("Error: Draw failed")
	}
}
func ExampleTestXwins() {
	drawBoard := [9]string{"X", "X", "O", "X", "X", "O", "O", "O", "X"}
	for i := range drawBoard {
		gameboard.Board[i].Slogic = drawBoard[i]
		Moves++
	}

	fmt.Println(CheckWin())
	//Output: true
}
func ExampleTestOwins() {
	drawBoard := [9]string{"X", "O", "X", "X", "O", "O", "7", "O", "X"}
	for i := range drawBoard {
		gameboard.Board[i].Slogic = drawBoard[i]
		Moves++
	}

	fmt.Println(CheckWin())
	//Output: true
}
func ExampleTestDraw() {
	drawBoard := [9]string{"O", "X", "O", "X", "O", "X", "X", "O", "X"}
	for i := range drawBoard {
		gameboard.Board[i].Slogic = drawBoard[i]
		Moves++
	}

	fmt.Println(CheckWin())

	// Output: true
}

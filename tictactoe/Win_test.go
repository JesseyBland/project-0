package main

import (
	"fmt"
	"testing"
)

//TestXwins tests all X Cases of win conditions in my checkWin function
func TestXwins(t *testing.T) {
	for i := range cells {
		cells[i].slogic = "X"
	}
	for range win {
		checkWin()
	}
	for i := 0; i < len(win)-1; i++ {

		if win[i] == false {
			t.Errorf("Error: The case %v X win condition statement is not working", i)
		} else {
			t.Logf("The case %v X win condition Passed!", i)
		}
	}
}

//TestOwins tests all O Cases of win conditions in my checkWin and wCondtions function
func TestOwins(t *testing.T) {
	for i := range cells {
		cells[i].slogic = "O"
	}
	for range win {
		checkWin()
	}
	for i := 0; i < len(win)-1; i++ {

		if win[i] == false {
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
		cells[i].slogic = drawBoard[i]
		moves++
	}
	checkWin()
	if win[8] == true {
		t.Logf("Draw : PASSED")
	} else {
		t.Error("Error: Draw failed")
	}
}
func ExampleTestXwins() {
	drawBoard := [9]string{"X", "X", "O", "X", "X", "O", "O", "O", "X"}
	for i := range drawBoard {
		cells[i].slogic = drawBoard[i]
		moves++
	}

	fmt.Println(checkWin())
	//Output: true
}
func ExampleTestOwins() {
	drawBoard := [9]string{"X", "O", "X", "X", "O", "O", "7", "O", "X"}
	for i := range drawBoard {
		cells[i].slogic = drawBoard[i]
		moves++
	}

	fmt.Println(checkWin())
	//Output: true
}
func ExampleTestDraw() {
	drawBoard := [9]string{"O", "X", "O", "X", "O", "X", "X", "O", "X"}
	for i := range drawBoard {
		cells[i].slogic = drawBoard[i]
		moves++
	}

	fmt.Println(checkWin())

	// Output: true
}

package main

import "testing"

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

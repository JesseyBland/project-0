package main

import (
	"math/rand"
)

func aiOmove() {
	checkRepeat := true
	for checkRepeat == true {

		n = rand.Intn(10)
		if moves == 8 {
			break
		}
		for i := 0; i < len(cells); i++ {
			if n == (i+1) && cells[i].fill == false {
				cells[i].slogic = "O"
				cells[i].fill = true
				moves++
				checkRepeat = false
			}
		}
	}
}

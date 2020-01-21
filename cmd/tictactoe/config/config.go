package config

import (
	"flag"

	"github.com/JesseyBland/project-0/gameboard"
)

//Av is the flag variable for AivsAi
var Av bool

//Tp is the flag variable for Two palyer
var Tp bool

func init() {
	gameboard.LoadCells("[", "]")
	flag.BoolVar(&Tp, "tp", false, "Two Player Enabled")
	flag.BoolVar(&Av, "av", false, "Ai vs Ai Enable")
	flag.Parse()

}

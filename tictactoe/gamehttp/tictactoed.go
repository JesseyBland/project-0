package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
)

//Grabs user number 1-9
func Number(w http.ResponseWriter, r *http.Request) {
	//get request method
	gameboard.LoadCells("|", "|")

	t, _ := template.ParseFiles("ttt1.html")

	t.Execute(w, nil)
	r.ParseForm()

	// logic part of log in

	fmt.Println("X moves to:", r.Form["quantity"])
	xread := strings.Join(r.Form["quantity"], "")

	x, _ := strconv.Atoi(xread)
	if x == 0 {
		fmt.Fprint(w, gameboard.PrintBoard())
	}
	for i := 0; i < len(gameboard.Board); i++ {
		if x == (i + 1) {

			gameboard.Board[i].Slogic = "X"
			fmt.Fprint(w, gameboard.PrintBoard())
		}
	}

}

func main() {

	http.HandleFunc("/ttt1", Number)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

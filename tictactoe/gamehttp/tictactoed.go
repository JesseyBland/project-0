package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/JesseyBland/project-0/tictactoe/gameboard"
	"github.com/JesseyBland/project-0/tictactoe/gameio/ai"
	"github.com/JesseyBland/project-0/tictactoe/gamewin"
)

//Grabs user number 1-9
func Number(w http.ResponseWriter, r *http.Request) {
	//get request method
	gameboard.LoadCells("[", "]")

	t, _ := template.ParseFiles("test.html")
	t2, _ := template.ParseFiles("ttt1.html")

	Cells := struct {
		One   string
		Two   string
		Three string
		Four  string
		Five  string
		Six   string
		Seven string
		Eight string
		Nine  string
	}{
		One:   fmt.Sprintf("%v\n", gameboard.Board[0]),
		Two:   fmt.Sprintf("%v\n", gameboard.Board[1]),
		Three: fmt.Sprintf("%v\n", gameboard.Board[2]),
		Four:  fmt.Sprintf("%v\n", gameboard.Board[3]),
		Five:  fmt.Sprintf("%v\n", gameboard.Board[4]),
		Six:   fmt.Sprintf("%v\n", gameboard.Board[5]),
		Seven: fmt.Sprintf("%v\n", gameboard.Board[6]),
		Eight: fmt.Sprintf("%v\n", gameboard.Board[7]),
		Nine:  fmt.Sprintf("%v\n", gameboard.Board[8]),
	}
	t.Execute(w, Cells)

	r.ParseForm()
	xread := strings.Join(r.Form["quantity"], "")
	x, _ := strconv.Atoi(xread)

	if xread != "" {
		for i := 0; i < len(gameboard.Board); i++ {
			if x == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "X"
				gameboard.Board[i].Fill = true
				gamewin.Moves++

				break
			}
		}
		gamewin.CheckWin()
		if gamewin.CheckWin() == false {
			ai.AiOmove()
			gamewin.CheckWin()
		} else {
			fmt.Print(gameboard.PrintBoard())
			fmt.Fprintln(w, gamewin.WonChat)
			t2.Execute(w, nil)

		}
	}

}

func main() {
	fmt.Println("Server Status:Listening Host:localhost Port:9090")
	http.HandleFunc("/ttt1", Number)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

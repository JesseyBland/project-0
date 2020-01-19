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

	type Cells struct {
		Cwin bool
		Wonc string
	}
	c := Cells{

		Cwin: gamewin.CheckWin(),
		Wonc: gamewin.WonChat,
	}

	t.Execute(w, c)
	r.ParseForm()
	xread := strings.Join(r.Form["quantity"], "")
	x, _ := strconv.Atoi(xread)
	xm := 0
	om := 0

	if x == 0 && gameboard.Board[x].Fill == false {
		fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[0], gameboard.Board[1], gameboard.Board[2])
		fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[3], gameboard.Board[4], gameboard.Board[5])
		fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[6], gameboard.Board[7], gameboard.Board[8])

	} else {

		for i := 0; i < len(gameboard.Board); i++ {
			if x == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "X"
				gameboard.Board[i].Fill = true
				gamewin.Moves++
				xm++
				fmt.Printf("X moves to %v\nTotal moves: %v\n", x, gamewin.Moves)
				break
			}

		}

		gamewin.CheckWin()
		if gamewin.CheckWin() == false && xm > om {
			ai.AiOmove()
			om++
			fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[0], gameboard.Board[1], gameboard.Board[2])
			fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[3], gameboard.Board[4], gameboard.Board[5])
			fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[6], gameboard.Board[7], gameboard.Board[8])

			gamewin.CheckWin()

		} else {
			fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[0], gameboard.Board[1], gameboard.Board[2])
			fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[3], gameboard.Board[4], gameboard.Board[5])
			fmt.Fprintf(w, "<p1>%v%v%v <br></p1>", gameboard.Board[6], gameboard.Board[7], gameboard.Board[8])
		}
		if gamewin.CheckWin() == true {
			fmt.Fprintf(w, "<br><p1>%v</p1>", gamewin.WonChat)
			fmt.Println(gamewin.WonChat)

		}
	}
}

func main() {
	fmt.Println("Server Status:Listening Host:localhost Port:8080")
	http.HandleFunc("/ttt1", Number)

	err := http.ListenAndServe(":8081", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

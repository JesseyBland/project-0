package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type cell struct {
	Slogic string
	Fill   bool
}

var (
	bg        *ebiten.Image
	x         *ebiten.Image
	o         *ebiten.Image
	winpic    *ebiten.Image
	grid      [9][2]float64
	board     [9]cell
	turn      = 0
	moveCount = 0
)

func pos(pos [2]float64) (float64, float64) {
	return pos[0], pos[1]
}
func init() {
	var err error
	bg, _, err = ebitenutil.NewImageFromFile("map.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	x, _, err = ebitenutil.NewImageFromFile("x.png", ebiten.FilterDefault)
	if err != nil {
		fmt.Println("Bad X image init Error:", err)
	}
	o, _, err = ebitenutil.NewImageFromFile("o.png", ebiten.FilterDefault)
	if err != nil {
		fmt.Println("Bad O image init Error:", err)
	}
	winpic, _, err = ebitenutil.NewImageFromFile("win.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	grid[0] = [2]float64{0, 0}
	grid[1] = [2]float64{75, 0}
	grid[2] = [2]float64{157, 0}
	grid[3] = [2]float64{0, 75}
	grid[4] = [2]float64{75, 75}
	grid[5] = [2]float64{157, 75}
	grid[6] = [2]float64{0, 155}
	grid[7] = [2]float64{75, 155}
	grid[8] = [2]float64{157, 155}

}

func tictactoe(screen *ebiten.Image) error {
	var gameText string

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	screen.DrawImage(bg, nil)

	switch {
	case ebiten.IsKeyPressed(ebiten.Key1):
		if board[0].Fill == false {
			board[0].Fill = true
			board[0].Slogic = "x"
			turn++
			moveCount++

		}
	case ebiten.IsKeyPressed(ebiten.Key2):
		if board[1].Fill == false {
			board[1].Fill = true
			board[1].Slogic = "x"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key3):
		if board[2].Fill == false {
			board[2].Fill = true
			board[2].Slogic = "x"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key4):
		if board[3].Fill == false {
			board[3].Fill = true
			board[3].Slogic = "x"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key5):
		if board[4].Fill == false {
			board[4].Fill = true
			board[4].Slogic = "x"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key6):
		if board[5].Fill == false {
			board[5].Fill = true
			board[5].Slogic = "x"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key7):
		if board[6].Fill == false {
			board[6].Fill = true
			board[6].Slogic = "x"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key8):
		if board[7].Fill == false {
			board[7].Fill = true
			board[7].Slogic = "x"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key9):
		if board[8].Fill == false {
			board[8].Fill = true
			board[8].Slogic = "x"
			moveCount++
			turn++
		}
	default:
		//turn = false

	}
	gameText = checkFill(screen, gameText)
	gameText = "Press 1-9 to place your move"
	if moveCount == 0 {
		ebitenutil.DebugPrintAt(screen, gameText, 0, 222)
	}
	return nil
}

//thats my chick fil a
func checkFill(screen *ebiten.Image, gameText string) string {
	for i := 0; i < 9; i++ {

		if board[i].Fill == true && board[i].Slogic == "x" {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(pos(grid[i]))
			screen.DrawImage(x, op)
			if turn == 1 && moveCount < 8 {

				gameText, turn = omove(gameText, turn)

			}
			if turn == 0 {
				gameText = fmt.Sprintf("X's Turn! Total Moves: %v", moveCount)
			} else {
				gameText = fmt.Sprintf("O's Turn! Total Moves: %v", moveCount)
			}
			ebitenutil.DebugPrintAt(screen, gameText, 0, 222)

		} else if board[i].Fill == true && board[i].Slogic == "o" {
			time.Sleep(300000 * time.Microsecond)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(pos(grid[i]))
			screen.DrawImage(o, op)
			// gameText = fmt.Sprint("turn: ", turn, " movecount:", moveCount)
		}
		if moveCount > 8 {

			screen.DrawImage(winpic, nil)
			ebitenutil.DebugPrintAt(screen, gameText, 0, 222)
		}
	}

	return gameText
}

func omove(gameText string, turn int) (string, int) {

	for turn == 1 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(9)

		for i := 0; i < 9; i++ {
			if n == i && board[i].Fill == false {

				board[i].Slogic = "o"
				board[i].Fill = true
				turn--
				moveCount++
				fmt.Println("n:", n, "turn: ", turn, "movecount:", moveCount)
				break

			}

		}

	}
	return gameText, turn
}

func main() {

	if err := ebiten.Run(tictactoe, 222, 245, 1, "TicTacBro"); err != nil {
		fmt.Println("Error happend on draw run. Error: ", err)

	}
}

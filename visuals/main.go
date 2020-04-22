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
	losepic   *ebiten.Image
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
	losepic, _, err = ebitenutil.NewImageFromFile("lose.png", ebiten.FilterDefault)
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
			board[0].Slogic = "X"
			turn++
			moveCount++

		}
	case ebiten.IsKeyPressed(ebiten.Key2):
		if board[1].Fill == false {
			board[1].Fill = true
			board[1].Slogic = "X"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key3):
		if board[2].Fill == false {
			board[2].Fill = true
			board[2].Slogic = "X"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key4):
		if board[3].Fill == false {
			board[3].Fill = true
			board[3].Slogic = "X"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key5):
		if board[4].Fill == false {
			board[4].Fill = true
			board[4].Slogic = "X"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key6):
		if board[5].Fill == false {
			board[5].Fill = true
			board[5].Slogic = "X"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key7):
		if board[6].Fill == false {
			board[6].Fill = true
			board[6].Slogic = "X"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key8):
		if board[7].Fill == false {
			board[7].Fill = true
			board[7].Slogic = "X"
			moveCount++
			turn++
		}
	case ebiten.IsKeyPressed(ebiten.Key9):
		if board[8].Fill == false {
			board[8].Fill = true
			board[8].Slogic = "X"
			moveCount++
			turn++
		}
	// case CheckWin() == true:
	// 	screen.DrawImage(winpic, nil)
	// 	ebitenutil.DebugPrintAt(screen, gameText, 0, 222)

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

		if board[i].Fill == true && board[i].Slogic == "X" {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(pos(grid[i]))
			screen.DrawImage(x, op)
			if turn == 1 && moveCount < 8 {

				gameText, turn = omove(gameText, turn)

			}
			if CheckWin() == true {
				gameText = WonChat
			} else if turn == 0 {
				gameText = fmt.Sprintf("X's Turn! Total Moves: %v", moveCount)
			} else if turn == 1 {
				gameText = fmt.Sprintf("O's Turn! Total Moves: %v", moveCount)
			}
			ebitenutil.DebugPrintAt(screen, gameText, 0, 222)

		} else if board[i].Fill == true && board[i].Slogic == "O" {
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

		if WonChat == "X Wins!" {
			screen.DrawImage(winpic, nil)
			gameText = WonChat
			ebitenutil.DebugPrintAt(screen, gameText, 0, 222)
		}
		if WonChat == "O Wins!" {
			screen.DrawImage(losepic, nil)
			ebitenutil.DebugPrintAt(screen, "", 0, 222)
			gameText = WonChat
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

				board[i].Slogic = "O"
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

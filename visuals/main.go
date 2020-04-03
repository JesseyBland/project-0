package main

import (
	"fmt"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type cell struct {
	Slogic string
	Fill   bool
}

var (
	bg    *ebiten.Image
	x     *ebiten.Image
	o     *ebiten.Image
	grid  [9][2]float64
	board = make([]cell, 0)
	//X     int
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
		fmt.Println("Bad X image init Error:", err)
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
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}
	// fmt.Printf("%v's move pick 1-9\n", strings.Title(user))
	screen.DrawImage(bg, nil)
	ebitenutil.DebugPrintAt(screen, "Press 1-9 to place your move", 0, 222)
	fmt.Println()

	// if ebiten.IsKeyPressed(ebiten.Key1) {
	// 	X := 1
	// 	for i := 0; i < 9; i++ {
	// 		if X == (i+1) && board[i].Fill == false {
	// 			board[i].Slogic = "X"
	// 			board[i].Fill = true
	// 			op := &ebiten.DrawImageOptions{}
	// 			op.GeoM.Translate(pos(grid[0]))
	// 			screen.DrawImage(x, op)
	// 			break

	// 		}
	// 	}
	// }

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(pos(grid[8]))
	screen.DrawImage(x, op2)
	return nil
}
func main() {

	if err := ebiten.Run(tictactoe, 222, 245, 1, "TicTacBro"); err != nil {
		fmt.Println("Error happend on draw run. Error: ", err)

	}
}

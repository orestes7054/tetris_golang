package board

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)
const (
	ScreenWidth  = 450
	ScreenHeight = 650
	BlockSize = 20
	BoardWidth = 10
	BoardHeight = 20
	Padding = 20
)

type Game struct {
	board [BoardHeight][BoardWidth]int

}

var (
	backGroundColor = color.RGBA{0, 0, 0, 255}
	blockColor = color.RGBA{255,255,255,255}
)


func (g *Game) Layout(outSideWidth, outSideHeigth int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Draw(screen *ebiten.Image){
	screen.Fill(backGroundColor)
	
	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
			if g.board[y][x] != 0 {
				vector.DrawFilledRect(screen, float32(Padding+x*BlockSize), float32(Padding+y*BlockSize), float32(BlockSize), float32(BlockSize), blockColor, true)
			}
		}
	}

}

func (g *Game) Update() error {
	return nil
}

func NewGame() *Game{
	return &Game{}
}

func CreateBoard() {
	file, err := os.Open("public/images/tetris.png") 

	if err != nil{
		fmt.Printf("Error while opening the icon file: %s", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil{
		fmt.Printf("Error while decoding the icon file: %s", err)
	}
	icon := []image.Image{img}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Tetris Game Golang")
	ebiten.SetWindowIcon(icon)
	game := NewGame()

	game.board[0][0] = 1
	game.board[1][1] = 1
	game.board[2][2] = 1
	
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

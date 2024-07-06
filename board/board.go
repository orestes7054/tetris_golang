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

type Game struct {}

const (
	ScreenWidth  = 450
	ScreenHeigth = 650
	boardRows    = 10
	boardCols    = 20
	Padding = 50
)

var (
	backGroundColor = color.RGBA{135, 132, 133, 80}
)


func (g *Game) Layout(outSideWidth, outSideHeigth int) (screenWidth, screenHeigth int) {
	return ScreenWidth - Padding, ScreenHeigth - Padding
}

func (g *Game) Draw(screen *ebiten.Image){
	screen.Fill(backGroundColor)
	tileSize := (ScreenHeigth - Padding) / boardRows
	for i:=0; i<=ScreenHeigth; i += tileSize {
		vector.StrokeLine(screen, float32(0), float32(i), float32(ScreenWidth - Padding), float32(i), float32(0.5), color.Black, true)
	}

	for i:=0; i<=ScreenWidth; i += tileSize {
		vector.StrokeLine(screen, float32(i), float32(0), float32(i), float32(ScreenHeigth - Padding), float32(0.5), color.Black, true)
	}


}

func (g *Game) Update() error {
	return nil
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

	ebiten.SetWindowSize(ScreenWidth, ScreenHeigth)
	ebiten.SetWindowTitle("Tetris Game Golang")
	ebiten.SetWindowIcon(icon)
	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

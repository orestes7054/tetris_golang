package board

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {}

const (
	ScreenWidth  = 600
	ScreenHeigth = 200
	boardRows    = 60
	boardCols    = 20
)

var (
	backGroundColor = color.RGBA{50, 100, 50, 50}
)


func (g *Game) Layout(outSideWidth, outSideHeigth int) (screenWidth, ScreenHeigth int) {
	return screenWidth, ScreenHeigth
}

func (g *Game) Draw(screen *ebiten.Image){
	screen.Fill(backGroundColor)
}

func (g *Game) Update() error {
	return nil
}

func createBoard() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeigth)
	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

package board

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	//"tetris_golang/tetrominos"
)


const (
	ScreenWidth  = 450
	ScreenHeight = 650
	BlockSize = 50
	BoardWidth = 10
	BoardHeight = 20
	Padding = 20
)

type Game struct {
	board [BoardHeight][BoardWidth]int

}

var (
	backGroundColor = color.RGBA{64, 64, 64, 0}
	blockColor = color.RGBA{255,128,0,255}
	lineColor = color.RGBA{0,0,0,255}
)


func (g *Game) Layout(outSideWidth, outSideHeigth int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Draw(screen *ebiten.Image){
	screen.Fill(backGroundColor)
	

	vector.DrawFilledRect(screen,float32(0), float32(0), float32(BlockSize), float32(BlockSize), blockColor, true)
	// the BlockSize * X moves the block two blocks to the right idem with BlockSize * Y
	vector.DrawFilledRect(screen,float32(BlockSize * 2), float32(BlockSize), float32(BlockSize), float32(BlockSize), blockColor, true)

	
	for x := 0; x < BoardWidth; x++ {
		for y := 0; y < BoardHeight; y++ {
			vector.StrokeRect(screen, float32(x * BlockSize), float32(y * BlockSize), float32(BlockSize), float32(BlockSize), float32(0.50), lineColor, true)
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


	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

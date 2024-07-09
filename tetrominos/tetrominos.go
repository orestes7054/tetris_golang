package tetrominos

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	//"golang.org/x/exp/shiny/screen"
)

type Point struct {
	X, Y int
}

type Tetromino struct {
	Rotations [][]Point
	Color     int
}

var Cube = Tetromino{
	Rotations: [][]Point{
		{
			{0, 0}, {1, 0}, {0, 1}, {1, 1},
		},
	},
	Color: 1,
}

func WriteFigure(screen *ebiten.Image, figure Tetromino, blockSize int, blockColor color.RGBA) {

	for _, rotation := range figure.Rotations{
		for _, point := range rotation{
			vector.DrawFilledRect(screen, float32(point.X * blockSize), float32(point.Y * blockSize), float32(blockSize), float32(blockSize), blockColor, true)
		}
	}


}

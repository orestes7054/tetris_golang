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
	Color     color.RGBA
}



var FiguresMap = map[string]Tetromino{
	"O": {
		Rotations: [][]Point{
			{
				{0, 0}, {1, 0}, {0, 1}, {1, 1},
			},
		},
		Color: color.RGBA{0, 255, 255, 255}, // Cian
	},
	"I": {
		Rotations: [][]Point{
			{
				{0, 1}, {1, 1}, {2, 1}, {3, 1},
			},
			{
				{1, 0}, {1, 1}, {1, 2}, {1, 3},
			},
		},
		Color: color.RGBA{255, 255, 0, 255}, // Amarillo
	},
	"T": {
		Rotations: [][]Point{
			{
				{1, 0}, {0, 1}, {1, 1}, {2, 1},
			},
			{
				{1, 0}, {1, 1}, {1, 2}, {0, 1},
			},
			{
				{0, 1}, {1, 1}, {2, 1}, {1, 2},
			},
			{
				{1, 0}, {1, 1}, {1, 2}, {2, 1},
			},
		},
		Color: color.RGBA{128, 0, 128, 255}, // Púrpura
	},
	"S": {
		Rotations: [][]Point{
			{
				{1, 0}, {2, 0}, {0, 1}, {1, 1},
			},
			{
				{0, 0}, {0, 1}, {1, 1}, {1, 2},
			},
		},
		Color: color.RGBA{0, 255, 0, 255}, // Verde
	},
	"Z": {
		Rotations: [][]Point{
			{
				{0, 0}, {1, 0}, {1, 1}, {2, 1},
			},
			{
				{1, 0}, {0, 1}, {1, 1}, {0, 2},
			},
		},
		Color: color.RGBA{255, 0, 0, 255}, // Rojo
	},
	"J": {
		Rotations: [][]Point{
			{
				{0, 0}, {0, 1}, {1, 1}, {2, 1},
			},
			{
				{1, 0}, {1, 1}, {1, 2}, {0, 2},
			},
			{
				{0, 1}, {1, 1}, {2, 1}, {2, 2},
			},
			{
				{1, 0}, {0, 0}, {0, 1}, {0, 2},
			},
		},
		Color: color.RGBA{0, 0, 255, 255}, // Azul
	},
	"L": {
		Rotations: [][]Point{
			{
				{2, 0}, {0, 1}, {1, 1}, {2, 1},
			},
			{
				{1, 0}, {1, 1}, {1, 2}, {2, 2},
			},
			{
				{0, 1}, {1, 1}, {2, 1}, {0, 2},
			},
			{
				{1, 0}, {1, 1}, {1, 2}, {0, 0},
			},
		},
		Color: color.RGBA{255, 165, 0, 255}, // Naranja
	},

}




type ActualPosition struct{
	X, Y int
}




func WriteFigure(screen *ebiten.Image, figure Tetromino, blockSize int, position ActualPosition, positionIndex int) {

	rotation := figure.Rotations[positionIndex]
	figureColor := figure.Color
	for _, point := range rotation{
		x := float32((point.X + position.X))
		y := float32((point.Y + position.Y))

		vector.DrawFilledRect(screen, x * float32(blockSize) ,  y * float32(blockSize), 
							  float32(blockSize), float32(blockSize), figureColor, true,
							  )
	}

}

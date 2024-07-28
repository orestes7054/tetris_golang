package board

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"tetris_golang/tetrominos"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	//"tetris_golang/tetrominos"
)


const (
	BlockSize = 35
	BoardWidth = 10
	BoardHeight = 20
	ScreenWidth  = BoardWidth * BlockSize 
	ScreenHeight = BoardHeight * BlockSize
)

type FieldSize struct {
    Width  int
    Height int
}

var ActualFieldSize = FieldSize{
	Width:  BoardWidth,
	Height: BoardHeight, 
}


type Game struct {
	board [BoardHeight][BoardWidth]int
	lastTime time.Time
	interval time.Duration
	selectedTetromino tetrominos.Tetromino
	tetrominoPosition int

}

var (
	backGroundColor = color.RGBA{64, 64, 64, 0}
	lineColor = color.RGBA{0,0,0,255}
	itialPosition = tetrominos.ActualPosition{
							X: int(float64(BoardWidth / 2)) - 1, 
							Y: int(float64(BoardHeight / 2)) - 1, 
					}
	figures = tetrominos.FiguresMap
	figuresNames = [5]string{"O", "I", "T", "S", "Z"}

)


func (g *Game) Layout(outSideWidth, outSideHeigth int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}


func (g *Game) Draw(screen *ebiten.Image){
	screen.Fill(backGroundColor)

	tetrominos.WriteFigure(screen, g.selectedTetromino, BlockSize,  itialPosition, int(g.tetrominoPosition))

	
	for x := 0; x < BoardWidth; x++ {
		for y := 0; y < BoardHeight; y++ {
			vector.StrokeRect(screen, float32(x * BlockSize), float32(y * BlockSize), float32(BlockSize), float32(BlockSize), float32(0.50), lineColor, true)
		}
	}

	
}

func returnTetrominoSelected() (tetrominos.Tetromino, int){
	pickedFigure := figuresNames[rand.Int31n(5)]
	selectedTetromino := figures[pickedFigure]
	startPosition :=  rand.Int31n(int32(len(selectedTetromino.Rotations)))
	return selectedTetromino, int(startPosition)
}

func (g *Game) Update() error {
	actualTime := time.Now()
	if(actualTime.Sub(g.lastTime) > g.interval){
		g.lastTime = actualTime
		g.selectedTetromino, g.tetrominoPosition = returnTetrominoSelected()
	}
	return nil
	
}

func NewGame() *Game{
	return &Game{
		lastTime: time.Now(),
		interval: 2 * time.Second,
		selectedTetromino: figures["O"],
		tetrominoPosition: 0,
	}
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

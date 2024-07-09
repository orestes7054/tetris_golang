package tetrominos

type Point struct{
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









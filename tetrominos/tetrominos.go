package tetrominos

type Tetrominos struct {
	rotations [][][]int
	color     int
}

var Shapes = []Tetrominos{
	//cube
	{
		rotations: [][][]int{
			{
				{1, 1},
				{1, 1},
			},
			
		},
		color: 1,
	},
	{ // I Tetromino
		rotations: [][][]int{
			{
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
			},
			{
				{0, 0, 0, 0},
				{1, 1, 1, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		color: 1,
	},

	
}
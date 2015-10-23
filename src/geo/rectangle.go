package geo

type Rectangle struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

type Point struct {
	RowIndex int
	ColumnIndex int
	X float64
	Y float64
	Role int  //-1 means empty, 1 means bean
}
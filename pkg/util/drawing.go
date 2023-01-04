package util

import "strings"

type Drawing struct {
	inner   [][]byte
	height  int
	width   int
	centerX int
	centerY int
}

func InitDrawing(width int, height int) *Drawing {
	inner := make([][]byte, height)
	for i := 0; i < height; i++ {
		inner[i] = make([]byte, width)
	}
	return &Drawing{
		inner:   inner,
		height:  height,
		width:   width,
		centerX: 0,
		centerY: 0,
	}
}

func InitDrawingWithCenter(width int, height int, centerX int, centerY int) *Drawing {
	inner := make([][]byte, height)
	for i := 0; i < height; i++ {
		inner[i] = make([]byte, width)
	}
	return &Drawing{
		inner:   inner,
		height:  height,
		width:   width,
		centerX: centerX,
		centerY: centerY,
	}
}

func (d *Drawing) Fill(char byte) *Drawing {
	for i := 0; i < d.height; i++ {
		for j := 0; j < d.width; j++ {
			d.inner[i][j] = char
		}
	}

	return d
}

func (d *Drawing) DrawAt(char byte, x int, y int) *Drawing {
	d.inner[d.centerY+y][d.centerX+x] = char

	return d
}

func (d *Drawing) String() string {
	var sb strings.Builder
	for i := 0; i < d.height; i++ {
		sb.Write(d.inner[d.height-1-i])
		if i < d.height-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

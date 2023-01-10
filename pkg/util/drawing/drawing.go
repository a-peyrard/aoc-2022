package drawing

import (
	"aoc2022/pkg/util"
	"aoc2022/pkg/util/collection"
	"aoc2022/pkg/util/geo"
	"fmt"
	"strings"
)

type Drawing struct {
	inner       [][]byte
	height      int
	width       int
	centerX     int
	centerY     int
	bottomToTop bool
}

func InitDrawing(width int, height int) *Drawing {
	inner := make([][]byte, height)
	for i := 0; i < height; i++ {
		inner[i] = make([]byte, width)
	}
	return &Drawing{
		inner:       inner,
		height:      height,
		width:       width,
		centerX:     0,
		centerY:     0,
		bottomToTop: true,
	}
}

func InitDrawingWithCenter(width int, height int, centerX int, centerY int) *Drawing {
	inner := make([][]byte, height)
	for i := 0; i < height; i++ {
		inner[i] = make([]byte, width)
	}
	return &Drawing{
		inner:       inner,
		height:      height,
		width:       width,
		centerX:     centerX,
		centerY:     centerY,
		bottomToTop: true,
	}
}

func InitDrawingTopToBottom(width int, height int) *Drawing {
	inner := make([][]byte, height)
	for i := 0; i < height; i++ {
		inner[i] = make([]byte, width)
	}
	return &Drawing{
		inner:       inner,
		height:      height,
		width:       width,
		centerX:     0,
		centerY:     0,
		bottomToTop: false,
	}
}

func InitDrawingTopToBottomWithCenter(width int, height int, centerX int, centerY int) *Drawing {
	inner := make([][]byte, height)
	for i := 0; i < height; i++ {
		inner[i] = make([]byte, width)
	}
	return &Drawing{
		inner:       inner,
		height:      height,
		width:       width,
		centerX:     centerX,
		centerY:     centerY,
		bottomToTop: false,
	}
}

func (d *Drawing) Height() int {
	return d.height
}

func (d *Drawing) Width() int {
	return d.width
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

func (d *Drawing) ValueAt(x int, y int) byte {
	return d.inner[d.centerY+y][d.centerX+x]
}

func (d *Drawing) String() string {
	var sb strings.Builder

	transformY := func(y int) int {
		if d.bottomToTop {
			return d.height - 1 - y
		}
		return y
	}

	for i := 0; i < d.height; i++ {
		sb.Write(d.inner[transformY(i)])
		if i < d.height-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func (d *Drawing) DrawLine(char byte, from, to geo.Coordinate) *Drawing {
	if from.X != to.X && from.Y != to.Y {
		panic(
			fmt.Sprintf(
				"Unable to draw line from %#v, to %#v, this neither an horizontal nor a vertical line",
				from,
				to,
			),
		)
	}

	if from.X == to.X {
		d.drawVerticalLine(char, from.X, from.Y, to.Y)
	} else if from.Y == to.Y {
		d.drawHorizontalLine(char, from.X, to.X, from.Y)
	}

	return d
}

func (d *Drawing) drawVerticalLine(char byte, x int, y1 int, y2 int) *Drawing {
	min := util.Min(y1, y2)
	max := util.Max(y1, y2)
	for i := min; i < max+1; i++ {
		d.DrawAt(char, x, i)
	}

	return d
}

func (d *Drawing) drawHorizontalLine(char byte, x1 int, x2 int, y int) *Drawing {
	min := util.Min(x1, x2)
	max := util.Max(x1, x2)
	for i := min; i < max+1; i++ {
		d.DrawAt(char, i, y)
	}

	return d
}

func DrawMatrix[T any](m *collection.Matrix[T], toByte func(T, int, int) byte) *Drawing {
	d := InitDrawingTopToBottom(m.Width(), m.Height())
	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			d.DrawAt(toByte(m.ValueAt(x, y), x, y), x, y)
		}
	}
	return d
}

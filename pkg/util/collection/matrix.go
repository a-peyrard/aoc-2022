package collection

type Matrix[T any] [][]T

type Coordinate struct {
	X int
	Y int
}

func CreateMatrix[T any](width int, height int, defaultValue T) Matrix[T] {
	matrix := make([][]T, height)
	for j := 0; j < height; j++ {
		matrix[j] = make([]T, width)
		for i := 0; i < width; i++ {
			matrix[j][i] = defaultValue
		}
	}

	return matrix
}

func (m *Matrix[T]) ValueAt(x, y int) T {
	return (*m)[y][x]
}

func (m *Matrix[T]) ValueAtC(c *Coordinate) T {
	return m.ValueAt(c.X, c.Y)
}

func (m *Matrix[T]) PutAt(x, y int, v T) {
	(*m)[y][x] = v
}

func (m *Matrix[T]) PutAtC(c *Coordinate, v T) {
	m.PutAt(c.X, c.Y, v)
}

func (m *Matrix[T]) Height() int {
	return len(*m)
}

func (m *Matrix[T]) Width() int {
	return len((*m)[0])
}

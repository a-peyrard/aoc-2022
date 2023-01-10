package geo

import "aoc2022/pkg/util"

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) ManhattanDistance(to Coordinate) int {
	return util.AbsInt(c.X-to.X) + util.AbsInt(c.Y-to.Y)
}

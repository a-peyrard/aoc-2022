package day12

import (
	"aoc2022/pkg/util"
	"aoc2022/pkg/util/collection"
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

/*
--- Day 12: Hill Climbing Algorithm ---

You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent
signal. You ask the device for a heightmap of the surrounding area (your puzzle input). The heightmap shows the local
area from above broken into a grid; the elevation of each square of the grid is given by a single lowercase letter,
where a is the lowest elevation, b is the next-lowest, and so on up to the highest elevation, z. Also included on the
heightmap are marks for your current position (S) and the location that should get the best signal (E). Your current
position (S) has elevation a, and the location that should get the best signal (E) has elevation z. You'd like to reach
E, but to save energy, you should do it in as few steps as possible. During each step, you can move exactly one square
up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be
at most one higher than the elevation of your current square; that is, if your current elevation is m, you could step to
elevation n, but not to elevation o. (This also means that the elevation of the destination square can be much lower
than the elevation of your current square.)

For example:

Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi

Here, you start in the top-left corner; your goal is near the middle. You could start by moving down or right, but
eventually you'll need to head toward the e at the bottom. From there, you can spiral around to the goal:

v..v<<<<
>v.vv<<^
.>vv>E^^
..v>>>^^
..>>>>>^

In the above diagram, the symbols indicate whether the path exits each square moving up (^), down (v), left (<), or
right (>). The location that should get the best signal is still E, and . marks unvisited squares. This path reaches the
goal in 31 steps, the fewest possible. What is the fewest steps required to move from your current position to the
location that should get the best signal?

--- Part Two ---

As you walk up the hill, you suspect that the Elves will want to turn this into a hiking trail. The beginning isn't very
scenic, though; perhaps you can find a better starting point. To maximize exercise while hiking, the trail should start
as low as possible: elevation a. The goal is still the square marked E. However, the trail should still be direct,
taking the fewest steps to reach its goal. So, you'll need to find the shortest path from any square at elevation a to
the square marked E. Again consider the example from above:

Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi

Now, there are six choices for starting position (five marked a, plus the square marked S that counts as being at
elevation a). If you start at the bottom-left square, you can reach the goal most quickly:

...v<<<<
...vv<<^
...v>E^^
.>v>>>^^
>^>>>>>^

This path reaches the goal in only 29 steps, the fewest possible. What is the fewest steps required to move starting
from any square with elevation a to the location that should get the best signal?

*/

const Debug = false

type heightmap collection.Matrix[byte]

func (h *heightmap) value(c *collection.Coordinate) byte {
	return h.valueAt(c.X, c.Y)
}

func (h *heightmap) height() int {
	return len(*h)
}

func (h *heightmap) width() int {
	return len((*h)[0])
}

func (h *heightmap) valueAt(x, y int) byte {
	val := (*h)[y][x]
	if val == 'S' {
		return 'a'
	}
	if val == 'E' {
		return 'z'
	}
	return val
}

func parse(input string) (heightmap, *collection.Coordinate, *collection.Coordinate) {
	var (
		res        heightmap
		line       string
		row        []byte
		start, end *collection.Coordinate
	)
	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		line = sc.Text()
		row = make([]byte, 0)
		for i := 0; i < len(line); i++ {
			row = append(row, line[i])
			if line[i] == 'S' {
				start = &collection.Coordinate{X: i, Y: len(res)}
			} else if line[i] == 'E' {
				end = &collection.Coordinate{X: i, Y: len(res)}
			}
		}
		res = append(res, row)
	}

	return res, start, end
}

func findShortestPath(heightmap heightmap, start *collection.Coordinate, target *collection.Coordinate) int {
	dist := collection.CreateMatrix[int](heightmap.width(), heightmap.height(), math.MaxInt)
	dist.PutAtC(start, 0)
	prev := collection.CreateMatrix[*collection.Coordinate](heightmap.width(), heightmap.height(), nil)

	queue := make(collection.PriorityQueue[*collection.Coordinate], 0)
	heap.Init(&queue)

	heap.Push(&queue, collection.CreateItem(start, dist.ValueAtC(start)))

	var (
		current   *collection.Item[*collection.Coordinate]
		neighbors []*collection.Coordinate
		alt       int
	)

	visited := collection.NewSet[collection.Coordinate]()
	for queue.Len() > 0 {
		current = heap.Pop(&queue).(*collection.Item[*collection.Coordinate])
		visited.Add(*current.GetValue())
		if *current.GetValue() == *target {
			break
		}

		neighbors = extractNeighbors(heightmap, current.GetValue())
		for _, neighbor := range neighbors {
			if visited.Contains(*neighbor) {
				continue
			}

			alt = dist.ValueAtC(current.GetValue()) + 1
			if alt < dist.ValueAtC(neighbor) {
				dist.PutAtC(neighbor, alt)
				prev.PutAtC(neighbor, current.GetValue())
				heap.Push(&queue, collection.CreateItem(neighbor, alt))
			}
		}
	}

	if Debug {
		d := util.DrawMatrix(&dist, func(v, x, y int) byte {
			character := heightmap.valueAt(x, y)
			if v == math.MaxInt {
				return character - 32 // print the letter in upper case
			}
			return character
		})
		fmt.Printf("\nhere are the analyzed nodes: \n%s\n\n", d.String())
	}

	return dist.ValueAtC(target)
}

func extractNeighbors(heightmap heightmap, c *collection.Coordinate) []*collection.Coordinate {
	neighbors := make([]*collection.Coordinate, 0)
	value := heightmap.value(c)
	// North
	if c.Y > 0 && isWalkable(heightmap.valueAt(c.X, c.Y-1), value) {
		neighbors = append(neighbors, &collection.Coordinate{X: c.X, Y: c.Y - 1})
	}
	// East
	if c.X < heightmap.width()-1 && isWalkable(heightmap.valueAt(c.X+1, c.Y), value) {
		neighbors = append(neighbors, &collection.Coordinate{X: c.X + 1, Y: c.Y})
	}
	// South
	if c.Y < heightmap.height()-1 && isWalkable(heightmap.valueAt(c.X, c.Y+1), value) {
		neighbors = append(neighbors, &collection.Coordinate{X: c.X, Y: c.Y + 1})
	}
	// West
	if c.X > 0 && isWalkable(heightmap.valueAt(c.X-1, c.Y), value) {
		neighbors = append(neighbors, &collection.Coordinate{X: c.X - 1, Y: c.Y})
	}

	return neighbors
}

func isWalkable(origin byte, dest byte) bool {
	return int(dest)-int(origin) <= 1
}

func doSolution1(raw string) int {
	heightmap, start, end := parse(raw)
	return findShortestPath(heightmap, end, start)
}

func Solution1() int {
	return doSolution1(util.GetInputContent())
}

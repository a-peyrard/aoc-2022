package day14

import (
	"aoc2022/pkg/util"
	"aoc2022/pkg/util/collection"
	"bufio"
	"fmt"
	"strings"
)

/*
--- Day 14: Regolith Reservoir ---

The distress signal leads you to a giant waterfall! Actually, hang on - the signal seems like it's coming from the
waterfall itself, and that doesn't make any sense. However, you do notice a little path that leads behind the waterfall.
Correction: the distress signal leads you behind a giant waterfall! There seems to be a large cave system here, and the
signal definitely leads further inside. As you begin to make your way deeper underground, you feel the ground rumble for
a moment. Sand begins pouring into the cave! If you don't quickly figure out where the sand is going, you could quickly
become trapped! Fortunately, your familiarity with analyzing the path of falling material will come in handy here. You
scan a two-dimensional vertical slice of the cave above you (your puzzle input) and discover that it is mostly air with
structures made of rock. Your scan traces the path of each solid rock structure and reports the x,y coordinates that
form the shape of the path, where x represents distance to the right and y represents distance down. Each path appears
as a single line of text in your scan. After the first point of each path, each point indicates the end of a straight
horizontal or vertical line to be drawn from the previous point. For example:

498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9

This scan means that there are two paths of rock; the first path consists of two straight lines, and the second path
consists of three straight lines. (Specifically, the first path consists of a line of rock from 498,4 through 498,6 and
another line of rock from 498,6 through 496,6.) The sand is pouring into the cave from point 500,0. Drawing rock as #,
air as ., and the source of the sand as +, this becomes:


  4     5  5
  9     0  0
  4     0  3
0 ......+...
1 ..........
2 ..........
3 ..........
4 ....#...##
5 ....#...#.
6 ..###...#.
7 ........#.
8 ........#.
9 #########.

Sand is produced one unit at a time, and the next unit of sand is not produced until the previous unit of sand comes to
rest. A unit of sand is large enough to fill one tile of air in your scan. A unit of sand always falls down one step if
possible. If the tile immediately below is blocked (by rock or sand), the unit of sand attempts to instead move
diagonally one step down and to the left. If that tile is blocked, the unit of sand attempts to instead move diagonally
one step down and to the right. Sand keeps moving as long as it is able to do so, at each step trying to move down, then
down-left, then down-right. If all three possible destinations are blocked, the unit of sand comes to rest and no longer
moves, at which point the next unit of sand is created back at the source. So, drawing sand that has come to rest as o,
the first unit of sand simply falls straight down and then stops:

......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
......o.#.
#########.

The second unit of sand then falls straight down, lands on the first one, and then comes to rest to its left:

......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
.....oo.#.
#########.

After a total of five units of sand have come to rest, they form this pattern:

......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
......o.#.
....oooo#.
#########.

After a total of 22 units of sand:

......+...
..........
......o...
.....ooo..
....#ooo##
....#ooo#.
..###ooo#.
....oooo#.
...ooooo#.
#########.

Finally, only two more units of sand can possibly come to rest:

......+...
..........
......o...
.....ooo..
....#ooo##
...o#ooo#.
..###ooo#.
....oooo#.
.o.ooooo#.
#########.

Once all 24 units of sand shown above have come to rest, all further sand flows out the bottom, falling into the endless
void. Just for fun, the path any new sand takes before falling forever is shown here with ~:

.......+...
.......~...
......~o...
.....~ooo..
....~#ooo##
...~o#ooo#.
..~###ooo#.
..~..oooo#.
.~o.ooooo#.
~#########.
~..........
~..........
~..........

Using your scan, simulate the falling sand. How many units of sand come to rest before sand starts flowing into the
abyss below?

*/

const Debug = true

func parse(input string) ([][]collection.Coordinate, int, int) {
	var (
		res            [][]collection.Coordinate
		row            []collection.Coordinate
		line           string
		rawCoordinates []string
		maxY, maxX     int
	)
	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		line = sc.Text()
		rawCoordinates = strings.Split(line, " -> ")
		row = make([]collection.Coordinate, len(rawCoordinates))
		for i := 0; i < len(rawCoordinates); i++ {
			coordinate := parseCoordinate(rawCoordinates[i])
			if maxX < coordinate.X {
				maxX = coordinate.X
			}
			if maxY < coordinate.Y {
				maxY = coordinate.Y
			}
			row[i] = coordinate
		}
		res = append(res, row)
	}

	return res, maxX, maxY
}

func parseCoordinate(rawCoordinate string) collection.Coordinate {
	tokens := strings.Split(rawCoordinate, ",")
	return collection.Coordinate{
		X: util.UnsafeParseInt(tokens[0]),
		Y: util.UnsafeParseInt(tokens[1]),
	}
}

func dropSand(drawing *util.Drawing, from collection.Coordinate) bool {
	var (
		next    = from
		canMove = true
		canRest = false
	)
	for canMove {
		next, canMove, canRest = advanceSand(drawing, next)
	}

	if canRest {
		drawing.DrawAt('o', next.X, next.Y)
	}

	return canRest
}

// return is:
// - the next coordinate where to go,
// - can we still move
// - did we found a place where to rest?
func advanceSand(drawing *util.Drawing, current collection.Coordinate) (collection.Coordinate, bool, bool) {
	if current.Y+1 > drawing.Height()-1 {
		return current, false, false // we reached the bottom
	}

	// first try to go down
	if drawing.ValueAt(current.X, current.Y+1) == '.' {
		return collection.Coordinate{X: current.X, Y: current.Y + 1}, true, false
	}
	// then try to go diagonal left
	if current.X > 0 && drawing.ValueAt(current.X-1, current.Y+1) == '.' {
		return collection.Coordinate{X: current.X - 1, Y: current.Y + 1}, true, false
	}
	// and diagonal right
	if current.X < drawing.Width()-1 && drawing.ValueAt(current.X+1, current.Y+1) == '.' {
		return collection.Coordinate{X: current.X + 1, Y: current.Y + 1}, true, false
	}

	// we can not move, so we will rest there
	return current, false, true
}

func doSolution1(raw string) int {
	lines, maxX, maxY := parse(raw)
	drawing := util.InitDrawingTopToBottom(maxX+1, maxY+1).
		Fill('.').
		DrawAt('+', 500, 0)

	var start collection.Coordinate
	for _, line := range lines {
		start = line[0]
		for i := 1; i < len(line); i++ {
			drawing.DrawLine('#', start, line[i])
			start = line[i]
		}
	}

	if Debug {
		fmt.Printf("\n====================== initial drawing:\n%s\n\n", drawing.String())
	}

	var (
		sandResting = true
		counter     = 0
		sandSource  = collection.Coordinate{X: 500, Y: 0}
	)
	for sandResting {
		sandResting = dropSand(drawing, sandSource)
		if sandResting {
			counter += 1
		}
	}

	if Debug {
		fmt.Printf("\n====================== final drawing:\n%s\n", drawing.String())
	}

	return counter
}

func Solution1() int {
	return doSolution1(util.GetInputContent())
}
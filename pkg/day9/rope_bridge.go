package day9

import (
	"aoc2022/pkg/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 9: Rope Bridge ---

This rope bridge creaks as you walk along it. You aren't sure how old it is, or whether it can even support your weight.
It seems to support the Elves just fine, though. The bridge spans a gorge which was carved out by the massive river far
below you. You step carefully; as you do, the ropes stretch and twist. You decide to distract yourself by modeling rope
physics; maybe you can even figure out where not to step. Consider a rope with a knot at each end; these knots mark the
head and the tail of the rope. If the head moves far enough away from the tail, the tail is pulled toward the head. Due
to nebulous reasoning involving Planck lengths, you should be able to model the positions of the knots on a
two-dimensional grid. Then, by following a hypothetical series of motions (your puzzle input) for the head, you can
determine how the tail will move. Due to the aforementioned Planck lengths, the rope must be quite short; in fact, the
head (H) and tail (T) must always be touching (diagonally adjacent and even overlapping both count as touching):

....
.TH.
....

....
.H..
..T.
....

...
.H. (H covers T)
...

If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step in that
direction so it remains close enough:

.....    .....    .....
.TH.. -> .T.H. -> ..TH.
.....    .....    .....

...    ...    ...
.T.    .T.    ...
.H. -> ... -> .T.
...    .H.    .H.
...    ...    ...

Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step
diagonally to keep up:

.....    .....    .....
.....    ..H..    ..H..
..H.. -> ..... -> ..T..
.T...    .T...    .....
.....    .....    .....

.....    .....    .....
.....    .....    .....
..H.. -> ...H. -> ..TH.
.T...    .T...    .....
.....    .....    .....

You just need to work out where the tail goes as the head follows a series of motions. Assume the head and the tail both
start at the same position, overlapping. For example:

R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2

This series of motions moves the head right four steps, then up four steps, then left three steps, then down one step,
and so on. After each step, you'll need to update the position of the tail if the step means the head is no longer
adjacent to the tail. Visually, these motions occur as follows (s marks the starting position as a reference point):

== Initial State ==

......
......
......
......
H.....  (H covers T, s)

== R 4 ==

......
......
......
......
TH....  (T covers s)

......
......
......
......
sTH...

......
......
......
......
s.TH..

......
......
......
......
s..TH.

== U 4 ==

......
......
......
....H.
s..T..

......
......
....H.
....T.
s.....

......
....H.
....T.
......
s.....

....H.
....T.
......
......
s.....

== L 3 ==

...H..
....T.
......
......
s.....

..HT..
......
......
......
s.....

.HT...
......
......
......
s.....

== D 1 ==

..T...
.H....
......
......
s.....

== R 4 ==

..T...
..H...
......
......
s.....

..T...
...H..
......
......
s.....

......
...TH.
......
......
s.....

......
....TH
......
......
s.....

== D 1 ==

......
....T.
.....H
......
s.....

== L 5 ==

......
....T.
....H.
......
s.....

......
....T.
...H..
......
s.....

......
......
..HT..
......
s.....

......
......
.HT...
......
s.....

......
......
HT....
......
s.....

== R 2 ==

......
......
.H....  (H covers T)
......
s.....

......
......
.TH...
......
s.....

After simulating the rope, you can count up all of the positions the tail visited at least once. In this diagram, s
again marks the starting position (which the tail also visited) and # marks other positions the tail visited:

..##..
...##.
.####.
....#.
s###..

So, there are 13 positions the tail visited at least once. Simulate your complete hypothetical series of motions. How
many positions does the tail of the rope visit at least once?

--- Part Two ---

A rope snaps! Suddenly, the river is getting a lot closer than you remember. The bridge is still there, but some of the
ropes that broke are now whipping toward you as you fall through the air! The ropes are moving too quickly to grab; you
only have a few seconds to choose how to arch your body to avoid being hit. Fortunately, your simulation can be extended
to support longer ropes. Rather than two knots, you now must simulate a rope consisting of ten knots. One knot is still
the head of the rope and moves according to the series of motions. Each knot further down the rope follows the knot in
front of it using the same rules as before. Using the same series of motions as the above example, but with the knots
marked H, 1, 2, ..., 9, the motions now occur as follows:

== Initial State ==

......
......
......
......
H.....  (H covers 1, 2, 3, 4, 5, 6, 7, 8, 9, s)

== R 4 ==

......
......
......
......
1H....  (1 covers 2, 3, 4, 5, 6, 7, 8, 9, s)

......
......
......
......
21H...  (2 covers 3, 4, 5, 6, 7, 8, 9, s)

......
......
......
......
321H..  (3 covers 4, 5, 6, 7, 8, 9, s)

......
......
......
......
4321H.  (4 covers 5, 6, 7, 8, 9, s)

== U 4 ==

......
......
......
....H.
4321..  (4 covers 5, 6, 7, 8, 9, s)

......
......
....H.
.4321.
5.....  (5 covers 6, 7, 8, 9, s)

......
....H.
....1.
.432..
5.....  (5 covers 6, 7, 8, 9, s)

....H.
....1.
..432.
.5....
6.....  (6 covers 7, 8, 9, s)

== L 3 ==

...H..
....1.
..432.
.5....
6.....  (6 covers 7, 8, 9, s)

..H1..
...2..
..43..
.5....
6.....  (6 covers 7, 8, 9, s)

.H1...
...2..
..43..
.5....
6.....  (6 covers 7, 8, 9, s)

== D 1 ==

..1...
.H.2..
..43..
.5....
6.....  (6 covers 7, 8, 9, s)

== R 4 ==

..1...
..H2..
..43..
.5....
6.....  (6 covers 7, 8, 9, s)

..1...
...H..  (H covers 2)
..43..
.5....
6.....  (6 covers 7, 8, 9, s)

......
...1H.  (1 covers 2)
..43..
.5....
6.....  (6 covers 7, 8, 9, s)

......
...21H
..43..
.5....
6.....  (6 covers 7, 8, 9, s)

== D 1 ==

......
...21.
..43.H
.5....
6.....  (6 covers 7, 8, 9, s)

== L 5 ==

......
...21.
..43H.
.5....
6.....  (6 covers 7, 8, 9, s)

......
...21.
..4H..  (H covers 3)
.5....
6.....  (6 covers 7, 8, 9, s)

......
...2..
..H1..  (H covers 4; 1 covers 3)
.5....
6.....  (6 covers 7, 8, 9, s)

......
...2..
.H13..  (1 covers 4)
.5....
6.....  (6 covers 7, 8, 9, s)

......
......
H123..  (2 covers 4)
.5....
6.....  (6 covers 7, 8, 9, s)

== R 2 ==

......
......
.H23..  (H covers 1; 2 covers 4)
.5....
6.....  (6 covers 7, 8, 9, s)

......
......
.1H3..  (H covers 2, 4)
.5....
6.....  (6 covers 7, 8, 9, s)

Now, you need to keep track of the positions the new tail, 9, visits. In this example, the tail never moves, and so it
only visits 1 position. However, be careful: more types of motion are possible than before, so you might want to
visually compare your simulated rope to the one above. Here's a larger example:

R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20

These motions occur as follows (individual steps are not shown):

== Initial State ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
...........H..............  (H covers 1, 2, 3, 4, 5, 6, 7, 8, 9, s)
..........................
..........................
..........................
..........................
..........................

== R 5 ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
...........54321H.........  (5 covers 6, 7, 8, 9, s)
..........................
..........................
..........................
..........................
..........................

== U 8 ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
................H.........
................1.........
................2.........
................3.........
...............54.........
..............6...........
.............7............
............8.............
...........9..............  (9 covers s)
..........................
..........................
..........................
..........................
..........................

== L 8 ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
........H1234.............
............5.............
............6.............
............7.............
............8.............
............9.............
..........................
..........................
...........s..............
..........................
..........................
..........................
..........................
..........................

== D 3 ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
.........2345.............
........1...6.............
........H...7.............
............8.............
............9.............
..........................
..........................
...........s..............
..........................
..........................
..........................
..........................
..........................

== R 17 ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
................987654321H
..........................
..........................
..........................
..........................
...........s..............
..........................
..........................
..........................
..........................
..........................

== D 10 ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
...........s.........98765
.........................4
.........................3
.........................2
.........................1
.........................H

== L 25 ==

..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
...........s..............
..........................
..........................
..........................
..........................
H123456789................

== U 20 ==

H.........................
1.........................
2.........................
3.........................
4.........................
5.........................
6.........................
7.........................
8.........................
9.........................
..........................
..........................
..........................
..........................
..........................
...........s..............
..........................
..........................
..........................
..........................
..........................

Now, the tail (9) visits 36 positions (including s) at least once:

..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
..........................
#.........................
#.............###.........
#............#...#........
.#..........#.....#.......
..#..........#.....#......
...#........#.......#.....
....#......s.........#....
.....#..............#.....
......#............#......
.......#..........#.......
........#........#........
.........########.........

Simulate your complete series of motions on a larger rope with ten knots. How many positions does the tail of the rope visit at least once?


*/

const DebugMove = false
const DebugInstruction = false
const DebugResult = false

type coordinate struct {
	x int
	y int
}

func (c *coordinate) moveToward(other *coordinate) bool {
	if util.AbsInt(other.x-c.x) <= 1 && util.AbsInt(other.y-c.y) <= 1 {
		// no need to move, we are close enough
		return false
	}
	if util.AbsInt(other.x-c.x) >= 1 {
		if c.x < other.x {
			c.x += 1
		} else {
			c.x -= 1
		}
	}
	if util.AbsInt(other.y-c.y) >= 1 {
		if c.y < other.y {
			c.y += 1
		} else {
			c.y -= 1
		}
	}
	return true
}

type rope struct {
	head    coordinate
	rest    []coordinate
	visited map[coordinate]struct{}
	drawing *util.Drawing
}

type direction byte

const (
	Up    direction = 'U'
	Right           = 'R'
	Down            = 'D'
	Left            = 'L'
)

type instruction struct {
	direction direction
	amount    int8
}

func initRope(numberOfKnots int) *rope {
	if numberOfKnots < 2 {
		panic(fmt.Sprintf("Number of knots, less than 2 😱: %drawing", numberOfKnots))
	}
	var rest []coordinate
	for i := 0; i < numberOfKnots-1; i++ {
		rest = append(rest, coordinate{0, 0})
	}
	return &rope{
		head:    coordinate{0, 0},
		rest:    rest,
		visited: map[coordinate]struct{}{coordinate{0, 0}: {}},
		drawing: util.InitDrawingWithCenter(50, 50, 25, 25).
			Fill('.').
			DrawAt('s', 0, 0),
	}
}

func (r *rope) draw(label string) {
	r.drawing.
		Fill('.').
		DrawAt('H', r.head.x, r.head.y)
	for i := 0; i < len(r.rest); i++ {
		r.drawing.DrawAt(byte('1'+i), r.rest[i].x, r.rest[i].y)
	}

	fmt.Printf("%s: \n%s\n\n", label, r.drawing.String())
}

func (r *rope) drawResult() {
	r.drawing.
		Fill('.').
		DrawAt('s', 0, 0)
	for k := range r.visited {
		r.drawing.DrawAt('X', k.x, k.y)
	}

	fmt.Printf("=== FINAL RESULT: \n%s\n", r.drawing.String())
}

func (r *rope) move(instruction *instruction) {
	for i := int8(0); i < instruction.amount; i++ {
		r.moveOne(instruction.direction)
	}
}

func (r *rope) moveOne(direction direction) {
	// do the move of the head
	switch direction {
	case Up:
		r.head.y += 1
	case Right:
		r.head.x += 1
	case Down:
		r.head.y -= 1
	case Left:
		r.head.x -= 1
	}

	// do the move of the other knots
	r.makeRestFollowHead()
	// store the new position of the tail
	r.visited[r.rest[len(r.rest)-1]] = struct{}{}

	if DebugMove {
		r.draw(fmt.Sprintf("+ intermediate %s: ", string(direction)))
	}
}

func (r *rope) makeRestFollowHead() {
	hasMoved := true
	previousKnot := &r.head
	for i := 0; i < len(r.rest) && hasMoved; i++ {
		hasMoved = r.rest[i].moveToward(previousKnot)
		previousKnot = &r.rest[i]
	}
}

func parse(input string) []instruction {
	var instructions []instruction
	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		tokens := strings.Fields(sc.Text())
		amount, _ := strconv.Atoi(tokens[1])
		instructions = append(
			instructions,
			instruction{
				direction: direction(tokens[0][0]),
				amount:    int8(amount),
			},
		)
	}
	return instructions
}

func doSolutionGen(raw string, numberOfKnots int) int {
	instructions := parse(raw)
	rope := initRope(numberOfKnots)

	for _, instruction := range instructions {
		rope.move(&instruction)

		if DebugInstruction {
			rope.draw(fmt.Sprintf("=== Result after instruction %#v: ", instruction))
		}
	}

	if DebugResult {
		rope.drawResult()
	}

	return len(rope.visited)
}

func doSolution1(raw string) int {
	return doSolutionGen(raw, 2)
}

func Solution1() int {
	return doSolution1(util.GetInputContent())
}

func doSolution2(raw string) int {
	return doSolutionGen(raw, 10)
}

func Solution2() int {
	return doSolution2(util.GetInputContent())
}

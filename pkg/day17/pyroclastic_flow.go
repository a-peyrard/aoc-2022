package day17

import (
	"aoc2022/pkg/util"
	"aoc2022/pkg/util/drawing"
	"aoc2022/pkg/util/geo"
	"fmt"
)

var emptyCharacter = byte('0')
var rockCharacter = byte('1')

type shape interface {
	moveLeft(drawing *drawing.Drawing) bool
	moveRight(drawing *drawing.Drawing) bool
	moveDown(drawing *drawing.Drawing) bool
	canMoveLeft(drawing *drawing.Drawing) bool
	canMoveRight(drawing *drawing.Drawing) bool
	canMoveDown(drawing *drawing.Drawing) bool
	draw(drawing *drawing.Drawing)
}

type abstractShape struct {
	position geo.Coordinate
	shape
}

func (s *abstractShape) moveLeft(drawing *drawing.Drawing) bool {
	canMove := false
	if s.canMoveLeft(drawing) {
		s.position.X -= 1
		canMove = true
	}

	return canMove
}

func (s *abstractShape) moveRight(drawing *drawing.Drawing) bool {
	canMove := false
	if s.canMoveRight(drawing) {
		s.position.X += 1
		canMove = true
	}

	return canMove
}

func (s *abstractShape) moveDown(drawing *drawing.Drawing) bool {
	canMove := false
	if s.canMoveDown(drawing) {
		s.position.Y -= 1
		canMove = true
	}

	return canMove
}

type flatLine struct {
	/*
		####
	*/
	abstractShape
}

func (s *flatLine) canMoveLeft(drawing *drawing.Drawing) bool {
	if s.position.X <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y) != emptyCharacter {
		return false
	}
	return true
}

func (s *flatLine) canMoveRight(drawing *drawing.Drawing) bool {
	if s.position.X >= 3 {
		return false
	}
	if drawing.ValueAt(s.position.X+4, s.position.Y) != emptyCharacter {
		return false
	}
	return true
}

func (s *flatLine) canMoveDown(drawing *drawing.Drawing) bool {
	if s.position.Y <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X, s.position.Y-1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y-1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+2, s.position.Y-1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+3, s.position.Y-1) != emptyCharacter {
		return false
	}
	return true
}

func (s *flatLine) draw(drawing *drawing.Drawing) {
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X+1, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X+2, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X+3, s.position.Y)
}

type plus struct {
	/*
	   .#.
	   ###
	   .#.
	*/
	abstractShape
}

func (s *plus) canMoveLeft(drawing *drawing.Drawing) bool {
	if s.position.X <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y+1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X, s.position.Y+2) != emptyCharacter {
		return false
	}
	return true
}

func (s *plus) canMoveRight(drawing *drawing.Drawing) bool {
	if s.position.X >= 4 {
		return false
	}
	if drawing.ValueAt(s.position.X+2, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+3, s.position.Y+1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+2, s.position.Y+2) != emptyCharacter {
		return false
	}
	return true
}

func (s *plus) canMoveDown(drawing *drawing.Drawing) bool {
	if s.position.Y <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y-1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+2, s.position.Y) != emptyCharacter {
		return false
	}
	return true
}

func (s *plus) draw(drawing *drawing.Drawing) {
	drawing.DrawAt(rockCharacter, s.position.X+1, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y+1)
	drawing.DrawAt(rockCharacter, s.position.X+1, s.position.Y+1)
	drawing.DrawAt(rockCharacter, s.position.X+2, s.position.Y+1)
	drawing.DrawAt(rockCharacter, s.position.X+1, s.position.Y+2)
}

type el struct {
	/*
		   	..#
			..#
			###
	*/
	abstractShape
}

func (s *el) canMoveLeft(drawing *drawing.Drawing) bool {
	if s.position.X <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y+1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y+2) != emptyCharacter {
		return false
	}
	return true
}

func (s *el) canMoveRight(drawing *drawing.Drawing) bool {
	if s.position.X >= 4 {
		return false
	}
	if drawing.ValueAt(s.position.X+3, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+3, s.position.Y+1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+3, s.position.Y+2) != emptyCharacter {
		return false
	}
	return true
}

func (s *el) canMoveDown(drawing *drawing.Drawing) bool {
	if s.position.Y <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X, s.position.Y-1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y-1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+2, s.position.Y-1) != emptyCharacter {
		return false
	}
	return true
}

func (s *el) draw(drawing *drawing.Drawing) {
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X+1, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X+2, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X+2, s.position.Y+1)
	drawing.DrawAt(rockCharacter, s.position.X+2, s.position.Y+2)
}

type verticalLine struct {
	/*
		   	#
			#
			#
			#
	*/
	abstractShape
}

func (s *verticalLine) canMoveLeft(drawing *drawing.Drawing) bool {
	if s.position.X <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y+1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y+2) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y+3) != emptyCharacter {
		return false
	}
	return true
}

func (s *verticalLine) canMoveRight(drawing *drawing.Drawing) bool {
	if s.position.X >= 6 {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y+1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y+2) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y+3) != emptyCharacter {
		return false
	}
	return true
}

func (s *verticalLine) canMoveDown(drawing *drawing.Drawing) bool {
	if s.position.Y <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X, s.position.Y-1) != emptyCharacter {
		return false
	}
	return true
}

func (s *verticalLine) draw(drawing *drawing.Drawing) {
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y+1)
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y+2)
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y+3)
}

type square struct {
	/*
		##
		##
	*/
	abstractShape
}

func (s *square) canMoveLeft(drawing *drawing.Drawing) bool {
	if s.position.X <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X-1, s.position.Y+1) != emptyCharacter {
		return false
	}
	return true
}

func (s *square) canMoveRight(drawing *drawing.Drawing) bool {
	if s.position.X >= 5 {
		return false
	}
	if drawing.ValueAt(s.position.X+2, s.position.Y) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+2, s.position.Y+1) != emptyCharacter {
		return false
	}
	return true
}

func (s *square) canMoveDown(drawing *drawing.Drawing) bool {
	if s.position.Y <= 0 {
		return false
	}
	if drawing.ValueAt(s.position.X, s.position.Y-1) != emptyCharacter {
		return false
	}
	if drawing.ValueAt(s.position.X+1, s.position.Y-1) != emptyCharacter {
		return false
	}
	return true
}

func (s *square) draw(drawing *drawing.Drawing) {
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X+1, s.position.Y)
	drawing.DrawAt(rockCharacter, s.position.X, s.position.Y+1)
	drawing.DrawAt(rockCharacter, s.position.X+1, s.position.Y+1)
}

var rocksFactories = []func(geo.Coordinate) shape{
	func(coordinate geo.Coordinate) shape {
		a := abstractShape{position: coordinate}
		s := &flatLine{a}
		s.abstractShape.shape = s
		return s
	},
	func(coordinate geo.Coordinate) shape {
		a := abstractShape{position: coordinate}
		s := &plus{a}
		s.abstractShape.shape = s
		return s
	},
	func(coordinate geo.Coordinate) shape {
		a := abstractShape{position: coordinate}
		s := &el{a}
		s.abstractShape.shape = s
		return s
	},
	func(coordinate geo.Coordinate) shape {
		a := abstractShape{position: coordinate}
		s := &verticalLine{a}
		s.abstractShape.shape = s
		return s
	},
	func(coordinate geo.Coordinate) shape {
		a := abstractShape{position: coordinate}
		s := &square{a}
		s.abstractShape.shape = s
		return s
	},
}

func findHighestRock(drawing *drawing.Drawing, startingFrom int) int {
	var noRocks bool
	for y := startingFrom; y < drawing.Height(); y++ {
		noRocks = true
		for x := 0; x < 7; x++ {
			if drawing.ValueAt(x, y) != emptyCharacter {
				noRocks = false
				break
			}
		}
		if noRocks {
			return y
		}
	}
	fmt.Printf("here is the drawing:\n%s\n\n", drawing.String())
	panic("Unable to find a row w/o rocks 😱")
}

func doSolution1(movements string) int {
	var (
		numberOfRocks     = 2022
		maxHeight         = numberOfRocks / 5 * 13
		d                 = drawing.InitDrawing(7, maxHeight).Fill(emptyCharacter)
		latestHighestRock = 0
		movementIdx       = 0
		canMove           bool
		rock              shape
		move              byte
	)
	for rockIdx := 0; rockIdx < numberOfRocks; rockIdx++ {
		latestHighestRock = findHighestRock(d, latestHighestRock)
		rock = rocksFactories[rockIdx%len(rocksFactories)](
			geo.Coordinate{X: 2, Y: latestHighestRock + 3},
		)
		canMove = true
		for canMove {
			move = movements[movementIdx%len(movements)]
			switch move {
			case '<':
				rock.moveLeft(d)
			case '>':
				rock.moveRight(d)
			}
			canMove = rock.moveDown(d)
			movementIdx++
		}
		rock.draw(d)
	}
	return findHighestRock(d, latestHighestRock)
}

func Solution1() int {
	return doSolution1(util.GetInputContent())
}

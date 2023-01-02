package day8

import (
	"aoc2022/pkg/util"
	"aoc2022/pkg/util/collection"
	"bufio"
	"strings"
)

/*
--- Day 8: Treetop Tree House ---

The expedition comes across a peculiar patch of tall trees all planted carefully in a grid. The Elves explain that a
previous expedition planted these trees as a reforestation effort. Now, they're curious if this would be a good location
for a tree house. First, determine whether there is enough tree cover here to keep a tree house hidden. To do this, you
need to count the number of trees that are visible from outside the grid when looking directly along a row or column.
The Elves have already launched a quadcopter to generate a map with the height of each tree (your puzzle input). For
example:

30373
25512
65332
33549
35390

Each tree is represented as a single digit whose value is its height, where 0 is the shortest and 9 is the tallest. A
tree is visible if all of the other trees between it and an edge of the grid are shorter than it. Only consider trees in
the same row or column; that is, only look up, down, left, or right from any given tree. All of the trees around the
edge of the grid are visible - since they are already on the edge, there are no trees to block the view. In this
example, that only leaves the interior nine trees to consider:

    The top-left 5 is visible from the left and top. (It isn't visible from the right or bottom since other trees of height 5 are in the way.)
    The top-middle 5 is visible from the top and right.
    The top-right 1 is not visible from any direction; for it to be visible, there would need to only be trees of height 0 between it and an edge.
    The left-middle 5 is visible, but only from the right.
    The center 3 is not visible from any direction; for it to be visible, there would need to be only trees of at most height 2 between it and an edge.
    The right-middle 3 is visible from the right.
    In the bottom row, the middle 5 is visible, but the 3 and 4 are not.

With 16 trees visible on the edge and another 5 visible in the interior, a total of 21 trees are visible in this
arrangement. Consider your map; how many trees are visible from outside the grid?

--- Part Two ---

Content with the amount of tree cover available, the Elves just need to know the best spot to build their tree house:
they would like to be able to see a lot of trees. To measure the viewing distance from a given tree, look up, down,
left, and right from that tree; stop if you reach an edge or at the first tree that is the same height or taller than
the tree under consideration. (If a tree is right on the edge, at least one of its viewing distances will be zero.) The
Elves don't care about distant trees taller than those found by the rules above; the proposed tree house has large eaves
to keep it dry, so they wouldn't be able to see higher than the tree house anyway. In the example above, consider the
middle 5 in the second row:

30373
25512
65332
33549
35390

    Looking up, its view is not blocked; it can see 1 tree (of height 3).
    Looking left, its view is blocked immediately; it can see only 1 tree (of height 5, right next to it).
    Looking right, its view is not blocked; it can see 2 trees.
    Looking down, its view is blocked eventually; it can see 2 trees (one of height 3, then the tree of height 5 that blocks its view).

A tree's scenic score is found by multiplying together its viewing distance in each of the four directions. For this
tree, this is 4 (found by multiplying 1 * 1 * 2 * 2). However, you can do even better: consider the tree of height 5 in
the middle of the fourth row:

30373
25512
65332
33549
35390

    Looking up, its view is blocked at 2 trees (by another tree with a height of 5).
    Looking left, its view is not blocked; it can see 2 trees.
    Looking down, its view is also not blocked; it can see 1 tree.
    Looking right, its view is blocked at 2 trees (by a massive tree of height 9).

This tree's scenic score is 8 (2 * 2 * 1 * 2); this is the ideal spot for the tree house.
Consider each tree on your map. What is the highest scenic score possible for any tree?


trees: 5 1 2 3 1 2 6
dist:  0 0 1 2 0 1 6

*/

type tree struct {
	x int
	y int
}

func parse(input string) [][]int8 {
	results := make([][]int8, 0)
	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		results = append(results, parseLine(sc.Text()))
	}

	return results
}

func parseLine(line string) []int8 {
	row := make([]int8, 0)
	for i := 0; i < len(line); i++ {
		row = append(row, int8(line[i]-'0'))
	}

	return row
}

func countVisibleTrees(forest [][]int8) int {
	visibleTrees := map[tree]struct{}{}

	var maxSoFar, maxSoFarReverse, value, valueReverse int8
	height, width := len(forest), len(forest[0])

	// analyze visible trees from the left side and right side
	for i := 1; i < height-1; i++ {
		maxSoFar = -1
		maxSoFarReverse = -1
		for j := 0; j < width; j++ {
			value = forest[i][j]
			valueReverse = forest[i][width-1-j]
			if value > maxSoFar {
				visibleTrees[tree{i, j}] = struct{}{}
				maxSoFar = value
			}
			if valueReverse > maxSoFarReverse {
				visibleTrees[tree{i, width - 1 - j}] = struct{}{}
				maxSoFarReverse = valueReverse
			}
		}
	}

	// analyze visible trees from top and bottom side
	for i := 1; i < width-1; i++ {
		maxSoFar = -1
		maxSoFarReverse = -1
		for j := 0; j < height; j++ {
			value = forest[j][i]
			valueReverse = forest[height-1-j][i]
			if value > maxSoFar {
				visibleTrees[tree{j, i}] = struct{}{}
				maxSoFar = value
			}
			if valueReverse > maxSoFarReverse {
				visibleTrees[tree{height - 1 - j, i}] = struct{}{}
				maxSoFarReverse = valueReverse
			}
		}
	}

	return len(visibleTrees) + 4 // the four corner of the forest that have not been analyzed
}

func searchTreeWithMostVisibility(forest [][]int8) int {
	treeVisibilities := map[tree]int{}

	type treeInLine struct {
		height int8
		index  int
	}

	maxStack := collection.CreateStack[*treeInLine]()
	maxStackReverse := collection.CreateStack[*treeInLine]()
	height, width := len(forest), len(forest[0])
	var value, valueReverse int8
	var visibility, treeValue, reverseIndex int
	var exist, empty bool

	// analyze visible trees from the left side and right side
	for i := 1; i < height-1; i++ {
		maxStack.Clear()
		maxStackReverse.Clear()
		for j := 0; j < width; j++ {
			value = forest[i][j]
			empty = maxStack.RemoveTill(func(t *treeInLine) bool { return t.height < value })
			if empty {
				visibility = j
			} else {
				visibility = j - maxStack.Peek().index
			}
			treeValue, exist = treeVisibilities[tree{i, j}]
			if exist {
				treeVisibilities[tree{i, j}] = treeValue * visibility
			} else {
				treeVisibilities[tree{i, j}] = visibility
			}
			maxStack.Push(&treeInLine{value, j})

			reverseIndex = width - 1 - j
			valueReverse = forest[i][reverseIndex]
			empty = maxStackReverse.RemoveTill(func(t *treeInLine) bool { return t.height < valueReverse })
			if empty {
				visibility = width - reverseIndex - 1
			} else {
				visibility = maxStackReverse.Peek().index - reverseIndex
			}
			treeValue, exist = treeVisibilities[tree{i, reverseIndex}]
			if exist {
				treeVisibilities[tree{i, reverseIndex}] = treeValue * visibility
			} else {
				treeVisibilities[tree{i, reverseIndex}] = visibility
			}
			maxStackReverse.Push(&treeInLine{valueReverse, reverseIndex})
		}
	}

	// analyze visible trees from top and bottom side
	for i := 1; i < width-1; i++ {
		maxStack.Clear()
		maxStackReverse.Clear()
		for j := 0; j < height; j++ {
			value = forest[j][i]
			empty = maxStack.RemoveTill(func(t *treeInLine) bool { return t.height < value })
			if empty {
				visibility = j
			} else {
				visibility = j - maxStack.Peek().index
			}
			treeValue, exist = treeVisibilities[tree{j, i}]
			if exist {
				treeVisibilities[tree{j, i}] = treeValue * visibility
			} else {
				treeVisibilities[tree{j, i}] = visibility
			}
			maxStack.Push(&treeInLine{value, j})

			reverseIndex = height - 1 - j
			valueReverse = forest[reverseIndex][i]
			empty = maxStackReverse.RemoveTill(func(t *treeInLine) bool { return t.height < valueReverse })
			if empty {
				visibility = width - reverseIndex - 1
			} else {
				visibility = maxStackReverse.Peek().index - reverseIndex
			}
			treeValue, exist = treeVisibilities[tree{reverseIndex, i}]
			if exist {
				treeVisibilities[tree{reverseIndex, i}] = treeValue * visibility
			} else {
				treeVisibilities[tree{reverseIndex, i}] = visibility
			}
			maxStackReverse.Push(&treeInLine{valueReverse, reverseIndex})
		}
	}

	maxVisibility := 0
	for _, visibility := range treeVisibilities {
		if visibility > maxVisibility {
			maxVisibility = visibility
		}
	}
	return maxVisibility
}

func doSolution1(raw string) int {
	return countVisibleTrees(parse(raw))
}

func Solution1() int {
	return doSolution1(util.GetInputContent())
}

func doSolution2(raw string) int {
	return searchTreeWithMostVisibility(parse(raw))
}

func Solution2() int {
	return doSolution2(util.GetInputContent())
}

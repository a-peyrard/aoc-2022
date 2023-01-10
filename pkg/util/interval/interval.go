package interval

import (
	"aoc2022/pkg/util"
	"fmt"
	"sort"
)

type Interval [2]int

func CreateInterval(from, to int) *Interval {
	if from > to {
		panic(fmt.Sprintf("From can not be greater than the to: [%d, %d]", from, to))
	}
	return &Interval{from, to}
}

func (i *Interval) Length() int {
	return i[1] - i[0] + 1
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i[0], i[1])
}

func (i *Interval) include(x int) bool {
	return i[0] <= x && x <= i[1]
}

func MergeIntervals(intervals []*Interval) []*Interval {
	// sort by start date
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	current := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[current][1] < intervals[i][0] {
			// no overlap with current
			current++
			intervals[current] = intervals[i]
		} else {
			intervals[current][1] = util.Max(intervals[current][1], intervals[i][1])
		}
	}
	return intervals[:current+1]
}

func IncludeInMergedIntervals(val int, intervals []*Interval) bool {
	for _, in := range intervals {
		if val < in[0] {
			// as the intervals are sorted and merged, if our value is lower than the start point
			// of the current interval, the value will be lower than all the other intervals
			break
		}
		if in.include(val) {
			return true
		}
	}
	return false
}

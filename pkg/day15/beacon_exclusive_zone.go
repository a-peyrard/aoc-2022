package day15

import (
	"aoc2022/pkg/util"
	"aoc2022/pkg/util/collection"
	"aoc2022/pkg/util/drawing"
	"aoc2022/pkg/util/geo"
	"aoc2022/pkg/util/interval"
	"bufio"
	"fmt"
	"math"
	"strings"
)

/*
--- Day 15: Beacon Exclusion Zone ---

You feel the ground rumble again as the distress signal leads you to a large network of subterranean tunnels. You don't
have time to search them all, but you don't need to: your pack contains a set of deployable sensors that you imagine
were originally built to locate lost Elves. The sensors aren't very powerful, but that's okay; your handheld device
indicates that you're close enough to the source of the distress signal to use them. You pull the emergency sensor
system out of your pack, hit the big button on top, and the sensors zoom off down the tunnels. Once a sensor finds a
spot it thinks will give it a good reading, it attaches itself to a hard surface and begins monitoring for the nearest
signal source beacon. Sensors and beacons always exist at integer coordinates. Each sensor knows its own position and
can determine the position of a beacon precisely; however, sensors can only lock on to the one beacon closest to the
sensor as measured by the Manhattan distance. (There is never a tie where two beacons are the same distance to a
sensor.) It doesn't take long for the sensors to report back their positions and closest beacons (your puzzle input).
For example:

Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3

So, consider the sensor at 2,18; the closest beacon to it is at -2,15. For the sensor at 9,16, the closest beacon to it
is at 10,16. Drawing sensors as S and beacons as B, the above arrangement of sensors and beacons looks like this:

               1    1    2    2
     0    5    0    5    0    5
 0 ....S.......................
 1 ......................S.....
 2 ...............S............
 3 ................SB..........
 4 ............................
 5 ............................
 6 ............................
 7 ..........S.......S.........
 8 ............................
 9 ............................
10 ....B.......................
11 ..S.........................
12 ............................
13 ............................
14 ..............S.......S.....
15 B...........................
16 ...........SB...............
17 ................S..........B
18 ....S.......................
19 ............................
20 ............S......S........
21 ............................
22 .......................B....

This isn't necessarily a comprehensive map of all beacons in the area, though. Because each sensor only identifies its
closest beacon, if a sensor detects a beacon, you know there are no other beacons that close or closer to that sensor.
There could still be beacons that just happen to not be the closest beacon to any sensor. Consider the sensor at 8,7:

               1    1    2    2
     0    5    0    5    0    5
-2 ..........#.................
-1 .........###................
 0 ....S...#####...............
 1 .......#######........S.....
 2 ......#########S............
 3 .....###########SB..........
 4 ....#############...........
 5 ...###############..........
 6 ..#################.........
 7 .#########S#######S#........
 8 ..#################.........
 9 ...###############..........
10 ....B############...........
11 ..S..###########............
12 ......#########.............
13 .......#######..............
14 ........#####.S.......S.....
15 B........###................
16 ..........#SB...............
17 ................S..........B
18 ....S.......................
19 ............................
20 ............S......S........
21 ............................
22 .......................B....

This sensor's closest beacon is at 2,10, and so you know there are no beacons that close or closer (in any positions
marked #). None of the detected beacons seem to be producing the distress signal, so you'll need to work out where the
distress beacon is by working out where it isn't. For now, keep things simple by counting the positions where a beacon
cannot possibly be along just a single row. So, suppose you have an arrangement of beacons and sensors like in the
example above and, just in the row where y=10, you'd like to count the number of positions a beacon cannot possibly
exist. The coverage from all sensors near that row looks like this:

                 1    1    2    2
       0    5    0    5    0    5
 9 ...#########################...
10 ..####B######################..
11 .###S#############.###########.

In this example, in the row where y=10, there are 26 positions where a beacon cannot be present. Consult the report from
the sensors you just deployed. In the row where y=2000000, how many positions cannot contain a beacon?

--- Part Two ---

Your handheld device indicates that the distress signal is coming from a beacon nearby. The distress beacon is not
detected by any sensor, but the distress beacon must have x and y coordinates each no lower than 0 and no larger than
4000000. To isolate the distress beacon's signal, you need to determine its tuning frequency, which can be found by
multiplying its x coordinate by 4000000 and then adding its y coordinate. In the example above, the search space is
smaller: instead, the x and y coordinates can each be at most 20. With this reduced search area, there is only a single
position that could have a beacon: x=14, y=11. The tuning frequency for this distress beacon is 56000011. Find the only
possible position for the distress beacon. What is its tuning frequency?

*/

const Debug = false

type sensorAndBeacon struct {
	sensor            geo.Coordinate
	beacon            geo.Coordinate
	manhattanDistance int
}

func createSensorAndBeacon(sensor geo.Coordinate, beacon geo.Coordinate) sensorAndBeacon {
	return sensorAndBeacon{
		sensor:            sensor,
		beacon:            beacon,
		manhattanDistance: sensor.ManhattanDistance(beacon),
	}
}

func (sb *sensorAndBeacon) updateMinAndMax(minX, maxX, minY, maxY *int) {
	updateMinAndMax(&sb.sensor, minX, maxX, minY, maxY, sb.manhattanDistance)
	updateMinAndMax(&sb.beacon, minX, maxX, minY, maxY, sb.manhattanDistance)
}

func (sb *sensorAndBeacon) coverageAt(y int) *interval.Interval {
	distanceToRow := util.AbsInt(sb.sensor.Y - y)
	if distanceToRow > sb.manhattanDistance {
		return nil // no coverage for this row
	}
	coverageWidth := sb.manhattanDistance - distanceToRow
	return interval.CreateInterval(
		sb.sensor.X-coverageWidth,
		sb.sensor.X+coverageWidth,
	)
}

func updateMinAndMax(c *geo.Coordinate, minX, maxX, minY, maxY *int, distance int) {
	if c.X-distance < *minX {
		*minX = c.X - distance
	}
	if c.X+distance > *maxX {
		*maxX = c.X + distance
	}
	if c.Y-distance < *minY {
		*minY = c.Y + distance
	}
	if c.Y+distance > *maxY {
		*maxY = c.Y + distance
	}
}

func parse(input string) ([]sensorAndBeacon, int, int, int, int) {
	var (
		res        []sensorAndBeacon
		tokens     []string
		sb         sensorAndBeacon
		maxX, maxY = 0, 0
		minX, minY = math.MaxInt, math.MaxInt
	)

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		tokens = strings.Split(sc.Text(), ": ")
		sb = createSensorAndBeacon(
			parseCoordinate(tokens[0][10:]),
			parseCoordinate(tokens[1][21:]),
		)
		sb.updateMinAndMax(&minX, &maxX, &minY, &maxY)
		res = append(res, sb)
	}
	return res, minX, maxX, minY, maxY
}

func parseCoordinate(rawCoordinate string) geo.Coordinate {
	tokens := strings.Split(rawCoordinate, ", ")
	return geo.Coordinate{
		X: util.UnsafeParseInt(tokens[0][2:]),
		Y: util.UnsafeParseInt(tokens[1][2:]),
	}
}

func doSolution1(raw string, row int) int {
	sbList, _, _, _, _ := parse(raw)

	var (
		intervals = make([]*interval.Interval, 0)
		coverage  *interval.Interval
	)
	for _, sb := range sbList {
		coverage = sb.coverageAt(row)
		if coverage != nil {
			intervals = append(intervals, coverage)
		}
	}

	mergedIntervals := interval.MergeIntervals(intervals)
	res := 0
	for _, in := range mergedIntervals {
		res += in.Length()
	}

	// remove existing beacons
	beaconToRemove := collection.NewSet[int]()
	for _, sb := range sbList {
		if sb.beacon.Y == row && !beaconToRemove.Contains(sb.beacon.X) {
			if interval.IncludeInMergedIntervals(sb.beacon.X, mergedIntervals) {
				beaconToRemove.Add(sb.beacon.X)
			}
		}
	}
	res -= beaconToRemove.Length()

	if Debug {
		d := drawing.InitDrawingWithCenter(31, 3, 4, 0).
			Fill(' ').
			DrawAt('0', 0, 1).
			DrawAt('5', 5, 1).
			DrawAt('0', 10, 1).
			DrawAt('5', 15, 1).
			DrawAt('0', 20, 1).
			DrawAt('5', 25, 1).
			DrawAt('1', 10, 2).
			DrawAt('1', 15, 2).
			DrawAt('2', 20, 2).
			DrawAt('2', 25, 2).
			DrawLine('.', geo.Coordinate{X: -4}, geo.Coordinate{X: 26})

		for _, i := range mergedIntervals {
			d.DrawLine('#', geo.Coordinate{X: i[0]}, geo.Coordinate{X: i[1]})
		}

		fmt.Printf("Here is a debug drawing...\n%s\n\n", d.String())
	}

	return res
}

func Solution1() int {
	return doSolution1(util.GetInputContent(), 2_000_000)
}

func doSolution2(raw string, maxCoordinate int) int {
	sbList, _, _, _, _ := parse(raw)

	var (
		intervals     []*interval.Interval
		coverage      *interval.Interval
		availableCell int
	)
	for i := 0; i <= maxCoordinate; i++ {
		intervals = make([]*interval.Interval, 0)
		for _, sb := range sbList {
			coverage = sb.coverageAt(i)
			if coverage != nil {
				intervals = append(intervals, coverage)
			}
		}
		availableCell = findAvailableCell(interval.MergeIntervals(intervals), maxCoordinate)
		if availableCell >= 0 {
			return tuningFrequency(availableCell, i)
		}
	}

	panic("Unable to find the beacon position :/")
}

func findAvailableCell(intervals []*interval.Interval, maxCoordinate int) int {
	if len(intervals) == 1 {
		if intervals[0][0] > 0 {
			return 0
		}
		if intervals[0][1] < maxCoordinate {
			return maxCoordinate
		}
	} else {
		for i := 0; i < len(intervals); i++ {
			if intervals[i][1] >= 0 && intervals[i][1] < maxCoordinate {
				return intervals[i][1] + 1
			}
		}
	}
	return -1
}

func tuningFrequency(x, y int) int {
	return x*4000000 + y
}

func Solution2() int {
	return doSolution2(util.GetInputContent(), 4_000_000)
}

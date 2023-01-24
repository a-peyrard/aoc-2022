package day16

import (
	"aoc2022/pkg/util"
	"aoc2022/pkg/util/collection"
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"math"
	"regexp"
	"sort"
	"strings"
)

const Debug = true

type valve struct {
	label  string
	flow   int
	leadTo []string
}

var valveRegex = regexp.MustCompile("^Valve (\\w+) has flow rate=(\\d+); tunnels? leads? to valves? (.*)$")

func parse(input string) map[string]*valve {
	sc := bufio.NewScanner(strings.NewReader(input))
	var (
		res     = map[string]*valve{}
		matches []string
		label   string
	)
	for sc.Scan() {
		matches = valveRegex.FindStringSubmatch(sc.Text())
		label = matches[1]
		res[label] = &valve{
			label:  label,
			flow:   util.UnsafeParseInt(matches[2]),
			leadTo: strings.Split(matches[3], ", "),
		}
	}

	return res
}

func buildGraph(valves map[string]*valve) (collection.Matrix[int], map[string]int, []string) {
	// start by building the mapping tables
	labelToIndex := map[string]int{}
	indexToLabel := make([]string, len(valves))

	index := 0
	for label := range valves {
		labelToIndex[label] = index
		indexToLabel[index] = label
		index++
	}

	// build the graph
	graph := collection.CreateMatrix[int](len(valves), len(valves), 0)
	for _, valve := range valves {
		for _, successor := range valve.leadTo {
			graph[labelToIndex[valve.label]][labelToIndex[successor]] = 1
		}
	}

	return graph, labelToIndex, indexToLabel
}

func shortestPaths(graph collection.Matrix[int]) collection.Matrix[int] {
	/*
		let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)
		for each edge (u, v) do
			dist[u][v] ← w(u, v)  // The weight of the edge (u, v)
		for each vertex v do
			dist[v][v] ← 0
		for k from 1 to |V|
			for i from 1 to |V|
				for j from 1 to |V|
					if dist[i][j] > dist[i][k] + dist[k][j]
						dist[i][j] ← dist[i][k] + dist[k][j]
					end if
	*/
	distances := collection.CreateMatrix[int](graph.Width(), graph.Height(), math.MaxInt)
	for i := 0; i < graph.Height(); i++ {
		for j := 0; j < graph.Height(); j++ {
			if i == j {
				distances[i][j] = 0
			}
			if graph[i][j] > 0 {
				distances[i][j] = graph[i][j]
			}
		}
	}
	for k := 0; k < graph.Height(); k++ {
		for i := 0; i < graph.Height(); i++ {
			for j := 0; j < graph.Height(); j++ {
				if distances[i][k] < math.MaxInt && // could have use float64 to have infinity?!
					distances[k][j] < math.MaxInt &&
					distances[i][j] > distances[i][k]+distances[k][j] {

					distances[i][j] = distances[i][k] + distances[k][j]
				}
			}
		}
	}
	return distances
}

func visit(path int, vertexIndex int) int {
	path |= 1 << vertexIndex
	return path
}

func isVisited(path int, vertexIndex int) bool {
	return path&(1<<vertexIndex) > 0
}

// check if the value for the best path could be beat
// analyzing the best scenario possible where all valves are adjacent
// (just 1 minute to travel from one to another)
func pathMayBeat(bestValue int,
	path int,
	currentValue int,
	elapsedTime int,
	maxTime int,
	valvesOrdered []*valve,
	labelToIndex map[string]int,
	uselessValves collection.Set[int]) bool {

	maybeValue := currentValue
	var (
		valveFound bool
		valveIndex int
	)
	for i := elapsedTime + 2; i <= maxTime; i += 2 {
		valveFound = false
		for _, valve := range valvesOrdered {
			valveIndex = labelToIndex[valve.label]
			if uselessValves.DoesNotContain(valveIndex) && !isVisited(path, valveIndex) {
				maybeValue += valve.flow * (maxTime - i)
				valveFound = true
				path = visit(path, valveIndex)
				break
			}
		}
		if !valveFound {
			break
		}
	}

	return maybeValue > bestValue
}

func valvesByFlowDesc(valves map[string]*valve) []*valve {
	values := maps.Values(valves)
	sort.Slice(values, func(i, j int) bool {
		return values[i].flow > values[j].flow
	})
	return values
}

func doSolution1(raw string) int {
	valves := parse(raw)
	graph, labelToIndex, indexToLabel := buildGraph(valves)
	distances := shortestPaths(graph)
	valvesOrdered := valvesByFlowDesc(valves)

	flow, _ := findBestPath(
		labelToIndex["AA"],
		distances,
		valves,
		valvesOrdered,
		labelToIndex,
		indexToLabel,
		30,
		collection.NewSet[int](),
	)

	return flow
}

func findBestPath(rootIndex int,
	distances collection.Matrix[int],
	valves map[string]*valve,
	valvesOrdered []*valve,
	labelToIndex map[string]int,
	indexToLabel []string,
	maxTime int,
	forbiddenValves collection.Set[int]) (int, int) {

	uselessValveIndexes := collection.NewSet[int]()
	uselessValveIndexes.AddAll(forbiddenValves)
	for i := 0; i < len(indexToLabel); i++ {
		if valves[indexToLabel[i]].flow == 0 {
			uselessValveIndexes.Add(i)
		}
	}
	// trigger to all first neighbors
	var (
		best                      = 0
		path                      = visit(0, rootIndex)
		branchValue, usedNeighbor int
	)

	for neighbor, distance := range distances[rootIndex] {
		if neighbor != rootIndex && distance < math.MaxInt && uselessValveIndexes.DoesNotContain(neighbor) {
			branchValue = findBestPathRec(
				neighbor,
				path,
				distance,
				0,
				distances,
				valves,
				valvesOrdered,
				uselessValveIndexes,
				labelToIndex,
				indexToLabel,
				maxTime,
				best,
			)
			if branchValue > best {
				best = branchValue
				usedNeighbor = neighbor
			}
		}
	}
	return best, usedNeighbor
}

func findBestPathRec(current int,
	path int,
	elapsedTime int,
	value int,
	distances collection.Matrix[int],
	valves map[string]*valve,
	valvesOrdered []*valve,
	uselessValveIndexes collection.Set[int],
	labelToIndex map[string]int,
	indexToLabel []string,
	maxTime int,
	best int) int {

	if elapsedTime >= maxTime {
		return value
	}

	// visit the current valve
	path = visit(path, current)

	// get current valve flow, and open it
	elapsedTime += 1
	value += valves[indexToLabel[current]].flow * (maxTime - elapsedTime)

	// do a DFS, and don't try to go to valves with 0 flow
	var (
		branchValue, localBest int
	)
	for next, distance := range distances[current] {
		if current != next && distance < math.MaxInt && uselessValveIndexes.DoesNotContain(next) {
			if !isVisited(path, next) &&
				pathMayBeat(util.Max(best, localBest), path, value, elapsedTime, maxTime, valvesOrdered, labelToIndex, uselessValveIndexes) {
				branchValue = findBestPathRec(
					next,
					path,
					elapsedTime+distance,
					value,
					distances,
					valves,
					valvesOrdered,
					uselessValveIndexes,
					labelToIndex,
					indexToLabel,
					maxTime,
					best,
				)
				if branchValue > localBest {
					localBest = branchValue
				}
			}
		}
	}

	return util.Max(localBest, value)
}

func Solution1() int {
	return doSolution1(util.GetInputContent())
}

type runner struct {
	label    string
	position int
	timeLeft int
	flow     int
}

func doSolution2(raw string) int {
	valves := parse(raw)
	graph, labelToIndex, indexToLabel := buildGraph(valves)
	distances := shortestPaths(graph)
	valvesOrdered := valvesByFlowDesc(valves)

	var (
		currentRunner                      *runner
		bestFlow, firstUsedValve, distance int
		me                                 = runner{
			label:    "me      ",
			position: labelToIndex["AA"],
			timeLeft: 26,
			flow:     0,
		}
		elephant = runner{
			label:    "elephant",
			position: labelToIndex["AA"],
			timeLeft: 26,
			flow:     0,
		}
		forbiddenValves = collection.NewSet[int]()
	)
	for me.timeLeft > 0 || elephant.timeLeft > 0 {
		if me.timeLeft >= elephant.timeLeft {
			currentRunner = &me
		} else {
			currentRunner = &elephant
		}
		bestFlow, firstUsedValve = findBestPath(
			currentRunner.position,
			distances,
			valves,
			valvesOrdered,
			labelToIndex,
			indexToLabel,
			currentRunner.timeLeft,
			forbiddenValves,
		)
		if bestFlow == 0 {
			currentRunner.timeLeft = 0
		} else {
			distance = distances[currentRunner.position][firstUsedValve]
			currentRunner.timeLeft -= distance + 1
			currentRunner.flow += currentRunner.timeLeft * valves[indexToLabel[firstUsedValve]].flow
			currentRunner.position = firstUsedValve
			forbiddenValves.Add(firstUsedValve)
		}

		if Debug {
			if bestFlow > 0 {
				fmt.Printf(
					"runner %s move to %s (distance %d), with %d time left, current flow is %d\n",
					currentRunner.label,
					indexToLabel[firstUsedValve],
					distance,
					currentRunner.timeLeft,
					currentRunner.flow,
				)
			} else {
				fmt.Printf(
					"runner %s has no more possible steps\n",
					currentRunner.label,
				)
			}
		}
	}

	return me.flow + elephant.flow
}

func Solution2() int {
	return doSolution2(util.GetInputContent())
}

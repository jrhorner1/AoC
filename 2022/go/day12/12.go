package day12

import (
	"image"
	"math"
	"sort"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	g := &graph{getNode: make(map[image.Point]*node)}
	var start, end image.Point
	var potentialStart []image.Point
	for y, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		for x, char := range line {
			node := &node{image.Point{x, y}, math.MaxInt, 0}
			switch char {
			case 'a':
				potentialStart = append(potentialStart, image.Point{x, y})
				node.height = 0
			case 'S': // start
				node.height = 0
				start = image.Point{x, y}
			case 'E': // end
				node.height = 25
				end = image.Point{x, y}
			default:
				node.height = int(char) - int('a')
			}
			g.addNode(node)
		}
	}
	if part2 {
		var potentialStartSteps []int
		for _, start := range potentialStart {
			if steps := g.shortestPath(start, end); steps > 0 {
				potentialStartSteps = append(potentialStartSteps, steps)
			}
		}
		sort.Ints(potentialStartSteps)
		return potentialStartSteps[0]
	}
	return g.shortestPath(start, end)
}

type node struct {
	id     image.Point
	steps  int
	height int
}

type graph struct {
	nodes   []*node
	getNode map[image.Point]*node
}

func (g *graph) addNode(n *node) {
	if _, found := g.getNode[n.id]; !found {
		g.nodes = append(g.nodes, n)
		g.getNode[n.id] = n
	}
}

type queue []*node

func (q *queue) push(n *node) { *q = append(*q, n) }
func (q *queue) pull() *node  { n := (*q)[0]; *q = (*q)[1:]; return n }

func (g *graph) shortestPath(start, end image.Point) int {
	q := queue{g.getNode[start]}
	visited := make(map[image.Point]bool)
	g.getNode[start].steps = 0
	visited[start] = true
	for len(q) > 0 {
		current := q.pull()
		if current.id == end {
			return current.steps
		}
		for _, delta := range []image.Point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} { // up, down, left, right
			if neighbor, exists := g.getNode[current.id.Add(delta)]; exists && !visited[neighbor.id] {
				if neighbor.height <= current.height+1 {
					neighbor.steps = current.steps + 1
					visited[neighbor.id] = true
					q.push(neighbor)
				}
			}
		}
	}
	return -1
}

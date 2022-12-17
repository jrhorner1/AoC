package day16

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	valves := newGraph()
	valveDests := make(map[string]destinations)
	for i, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		valve := node{visited: false}
		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &valve.id, &valve.value)
		if err != nil {
			logrus.Errorf("Line %d: %s", i, err)
		}
		re := regexp.MustCompile("[A-Z]{2}")
		tunnelDests := re.FindAll([]byte(line), -1)
		dests := []string{}
		for _, dest := range tunnelDests {
			if string(dest) != valve.id {
				dests = append(dests, string(dest))
			}
		}
		valveDests[valve.id] = dests
		valves.addNode(&valve)
	}
	defaultWeight := len(valves.nodes) + 1
	for _, v1 := range valves.nodes {
		for _, v2 := range valves.nodes {
			if v1 == v2 { // set weight of node to itself to 0
				valves.addEdge(v1, v2, 0)
			} else if valveDests[v1.id].contains(v2.id) { // set the weight of node to all ajacent nodes to 1
				valves.addEdge(v1, v2, 1)
			} else { // set weight of node to all other nodes to MaxInt
				valves.addEdge(v1, v2, defaultWeight)
			}
		}
	}
	// floydWarshall algorithm https://en.wikipedia.org/wiki/Floyd–Warshall_algorithm
	for _, k := range valves.nodes {
		for _, i := range valves.nodes {
			for _, j := range valves.nodes {
				ik := valves.edges[i].getEdge(k)
				kj := valves.edges[k].getEdge(j)
				ij := valves.edges[i].getEdge(j)
				if ij.weight > ik.weight+kj.weight {
					ij.weight = ik.weight + kj.weight
				}
			}
		}
	}
	for _, node := range valves.nodes {
		if node.value == 0 && node.id != "AA" {
			delete(valves.nodes, node.id)
			delete(valves.edges, node)
			for _, others := range valves.nodes {
				for i, edge := range valves.edges[others] {
					if edge.dest == node {
						temp := edges{}
						if i < len(valves.edges[others])-1 {
							temp = valves.edges[others][i+1:]
						}
						valves.edges[others] = valves.edges[others][:i]
						valves.edges[others] = append(valves.edges[others], temp...)
					}
				}
			}
		}
	}
	start := valves.nodes["AA"]
	maxPressure := valves.depthFirstSearch(0, 0, start)
	if part2 {
		paths := valves.depthFirstSearchPaths(0, 0, start, "")
		for i := 0; i < len(paths); i++ {
		loop:
			for j := i + 1; j < len(paths); j++ {
				for si := 0; si+1 < len(paths[i][1].(string)); si += 2 {
					substring := paths[i][1].(string)[si : si+2]
					if strings.Contains(paths[j][1].(string), substring) {
						continue loop
					}
				}
				if max := paths[i][0].(int) + paths[j][0].(int); max > maxPressure {
					maxPressure = max
				}
			}
		}
		return maxPressure
	}
	return maxPressure
}

type destinations []string

func (d destinations) contains(id string) bool {
	for _, destId := range d {
		if destId == id {
			return true
		}
	}
	return false
}

type node struct {
	id      string
	value   int
	visited bool
}

type edge struct {
	dest   *node
	weight int
}

type edges []*edge

func (e edges) getEdge(d *node) *edge {
	for _, edge := range e {
		if edge.dest == d {
			return edge
		}
	}
	return &edge{}
}

type graph struct {
	nodes map[string]*node
	edges map[*node]edges
}

func newGraph() *graph {
	return &graph{
		nodes: make(map[string]*node),
		edges: make(map[*node]edges),
	}
}

func (g *graph) addNode(n *node) {
	if _, found := g.nodes[n.id]; !found {
		g.nodes[n.id] = n
	}
}

func (g *graph) addEdge(n1, n2 *node, weight int) {
	g.edges[n1] = append(g.edges[n1], &edge{n2, weight})
}

func (g *graph) depthFirstSearch(pressure, minute int, currentNode *node) int {
	max := pressure
	for _, node := range g.nodes {
		if node == currentNode || node.id == "AA" || node.visited {
			continue
		}
		openTime := g.edges[currentNode].getEdge(node).weight + 1
		if minute+openTime > 30 {
			continue
		}
		node.visited = true
		next := g.depthFirstSearch(pressure+(30-minute-openTime)*node.value, minute+openTime, node)
		if next > max {
			max = next
		}
		node.visited = false
	}
	return max
}

func (g *graph) depthFirstSearchPaths(pressure, minute int, currentNode *node, path string) [][2]any {
	paths := [][2]any{{pressure, path}}
	for _, node := range g.nodes {
		if node == currentNode || node.id == "AA" || node.visited {
			continue
		}
		openTime := g.edges[currentNode].getEdge(node).weight + 1
		if minute+openTime > 26 {
			continue
		}
		node.visited = true
		paths = append(paths, g.depthFirstSearchPaths(pressure+(26-minute-openTime)*node.value, minute+openTime, node, path+node.id)...)
		node.visited = false
	}
	return paths
}

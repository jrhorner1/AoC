package day12

import (
	"strings"
	"unicode"
)

type node struct {
	id      string
	isBig   bool
	visited int
}

type graph struct {
	nodes   []*node
	getNode map[string]*node
	edges   map[*node][]*node
}

func (g *graph) addNode(n *node) {
	if _, found := g.getNode[n.id]; !found {
		g.nodes = append(g.nodes, n)
		if g.getNode == nil {
			g.getNode = make(map[string]*node)
		}
		g.getNode[n.id] = n
	}
}

func (g *graph) addEdge(n1, n2 string) {
	a, b := g.getNode[n1], g.getNode[n2]
	if g.edges == nil {
		g.edges = make(map[*node][]*node)
	}
	g.edges[a] = append(g.edges[a], b)
	g.edges[b] = append(g.edges[b], a)
}

func (g *graph) distinctPaths(current *node, count int, secondVisit bool) int {
	if current.id == "end" {
		return count + 1
	}
	for _, next := range g.edges[current] {
		if next.id == "start" {
			continue
		}
		if next.isBig || next.visited == 0 || !secondVisit {
			current.visited++
			secondVisitUsed := !next.isBig && next.visited == 1
			count = g.distinctPaths(next, count, secondVisit || secondVisitUsed)
			current.visited--
		}
	}
	return count
}

func Puzzle(input *[]byte, part2 bool) int {
	var g graph
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		connection := strings.Split(line, "-")
		for _, id := range connection {
			g.addNode(&node{id, unicode.IsUpper(rune(id[0])), 0})
		}
		g.addEdge(connection[0], connection[1])
	}
	return g.distinctPaths(g.getNode["start"], 0, !part2)
}

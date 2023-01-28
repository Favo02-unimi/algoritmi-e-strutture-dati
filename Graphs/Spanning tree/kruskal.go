package main

import (
	"fmt"
	"sort"
)

// graph as weightMatrix
type graph struct {
	map_ map[string]map[string]int
}

// weighted edge
type wedge struct {
	v1, v2 string
	weight int
}

func kruskal(g graph) graph {
	// save edges of graph in a set (no duplicates)
	edgesSet := make(map[wedge]bool)
	for v1, m := range g.map_ {
		for v2, w := range m {
			e1 := wedge{v1, v2, w}
			e2 := wedge{v2, v1, w}
			if !(edgesSet[e1] || edgesSet[e2]) {
				edgesSet[e1] = true
			}
		}
	}
	// convert set to list
	edgesList := make([]wedge, 0, len(g.map_))
	for k := range edgesSet {
		edgesList = append(edgesList, k)
	}

	// sort list
	sort.Slice(edgesList, func(i, j int) bool {
		return edgesList[i].weight < edgesList[j].weight
	})

	spanning := newGraph()
	// scan each edge in list (starting from smallest weight)
	for _, e := range edgesList {
		// spanning tree has all nodes
		if len(spanning.map_) == len(g.map_) {
			break
		}

		// add edge to spanning
		spanning.addEdge(e.v1, e.v2, e.weight)
		// if form cycle, remove it
		if hasCycle(spanning, e.v1, e.v1, make(map[string]bool)) {
			spanning.removeEdge(e.v1, e.v2)
		}
	}

	return spanning
}

// GRAPH
func newGraph() graph {
	var g graph = graph{make(map[string]map[string]int)}
	return g
}

func (g graph) addEdge(v1, v2 string, weight int) {
	if g.map_[v1] == nil {
		g.map_[v1] = make(map[string]int)
	}
	if g.map_[v2] == nil {
		g.map_[v2] = make(map[string]int)
	}

	g.map_[v1][v2] = weight
	g.map_[v2][v1] = weight
}

func (g graph) removeEdge(v1, v2 string) {
	delete(g.map_[v1], v2)
	delete(g.map_[v2], v1)
}

func hasCycle(g graph, start, last string, reached map[string]bool) bool {
	reached[start] = true
	for v := range g.map_[start] {
		if !reached[v] {
			if hasCycle(g, v, start, reached) {
				return true
			}
		} else {
			if v != last {
				return true
			}
		}
	}
	return false
}

func main() {
	g := newGraph()
	g.addEdge("a", "b", 3)
	g.addEdge("a", "c", 2)
	g.addEdge("b", "c", 4)
	g.addEdge("b", "f", 5)
	g.addEdge("c", "f", 5)
	g.addEdge("b", "d", 6)
	g.addEdge("b", "e", 6)
	g.addEdge("f", "e", 1)

	fmt.Println(kruskal(g))
}

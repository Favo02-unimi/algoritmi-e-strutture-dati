package main

import (
	"fmt"
	"sort"
)

// kruskal with sets implemented as maps: theta(V) + theta(E) + O(E log E) + O(E)*O(1)*O(V) = O(E log E)
func kruskal(g graph, start string) graph {

	// each point to parent set (or nil if own set)
	sets := make(map[string]*string)

	// each point in his own set (parent set = nothing)
	for v := range g.map_ {
		sets[v] = nil
	}

	// list of edges
	edges := make([]wedge, 0, len(g.map_)) // intialize the underlying array to at least the number of nodes of the graph
	for a, m := range g.map_ {
		for b, weight := range m {
			edges = append(edges, wedge{a, b, weight})
		}
	}

	// sort edges
	sort.Slice(edges, func(a, b int) bool {
		return edges[a].weight < edges[b].weight
	})

	// initialize spanning tree
	spanning := newGraph()

	// scan edges, starting from shortest one
	for _, edge := range edges {
		// spanning tree formed, stop
		if len(spanning.map_) == len(g.map_) {
			break
		}

		// check if edge forms cycle checking sets
		setStart := findSet(sets, edge.v1)
		setEnd := findSet(sets, edge.v2)
		// no cycle, add edge
		if setStart != setEnd {
			spanning.addUndirectedEdge(edge.v1, edge.v2, edge.weight)
			// union of sets
			sets[setEnd] = &setStart
		}
	}

	return spanning
}

func findSet(sets map[string]*string, key string) string {
	if sets[key] == nil {
		return key
	}
	return findSet(sets, *sets[key])
}

func main() {
	ug := newGraph()
	ug.addUndirectedEdge("a", "b", 3)
	ug.addUndirectedEdge("a", "c", 2)
	ug.addUndirectedEdge("b", "c", 4)
	ug.addUndirectedEdge("c", "e", 5)
	ug.addUndirectedEdge("d", "b", 6)
	ug.addUndirectedEdge("f", "d", 6)
	ug.addUndirectedEdge("e", "f", 1)

	distug := kruskal(ug, "f")
	fmt.Println(distug)
}

// GRAPH
type wedge struct {
	v1, v2 string
	weight int
}

type graph struct {
	map_ map[string]map[string]int
}

func newGraph() graph {
	var g graph = graph{make(map[string]map[string]int)}
	return g
}

func (g graph) addUndirectedEdge(v1, v2 string, weight int) {
	if g.map_[v1] == nil {
		g.map_[v1] = make(map[string]int)
	}
	if g.map_[v2] == nil {
		g.map_[v2] = make(map[string]int)
	}

	g.map_[v1][v2] = weight
	g.map_[v2][v1] = weight
}

func (g graph) addDirectedEdge(v1, v2 string, weight int) {
	if g.map_[v1] == nil {
		g.map_[v1] = make(map[string]int)
	}

	g.map_[v1][v2] = weight
}

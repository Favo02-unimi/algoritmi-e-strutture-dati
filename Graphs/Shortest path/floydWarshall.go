package main

import (
	"fmt"
)

const INF = 100000

// works with negative weight, not with negative cycle
func floydWarshall(g graph) map[string]map[string]int {
	dist := make(map[string]map[string]int)

	// check if directly connected
	for i := range g.map_ {
		for j := range g.map_ {
			weight, reachable := g.map_[i][j]
			if dist[i] == nil {
				dist[i] = make(map[string]int)
			}
			if i == j {
				dist[i][j] = 0
			} else if reachable {
				dist[i][j] = weight
			} else {
				dist[i][j] = INF
			}
		}
	}

	// try every possible path
	for k := range g.map_ { // mid point
		for i := range g.map_ { // start
			for j := range g.map_ { // end
				// if this path shorter than path saved
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j] // save this path
				}
			}
		}
	}

	return dist
}

func main() {
	dg := newGraph()
	dg.addDirectedEdge("a", "b", 3)
	dg.addDirectedEdge("a", "c", 2)
	dg.addDirectedEdge("b", "c", 4)
	dg.addDirectedEdge("c", "e", 5)
	dg.addDirectedEdge("d", "b", 6)
	dg.addDirectedEdge("f", "d", 6)
	dg.addDirectedEdge("e", "f", 1)

	distdg := floydWarshall(dg)
	fmt.Println(distdg)

	ug := newGraph()
	ug.addUndirectedEdge("a", "b", 3)
	ug.addUndirectedEdge("a", "c", 2)
	ug.addUndirectedEdge("b", "c", 4)
	ug.addUndirectedEdge("c", "e", 5)
	ug.addUndirectedEdge("d", "b", 6)
	ug.addUndirectedEdge("f", "d", 6)
	ug.addUndirectedEdge("e", "f", 1)

	distug := floydWarshall(ug)
	fmt.Println(distug)
}

// GRAPH
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

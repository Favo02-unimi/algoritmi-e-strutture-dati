package main

import (
	"fmt"
)

type graph struct {
	map_ map[string]map[string]int
}

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
				dist[i][j] = 100000 // inf
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
				fmt.Println()
			}
		}
	}

	return dist
}

func main() {
	g := newGraph()
	g.addEdge("a", "b", 3)
	g.addEdge("a", "c", 2)
	g.addEdge("b", "c", 4)
	g.addEdge("b", "f", 5)
	g.addEdge("c", "f", 5)
	g.addEdge("b", "d", 6)
	g.addEdge("d", "e", 6)
	g.addEdge("f", "e", 1)

	dist := floydWarshall(g)
	for k, v := range dist {
		fmt.Println(k, ":", v)
	}
}

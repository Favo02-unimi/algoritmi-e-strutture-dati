package main

import (
	"fmt"
	"math"
)

// graph as weightMatrix
type graph struct {
	map_ map[string]map[string]int
}

// edge with no weight
type edge struct {
	v1, v2 string
}

func prim(g graph, start string) graph {
	spanning := newGraph() // spanning tree

	queue := make(map[edge]int) // priority queue (should use minHeap)

	// while spanning tree not full
	for len(spanning.map_) < len(g.map_) {

		// edges out from start
		fromStart := g.map_[start]
		// add edges out from start to queue
		for k, v := range fromStart {
			queue[edge{start, k}] = v
		}

		// extract minweight not already in spanning from queue
		var minW int = math.MaxInt
		var minE edge
		for k, v := range queue {
			_, inSpanning := spanning.map_[k.v2] // check in spanning
			if !inSpanning && v < minW {
				minE = k
				minW = v
			}
		}

		// add to spanning
		spanning.addEdge(minE.v1, minE.v2, minW)
		// remove from queue
		delete(queue, minE)

		// set as start node reached by new edge
		// (add to queue edges out from this new node)
		start = minE.v2
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

	fmt.Println(prim(g, "a"))
}

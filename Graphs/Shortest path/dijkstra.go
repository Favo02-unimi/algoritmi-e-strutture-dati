package main

import (
	"fmt"
)

const INF = 100000

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

// does not works with negative weights
func dijkstra(g graph, start string) map[string]int {

	dist := make(map[string]int)
	notLocked := make(map[string]bool) // should be a minheap

	// initialize each point to inf
	for k := range g.map_ {
		dist[k] = INF
		notLocked[k] = true
	}

	// starting point to 0
	dist[start] = 0

	// while notLocked not empty
	for len(notLocked) > 0 {

		// cur = min distance in notLocked
		// should be a minheap, so extract the min from the heap
		var min int = INF
		var cur string
		for k := range notLocked {
			if dist[k] < min {
				min = dist[k]
				cur = k
			}
		}

		// remove minPoint from notLocked
		delete(notLocked, cur)

		// scan each point reachable from current point (cur)
		for v, weight := range g.map_[cur] {
			// if this path to reach v is less than old path, save this one
			if (dist[cur] + weight) < dist[v] {
				dist[v] = dist[cur] + weight
			}
		}

		// ignore unreachable nodes
		unreachable := true
		for k := range notLocked {
			if dist[k] != INF {
				unreachable = false
			}
		}
		if unreachable {
			break
		}

	}

	return dist
}

func main() {
	g := newGraph()
	g.addDirectedEdge("a", "b", 3)
	g.addDirectedEdge("a", "c", 2)
	g.addDirectedEdge("b", "c", 4)
	g.addDirectedEdge("c", "e", 5)
	g.addDirectedEdge("d", "b", 6)
	g.addDirectedEdge("f", "d", 6)
	g.addDirectedEdge("e", "f", 1)

	dist := dijkstra(g, "f")
	fmt.Println(dist)
}

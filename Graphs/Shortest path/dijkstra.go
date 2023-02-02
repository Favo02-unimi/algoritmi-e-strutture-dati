package main

import (
	"fmt"
)

const INF = 100000

// does not works with negative weights
func dijkstra(g graph, start string) map[string]int {

	dist := make(map[string]int)
	queue := make([]string, 0, len(g.map_)) // should be a minheap

	// initialize each distance to inf, add all to queue
	for k := range g.map_ {
		dist[k] = INF
		queue = append(queue, k)
	}

	// starting point to 0
	dist[start] = 0

	// while notLocked not empty
	for len(queue) > 0 {

		// cur = element extracted from queue
		// should be a minheap, so extract the min from the heap
		var cur string = queue[0]
		var curI = 0
		for i, k := range queue {
			if dist[k] < dist[cur] {
				cur = k
				curI = i
			}
		}

		// remove element extracted from queue
		// should be a minheap, so extract will remove
		queue = append(queue[:curI], queue[curI+1:]...)

		// scan each point reachable from current point (cur)
		for v, weight := range g.map_[cur] {
			// if this path to reach v is less than old path, save this one
			if (dist[cur] + weight) < dist[v] {
				dist[v] = dist[cur] + weight
			}
		}

		// ignore unreachable nodes
		unreachable := true
		for _, k := range queue {
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
	dg := newGraph()
	dg.addDirectedEdge("a", "b", 3)
	dg.addDirectedEdge("a", "c", 2)
	dg.addDirectedEdge("b", "c", 4)
	dg.addDirectedEdge("c", "e", 5)
	dg.addDirectedEdge("d", "b", 6)
	dg.addDirectedEdge("f", "d", 6)
	dg.addDirectedEdge("e", "f", 1)

	distdg := dijkstra(dg, "f")
	fmt.Println(distdg)

	ug := newGraph()
	ug.addUndirectedEdge("a", "b", 3)
	ug.addUndirectedEdge("a", "c", 2)
	ug.addUndirectedEdge("b", "c", 4)
	ug.addUndirectedEdge("c", "e", 5)
	ug.addUndirectedEdge("d", "b", 6)
	ug.addUndirectedEdge("f", "d", 6)
	ug.addUndirectedEdge("e", "f", 1)

	distug := dijkstra(ug, "f")
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

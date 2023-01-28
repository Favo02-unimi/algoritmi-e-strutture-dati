package main

import "fmt"

const INF = 100000

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

// works with negative weight, not with negative cycle
func bellmanFord(g graph, start string) map[string]int {
	dist := make(map[string]int)

	// initialize all distances but start to inf
	for k := range g.map_ {
		dist[k] = INF
	}
	dist[start] = 0

	// scan n-1 times every edge of the graph
	for i := 1; i < len(g.map_); i++ {
		// scan each edge of the graph
		for cur, v := range g.map_ {
			for dest, wei := range v {
				// if dist to reach current + weight between current and dest is < than dest distance then swap
				if (dist[cur] + wei) < dist[dest] {
					dist[dest] = dist[cur] + wei
				}
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

	dist := bellmanFord(g, "a")
	fmt.Println(dist)
}

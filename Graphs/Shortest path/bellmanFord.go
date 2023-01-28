package main

import "fmt"

const INF = 100000

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
	dg := newGraph()
	dg.addDirectedEdge("a", "b", 3)
	dg.addDirectedEdge("a", "c", 2)
	dg.addDirectedEdge("b", "c", 4)
	dg.addDirectedEdge("c", "e", 5)
	dg.addDirectedEdge("d", "b", 6)
	dg.addDirectedEdge("f", "d", 6)
	dg.addDirectedEdge("e", "f", 1)

	distdg := bellmanFord(dg, "f")
	fmt.Println(distdg)

	ug := newGraph()
	ug.addUndirectedEdge("a", "b", 3)
	ug.addUndirectedEdge("a", "c", 2)
	ug.addUndirectedEdge("b", "c", 4)
	ug.addUndirectedEdge("c", "e", 5)
	ug.addUndirectedEdge("d", "b", 6)
	ug.addUndirectedEdge("f", "d", 6)
	ug.addUndirectedEdge("e", "f", 1)

	distug := bellmanFord(ug, "f")
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

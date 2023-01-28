package main

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

func (g graph) removeEdge(v1, v2 string) {
	delete(g.map_[v1], v2)
	delete(g.map_[v2], v1)
}

func depthFirstSearch(g graph, start string, reached map[string]bool) {
	reached[start] = true
	for v := range g.map_[start] {
		if !reached[v] {
			depthFirstSearch(g, v, reached)
		}
	}
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

	depthFirstSearch(g, "a", make(map[string]bool))
}

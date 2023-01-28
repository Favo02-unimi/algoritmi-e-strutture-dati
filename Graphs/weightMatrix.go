package main

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

func (g graph) removeUndirectedEdge(v1, v2 string) {
	delete(g.map_[v1], v2)
	delete(g.map_[v2], v1)
}

func (g graph) removeDirectedEdge(v1, v2 string) {
	delete(g.map_[v1], v2)
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
	dg := newGraph()
	dg.addDirectedEdge("a", "b", 3)
	dg.addDirectedEdge("a", "c", 2)
	dg.addDirectedEdge("b", "c", 4)
	dg.addDirectedEdge("c", "e", 5)
	dg.addDirectedEdge("d", "b", 6)
	dg.addDirectedEdge("f", "d", 6)
	dg.addDirectedEdge("e", "f", 1)

	depthFirstSearch(dg, "a", make(map[string]bool))

	ug := newGraph()
	ug.addUndirectedEdge("a", "b", 3)
	ug.addUndirectedEdge("a", "c", 2)
	ug.addUndirectedEdge("b", "c", 4)
	ug.addUndirectedEdge("c", "e", 5)
	ug.addUndirectedEdge("d", "b", 6)
	ug.addUndirectedEdge("f", "d", 6)
	ug.addUndirectedEdge("e", "f", 1)

	depthFirstSearch(ug, "a", make(map[string]bool))
}

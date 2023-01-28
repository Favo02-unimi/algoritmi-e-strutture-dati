package main

import (
	"fmt"
)

type graph struct {
	map_ map[string][]string
}

func newGraph() graph {
	return graph{make(map[string][]string)}
}

func main() {
	graph := newGraph()

	graph.map_["a"] = []string{"b", "c"}
	graph.map_["b"] = []string{"a", "c", "f", "d"}
	graph.map_["c"] = []string{"a", "b", "f"}
	graph.map_["d"] = []string{"b", "e"}
	graph.map_["e"] = []string{"d", "f"}
	graph.map_["f"] = []string{"b", "c", "e"}

	fmt.Println("BFS:")
	breadthFirstSearch(graph, "a")
	fmt.Println("DFS:")
	depthFirstSearch(graph, "a", make(map[string]bool))
}

func breadthFirstSearch(g graph, start string) {
	var q []string
	reached := make(map[string]bool)
	fmt.Println(start)
	reached[start] = true
	q = append(q, start)
	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g.map_[u] {
			if !reached[v] {
				fmt.Println(v)
				reached[v] = true
				q = append(q, v)
			}
		}
	}
}

func depthFirstSearch(g graph, start string, reached map[string]bool) {
	fmt.Println(start)
	reached[start] = true
	for _, v := range g.map_[start] {
		if !reached[v] {
			depthFirstSearch(g, v, reached)
		}
	}
}

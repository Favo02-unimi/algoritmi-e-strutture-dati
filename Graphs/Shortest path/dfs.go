package main

import "fmt"

type graph struct {
	map_ map[string][]string
}

func newGraph() graph {
	return graph{make(map[string][]string)}
}

const WEIGHT = 1

func dfs(g graph, start string, distances map[string]int) map[string]int {
	// initialize start
	if _, startReached := distances[start]; !startReached {
		distances[start] = 0
	}

	// scan each node reachable from start
	for _, v := range g.map_[start] {

		_, reached := distances[v]
		// if node not already reached
		if !reached || distances[v] > distances[start]+WEIGHT {
			// save distance: distance to reach start + weight
			distances[v] = distances[start] + WEIGHT
			// explore from reached node
			distances = dfs(g, v, distances)
		}
	}

	return distances
}

func main() {
	graph := newGraph()

	graph.map_["a"] = []string{"b", "c"}
	graph.map_["b"] = []string{"a", "c", "f", "d"}
	graph.map_["c"] = []string{"a", "b", "f"}
	graph.map_["d"] = []string{"b", "e"}
	graph.map_["e"] = []string{"d", "f"}
	graph.map_["f"] = []string{"b", "c", "e"}

	fmt.Println(dfs(graph, "a", make(map[string]int)))
}

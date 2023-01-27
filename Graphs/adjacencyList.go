package main

import (
	"fmt"
)

func main() {
	var graph map[string][]string = make(map[string][]string)

	graph["a"] = []string{"b", "c"}
	graph["b"] = []string{"a", "c", "f", "d"}
	graph["c"] = []string{"a", "b", "f"}
	graph["d"] = []string{"b", "e"}
	graph["e"] = []string{"d", "f"}
	graph["f"] = []string{"b", "c", "e"}

	fmt.Println("BFS:")
	breadthFirstSearch(graph, "a")
	fmt.Println("DFS:")
	depthFirstSearch(graph, "a", make(map[string]bool))
}

func breadthFirstSearch(graph map[string][]string, start string) {
	var q []string
	reached := make(map[string]bool)
	fmt.Println(start)
	reached[start] = true
	q = append(q, start)
	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		for _, v := range graph[u] {
			if !reached[v] {
				fmt.Println(v)
				reached[v] = true
				q = append(q, v)
			}
		}
	}
}

func depthFirstSearch(graph map[string][]string, start string, reached map[string]bool) {
	fmt.Println(start)
	reached[start] = true
	for _, v := range graph[start] {
		if !reached[v] {
			depthFirstSearch(graph, v, reached)
		}
	}
}

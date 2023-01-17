package main

import "math"

type Point struct {
	x, y int
}

func dijkstra(graph map[Point][]Point, start Point) map[Point]int {

	// create map of size number of nodes
	distances := make(map[Point]int)
	c := make(map[Point]int)

	// initialize each point to inf
	for k := range graph {
		distances[k] = math.MaxInt
		c[k] = math.MaxInt
	}

	// starting point to 0
	distances[start] = 0
	c[start] = 0

	// while c not empty
	for len(c) > 0 {

		// cur = min v
		var min int = math.MaxInt
		var cur Point
		for k := range c {
			if distances[k] < min {
				min = distances[k]
				cur = k
			}
		}

		// remove minPoint from c
		delete(c, cur)

		// scan each point reachable from current point (cur)
		for _, v := range graph[cur] {
			weight := 1 // weight of move from cur to v

			// if this path to reach v is less than old path, save this one
			if (distances[cur] + weight) < distances[v] {
				distances[v] = distances[cur] + weight
			}
		}

		// ignore unreachable nodes
		unreachable := true
		for cLeft := range c {
			if distances[cLeft] != math.MaxInt {
				unreachable = false
			}
		}
		if unreachable {
			break
		}

	}
	return distances
}

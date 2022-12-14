package main

import "math"

type Point struct {
	x, y int
}

func dijkstra(grid [][]int, cur Point) map[Point]int {

	// create map of size number of nodes
	distances := make(map[Point]int)
	queue := make(map[Point]int)

	// initialize each point to inf
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			distances[Point{x, y}] = math.MaxInt
			queue[Point{x, y}] = math.MaxInt

			// starting point to 0
			if x == cur.x && y == cur.y {
				distances[Point{x, y}] = 0
				queue[Point{x, y}] = 0
			}
		}
	}

	// while c not empty
	for len(queue) > 0 {

		// cur = min v
		var min int = math.MaxInt
		var cur Point
		for k := range queue {
			if distances[k] < min {
				min = distances[k]
				cur = k
			}
		}

		// remove minPoint from c
		delete(queue, cur)

		// scan each point reachable from current point (cur)
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				// do not analyse self
				if x == 0 && y == 0 {
					continue
				}

				// point reached from cur
				v := Point{cur.x + x, cur.y + y}

				weight := 1 // weight of move from cur to v

				// if this path to reach v is less than old path, save this one
				if (distances[cur] + weight) < distances[v] {
					distances[v] = distances[cur] + weight
				}
			}
		}

		// ignore unreachable nodes
		unreachable := true
		for cLeft := range queue {
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

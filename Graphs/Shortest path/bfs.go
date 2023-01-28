package main

import "fmt"

type graph struct {
	map_ map[string][]string
}

func newGraph() graph {
	return graph{make(map[string][]string)}
}

const WEIGHT = 1

func bfs(g graph, start string) map[string]int {
	queue := queue{nil}
	distances := make(map[string]int)
	distances[start] = 0

	queue.enqueue(start)

	for !queue.isEmpty() {
		u := queue.dequeue()

		for _, v := range g.map_[u] {
			if _, reached := distances[v]; !reached {
				distances[v] = distances[u] + 1
				queue.enqueue(v)
			}
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

	fmt.Println(bfs(graph, "a"))
}

// QUEUE

type queue struct {
	head *queueNode
}

type queueNode struct {
	next    *queueNode
	payload string
}

func (q *queue) enqueue(p string) {
	if q.head == nil {
		q.head = &queueNode{nil, p}
		return
	}
	node := q.head
	for node.next != nil {
		node = node.next
	}
	newNode := queueNode{nil, p}
	node.next = &newNode
}

func (q *queue) dequeue() string {
	head := q.head
	q.head = q.head.next
	return head.payload
}

func (q *queue) isEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}

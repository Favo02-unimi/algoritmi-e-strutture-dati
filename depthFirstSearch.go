package main

func depthFirstSearch(map_ map[int][]int, cur int) (map[int]int, map[int][]int) {
	queue := queue{nil}
	distances := make(map[int]int)
	distances[cur] = 0
	reached := make(map[int]bool)
	reached[cur] = true
	visitedToReach := make(map[int][]int)

	queue.enqueue(cur)

	for !queue.isEmpty() {
		u := queue.dequeue()

		for _, v := range map_[u] {
			if !reached[v] {
				distances[v] = distances[u] + 1
				visitedToReach[v] = visitedToReach[u]
				visitedToReach[v] = append(visitedToReach[v], u)
				reached[v] = true
				queue.enqueue(v)
			}
		}

	}
	return distances, visitedToReach
}

// QUEUE

type queue struct {
	head *queueNode
}

type queueNode struct {
	next    *queueNode
	payload int
}

func (q *queue) enqueue(p int) {
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

func (q *queue) dequeue() int {
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


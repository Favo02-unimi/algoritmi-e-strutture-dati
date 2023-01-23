package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	table := parseInput()

	distances, nodesVisited := depthFirstSearch(table, 1)
	for i := 2; i <= len(table)+1; i++ {
		fmt.Println("distanza (passi, non mosse) da", i, ":", distances[i])
		moves := calulateMoves(i, nodesVisited)
		fmt.Println("mosse per arrivare a", i, ":", moves)
		fmt.Println()
	}
}

func parseInput() map[int][]int {
	table := make(map[int][]int)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		var rows, columns int
		if i == 0 {
			rc := scanner.Text()
			tokens := strings.Split(rc, " ")
			rows, _ = strconv.Atoi(tokens[0])
			columns, _ = strconv.Atoi(tokens[1])

			for i := 1; i < rows*columns; i++ {
				table[i] = append(table[i], i+1)
			}
		} else {
			sd := scanner.Text()
			tokens := strings.Split(sd, " ")
			source, _ := strconv.Atoi(tokens[0])
			destination, _ := strconv.Atoi(tokens[1])
			table[source-1] = append(table[source-1], destination)
		}

	}
	return table
}

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

func calulateMoves(dest int, nodesVisited map[int][]int) []int {
	moves := nodesVisited[dest]
	var consecutive []int
	var consCount int = 1
	for i := 0; i < len(moves)-1; i++ {
		if moves[i] != moves[i+1]-1 {
			consecutive = append(consecutive, consCount)
			consCount = 0
		}
		consCount++
	}
	consecutive = append(consecutive, consCount)

	moves = make([]int, 0)
	for _, c := range consecutive {
		if c > 6 {
			for c > 6 {
				c = c - 6
				moves = append(moves, 6)
			}
			moves = append(moves, c)
		} else {
			moves = append(moves, c)
		}
	}
	return moves
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

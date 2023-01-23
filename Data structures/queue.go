package main

import "fmt"

type queue struct {
	head *node
	last *node
}

type node struct {
	next  *node
	value int
}

func (q *queue) enqueue(value int) {
	newNode := &node{nil, value}
	if q.head == nil && q.last == nil {
		q.head = newNode
		q.last = newNode
		return
	}
	q.last.next = newNode
	q.last = newNode
}

func (q *queue) dequeue() *node {
	if q.head == nil && q.last == nil {
		return nil
	}
	node := q.head
	if q.head == q.last {
		q.head = nil
		q.last = nil
	}
	q.head = node.next
	return node
}

func (q queue) first() *node {
	return q.head
}

func (q queue) isEmpty() bool {
	return q.head == nil
}

func main() {
	var queue queue
	fmt.Println(queue.isEmpty())
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(4)
	fmt.Println(queue.first())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.first())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.first())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.isEmpty())
	fmt.Println(queue.first())
	fmt.Println(queue.dequeue())
	queue.enqueue(9)
	fmt.Println(queue.dequeue())
	fmt.Println(queue.first())
	fmt.Println(queue.isEmpty())
}

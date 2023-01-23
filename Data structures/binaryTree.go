package main

import "fmt"

type binaryTree struct {
	root *node
}

type node struct {
	value int
	left  *node
	right *node
}

func newNode(value int) *node {
	return &node{value, nil, nil}
}

func (tree binaryTree) breadthFirstSearch() {
	var queue queue
	queue.enqueue(tree.root)
	for !queue.isEmpty() {
		x := queue.dequeue().value
		if x != nil {
			fmt.Print(x.value, " ")
			queue.enqueue(x.left)
			queue.enqueue(x.right)
		}
	}
	fmt.Println()
}

func (tree binaryTree) depthFirstSearch() {
	var stack stack
	stack.push(tree.root)
	for !stack.isEmpty() {
		x := stack.pop().value
		if x != nil {
			fmt.Print(x.value, " ")
			stack.push(x.right)
			stack.push(x.left)
		}
	}
	fmt.Println()
}

func recursiveDFSpreorder(root *node) {
	if root != nil {
		fmt.Print(root.value, " ")
		recursiveDFSpreorder(root.left)
		recursiveDFSpreorder(root.right)
	}
}

func recursiveDFSinorder(root *node) {
	if root != nil {
		recursiveDFSinorder(root.left)
		fmt.Print(root.value, " ")
		recursiveDFSinorder(root.right)
	}
}

func recursiveDFSpostorder(root *node) {
	if root != nil {
		recursiveDFSpostorder(root.left)
		recursiveDFSpostorder(root.right)
		fmt.Print(root.value, " ")
	}
}

func main() {
	tree := binaryTree{}

	// tree.root = newNode(1)
	// tree.root.left = newNode(2)
	// tree.root.right = newNode(3)
	// tree.root.left.left = newNode(4)
	// tree.root.left.right = newNode(5)
	// tree.root.right.left = newNode(6)
	// tree.root.right.right = newNode(7)
	// tree.root.left.left.left = newNode(8)
	// tree.root.left.right.right = newNode(9)
	// tree.root.left.right.right.right = newNode(10)

	tree.root = newNode(8)
	tree.root.left = newNode(14)
	tree.root.left.left = newNode(22)
	tree.root.left.right = newNode(12)
	tree.root.left.right.left = newNode(1)
	tree.root.left.right.right = newNode(3)
	tree.root.right = newNode(1)
	tree.root.right.right = newNode(5)

	tree.breadthFirstSearch()
	tree.depthFirstSearch()
	fmt.Println()
	recursiveDFSpreorder(tree.root)
	fmt.Println()
	recursiveDFSinorder(tree.root)
	fmt.Println()
	recursiveDFSpostorder(tree.root)
	fmt.Println()
}

// QUEUE
type queue struct {
	head *queueNode
	last *queueNode
}

type queueNode struct {
	next  *queueNode
	value *node
}

func (q *queue) enqueue(node *node) {
	newNode := &queueNode{nil, node}
	if q.head == nil && q.last == nil {
		q.head = newNode
		q.last = newNode
		return
	}
	q.last.next = newNode
	q.last = newNode
}

func (q *queue) dequeue() *queueNode {
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

func (q queue) first() *queueNode {
	return q.head
}

func (q queue) isEmpty() bool {
	return q.head == nil
}

// STACK
type stack struct {
	head *stackNode
}

type stackNode struct {
	next  *stackNode
	value *node
}

func (s *stack) push(value *node) {
	newstackNode := &stackNode{s.head, value}
	s.head = newstackNode
}

func (s *stack) pop() *stackNode {
	if s.head == nil {
		return nil
	}
	stackNode := s.head
	s.head = stackNode.next
	return stackNode
}

func (s stack) peek() *stackNode {
	return s.head
}

func (s stack) isEmpty() bool {
	return s.head == nil
}

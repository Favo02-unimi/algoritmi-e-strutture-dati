package main

import "fmt"

type stack struct {
	head *node
}

type node struct {
	next  *node
	value int
}

func (s *stack) push(value int) {
	newNode := &node{s.head, value}
	s.head = newNode
}

func (s *stack) pop() *node {
	if s.head == nil {
		return nil
	}
	node := s.head
	s.head = node.next
	return node
}

func (s stack) peek() *node {
	return s.head
}

func (s stack) isEmpty() bool {
	return s.head == nil
}

func main() {
	var stack stack
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	stack.push(9)
	fmt.Println(stack.peek())
}

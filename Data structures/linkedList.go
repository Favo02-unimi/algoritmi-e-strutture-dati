package main

import "fmt"

type linkedList struct {
	head *node
	size int
}

type node struct {
	next  *node
	value int
}

func (list linkedList) getSize() int {
	return list.size
}

func (list *linkedList) add(value int) {
	newNode := &node{nil, value}
	if list.head == nil {
		list.head = newNode
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newNode
	}
	list.size++
}

func (list *linkedList) addAtIndex(value, index int) {
	newNode := &node{nil, value}
	if list.head == nil {
		list.head = newNode
	} else if index == 0 {
		newNode.next = list.head
		list.head = newNode
	} else {
		cur := list.head
		for cur.next != nil && index > 1 {
			cur = cur.next
			index--
		}
		newNode.next = cur.next
		cur.next = newNode
	}
	list.size++
}

func (list *linkedList) removeByIndex(index int) {
	if list.head == nil {
		return
	} else if index == 0 {
		list.head = list.head.next
		list.size--
	} else {
		cur := list.head
		for cur.next != nil && index > 1 {
			cur = cur.next
			index--
		}
		cur.next = cur.next.next
		list.size--
	}
}

func (list *linkedList) removeByValue(value int) {
	if list.head == nil {
		return
	} else if list.head.value == value {
		list.head = list.head.next
		list.size--
	} else {
		cur := list.head
		for cur.next != nil && cur.next.value != value {
			cur = cur.next
		}
		cur.next = cur.next.next
		list.size--
	}
}

func (list linkedList) getByIndex(index int) *node {
	cur := list.head
	for cur != nil && index > 0 {
		cur = cur.next
		index--
	}
	return cur
}

func (list linkedList) getByValue(value int) *node {
	cur := list.head
	for cur.next != nil {
		if cur.value == value {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func (list linkedList) printList() {
	if list.size == 0 {
		fmt.Println("[ ]")
		return
	}

	fmt.Print("[ ", list.head.value)

	cur := list.head.next
	for cur != nil {
		fmt.Print(", ", cur.value)
		cur = cur.next
	}

	fmt.Println(" ]")
}

func main() {
	list := linkedList{nil, 0}
	list.add(4)
	list.add(9)
	list.add(5)
	list.add(5)
	list.add(5)
	list.add(5)
	list.add(5)
	list.add(5)
	list.addAtIndex(7, 5)
	list.addAtIndex(0, 1)
	list.printList()
	fmt.Println(list.size)
	list.removeByIndex(6)
	list.printList()
	fmt.Println(list.size)
	list.removeByIndex(0)
	list.printList()
	fmt.Println(list.size)

	list.removeByValue(0)
	list.printList()
	fmt.Println(list.size)
	fmt.Println(list.getByIndex(1))
	fmt.Println(list.getByValue(9))
}

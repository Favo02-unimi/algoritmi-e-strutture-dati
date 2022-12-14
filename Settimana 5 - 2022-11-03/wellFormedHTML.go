package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Insert html document (one line): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	doc := scanner.Text()
	wellFormed, errPos, remainingTags := checkTags(doc)
	if wellFormed {
		fmt.Println("The document is well formed")
	} else {
		fmt.Println("The document is NOT well formed:")
		if len(remainingTags) > 0 {
			fmt.Println("\tTags not closed:", remainingTags)
		} else {
			fmt.Println("\tError position:", errPos)
		}
	}
}

func checkTags(doc string) (bool, int, []string) {
	var stack stack = stack{nil}
	tags := strings.Split(doc, " ")
	for i, t := range tags {
		if (!strings.Contains(t, "/")) {
			stack.push(t)
		} else if strings.Contains(t, "/") {
			openingTag := stack.pop()
			if (openingTag[1:] != t[2:]) {
				return false, i+1, nil
			}
		}
	}
	if (stack.isEmpty()) {
		return true, -1, nil
	} else {
		var remainingTags = make([]string, 0)
		for (!stack.isEmpty()) {
			remainingTags = append(remainingTags, stack.pop())
		}
		for i, j := 0, len(remainingTags)-1; i < j; i, j = i+1, j-1 {
			remainingTags[i], remainingTags[j] = remainingTags[j], remainingTags[i]
		}
		return false, -1, remainingTags
	}
}


// IMPLEMENTAZIONE PILA TRAMITE LISTA //

type listNode struct {
	next *listNode
	item string
}

type stack struct {
	head *listNode
}

func newNode(item string) *listNode {
	return &listNode{nil, item}
}

func (list *stack) push(item string) {
	newNode := newNode(item)
	newNode.next = list.head
	list.head = newNode
}

func (list *stack) pop() string {
	node := list.head
	list.head = node.next
	return node.item
}

func (list stack) isEmpty() bool {
	if list.head == nil {
		return true
	}
	return false
}

func (list stack) print() {
	fmt.Print("[ ")
	var node *listNode = list.head
	for node != nil {
		fmt.Print(node.item, " ")
		node = node.next
	}
	fmt.Println("]")
}

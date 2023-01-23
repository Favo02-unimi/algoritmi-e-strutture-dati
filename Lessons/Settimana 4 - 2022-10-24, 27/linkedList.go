package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type listNode struct {
	next *listNode
	prec *listNode
	item int
}

type linkedList struct {
	head *listNode
	tail *listNode
}

func newNode(item int) *listNode {
	return &listNode{nil, nil, item}
}

func (list linkedList) print() {
	fmt.Print("[ ")
	var node *listNode = list.head
	for node != nil {
		fmt.Print(node.item, " ")
		node = node.next
	}
	fmt.Println("]")
}

func (list linkedList) printReverse() {
	fmt.Print("[ ")
	var node *listNode = list.tail
	for node != nil {
		fmt.Print(node.item, " ")
		node = node.prec
	}
	fmt.Println("]")
}

func (list linkedList) length() int {
	var node *listNode = list.head
	var count int = 0
	for node != nil {
		count++
		node = node.next
	}
	return count
}

func (list *linkedList) addNewNode(item int) {
	newNode := newNode(item)
	if list.tail == nil {
		list.tail = newNode
	} else {
		list.head.prec = newNode
	}
	newNode.next = list.head
	list.head = newNode
}

func (list linkedList) searchIndex(index int) (bool, *listNode) {
	var node *listNode = list.head
	for node != nil && index > 0 {
		node = node.next
		index--
	}
	if node == nil {
		return false, nil
	} else {
		return true, node
	}
}

func (list linkedList) searchItem(item int) (bool, *listNode) {
	var node *listNode = list.head
	for node != nil && item != node.item {
		node = node.next
	}
	if node == nil {
		return false, nil
	} else {
		return true, node
	}
}

func (list *linkedList) removeIndex(index int) bool {
	var node, prec *listNode = list.head, nil
	for node != nil && index > 0 {
		prec = node
		node = node.next
		index--
	}
	if prec == nil {
		if node.next == nil {
			list.tail = prec
		}
		list.head = node.next
		return true
	} else if node != nil {
		if node.next == nil {
			list.tail = prec
		}
		prec.next = node.next
		node.next = nil
		return true
	} else {
		return false
	}
}

func (list *linkedList) removeItem(item int) bool {
	var node, prec *listNode = list.head, nil
	for node != nil && item != node.item {
		prec = node
		node = node.next
	}
	if prec == nil {
		if node.next == nil {
			list.tail = prec
		}
		list.head = node.next
		return true
	} else if node != nil {
		if node.next == nil {
			list.tail = prec
		}
		prec.next = node.next
		node.next = nil
		return true
	} else {
		return false
	}
}

func main() {
	list := linkedList{nil, nil}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		executeOperation(&list, scanner.Text())
	}

}

func executeOperation(list *linkedList, operation string) {
	switch operation[0] {
	case '+':
		n, _ := strconv.Atoi(operation[1:])
		found, _ := list.searchItem(n)
		if !found {
			list.addNewNode(n)
		}
		list.print()
	case '-':
		n, _ := strconv.Atoi(operation[1:])
		found, _ := list.searchItem(n)
		if found {
			list.removeItem(n)
		}
		list.print()
	case '?':
		n, _ := strconv.Atoi(operation[1:])
		found, _ := list.searchItem(n)
		if found {
			fmt.Println(n, " appartiene all'insieme")
		} else {
			fmt.Println(n, " non appartiene all'insieme")
		}
	case 'c':
		fmt.Println("grandezza insieme:", list.length())
	case 'p':
		list.print()
	case 'o':
		list.printReverse()
	case 'd':
		list.head = nil
		list.tail = nil
		list.print()
	case 'f':
		os.Exit(0)
	default:
		fmt.Println("Invalid operation!" +
			"\nUSAGE:" +
			"\n\t+n aggiungere elemento n se non esiste" +
			"\n\t-n rimuovere elemento n se esiste" +
			"\n\t?n controllare esistenza di n" +
			"\n\tc grandezza insieme" +
			"\n\tp stampa insieme" +
			"\n\to stampa insieme in ordine inverso" +
			"\n\td svuota insieme" +
			"\n\tf termina esecuzione")
	}
}

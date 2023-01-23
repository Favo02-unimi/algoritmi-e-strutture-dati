package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: Command line argument:\n EVAL to evaluate a postfix expression\n CONV to convert a infix expression to postfix")
		return
	}

	fmt.Print("Insert expression: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	espr := scanner.Text()
	fmt.Println("Expression:", espr)
	if os.Args[1] == "EVAL" {
		fmt.Println("Result:", evaluate(espr))
	} else if os.Args[1] == "CONV" {
		fmt.Println("Result:", convert(espr))
	} else {
		fmt.Println("Wrong usage. Command line argument EVAL or CONV")
	}
}

func evaluate(espr string) int {
	var stack stack = stack{nil}
	tokens := strings.Split(espr, " ")
	for _, t := range tokens {
		if t[0] >= '0' && t[0] <= '9' {
			stack.push(t)
		} else {
			op := t
			v1 := stack.pop()
			v2 := stack.pop()
			res := fmt.Sprint(operation(op, strToInt(v1), strToInt(v2)))
			stack.push(res)
		}
	}
	return strToInt(stack.pop())
}

func convert(espr string) string {
	var stack2 stack = stack{nil}
	tokens := strings.Split(espr, " ")
	var res string
	for _, t := range tokens {
		if t[0] >= '0' && t[0] <= '9' {
			res += t + " "
		} else if strings.Contains("+-/*", t) {
			stack2.push(t)
		} else if t == "(" {
			continue
		} else if t == ")" {
			res += stack2.pop() + " "
		}
	}
	return res
}

// FUNZIONI DI APPOGGIO //

func operation(op string, v1, v2 int) int {
	switch op {
	case "+":
		return v2 + v1
	case "-":
		return v2 - v1
	case "*":
		return v2 * v1
	case "/":
		return v2 / v1
	default:
		fmt.Println("Operation not supported")
		return 0
	}
}

func strToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
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

func (list stack) print() {
	fmt.Print("[ ")
	var node *listNode = list.head
	for node != nil {
		fmt.Print(node.item, " ")
		node = node.next
	}
	fmt.Println("]")
}

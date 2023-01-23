package main

import "fmt"

func main() {
	//t := & bitree {nil}
	//t. root = & bitreeNode {nil, nil, 8}
	//t. root . left = newNode (14)
	//t. root . left.left = newNode (22)
	//t. root . left.right = newNode (12)
	//t. root . left.right.left = newNode (1)
	//t. root . left.right.right = newNode (3)
	//t. root . right = newNode (1)
	//t. root . right.right = newNode (5)

	/*
		t := &bitree{nil}
		t.root = &bitreeNode{nil, nil, 78}
		t.root.left = newNode(54)
		t.root.left.right = newNode(90)
		t.root.left.right.left = newNode(19)
		t.root.left.right.right = newNode(95)
		t.root.right = newNode(21)
		t.root.right.left = newNode(16)
		t.root.right.left.left = newNode(5)
		t.root.right.right = newNode(19)
		t.root.right.right.left = newNode(56)
		t.root.right.right.right = newNode(43)

		preorder(t.root) // Ordine di visita: 8, 14, 22, 12, 1, 3, 1, 5
		fmt.Println()
		inorder(t.root) // Ordine di visita: 22, 14, 1, 12, 3, 8, 1, 5
		fmt.Println()
		postorder(t.root) // Ordine di visita: 22, 1, 3, 12, 14, 5, 1, 8
		fmt.Println()
		stampaAlberoASommario(t.root, 0)
		fmt.Println()
		stampaAlbero(t.root)
		fmt.Println()
		fmt.Println()
	*/

	arr := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	tree := arr2tree(arr, 0)
	stampaAlberoASommario(tree, 0)
}

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}

type bitree struct {
	root *bitreeNode
}

func newNode(val int) *bitreeNode {
	return &bitreeNode{nil, nil, val}
}

func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	preorder(node.left)
	preorder(node.right)
}

func inorder(node *bitreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Println(node.val)
	inorder(node.right)
}

func postorder(node *bitreeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Println(node.val)
}

func stampaAlberoASommario(node *bitreeNode, spaces int) {
	if node == nil {
		printSpaces(spaces)
		fmt.Println("*")
		return
	}
	printSpaces(spaces)
	fmt.Print("*", node.val, "\n")
	if node.left == nil && node.right == nil {
		return
	}
	stampaAlberoASommario(node.left, spaces+1)
	stampaAlberoASommario(node.right, spaces+1)
}

func printSpaces(spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
}

func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}

	root = newNode(a[i])
	root.left = arr2tree(a, i*2+1)
	root.right = arr2tree(a, i*2+2)
	return root
}

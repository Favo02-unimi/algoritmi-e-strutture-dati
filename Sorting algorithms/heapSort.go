package main

import (
	"fmt"
)

func heapSort(arr []int) []int {
	h := heap{arr}
	generateHeap(h)
	for i := len(arr)-1; i > 0; i-- {
		arr[i] = h.arr[i]
		h.arr[0], h.arr[i] = h.arr[i], h.arr[0]
		h = reorder(heap{h.arr[:i]}, 0)
	}
	return arr
}

func main() {
	sorted := heapSort([]int{7, 9, 12, 2, 15, 25, 3, 14, 10})
	fmt.Println("f", sorted)
}

// TREE
type heap struct {
	arr []int
}

func (t heap) get(index int) int {
	if index == -1 {
		return 0
	}
	return t.arr[index]
}

func (t *heap) set(index, val int) {
	t.arr[index] = val
}

func (t heap) getChildrenIndex(index int) (int, int) {
	if index*2+1 >= len(t.arr) {
		return -1, -1
	} else if index*2+2 >= len(t.arr) {
		return (index * 2) + 1, -1
	} else {
		return (index * 2) + 1, (index * 2) + 2
	}
}

func reorder(t heap, root int) heap {

	v := root
	x := t.get(v)
	for true {
		c1, c2 := t.getChildrenIndex(v)

		// both leaf, stop
		if c1 == -1 {
			break
		}

		maxChild := c1
		if c2 != -1 && t.get(c2) > t.get(c1) { // check second child
			maxChild = c2
		}

		// swap
		if t.get(maxChild) > x {
			t.set(v, t.get(maxChild))
			v = maxChild
		} else {
			break
		}
	}

	// move root
	t.set(v, x)

	return t
}

func generateHeap(tree heap) heap {
	for i := len(tree.arr) -1; i >= 0; i-- {
		tree = reorder(tree, i)
	}
	return tree
}

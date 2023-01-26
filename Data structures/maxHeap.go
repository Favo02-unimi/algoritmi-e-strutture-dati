package main

import (
	"fmt"
	"math"
)

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

func (t heap) getParentIndex(index int) int {
	if index%2 == 0 {
		return (index / 2) - 1
	} else {
		return index / 2
	}
}

func (t heap) getParent(index int) int {
	return t.arr[t.getParentIndex(index)]
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

func (t heap) getChildren(index int) (int, int, bool, bool) {
	c1, c2 := t.getChildrenIndex(index)
	if c1 == -1 {
		return 0, 0, false, false
	} else if c2 == -1 {
		return t.arr[c1], 0, true, false
	} else {
		return t.arr[c1], t.arr[c2], true, true
	}
}

func (t heap) getHeight() int {
	return int(math.Log2(float64(len(t.arr))))
}

func (t heap) getNodesByDepth(depth int) []int {
	var nodes []int
	nodesNum := int(math.Pow(2, float64(depth)))
	for i := nodesNum - 1; i < nodesNum; i++ {
		nodes = append(nodes, t.get(i))
	}
	return nodes
}

func topReorder(t heap, root int) heap {

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

func bottomReorder(t heap, leaf int) heap {
	v := leaf
	x := t.get(v)
	for true {
		parent := t.getParentIndex(v)
		fmt.Println(parent, t.get(parent))

		if (v == 0) {
			break
		}

		if t.get(parent) < x {
			t.set(v, t.get(parent))
			v = parent
		} else {
			break
		}
	}

	t.set(v, x)

	return t
}

func recGenerateHeap(heap heap, root int) heap {
	if root != -1 {
		cl, cr := heap.getChildrenIndex(root)
		recGenerateHeap(heap, cl)
		recGenerateHeap(heap, cr)
		topReorder(heap, root)
	}
	return heap
}

func generateHeap(heap heap) heap {
	for i := len(heap.arr) -1; i >= 0; i-- {
		heap = topReorder(heap, i)
	}
	return heap
}

func (t heap) findMax() int {
	return t.get(0)
}

func (t *heap) deleteMax() int {
	max := t.findMax()
	t.arr[0], t.arr[len(t.arr)-1] = t.arr[len(t.arr)-1], t.arr[0]
	t.arr = t.arr[:len(t.arr)-1]
	topReorder(*t, 0)
	return max
}

func (t *heap) insert(val int) {
	t.arr = append(t.arr, val)
	bottomReorder(*t, len(t.arr)-1)
}

func (t *heap) delete(index int) {
	last := t.arr[len(t.arr)-1]
	del := t.arr[index]
	t.arr[index], t.arr[len(t.arr)-1] = t.arr[len(t.arr)-1], t.arr[index]
	t.arr = t.arr[:len(t.arr)-1]

	if (last > del) {
		bottomReorder(*t, index)
	} else {
		topReorder(*t, index)
	}
}

func (t *heap) edit(index, val int) {
	old := t.get(index)
	t.set(index, val)
	if val > old {
		bottomReorder(*t, index)
	} else {
		topReorder(*t, index)
	}
}

func main() {
	heap := heap{[]int{7, 9, 12, 2, 15, 25, 3, 14, 10, 34, 5}}
	fmt.Println(generateHeap(heap))
	heap.insert(18)
	fmt.Println(heap)
	heap.delete(3)
	fmt.Println(heap)
	heap.edit(0, 0)
	fmt.Println(heap)
}

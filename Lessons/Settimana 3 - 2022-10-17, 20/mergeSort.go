package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	slice1 := generateRandomSlice(10, seed)
	fmt.Println("slice:\n", slice1)
	fmt.Println("slice ordinata con mergeSort:\n", mergeSort(slice1))
}

func generateRandomSlice(len int, seed int64) []int {
	rand.Seed(seed)
	var slice []int
	for i := 0; i < len; i++ {
		slice = append(slice, rand.Intn(100))
	}
	return slice
}

func mergeSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}
	var m int = len(slice)/2
	return merge(mergeSort(slice[:m]), mergeSort(slice[m:]))
}

func merge(a, b []int) []int {
	// dato che conosco già la lunghezza finale della slice mergeata
	// la inizializzo direttamente senza utilizzare append per evitare
	// inutili relocazioni
	var slice []int = make([]int, len(a)+len(b))
	var ai, bi, ci int
	for ai < len(a) && bi < len(b) {
		if a[ai] < b[bi] {
			slice[ci] = a[ai]
			ai++
		} else {
			slice[ci] = b[bi]
			bi++
		}
		ci++
	}
	if ai < len(a) {
		for ; ai < len(a); ai++ {
			slice[ci] = a[ai]
			ci++
		}
	}
	if bi < len(b) {
		for ; bi < len(b); bi++ {
			slice[ci] = b[bi]
			ci++
		}
	}
	return slice
}

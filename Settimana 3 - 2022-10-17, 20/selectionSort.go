package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	slice1 := generateRandomSlice(10, seed)
	slice2 := generateRandomSlice(10, seed)
	fmt.Println("slice:\t", slice1)
	fmt.Println("selection sort iter:\t", selectionSortIter(slice1))
	fmt.Println("selection sort rec:\t", selectionSortIter(slice2))
}

func generateRandomSlice(len int, seed int64) []int {
	rand.Seed(seed)
	var slice []int
	for i := 0; i < len; i++ {
		slice = append(slice, rand.Intn(100))
	}
	return slice
}

func selectionSortIter(slice []int) []int {
	for parteDis := len(slice); parteDis > 1; parteDis-- {

		var maxIndex int = 0
		for i := 0; i < parteDis; i++ {
			if slice[i] > slice[maxIndex] {
				maxIndex = i
			}
		}
		slice[parteDis-1], slice[maxIndex] = slice[maxIndex], slice[parteDis-1]
		
	}
	return slice
}

func selectionSortRec(slice []int) []int {
	if len(slice) == 0 || len(slice) == 1 {
		return slice
	}

	var maxIndex int = 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > slice[maxIndex] {
			maxIndex = i
		}
	}
	slice[maxIndex], slice[len(slice)-1] = slice[len(slice)-1], slice[maxIndex]

	return selectionSortRec(slice[:len(slice)-1])
}

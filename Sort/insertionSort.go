package main

import "fmt"

func main() {
	arr := []int{5, 4, 8, 9, 2, 10, 1}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(array []int) {
	for sortedPart := 1; sortedPart < len(array); sortedPart++ {
		var elemToSort int = array[sortedPart]
		var i int
		for i = sortedPart - 1; i >= 0 && array[i] > elemToSort; i-- {
			array[i+1] = array[i]
		}
		array[i+1] = elemToSort
	}
}

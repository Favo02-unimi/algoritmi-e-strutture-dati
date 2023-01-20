package main

import "fmt"

func main() {
	arr := []int{5, 4, 8, 9, 2, 10, 1}
	fmt.Println(arr)
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		var minIndex int = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[minIndex], array[i] = array[i], array[minIndex]
	}
}

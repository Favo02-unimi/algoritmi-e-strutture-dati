package main

import "fmt"

func main() {
	arr := []int{5, 4, 8, 9, 2, 10, 1}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

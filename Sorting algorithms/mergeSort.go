package main

import "fmt"

func main() {
	arr1 := []int{3, 5, 7, 8, 14, 40, 50, 60}
	arr2 := []int{1, 2, 3, 6, 9, 15, 20, 21}
	merged := merge(arr1, arr2)
	fmt.Println(merged)

	arr3 := []int{100, 24, 3, 63, 8, 15, 20, 21, 1, -10, 1000}
	sorted := mergeSort(arr3)
	fmt.Println(sorted)
}

func merge(arr1, arr2 []int) []int {
	var i1, i2, imerge int
	var merge []int = make([]int, len(arr1)+len(arr2))
	for i1 < len(arr1) && i2 < len(arr2) {
		if arr1[i1] <= arr2[i2] {
			merge[imerge] = arr1[i1]
			i1++
		} else {
			merge[imerge] = arr2[i2]
			i2++
		}
		imerge++
	}
	if i1 != len(arr1) {
		for i := i1; i < len(arr1); i++ {
			merge[imerge] = arr1[i]
			imerge++
		}
	}
	if i2 != len(arr2) {
		for i := i2; i < len(arr2); i++ {
			merge[imerge] = arr2[i]
			imerge++
		}
	}
	return merge
}

func mergeSort(array []int) []int {
	if len(array) > 1 {
		mid := len(array) / 2
		part1 := mergeSort(array[:mid])
		part2 := mergeSort(array[mid:])
		array = merge(part1, part2)
	}
	return array
}

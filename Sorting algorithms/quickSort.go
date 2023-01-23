package main

import "fmt"

func main() {
	arr := []int{4, 5, 32, 5, 65, 1, 3, 0, 99}
	arr = quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) []int {
	if len(arr) > 1 {
		perno := partiziona(arr, 0, len(arr)-1)
		quickSort(arr[:perno])
		quickSort(arr[perno:])
	}
	return arr
}

func partiziona(arr []int, i, f int) int {
	var perno int = i

	for i != f { // indici non incrociati
		for i != f && arr[i] <= arr[perno] {
			i++
		}
		for f != i && arr[f] > arr[perno] {
			f--
		}
		arr[i], arr[f] = arr[f], arr[i] // scambiare elementi
	}

	arr[perno], arr[i-1] = arr[i-1], arr[perno] // spostare persno in posizione corretta
	return i
}

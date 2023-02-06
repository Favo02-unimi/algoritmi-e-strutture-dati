package main

import "fmt"

func main() {
	fmt.Println(seqCrescente([]int{12, 3, 5, 7, 8, 4, 6, 1, 2, 3, 4, 5, 6, 7, 8}))
}

// la complessità di questa funzione è theta(n), dove n è la lunghezza dell'array da cui estrarre la sequenza crescente più lunga
// ogni elemento viene controllato una sola volta
func seqCrescente(arr []int) []int {

	var maxSeq []int

	if len(arr) == 0 {
		return maxSeq
	}

	last := arr[0]
	curSeq := make([]int, 0)

	for i := 1; i < len(arr); i++ {
		if arr[i] < last {
			curSeq = make([]int, 0, 1)
			curSeq = append(curSeq, arr[i])
		} else {
			curSeq = append(curSeq, arr[i])
			if len(curSeq) > len(maxSeq) {
				maxSeq = curSeq
			}
		}
		last = arr[i]
	}

	return maxSeq
}

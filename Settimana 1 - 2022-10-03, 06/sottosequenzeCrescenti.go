package main

import . "fmt"

/*
idea:
contate i minimi contanto le "v" (i-1 > i < i+1) al posto delle salite vere e proprie
*/

func main() {
	n := []int{9, 1, 3, 5, 2, 0, 8, 6, 5, 8}
	var count int

	for i := 1; i < len(n)-1; i++ {
		if n[i-1] > n[i] && n[i] < n[i+1] {
			count++
		}
	}

	Println(count)
}

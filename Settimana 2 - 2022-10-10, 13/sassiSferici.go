package main

import (
	. "fmt"
)

func main() {
	Print("altezza piramide: ")
	var h int
	Scan(&h)
	Println("sassi:", sassi(h))
}

func sassi(height int) int {
	if (height == 1) {
		return 1
	}
	return height*height + sassi(height-1)
}

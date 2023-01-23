package main

import (
	. "fmt"
)

/*
Le torri devono essere globali perchè nella chiamate ricorsive
vengono "mischiati" i paletti, in base a cosa si vuole spostare.
Nella visualizzazione invece è necessario che siano sempre nello
stesso ordine, quindi non è possibile passarle (o farle resituire)
in ordine alla funzione ricorsiva
*/

var towers [3][]rune

func main() {
	var n int
	Print("numero dischi: ")
	Scan(&n)
	Println("numero di mosse:", countHanoi(n, 0, 1, 2))

	var char rune = 'A'
	for i := 0; i < n; i++ {
		towers[0] = append(towers[0], char)
		char++
	}
	Println("mosse:")
	printRune(towers[0], towers[1], towers[2])
	hanoi(n, 0, 1, 2)
}

func countHanoi(n, from, temp, to int) int {
	if (n == 1) {
		return 1
	} else {
		return countHanoi(n-1, from, to, temp) + 1 + countHanoi(n-1, temp, from, to)
	}
}

func hanoi(n, from, temp, to int) {
	if (n == 1) {
		Println(from, "->", to)
		move(from, to)
	} else {
		hanoi(n-1, from, to, temp)
		Println(from, "->", to)
		move(from, to)
		hanoi(n-1, temp, from, to)
	}
}

func move(from, to int) {
	towers[to] = append(towers[to], towers[from][len(towers[from])-1])
	towers[from] = towers[from][:len(towers[from])-1]
	printRune(towers[0], towers[1], towers[2])
}

func printRune(from, temp, to []rune) {
	for _, v := range from {
		Print(string(v))
	}
	Print(", ")
	for _, v := range temp {
		Print(string(v))
	}
	Print(", ")
	for _, v := range to {
		Print(string(v))
	}
	Println()
}

package main

import . "fmt"

/*
idea:
ogni b che trovo chiude tante sottostringhe che iniziano con a quante a ci sono prima, quindi basta un solo ciclo che conta le a: ogni volta che trovo una b sommo al contatore finale il numero di a trovate fino a lì
*/

func main() {
	var str string
	Scan(&str)

	var countA, count int

	for _, v := range str {
		if v == 'a' {
			countA++
		}
		if v == 'b' {
			count += countA
		}
	}

	Println(count)
}

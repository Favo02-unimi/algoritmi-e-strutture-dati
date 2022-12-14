package main

import . "fmt"

func main() {
	var saldo, spesa int
	Scan(&saldo)

	for saldo > 0 {
		Scan(&spesa)
		saldo -= spesa
		Println("saldo:", saldo)
	}

}

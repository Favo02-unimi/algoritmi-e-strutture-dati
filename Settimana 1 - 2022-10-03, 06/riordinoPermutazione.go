package main

import . "fmt"

/*
idea:
dato che so che è una permutazione, quindi non ci sono salti e partono da 1, posso mettere direttamente al posto giusto
scorro l'array 1 sola volta: ad ogni elemento lo metto nella posizione corretta (length - id) scambiandolo con quello che c'è in quella posizione.
Solo se l'elemento nella posizione che sto analizzando continuo, altrimenti metto al posto giusto quello alla posizione corrente.
Faccio solo n swap, dato che tutti verranno messi direttamente al posto giusto
*/

func main() {

	// al posto di una mappa uso un array che sono le chiavi della mappa, solo per semplicità

	keys := []int{6, 7, 4, 2, 3, 1, 8, 5, 9}

	var i int

	for i < len(keys) {

		if i == len(keys)-keys[i] {
			i++
		} else {
			keys[i], keys[len(keys)-keys[i]] = keys[len(keys)-keys[i]], keys[i]
		}

	}

	Print(keys)
}

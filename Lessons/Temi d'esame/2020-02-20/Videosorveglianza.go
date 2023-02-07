package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
La struttura dati per rappresentare i depositi di merce è un grafo senza cicli e
con un punto di ingresso (quindi un albero).
Ogni stanza del deposito è un nodo, mentre ogni corridoio è un arco che collega
due vertici (due stanze)
*/

/*
Formato INPUT:
s S0 1
s S1 1
s S2 1
s S3 100
s S4 1
s S5 99
s S6 2
s S7 2
c S0 S1
c S1 S2
c S1 S3
c S1 S4
c S3 S5
c S3 S6
c S3 S7
*/

func main() {
	d := costruisciDeposito()
	impianto := d.impianto()
	fmt.Println(impianto.valoreImpianto())
}

type deposito struct {
	payloads map[string]*stanza
	albero   map[string][]string // lista di adiacenza
}

type stanza struct {
	nome       string
	valore     int
	telecamera bool
	locked     bool // se la presenza o meno di telecamera è definitiva
}

// costruisce un deposito senza impianto dall'input
func costruisciDeposito() deposito {
	payloads := make(map[string]*stanza)
	albero := make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		if tokens[0] == "s" { // stanza
			val, _ := strconv.Atoi(tokens[2])
			stanza := &stanza{tokens[1], val, false, false}
			payloads[tokens[1]] = stanza
		}

		if tokens[0] == "c" { // corridoio
			s1 := tokens[1]
			s2 := tokens[2]
			albero[s1] = append(albero[s1], s2)
			albero[s2] = append(albero[s2], s1)
		}
	}

	return deposito{payloads, albero}
}

// restituisce true se l'impianto è valido, false altrimenti
func (d deposito) impiantoValido() bool {
	for nome, stanza := range d.payloads {

		// ha telecamera, adiacenti non devono averla
		if stanza.telecamera {
			for _, adiacente := range d.albero[nome] {
				if d.payloads[adiacente].telecamera {
					return false
				}
			}
		}

		// non ha telecamera, almeno 1 adiacente deve averla
		if !stanza.telecamera {
			trovata := false
			for _, adiacente := range d.albero[nome] {
				if d.payloads[adiacente].telecamera {
					trovata = true
					break
				}
			}
			if !trovata {
				return false
			}
		}

	}
	return true
}

// costruisce un deposito con impianto di videosorveglianza valido
func (d *deposito) impianto() deposito {
	for nome, stanza := range d.payloads {
		fmt.Println(nome)
		if stanza.locked {
			continue
		}

		d.piazzaTelecamera(nome)
	}

	if !d.impiantoValido() {
		fmt.Println("ERRORE IMPIANTO INVALIDO")
	}

	return *d
}

// piazza una o più telecamere partendo dal nodo "nome"
func (d *deposito) piazzaTelecamera(nome string) {

	if nodoMaggioreAdiacenti(*d, nome) {
		// se il nodo è maggiore degli adiacenti, allora prendi lui e no gli adiacenti
		d.payloads[nome].telecamera = true
		d.payloads[nome].locked = true

		for _, ad := range d.albero[nome] {
			d.payloads[ad].telecamera = false
			d.payloads[ad].locked = true
		}
	} else {
		// se gli adiacenti sono maggiori del nodo E gli adiacenti degli adiacenti
		// sono minori degli adiacenti allora prendi tutti gli adiacenti

		// se almeno un adiacente ha gli adiacenti maggiori di lui, allora prendi
		// il nodo originale e non gli adiacenti
		adiacentiMaggioriAdiacenti := true
		for _, ad := range d.albero[nome] {
			if !nodoMaggioreAdiacenti(*d, ad) {
				adiacentiMaggioriAdiacenti = false
			}
		}

		if adiacentiMaggioriAdiacenti {
			// prendo adiacenti e non il nodo
			d.payloads[nome].telecamera = false
			d.payloads[nome].locked = true
			for _, ad := range d.albero[nome] {
				d.payloads[ad].telecamera = true
				d.payloads[ad].locked = true
			}
		} else {
			// prendo il nodo e non gli adiacenti
			d.payloads[nome].telecamera = true
			d.payloads[nome].locked = true
			for _, ad := range d.albero[nome] {
				d.payloads[ad].telecamera = false
				d.payloads[ad].locked = true
			}
		}
	}

}

// restituisce true se il nodo è maggiore della somma degli adiacenti,
// false se la somma degli adiacenti è maggiore del nodo
func nodoMaggioreAdiacenti(d deposito, padre string) bool {
	adiacenti := d.albero[padre]

	valAdicenti := 0
	for _, s := range adiacenti {
		valAdicenti += d.payloads[s].valore
	}

	return d.payloads[padre].valore >= valAdicenti
}

// restituisce il valore dell'impianto, -1 se non è valido
func (d deposito) valoreImpianto() int {
	fmt.Println("impianto:")
	if !d.impiantoValido() {
		return -1
	}
	valTot := 0
	for _, v := range d.payloads {
		if v.telecamera {
			fmt.Println(v.nome)
			valTot += v.valore
		}
	}
	return valTot
}

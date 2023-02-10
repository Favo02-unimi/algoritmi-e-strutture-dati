package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var p pulsanti

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		switch tokens[0] {
		case "+": // + 7
			n, _ := strconv.Atoi(tokens[1])
			p = inizializzaPulsantiSpenti(n)

		case "o": // o 0100011
			p = inizializzaPulsanti(tokens[1])

		case "p": // p
			fmt.Println(p.cur.toString())

		case "s": // s 4
			n, _ := strconv.Atoi(tokens[1])
			p = p.premiInterruttore(n)

		case "S":
			for i := 1; i < len(tokens); i++ {
				n, _ := strconv.Atoi(tokens[i])
				p = p.premiInterruttore(n)
			}

		case "x":
			t := make([]bool, len(p.cur.interruttori))
			target := stato{t}

			dist, paths := p.numeroSpegniTutto()
			fmt.Println("distanza a", target.toString(), ":", dist[target.toString()])

			path := ricostruisciPath(p.cur, target, paths)

			cur := p.cur
			for i := len(path) - 1; i >= 0; i-- {
				cur = cur.premiInterruttore(path[i])
				fmt.Println("premi", path[i], ":", cur.toString())
			}

		case "f":
			return

		}
	}

}

// PULSANTI

type pulsanti struct {
	cur stato
}

// l'operatore % di go si comporta in modo "strano" sui moduli di numeri negativi, quindi non va bene per lista circolare
func mod(a, b int) int {
	return (a%b + b) % b
}

// inizializza "numeri" pulsanti tutti spenti
func inizializzaPulsantiSpenti(numero int) pulsanti {
	// creazione spenti
	interruttori := make([]bool, numero)
	s := stato{interruttori}

	return pulsanti{s}
}

// inizializza len(p) pulsanti impostati allo stato della stringa p (composta da 0 e 1)
func inizializzaPulsanti(p string) pulsanti {
	s := fromString(p)

	grafo := make(map[*stato]map[*stato]int)
	grafo[&s] = make(map[*stato]int)

	return pulsanti{s}
}

// modifica e restituisce i pulsanti premendo il pulsante "numeroInterruttore"
func (p *pulsanti) premiInterruttore(numeroInterruttore int) pulsanti {
	newStato := p.cur.premiInterruttore(numeroInterruttore)
	p.cur = newStato

	return *p
}

// restituisce il numero di pulsanti da premere per raggiungere qualsiasi stato di pulsanti
func (p pulsanti) numeroSpegniTutto() (map[string]int, map[string]int) {
	dist := make(map[string]int) // distanza da stato curr di p ad ogni stato raggiungibile
	last := make(map[string]int) // stato da cui si arriva per ogni possibile stato (precedente)

	queue := make([]*stato, 0, 1)

	dist[p.cur.toString()] = 0

	queue = append(queue, &p.cur)

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		raggiungibili := p.generaStati(c)

		for st, int := range raggiungibili {
			if _, found := dist[st.toString()]; !found {
				dist[st.toString()] = dist[c.toString()] + 1
				last[st.toString()] = int
				queue = append(queue, st)
			}
		}
	}

	return dist, last
}

func (p *pulsanti) generaStati(cur *stato) map[*stato]int {
	var LEN = len(p.cur.interruttori)
	raggiungibili := make(map[*stato]int)

	for i := 1; i <= LEN; i++ {
		s := cur.premiInterruttore(i)
		raggiungibili[&s] = i
	}

	return raggiungibili
}

func ricostruisciPath(start, end stato, paths map[string]int) []int {
	path := make([]int, 0)

	curNode := end
	for curNode.toString() != start.toString() {
		intPremuto := paths[curNode.toString()]

		path = append(path, intPremuto)
		curNode = curNode.premiInterruttore(intPremuto)
	}

	return path
}

// STATO

type stato struct {
	interruttori []bool
}

// crea uno stato data una stringa composta da 0 e 1
func fromString(s string) stato {
	newInterruttori := make([]bool, 0)
	for _, v := range s {
		if v == '0' {
			newInterruttori = append(newInterruttori, false)
		} else {
			newInterruttori = append(newInterruttori, true)
		}
	}

	return stato{newInterruttori}
}

// converte uno stato da una stringa composta da 0 e 1
func (s stato) toString() string {
	var res string

	for _, v := range s.interruttori {
		if v == true {
			res += "1"
		} else {
			res += "0"
		}
	}

	return res
}

// restituisce un NUOVO stato (senza modificare il precedente) ottenuto premendo il pulsante passato
func (s stato) premiInterruttore(numeroInterruttore int) stato {
	var DIM = len(s.interruttori)

	// clonare vecchio stato
	newInterruttori := make([]bool, 0, DIM)
	for _, v := range s.interruttori {
		newInterruttori = append(newInterruttori, v)
	}

	// numero interrutore da 1-n a 0-(n-1)
	indexInterruttore := mod(numeroInterruttore-1, DIM)

	newInterruttori[mod(indexInterruttore-1, DIM)] = !newInterruttori[mod(indexInterruttore-1, DIM)]
	newInterruttori[mod(indexInterruttore, DIM)] = !newInterruttori[mod(indexInterruttore, DIM)]
	newInterruttori[mod(indexInterruttore+1, DIM)] = !newInterruttori[mod(indexInterruttore+1, DIM)]

	return stato{newInterruttori}
}

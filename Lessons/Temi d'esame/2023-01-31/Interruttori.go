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
			target := &stato{t}

			dist := p.numeroSpegniTutto(*target)
			for st, di := range dist {
				fmt.Println(st.toString(), di)
			}

			fmt.Println(dist[target])

		case "f":
			return

		}
	}

}

// PULSANTI

type pulsanti struct {
	cur   stato
	grafo map[*stato]map[*stato]int // lista di adiacenza, ogni stato punta agli stati che può raggiungere. ad ogni arco è anche collegato il numero di interruttore da premere per ottenere quella transizione
}

// l'operatore % di go si comporta in modo "strano" sui moduli di numeri negativi, quindi non va bene per lista circolare
func mod(a, b int) int {
	return (a%b + b) % b
}

func inizializzaPulsantiSpenti(numero int) pulsanti {
	// creazione spenti
	interruttori := make([]bool, numero)
	s := stato{interruttori}

	grafo := make(map[*stato]map[*stato]int)
	grafo[&s] = make(map[*stato]int)

	return pulsanti{s, grafo}
}

func inizializzaPulsanti(p string) pulsanti {
	s := fromString(p)

	grafo := make(map[*stato]map[*stato]int)
	grafo[&s] = make(map[*stato]int)

	return pulsanti{s, grafo}
}

func (p *pulsanti) premiInterruttore(numeroInterruttore int) pulsanti {
	newStato := p.cur.premiInterruttore(numeroInterruttore)

	if p.grafo[&p.cur] == nil {
		p.grafo[&p.cur] = make(map[*stato]int)
	}
	p.grafo[&p.cur][&newStato] = numeroInterruttore

	if p.grafo[&newStato] == nil {
		p.grafo[&newStato] = make(map[*stato]int)
	}
	p.grafo[&newStato][&p.cur] = numeroInterruttore

	p.cur = newStato

	return *p
}

func (p pulsanti) numeroSpegniTutto(target stato) map[*stato]int {
	dist := make(map[*stato]int)
	queue := make([]*stato, 0, 1)

	dist[&p.cur] = 0

	queue = append(queue, &p.cur)

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		p.generaStati(c)

		for s := range p.grafo[c] {
			if _, found := dist[s]; !found {
				dist[s] = dist[c] + 1
				queue = append(queue, s)
			}

			if s.toString() == target.toString() { // target found
				return dist
			}
		}
	}

	return dist
}

func (p *pulsanti) generaStati(cur *stato) pulsanti {
	var LEN = len(p.cur.interruttori)

	for i := 1; i <= LEN; i++ {
		s := cur.premiInterruttore(i)

		if p.grafo[cur] == nil {
			p.grafo[cur] = make(map[*stato]int)
		}
		p.grafo[cur][&s] = i

		if p.grafo[&s] == nil {
			p.grafo[&s] = make(map[*stato]int)
		}
		p.grafo[&s][cur] = i
	}

	return *p
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

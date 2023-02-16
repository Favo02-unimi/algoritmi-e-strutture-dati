package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	rete := creaRete()

	dist, path := rete.distanza("Repubblica", "Lima")

	fmt.Println(formattaRisultato(path, dist))
}

type rete struct {
	// lista di adiacenza (chiave: "<nomestazione>[<numerolinea>]")
	grafo    map[string][]string
	payloads map[string]stazione
}

type stazione struct {
	nome  string
	linee []int
}

// costruisce la rete di linee metropolitane da stdin
func creaRete() rete {
	payloads := make(map[string]stazione)
	grafo := make(map[string][]string)

	// creare linee separate
	scanner := bufio.NewScanner(os.Stdin)
	for linea := 1; scanner.Scan(); linea++ {
		line := strings.ReplaceAll(scanner.Text()[9:], ".", "")
		stazioni := strings.Split(line, ", ")

		// creare tutte le stazioni
		for i := 0; i < len(stazioni); i++ {
			nome := stazioni[i]

			// creare payload per ogni stazione
			if _, found := payloads[nome]; found {
				old := payloads[nome]
				old.linee = append(old.linee, linea)
				payloads[nome] = old
			} else {
				payloads[nome] = stazione{nome, []int{linea}}
			}

			// creare collegamenti (grafo)

			// caso speciale prima stazione (1 collegamento)
			if i == 0 {
				next := nomeConLinea(stazioni[1], linea)
				grafo[nomeConLinea(nome, linea)] = []string{next}
			}

			// caso speciale ultima stazione (1 collegamento)
			if i == len(stazioni)-1 {
				prec := nomeConLinea(stazioni[len(stazioni)-2], linea)
				grafo[nomeConLinea(nome, linea)] = []string{prec}
			}

			// stazioni in mezzo (2 collegamenti)
			if i > 0 && i < len(stazioni)-1 {
				next := nomeConLinea(stazioni[i-1], linea)
				prec := nomeConLinea(stazioni[i+1], linea)
				grafo[nomeConLinea(nome, linea)] = []string{next, prec}
			}
		}
	}

	// unire stazioni interscambio
	for nome, s := range payloads {
		if len(s.linee) > 1 { // le stazioni con più di 1 linea

			// aggiungere ad ogni stazione di ogni lina il collegamento alla stessa
			// stazione di tutte le altre linee connesse
			for i := 0; i < len(s.linee); i++ {
				for j := 0; j < len(s.linee); j++ {
					if s.linee[i] == s.linee[j] { // non aggiungere a Cadorna[1] anche Cadorna[1]
						continue
					}
					cur := nomeConLinea(nome, s.linee[i])
					conn := nomeConLinea(nome, s.linee[j])
					grafo[cur] = append(grafo[cur], conn)
				}
			}

		}
	}

	return rete{grafo, payloads}
}

// restituisce il nome della stazione formattato con la sua linea: <nome>[<linea>]
func nomeConLinea(nome string, linea int) string {
	return fmt.Sprint(nome, "[", linea, "]")
}

// restituisce la distanza minima tra stazione "start" e stazione "end"
func (r rete) distanza(start, end string) (int, []string) {

	minDist := math.MaxInt
	var path []string

	// la stazione di partenza e di fine possono essere più di una se sono uno scambio
	starts := r.payloads[start]
	ends := r.payloads[end]

	// provo come start tutte le linee diverse della stazione start
	for _, lineaStart := range starts.linee {
		dist, prec := r.bfs(nomeConLinea(start, lineaStart))

		// provo come end tutte le linee diverse della stazione end
		for _, lineaEnd := range ends.linee {

			// salvo solo la più vicina
			if dist[nomeConLinea(end, lineaEnd)] < minDist {
				minDist = dist[nomeConLinea(end, lineaEnd)]
				path = ricostruisciPath(prec, nomeConLinea(end, lineaEnd), nomeConLinea(start, lineaStart))
			}
		}
	}

	return minDist, path
}

// effettua una bfs e calcola il cammino minimo tra "start" e tutte le altre stazioni
// restituisce anche il precedente di ogni stazione per poter ricostruire il percorso
func (r rete) bfs(start string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	dist[start] = 0

	prec := make(map[string]string)

	queue := []string{start}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		for _, ad := range r.grafo[c] {
			if _, found := dist[ad]; !found {
				dist[ad] = dist[c] + 1
				prec[ad] = c
				queue = append(queue, ad)
			}
		}
	}

	return dist, prec
}

// ricostruisce la path conoscendo arrivo, inizio e il precedente di ogni stazione
func ricostruisciPath(paths map[string]string, end, start string) []string {
	path := make([]string, 0)
	for end != start {
		path = append([]string{end}, path...) // prepend
		end = paths[end]
	}
	path = append([]string{end}, path...) // prepend

	return path
}

// formmata la distanza tra due stazioni conoscendo la path
func formattaRisultato(path []string, dist int) string {
	res := fmt.Sprint("Da ", path[0], " a ", path[len(path)-1], ": ", dist, " cambi\n")
	res = fmt.Sprint(res, path[0])
	for i := 1; i < len(path); i++ {
		res = fmt.Sprint(res, " --> ", path[i])
	}
	return res
}

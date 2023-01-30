package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
La situazione descritta nella traccia è rappresentabile con un grafo non direzionato pesato in cui ogni coppia di vertici è connessa (cricca), da cui è necessario ricavare l'albero ricoprente minimo. Ogni punto è un vertice del grafo, mentre gli archi del grafo sono tutti i collegamenti da un punto a tutti gli altri. Il peso di ogni arco è dato dalla distanza di manhattan tra i due vertici toccati dall'arco.
*/

type grafo struct {
	m map[point]map[point]int // lista di adiacenza, al posto di liste utilizzo delle mappe per comodità
	// ad ogni punto della prima mappa è associata una mappa che contiene tutti i putni che può raggiungere, ognuno associato al peso necessario per raggiungerlo
}

type point struct {
	x, y int
}

type wedge struct {
	start, end point
	peso       int
}

func newGrafo() grafo {
	return grafo{make(map[point]map[point]int)}
}

func (g *grafo) addArcoNonDirezionato(start, end point, peso int) {
	if g.m[start] == nil {
		g.m[start] = make(map[point]int)
	}
	if g.m[end] == nil {
		g.m[end] = make(map[point]int)
	}

	g.m[start][end] = peso
	g.m[end][start] = peso
}

// kruskal con set implementati come mappe:
// - inizializzare set per ogni vertice: theta(V)
// - aggiungere tutti gli archi in una lista: theta(E) = in questo caso theta(V^2) dato che cricca
// - ordinare lista archi: O(E log E) = in questo caso O(V^2 log V^2) dato che cricca
// - scorrere tutti gli archi fino ad albero ricoprente finito: O(E) = O(V^2)
// - - unire set: O(1)
// - - cercare in set: O(V)
// sarebbe opportuno utilizzare quickunion con compressione di cammino per ridurre la complessità
func alberoRicoprenteMinimo(g grafo, start point) (grafo, int) {

	// each point to parent set (or nil if own set)
	sets := make(map[point]*point)

	// each point in his own set (parent set = nothing)
	for v := range g.m {
		sets[v] = nil
	}

	// list of edges
	edges := make([]wedge, 0, len(g.m)*len(g.m))
	for a, m := range g.m {
		for b, peso := range m {
			edges = append(edges, wedge{a, b, peso})
		}
	}

	// sort edges
	sort.Slice(edges, func(i2, j int) bool {
		return edges[i2].peso < edges[j].peso
	})

	spanning := newGrafo()
	pesoTot := 0

	for _, arco := range edges {
		if len(spanning.m) == len(g.m) {
			break
		}

		setStart := nomeSet(sets, arco.start)
		setEnd := nomeSet(sets, arco.end)

		if setStart != setEnd {
			spanning.addArcoNonDirezionato(arco.start, arco.end, arco.peso)
			pesoTot += arco.peso
			sets[setEnd] = &setStart
		}
	}

	return spanning, pesoTot
}

func nomeSet(sets map[point]*point, key point) point {
	if sets[key] == nil {
		return key
	}
	return nomeSet(sets, *sets[key])
}

func main() {
	var nPunti int
	fmt.Scan(&nPunti)
	scanner := bufio.NewScanner(os.Stdin)
	punti := make([]point, 0, nPunti)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		punti = append(punti, point{x, y})
	}

	// generaione grafo (cricca): theta(v^2), v = #vertici = #punti
	g := newGrafo()
	for _, i := range punti {
		for _, j := range punti {
			if i == j {
				continue
			}
			g.addArcoNonDirezionato(i, j, manhattanDistance(i, j))
		}
	}

	// generazione albero ricoprente minimo: O(V^2 log V^2)
	ricoprente, pesoTot := alberoRicoprenteMinimo(g, punti[0])
	for p, m := range ricoprente.m {
		fmt.Println(p, ":", m)
	}
	fmt.Println(pesoTot)
}

func manhattanDistance(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

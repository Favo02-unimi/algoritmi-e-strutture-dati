package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	oggetti := leggiInput()

	foresta := costruisciForesta(oggetti)
	stampaAlbero(foresta, "frigorifero")
	res := calcolaPrezzo(foresta, "frigorifero")
	fmt.Println(res)
}

type oggetto struct {
	name string
	val  int
	op   rune
	sx   string
	dx   string
	tipo string
}

func leggiInput() map[string]*oggetto {
	oggetti := make(map[string]*oggetto)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		nome := tokens[0][:len(tokens[0])-1] // remove : at end of name

		if len(tokens) == 2 { // nome: val
			val, _ := strconv.Atoi(tokens[1])
			o := oggetto{nome, val, 0, "", "", "val"}
			oggetti[nome] = &o
		} else { // nome: sx op dx
			o := oggetto{nome, 0, rune(tokens[2][0]), tokens[1], tokens[3], "op"}
			oggetti[nome] = &o
		}
	}

	return oggetti
}

type foresta struct {
	payloads map[string]*oggetto
	radici   []*nodo
}

type nodo struct {
	nome               string
	figliosx, figliodx *nodo
}

func costruisciForesta(oggetti map[string]*oggetto) foresta {
	notRoot := make(map[string]bool)
	radici := make([]*nodo, 0, len(oggetti))

	for nome, ogg := range oggetti {
		var n *nodo
		if ogg.tipo == "op" {
			n, notRoot = costruisciSottoalbero(oggetti, notRoot, nome)
		} else {
			n = &nodo{nome, nil, nil}
		}
		radici = append(radici, n)
	}

	r := make([]*nodo, 0)
	// tenere come radici solo quelli che non sono stati usati come nodi interni
	for _, rad := range radici {
		if !notRoot[rad.nome] {
			r = append(r, rad)
		}
	}

	return foresta{oggetti, r}
}

func costruisciSottoalbero(oggetti map[string]*oggetto, notRoot map[string]bool, radice string) (*nodo, map[string]bool) {

	var sx, dx *nodo

	if oggetti[radice].tipo == "op" {
		notRoot[oggetti[radice].sx] = true
		sx, notRoot = costruisciSottoalbero(oggetti, notRoot, oggetti[radice].sx)
		notRoot[oggetti[radice].dx] = true
		dx, notRoot = costruisciSottoalbero(oggetti, notRoot, oggetti[radice].dx)
	} else {
		sx = nil
		dx = nil
	}

	return &nodo{radice, sx, dx}, notRoot
}

// stampa ogni albero della foresta
func (f foresta) stampaForesta() {
	for _, v := range f.radici {
		f.stampaAlbero(v)
		fmt.Println()
	}
}

// stampa l'albero che ha come root "root"
func (f foresta) stampaAlbero(root *nodo) { // dfs
	if root.figliosx != nil {
		f.stampaAlbero(root.figliosx)
	}

	if root.figliosx == nil && root.figliodx == nil {
		fmt.Print(root.nome, " (val = ", f.payloads[root.nome].val, ")\n")
	} else {
		fmt.Println(root.nome)
	}

	if root.figliodx != nil {
		f.stampaAlbero(root.figliodx)
	}
}

// stampa l'albero del nodo di nome "nome"
func stampaAlbero(f foresta, nome string) {
	root, _ := f.nodoByStringForesta(nome)
	f.stampaAlbero(root)
}

// restituisce il puntatore al nodo di nome "nome" cercandolo in una foresta
// e restituisce il padre al nodo restituito
func (f foresta) nodoByStringForesta(nome string) (*nodo, *nodo) {
	for _, v := range f.radici {
		if v.nome == nome {
			return v, nil
		}

		nodo, padre := nodoByStringAlbero(v, nome)
		if nodo != nil {
			return nodo, padre
		}
	}
	return nil, nil
}

// restituisce il puntatore al nodo di nome "nome" cercandolo in un albero
// e restituisce il padre al nodo restituito
func nodoByStringAlbero(root *nodo, nome string) (*nodo, *nodo) { // dfs
	if root.figliosx != nil {
		if root.figliosx.nome == nome {
			return root.figliosx, root
		}
		return nodoByStringAlbero(root.figliosx, nome)
	}

	if root.figliodx != nil {
		if root.figliodx.nome == nome {
			return root.figliodx, root
		}
		return nodoByStringAlbero(root.figliodx, nome)
	}

	return nil, nil
}

func up(f foresta, n string) (string, bool) {
	_, padre := f.nodoByStringForesta(n)
	if padre == nil {
		return "", false
	}
	return padre.nome, true
}

func sx(f foresta, n string) (string, bool) {
	nodo, _ := f.nodoByStringForesta(n)
	if nodo == nil || nodo.figliosx == nil {
		return "", false
	}
	return nodo.figliosx.nome, true
}

func dx(f foresta, n string) (string, bool) {
	nodo, _ := f.nodoByStringForesta(n)
	if nodo == nil || nodo.figliodx == nil {
		return "", false
	}
	return nodo.figliodx.nome, true
}

func calcolaPrezzo(f foresta, s string) int {
	ogg := f.payloads[s]

	if ogg.tipo == "val" {
		return ogg.val
	}

	switch ogg.op {
	case '+':
		return calcolaPrezzo(f, ogg.sx) + calcolaPrezzo(f, ogg.dx)
	case '-':
		return calcolaPrezzo(f, ogg.sx) - calcolaPrezzo(f, ogg.dx)
	case '*':
		return calcolaPrezzo(f, ogg.sx) * calcolaPrezzo(f, ogg.dx)
	case '/':
		return calcolaPrezzo(f, ogg.sx) / calcolaPrezzo(f, ogg.dx)
	}
	return 0
}

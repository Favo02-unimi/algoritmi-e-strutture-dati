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
	padri map[string]string
	// lista dei padri (la struct oggetto contiene i figli, quindi è possibile muoversi sia salendo grazie alla mappa dei padri che scendendo grazie all'oggetto in sè salvato nella mappa payloads)
	// se il padre è vuoto, allora vuol dire che è la root dell'albero
}

func costruisciForesta(oggetti map[string]*oggetto) foresta {
	padri := make(map[string]string)

	// inizializzare tutti gli oggetti nella mappa
	for k := range oggetti {
		padri[k] = ""
	}

	// impostare i padri, non era possibile farlo durante l'inizializzazione perchè non è detto che quando viene incontrato il padre esistano già i figli, invece adesso esistono già tutti
	for k, v := range oggetti {
		if (v.tipo == "op") {
			padri[v.sx] = k
			padri[v.dx] = k
		}
	}

	return foresta{oggetti, padri}
}

func up(f foresta, n string) (string, bool) {
	padre := f.padri[n]
	if padre == "" {
		return "", false
	}
	return padre, true
}

func sx(f foresta, n string) (string, bool) {
	sx := f.payloads[n].sx
	if sx == "" {
		return "", false
	}
	return sx, true
}

func dx(f foresta, n string) (string, bool) {
	dx := f.payloads[n].dx
	if dx == "" {
		return "", false
	}
	return dx, true
}

func stampaAlbero(f foresta, n string) { // dfs
	sxx, founddx := sx(f, n)
	if founddx {
		stampaAlbero(f, sxx)
	}

	if f.payloads[n].tipo == "val" {
		fmt.Print(n, " (val = ", f.payloads[n].val, ")\n")
	} else {
		fmt.Println(n)
	}

	dxx, foundsx := dx(f, n)
	if foundsx {
		stampaAlbero(f, dxx)
	}
}

func calcolaPrezzo(f foresta, n string) int {
	if (f.payloads[n].tipo == "val") {
		return f.payloads[n].val
	}

	sx := calcolaPrezzo(f, f.payloads[n].sx)
	dx := calcolaPrezzo(f, f.payloads[n].sx)
	var res int

	switch f.payloads[n].op {
	case '+':
		res = sx + dx
	case '-':
		res = sx - dx
	case '*':
		res = sx * dx
	case '/':
		res = sx / dx
	}

	return res
}
 
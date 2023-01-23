package main

import . "fmt"

func main() {
	res, num := 0, 0
	for num != -1 {
		num = 0
		Scan(&num)
		if num > 100 && res == 0 {
			res = num
		}
	}
	if res != 0 {
		Print(res)
	} else {
		Print("nessun numero maggiore di 100")
	}
}

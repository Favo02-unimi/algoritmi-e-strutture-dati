package main

import . "fmt"

func main() {

	num := -1
	var prec int

	for true {

		prec = num
		Scan(&num)

		if num == 0 {
			return
		}

		if num > prec {
			Print("+")
		} else if num < prec {
			Print("-")
		}

	}

}

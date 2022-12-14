package main

import (
	. "fmt"
	"strconv"
	"strings"
)

func main() {
	var numStr string
	Scan(&numStr)
	numsStr := strings.Split(numStr, ",")

	// converto lo slice di stringhe a slice di interi
	var nums = []int{}
	for _, i := range numsStr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		nums = append(nums, j)
	}

	daTrovare := 1
	ritorniASinistra := 0
	for i := 0; daTrovare <= len(nums); i++ {

		if(ritorniASinistra > len(nums)) {
			break
		}
		if (daTrovare == nums[i]) {
			daTrovare++
		}
		if (i == len(nums)-1) {
			i = -1
			ritorniASinistra++
			Println("-------ritorno-------")
		}

		Println("datrovare: ", daTrovare)

	}

	Println(ritorniASinistra)

}

package main

import "fmt"

type structure struct {
	list []int
}

/*
	create map of a complex structure:
		map[structure]bool 		is not allowed
		map[*structure]bool 	is allowed but not working as intended
*/

func main() {

	a := structure{[]int{0, 1, 2}}

	map_ := make(map[*structure]bool) // map on the pointer
	map_[&a] = true

	fmt.Println(map_[&a]) // true

	a2 := structure{[]int{0, 1, 2}} // generate the same structure

	fmt.Println(map_[&a2]) // false
	// the map is on the pointer, the new structure points to a new structure

	// so a solution is to serialize (with a toString method) and map on a string
	map2 := make(map[string]bool)
	map2[a.toString()] = true

	fmt.Println(map2[a.toString()])  // true
	fmt.Println(map2[a2.toString()]) // true

}

func (s structure) toString() string {
	if len(s.list) == 0 {
		return ""
	}
	str := fmt.Sprint(s.list[0])
	for i := 1; i < len(s.list); i++ {
		str += fmt.Sprint(s.list[i])
	}
	return str
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	h, s, tunnels := parseInput()
	fmt.Println(h, s, tunnels)
	fmt.Println(walk(h, s, tunnels))
}

type tunnel struct {
	a, b int
	lum  int
}

func parseInput() (int, int, []tunnel) {
	// first line: "instructions"
	var temp, tunnelsN, h, s int
	fmt.Scan(&temp)
	fmt.Scan(&tunnelsN)
	fmt.Scan(&h)
	fmt.Scan(&s)

	// next lines: tunnels
	scanner := bufio.NewScanner(os.Stdin)
	var tunnels []tunnel

	for scanner.Scan() {
		if len(tunnels) == tunnelsN {
			break
		}
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		a, _ := strconv.Atoi(tokens[0])
		b, _ := strconv.Atoi(tokens[1])
		lum, _ := strconv.Atoi(tokens[2])
		tunnels = append(tunnels, tunnel{a, b, lum})
	}
	return h, s, tunnels
}

func walk(source, destination int, tunnels []tunnel) int {
	crossed := make(map[tunnel]bool) // tunnels already crossed

	for true {
		// get tunnels possible to take from source
		possible := findTunnels(source, tunnels)

		// select tunnel with the least luminosity
		var taken tunnel = possible[0]
		for _, t := range possible {
			if t.lum < taken.lum {
				taken = t
			}
		}
		fmt.Println("tunnel taken:", taken)

		// set new source point to end of tunnel
		if taken.a == source {
			source = taken.b
		} else {
			source = taken.a
		}

		// check if destination reached
		if taken.a == destination || taken.b == destination {
			crossed[taken] = true
			return len(crossed)
		}

		// check if in a loop
		if crossed[taken] {
			return -1
		}

		// set current tunnel crossed
		crossed[taken] = true
	}
	return -1
}

// returns all tunnels starting or arriving to "start"
func findTunnels(start int, tunnels []tunnel) (res []tunnel) {
	for _, t := range tunnels {
		if t.a == start || t.b == start {
			res = append(res, t)
		}
	}
	return res
}

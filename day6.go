package main

import (
	"fmt"
	"strings"
)

var orbitMap map[string]Orbit = make(map[string]Orbit)

func main() {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}

	for i := 0; i < len(input); i++ {
		orbits := strings.Split(input[i], ")")
		if orbit, ok := orbitMap[orbits[0]]; ok {
			orbit.orbitedBy = append(orbit.orbitedBy, orbits[1])
			orbitMap[orbits[0]] = orbit
		} else {
			orbitMap[orbits[0]] = Orbit{"", []string{orbits[1]}}
		}

		if orbit, ok := orbitMap[orbits[1]]; ok {
			orbit.orbits = orbits[0]
			orbitMap[orbits[1]] = orbit
		} else {
			orbitMap[orbits[1]] = Orbit{orbits[0], []string{}}
		}
	}
	fmt.Println("Shortest path: ", getShortestPathToSanta())
}

func getShortestPathToSanta() int {
	distance := map[string]int{"SAN": 0}
	queue := []string{"SAN"}
	visited := map[string]bool{"SAN": true}
	for len(queue) > 0 {
		orbit := orbitMap[queue[0]]
		neighbors := orbit.orbitedBy
		neighbors = append(neighbors, orbit.orbits)
		for i := 0; i < len(neighbors); i++ {
			if !visited[neighbors[i]] {
				queue = append(queue, neighbors[i])

				distance[neighbors[i]] = distance[queue[0]] + 1
				visited[neighbors[i]] = true
			}
		}

		queue = queue[1:]
	}
	// account for map adding 1 for transitions away from SAN node and onto YOU node
	// these are superfluous since we only need to reach the same orbital
	return distance["YOU"] - 2
}

type Orbit struct {
	orbits    string
	orbitedBy []string
}

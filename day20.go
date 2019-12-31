package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := getInput()
	paths := buildMaze(input)
	shortestPath := getShortestPath(paths)
	fmt.Println("Shortest path: ", shortestPath)
}

func getShortestPath(paths map[string][]Neighbor) int {
	visited := map[string]bool{"AA": true}
	neighbors := map[string]int{}

	for _, neighbor := range paths["AA"] {
		neighbors[neighbor.label] = neighbor.distance
		visited[neighbor.label] = true
	}

	isFinished := false
	for !isFinished {
		isFinished = true

		for label, dist := range neighbors {
			if label != "ZZ" {
				extendedNeighbors := paths[label]
				for _, extendedNeighbor := range extendedNeighbors {
					newDist := dist + 1 + extendedNeighbor.distance
					if currentDist, ok := neighbors[extendedNeighbor.label]; !ok || (ok && newDist < currentDist) {
						neighbors[extendedNeighbor.label] = newDist
						isFinished = false
					}
				}
			}
		}
	}
	return neighbors["ZZ"]
}

func buildMaze(input []string) map[string][]Neighbor {
	paths := make(map[string][]Neighbor)
	// pad two lines on all sides to skip jump-point labels
	for row := 2; row < len(input)-2; row++ {
		for col := 2; col < len(input[0])-2; col++ {
			if label, ok := isJumppoint(input, &Coordinates{row, col}); ok {
				neighbors := getPaths(input, label, &Coordinates{row, col})
				for neighborLabel, neighborDistance := range neighbors {
					paths[label] = append(paths[label], Neighbor{neighborLabel, neighborDistance})
				}
			}
		}
	}
	return paths
}

func isJumppoint(input []string, coords *Coordinates) (string, bool) {
	if unicode.IsLetter(rune(input[coords.row+1][coords.col])) {
		return string(input[coords.row+1][coords.col]) + string(input[coords.row+2][coords.col]), true
	}
	if unicode.IsLetter(rune(input[coords.row-1][coords.col])) {
		return (string(input[coords.row-2][coords.col]) + string(input[coords.row-1][coords.col])), true
	}
	if unicode.IsLetter(rune(input[coords.row][coords.col-1])) {
		return (string(input[coords.row][coords.col-2]) + string(input[coords.row][coords.col-1])), true
	}
	if unicode.IsLetter(rune(input[coords.row][coords.col+1])) {
		return (string(input[coords.row][coords.col+1]) + string(input[coords.row][coords.col+2])), true
	}

	return "", false
}

func getPaths(input []string, label string, coords *Coordinates) map[string]int {
	neighbors := make(map[string]int)
	visited := map[Coordinates]bool{*coords: true}
	queue := []*Coordinates{coords}
	distance := 0

	for len(queue) > 0 {
		distance++
		newQueue := []*Coordinates{}
		for _, queueCoords := range queue {
			for _, neighbor := range []*Coordinates{
				&Coordinates{queueCoords.row, queueCoords.col + 1},
				&Coordinates{queueCoords.row, queueCoords.col - 1},
				&Coordinates{queueCoords.row + 1, queueCoords.col},
				&Coordinates{queueCoords.row - 1, queueCoords.col},
			} {
				if visited[*neighbor] {
					continue
				}
				if input[neighbor.row][neighbor.col] == '.' {
					if jumpLabel, ok := isJumppoint(input, neighbor); ok && jumpLabel != label {
						neighbors[jumpLabel] = distance
					} else {
						newQueue = append(newQueue, neighbor)
					}
				}
				visited[*neighbor] = true
			}
		}
		queue = newQueue
	}
	return neighbors
}
func getInput() []string {
	input := []string{}
	file, _ := os.Open("day20input.txt")
	reader := bufio.NewReader(file)
	for true {
		str, err := reader.ReadString('\n')
		input = append(input, string(str)[:len(str)-1])
		if err != nil {
			break
		}
	}
	return input
}

type Coordinates struct {
	row int
	col int
}

type Neighbor struct {
	label    string
	distance int
}

package main

import (
	"fmt"
	"math"
)

func main() {
	input := []string{
		".##.#.#....#.#.#..##..#.#.",
		"#.##.#..#.####.##....##.#.",
		"###.##.##.#.#...#..###....",
		"####.##..###.#.#...####..#",
		"..#####..#.#.#..#######..#",
		".###..##..###.####.#######",
		".##..##.###..##.##.....###",
		"#..#..###..##.#...#..####.",
		"....#.#...##.##....#.#..##",
		"..#.#.###.####..##.###.#.#",
		".#..##.#####.##.####..#.#.",
		"#..##.#.#.###.#..##.##....",
		"#.#.##.#.##.##......###.#.",
		"#####...###.####..#.##....",
		".#####.#.#..#.##.#.#...###",
		".#..#.##.#.#.##.#....###.#",
		".......###.#....##.....###",
		"#..#####.#..#..##..##.#.##",
		"##.#.###..######.###..#..#",
		"#.#....####.##.###....####",
		"..#.#.#.########.....#.#.#",
		".##.#.#..#...###.####..##.",
		"##...###....#.##.##..#....",
		"..##.##.##.#######..#...#.",
		".###..#.#..#...###..###.#.",
		"#..#..#######..#.#..#..#.#",
	}

	mostAsteroidsDetected := -1
	for rowIndex, row := range input {
		for colIndex := 0; colIndex < len(row); colIndex++ {
			if string(row[colIndex]) == "#" {
				asteroidsDetected := getAsteroidsDetected(input, rowIndex, colIndex)
				if asteroidsDetected > mostAsteroidsDetected {
					mostAsteroidsDetected = asteroidsDetected
				}
			}
		}
	}

	fmt.Println("Detected: ", mostAsteroidsDetected)
}

func getAsteroidsDetected(input []string, row int, col int) int {
	visited := make(map[Coordinates]bool)
	asteroidsDetected := []Coordinates{}

	startingPoint := Coordinates{row, col}
	queue := []Coordinates{startingPoint}
	visited[startingPoint] = true
	for len(queue) > 0 {
		coords := queue[0]
		queue = queue[1:]
		for rowOffset := -1; rowOffset < 2; rowOffset++ {
			for colOffset := -1; colOffset < 2; colOffset++ {
				if (rowOffset == 0 && colOffset == 0) ||
					coords.col+colOffset >= len(input[0]) ||
					coords.col+colOffset < 0 ||
					coords.row+rowOffset >= len(input) ||
					coords.row+rowOffset < 0 {
					continue
				}
				neighbor := Coordinates{(coords.row + rowOffset), (coords.col + colOffset)}
				if !visited[neighbor] {
					if string(input[neighbor.row][neighbor.col]) == "#" {
						if isAsteroidVisible(asteroidsDetected, neighbor, startingPoint) {
							asteroidsDetected = append(asteroidsDetected, neighbor)
						}
					}
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
		}
	}
	return len(asteroidsDetected)
}

func isAsteroidVisible(asteroidsDetected []Coordinates, neighbor Coordinates, loc Coordinates) bool {
	for _, asteroid := range asteroidsDetected {
		if asteroid.col == neighbor.col && neighbor.col != loc.col {
			continue
		} else if asteroid.row == neighbor.row && neighbor.row != loc.row {
			continue
		} else if asteroid.row == neighbor.row && neighbor.row == loc.row {
			if (asteroid.col < loc.col && neighbor.col < asteroid.col) ||
				(asteroid.col > loc.col && neighbor.col > asteroid.col) {
				return false
			}
		} else if asteroid.col == neighbor.col && neighbor.col == loc.col {
			if (asteroid.row < loc.row && neighbor.row < asteroid.row) ||
				(asteroid.row > loc.row && neighbor.row > asteroid.row) {
				return false
			}
		} else if ((asteroid.row > loc.row && neighbor.row > loc.row) ||
			(asteroid.row < loc.row && neighbor.row < loc.row)) &&
			((asteroid.col > loc.col && neighbor.col > loc.col) ||
				(asteroid.col < loc.col && neighbor.col < loc.col)) {

			neighborRowDiff := math.Abs(float64(neighbor.row) - float64(loc.row))
			asteroidRowDiff := math.Abs(float64(asteroid.row) - float64(loc.row))

			neighborColDiff := math.Abs(float64(neighbor.col) - float64(loc.col))
			asteroidColDiff := math.Abs(float64(asteroid.col) - float64(loc.col))

			if (neighborRowDiff / asteroidRowDiff) == (neighborColDiff / asteroidColDiff) {
				return false
			}
		}
	}
	return true
}

type Coordinates struct {
	row int
	col int
}

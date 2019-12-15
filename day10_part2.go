package main

import (
	"fmt"
	"math"
	"sort"
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
	monitoringStation := Coordinates{0, 0}
	for rowIndex, row := range input {
		for colIndex := 0; colIndex < len(row); colIndex++ {
			if string(row[colIndex]) == "#" {
				asteroidsDetected := getAsteroidsDetected(input, rowIndex, colIndex)
				if asteroidsDetected > mostAsteroidsDetected {
					mostAsteroidsDetected = asteroidsDetected
					monitoringStation = Coordinates{rowIndex, colIndex}
				}
			}
		}
	}

	fmt.Println("Monitoring station coords ", monitoringStation.col, monitoringStation.row)
	asteroidAngles := *getAsteroidAngles(input, &monitoringStation)
	vaporize(asteroidAngles)
}

func vaporize(asteroids []AsteroidDetails) {
	angleMap := map[float64][]AsteroidDetails{}
	angles := []float64{}
	for _, asteroid := range asteroids {
		if asteroids, ok := angleMap[asteroid.angle]; ok {
			angleMap[asteroid.angle] = append(asteroids, asteroid)
		} else {
			angleMap[asteroid.angle] = []AsteroidDetails{asteroid}
			angles = append(angles, asteroid.angle)
		}
	}
	sort.Float64s(angles)

	// begin vaporizing
	vaporizationCount := 1
	index := 0
	for len(angleMap) > 0 {
		angle := angles[index]
		if asteroids, ok := angleMap[angle]; ok {
			var minDistance float64 = -1
			minDistanceIndex := -1
			for i := 0; i < len(asteroids); i++ {
				if asteroids[i].distance < minDistance || minDistance == -1 {
					minDistanceIndex = i
				}
			}
			fmt.Println("Vaporizing asteroid #", vaporizationCount, " at ", asteroids[minDistanceIndex].coords.col, asteroids[minDistanceIndex].coords.row)
			vaporizationCount++
			asteroids = append(asteroids[:minDistanceIndex], asteroids[minDistanceIndex+1:]...)
			if len(asteroids) == 0 {
				delete(angleMap, angle)
			} else {
				angleMap[angle] = asteroids
			}
		}
		index = (index + 1) % len(angles)
	}
}

func getAsteroidAngles(asteroidMap []string, loc *Coordinates) *[]AsteroidDetails {
	asteroidAngles := []AsteroidDetails{}
	for rowIndex := range asteroidMap {
		for colIndex := 0; colIndex < len(asteroidMap[0]); colIndex++ {
			asteroid := Coordinates{rowIndex, colIndex}
			if asteroid.row == loc.row && asteroid.col == loc.col ||
				string(asteroidMap[asteroid.row][asteroid.col]) != "#" {
				continue
			}
			rowDiff := math.Abs(float64(asteroid.row) - float64(loc.row))
			colDiff := math.Abs(float64(asteroid.col) - float64(loc.col))

			distance := math.Sqrt((rowDiff * rowDiff) + (colDiff * colDiff))
			angle := math.Atan2(rowDiff, colDiff) * (180 / math.Pi)

			quadrant := getQuadrant(&asteroid, loc)
			switch quadrant {
			case 1:
				angle = 90 - angle
			case 2:
				angle += 90
			case 3:
				angle = 180 + (90 - angle)
			case 4:
				angle += 270
			default:
				break
			}
			asteroidAngles = append(asteroidAngles, AsteroidDetails{distance, angle, asteroid})
		}
	}
	return &asteroidAngles
}

func getQuadrant(obj *Coordinates, center *Coordinates) int {
	if center.row == obj.row {
		if center.col > obj.col {
			return 3
		}
		return 1
	}
	if center.col == obj.col {
		if center.row > obj.row {
			return 1
		}
		return 3
	}
	if center.row > obj.row {
		if center.col > obj.col {
			return 4
		}
		return 1
	} else {
		if center.col > obj.col {
			return 3
		}
		return 2
	}
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

type AsteroidDetails struct {
	distance float64
	angle    float64
	coords   Coordinates
}

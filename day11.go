package main

import "fmt"

func main() {
	input := []int{3, 8, 1005, 8, 329, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1002, 8, 1, 29, 2, 1102, 1, 10, 1, 1009, 16, 10, 2, 4, 4, 10, 1, 9, 5, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 66, 2, 106, 7, 10, 1006, 0, 49, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1002, 8, 1, 95, 1006, 0, 93, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 120, 1006, 0, 61, 2, 1108, 19, 10, 2, 1003, 2, 10, 1006, 0, 99, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 157, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 179, 2, 1108, 11, 10, 1, 1102, 19, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 209, 2, 108, 20, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 234, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 256, 2, 1102, 1, 10, 1006, 0, 69, 2, 108, 6, 10, 2, 4, 13, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1002, 8, 1, 294, 1, 1107, 9, 10, 1006, 0, 87, 2, 1006, 8, 10, 2, 1001, 16, 10, 101, 1, 9, 9, 1007, 9, 997, 10, 1005, 10, 15, 99, 109, 651, 104, 0, 104, 1, 21101, 387395195796, 0, 1, 21101, 346, 0, 0, 1105, 1, 450, 21101, 0, 48210129704, 1, 21101, 0, 357, 0, 1105, 1, 450, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21101, 0, 46413147328, 1, 21102, 404, 1, 0, 1106, 0, 450, 21102, 179355823323, 1, 1, 21101, 415, 0, 0, 1105, 1, 450, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21102, 1, 838345843476, 1, 21101, 0, 438, 0, 1105, 1, 450, 21101, 709475709716, 0, 1, 21101, 449, 0, 0, 1105, 1, 450, 99, 109, 2, 22102, 1, -1, 1, 21102, 40, 1, 2, 21101, 0, 481, 3, 21101, 0, 471, 0, 1105, 1, 514, 109, -2, 2105, 1, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 476, 477, 492, 4, 0, 1001, 476, 1, 476, 108, 4, 476, 10, 1006, 10, 508, 1101, 0, 0, 476, 109, -2, 2106, 0, 0, 0, 109, 4, 2101, 0, -1, 513, 1207, -3, 0, 10, 1006, 10, 531, 21101, 0, 0, -3, 21201, -3, 0, 1, 21201, -2, 0, 2, 21101, 1, 0, 3, 21101, 550, 0, 0, 1105, 1, 555, 109, -4, 2106, 0, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 578, 2207, -4, -2, 10, 1006, 10, 578, 21201, -4, 0, -4, 1105, 1, 646, 22101, 0, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 597, 0, 0, 1105, 1, 555, 22102, 1, 1, -4, 21101, 0, 1, -1, 2207, -4, -2, 10, 1006, 10, 616, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 638, 22102, 1, -1, 1, 21101, 638, 0, 0, 106, 0, 513, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2106, 0, 0}
	robot := robot()
	robot.runIntcode(input)
	fmt.Println(len(robot.paintMap))
}

func (r Robot) runIntcode(inputArr []int) {
	input := *getInputMap(inputArr)
	i := 0
	relativeOffset := 0
	colorToPaint := -1
	for true {
		instruction := input[i]

		opcode := instruction % 100
		instruction /= 100

		if opcode == 99 {
			return
		} else if opcode == 3 {
			targetIndex := input[i+1]
			if instruction%10 == 2 {
				targetIndex += relativeOffset
			}

			input[targetIndex] = r.input

			i += 2
		} else if opcode == 4 {
			index := input[i+1]
			val := index
			if instruction%10 == 0 {
				val = input[index]
			} else if instruction%10 == 2 {
				val = input[index+relativeOffset]
			}
			// fmt.Print(val)
			if colorToPaint != -1 {
				r.paint(colorToPaint)
				r.turn(val)
				colorToPaint = -1
			} else {
				colorToPaint = val
			}

			i += 2
		} else if opcode == 9 {
			index := input[i+1]
			mode := instruction % 10
			val := index
			if mode == 0 {
				val = input[index]
			} else if mode == 2 {
				val = input[val+relativeOffset]
			}
			relativeOffset += val
			i += 2
		} else {
			param1 := input[i+1]
			mode1 := instruction % 10

			instruction /= 10

			param2 := input[i+2]
			mode2 := instruction % 10

			instruction /= 10

			targetIndex := input[i+3]
			mode3 := instruction % 10

			if mode1 == 0 {
				param1 = input[param1]
			} else if mode1 == 2 {
				param1 = input[param1+relativeOffset]
			}

			if mode2 == 0 {
				param2 = input[param2]
			} else if mode2 == 2 {
				param2 = input[param2+relativeOffset]
			}

			if mode3 == 2 {
				targetIndex += relativeOffset
			}

			if opcode == 1 {
				input[targetIndex] = param1 + param2
				i += 4
			} else if opcode == 2 {
				input[targetIndex] = param1 * param2
				i += 4
			} else if opcode == 5 {
				if param1 != 0 {
					i = param2
				} else {
					i += 3
				}
			} else if opcode == 6 {
				if param1 == 0 {
					i = param2
				} else {
					i += 3
				}
			} else if opcode == 7 {
				if param1 < param2 {
					input[targetIndex] = 1
				} else {
					input[targetIndex] = 0
				}
				i += 4
			} else if opcode == 8 {
				if param1 == param2 {
					input[targetIndex] = 1
				} else {
					input[targetIndex] = 0
				}
				i += 4
			}
		}
	}
}

func getInputMap(arr []int) *map[int]int {
	m := make(map[int]int)
	for i, val := range arr {
		m[i] = val
	}
	return &m
}

func (r *Robot) paint(color int) {
	r.paintMap[r.currentLocation] = color
	r.coordinatesSeen[r.currentLocation] = true
}

func (r *Robot) turn(direction int) {
	switch r.currentDirection {
	case "N":
		if direction == 0 {
			r.currentLocation.col = r.currentLocation.col - 1
			r.currentDirection = "W"
		} else {
			r.currentLocation.col = r.currentLocation.col + 1
			r.currentDirection = "E"
		}
	case "W":
		if direction == 0 {
			r.currentLocation.row = r.currentLocation.row - 1
			r.currentDirection = "S"
		} else {
			r.currentLocation.row = r.currentLocation.row + 1
			r.currentDirection = "N"
		}
	case "S":
		if direction == 0 {
			r.currentLocation.col = r.currentLocation.col + 1
			r.currentDirection = "E"
		} else {
			r.currentLocation.col = r.currentLocation.col - 1
			r.currentDirection = "W"
		}
	case "E":
		if direction == 0 {
			r.currentLocation.row = r.currentLocation.row + 1
			r.currentDirection = "N"
		} else {
			r.currentLocation.row = r.currentLocation.row - 1
			r.currentDirection = "S"
		}
	}
	r.input = r.paintMap[r.currentLocation]
}

func robot() *Robot {
	return &Robot{
		Coordinates{0, 0},
		"N",
		map[Coordinates]int{},
		map[Coordinates]bool{},
		0,
	}
}

type Robot struct {
	currentLocation  Coordinates
	currentDirection string
	paintMap         map[Coordinates]int
	coordinatesSeen  map[Coordinates]bool
	input            int
}

type Coordinates struct {
	row int
	col int
}

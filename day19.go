package main

import "fmt"

func main() {
	input := []int{109, 424, 203, 1, 21101, 11, 0, 0, 1105, 1, 282, 21102, 1, 18, 0, 1105, 1, 259, 1201, 1, 0, 221, 203, 1, 21102, 31, 1, 0, 1106, 0, 282, 21101, 38, 0, 0, 1106, 0, 259, 21001, 23, 0, 2, 22102, 1, 1, 3, 21101, 0, 1, 1, 21102, 57, 1, 0, 1106, 0, 303, 1202, 1, 1, 222, 20102, 1, 221, 3, 21002, 221, 1, 2, 21101, 259, 0, 1, 21102, 80, 1, 0, 1106, 0, 225, 21101, 0, 51, 2, 21101, 0, 91, 0, 1106, 0, 303, 1202, 1, 1, 223, 20101, 0, 222, 4, 21101, 259, 0, 3, 21102, 225, 1, 2, 21101, 225, 0, 1, 21101, 118, 0, 0, 1105, 1, 225, 20102, 1, 222, 3, 21102, 1, 152, 2, 21102, 133, 1, 0, 1105, 1, 303, 21202, 1, -1, 1, 22001, 223, 1, 1, 21102, 1, 148, 0, 1105, 1, 259, 1202, 1, 1, 223, 20101, 0, 221, 4, 21002, 222, 1, 3, 21102, 1, 17, 2, 1001, 132, -2, 224, 1002, 224, 2, 224, 1001, 224, 3, 224, 1002, 132, -1, 132, 1, 224, 132, 224, 21001, 224, 1, 1, 21101, 195, 0, 0, 105, 1, 108, 20207, 1, 223, 2, 21002, 23, 1, 1, 21102, 1, -1, 3, 21102, 214, 1, 0, 1105, 1, 303, 22101, 1, 1, 1, 204, 1, 99, 0, 0, 0, 0, 109, 5, 1202, -4, 1, 249, 22101, 0, -3, 1, 21202, -2, 1, 2, 22102, 1, -1, 3, 21101, 250, 0, 0, 1106, 0, 225, 22101, 0, 1, -4, 109, -5, 2105, 1, 0, 109, 3, 22107, 0, -2, -1, 21202, -1, 2, -1, 21201, -1, -1, -1, 22202, -1, -2, -2, 109, -3, 2106, 0, 0, 109, 3, 21207, -2, 0, -1, 1206, -1, 294, 104, 0, 99, 22101, 0, -2, -2, 109, -3, 2105, 1, 0, 109, 5, 22207, -3, -4, -1, 1206, -1, 346, 22201, -4, -3, -4, 21202, -3, -1, -1, 22201, -4, -1, 2, 21202, 2, -1, -1, 22201, -4, -1, 1, 21201, -2, 0, 3, 21102, 1, 343, 0, 1105, 1, 303, 1106, 0, 415, 22207, -2, -3, -1, 1206, -1, 387, 22201, -3, -2, -3, 21202, -2, -1, -1, 22201, -3, -1, 3, 21202, 3, -1, -1, 22201, -3, -1, 2, 21202, -4, 1, 1, 21102, 1, 384, 0, 1105, 1, 303, 1105, 1, 415, 21202, -4, -1, -4, 22201, -4, -3, -4, 22202, -3, -2, -2, 22202, -2, -4, -4, 22202, -3, -2, -3, 21202, -4, -1, -2, 22201, -3, -2, 1, 22102, 1, 1, -4, 109, -5, 2105, 1, 0}

	detected := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if intcode(getInputMap(input), x, y) {
				detected++
			}
		}
	}
	fmt.Println("Part 1 - ", detected)

	for y := 0; y < 10000; y++ {
		for x := 0; x < 10000; x++ {
			if intcode(getInputMap(input), x, y) {
				if intcode(getInputMap(input), x+99, y) {
					if intcode(getInputMap(input), x+99, y+99) &&
						intcode(getInputMap(input), x, y+99) {
						fmt.Println("Part 2 - Detected at (%d, %d)\n", x, y)
						return
					}
				} else {
					break
				}
			}
		}
	}
}

func intcode(input map[int]int, x int, y int) bool {
	i := 0
	relativeOffset := 0
	usedX := false
	for true {
		instruction := input[i]

		opcode := instruction % 100
		instruction /= 100

		if opcode == 99 {
			return false
		} else if opcode == 3 {
			targetIndex := input[i+1]
			if instruction%10 == 2 {
				targetIndex += relativeOffset
			}

			val := x
			if usedX {
				val = y
			}
			usedX = true

			input[targetIndex] = val

			i += 2
		} else if opcode == 4 {
			index := input[i+1]
			val := index
			if instruction%10 == 0 {
				val = input[index]
			} else if instruction%10 == 2 {
				val = input[index+relativeOffset]
			}

			if val == 1 {
				return true
			}
			return false
			// switch val {
			// case 35:
			// 	result[row] = append(result[row], "#")
			// case 46:
			// 	result[row] = append(result[row], ".")
			// case 10:
			// 	row--
			// default:
			// 	result[row] = append(result[row], "x")
			// }

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
	return false
}

func getInputMap(arr []int) map[int]int {
	m := make(map[int]int)
	for i, val := range arr {
		m[i] = val
	}
	return m
}

type Coordinates struct {
	row int
	col int
}

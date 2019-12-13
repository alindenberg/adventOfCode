package main

import (
	"fmt"
	"strconv"
)

func main() {
	signalPerms := []string{"98765"}
	Perm([]rune("56789"), func(a []rune) {
		signalPerms = append(signalPerms, string(a))
	})
	maxVal := -1
	for i := 0; i < len(signalPerms); i++ {
		val := runAmplifiers(signalPerms[i])
		if val > maxVal {
			maxVal = val
		}
	}
	fmt.Println("Final value ", maxVal)
}

func runAmplifiers(signalInput string) int {
	phaseInts := []int{}
	// parse signal input into phases
	for i := 0; i < len(signalInput); i++ {
		phase, _ := strconv.Atoi(string([]rune(signalInput)[i]))
		phaseInts = append(phaseInts, phase)
	}
	// run amplifiers
	input := [515]int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 46, 55, 72, 85, 110, 191, 272, 353, 434, 99999, 3, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 101, 2, 9, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 101, 2, 9, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 101, 5, 9, 9, 1002, 9, 3, 9, 101, 3, 9, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99}
	amp1 := amplifier(input, phaseInts[0])
	amp2 := amplifier(input, phaseInts[1])
	amp3 := amplifier(input, phaseInts[2])
	amp4 := amplifier(input, phaseInts[3])
	amp5 := amplifier(input, phaseInts[4])
	amp1.run(0)
	amp2.run(amp1.getOutput())
	amp3.run(amp2.getOutput())
	amp4.run(amp3.getOutput())
	amp5.run(amp4.getOutput())
	for !amp5.isHalted {
		amp1.run(amp5.getOutput())
		amp2.run(amp1.getOutput())
		amp3.run(amp2.getOutput())
		amp4.run(amp3.getOutput())
		amp5.run(amp4.getOutput())
	}
	return amp5.getOutput()
}

type Amplifier struct {
	index       int
	phase       int
	inputArr    [515]int
	outputQueue []int
	usedPhase   bool
	isHalted    bool
}

func amplifier(inputArr [515]int, phase int) *Amplifier {
	return &Amplifier{
		0,
		phase,
		inputArr,
		[]int{},
		false,
		false}
}

func (amp *Amplifier) getOutput() int {
	return amp.outputQueue[len(amp.outputQueue)-1]
}

func (amp *Amplifier) run(inputVal int) {
	for !amp.isHalted {
		instruction := amp.inputArr[amp.index]

		opcode := instruction % 100
		instruction /= 100

		switch opcode {
		case 99:
			amp.isHalted = true
		case 1:
			index1 := amp.inputArr[amp.index+1]
			mode1 := instruction % 10
			val1 := index1
			if mode1 == 0 {
				val1 = amp.inputArr[index1]
			}
			instruction /= 10

			index2 := amp.inputArr[amp.index+2]
			mode2 := instruction % 10
			val2 := index2
			if mode2 == 0 {
				val2 = amp.inputArr[index2]
			}
			instruction /= 10

			targetIndex := amp.inputArr[amp.index+3]

			sum := val1 + val2
			amp.inputArr[targetIndex] = sum

			amp.index += 4
			break
		case 2:
			index1 := amp.inputArr[amp.index+1]
			mode1 := instruction % 10
			val1 := index1
			if mode1 == 0 {
				val1 = amp.inputArr[index1]
			}
			instruction /= 10

			index2 := amp.inputArr[amp.index+2]
			mode2 := instruction % 10
			val2 := index2
			if mode2 == 0 {
				val2 = amp.inputArr[index2]
			}
			instruction /= 10

			targetIndex := amp.inputArr[amp.index+3]

			product := val1 * val2
			amp.inputArr[targetIndex] = product

			amp.index += 4
			break
		case 3:
			val := inputVal
			if !amp.usedPhase {
				val = amp.phase
				amp.usedPhase = true
			}

			targetIndex := amp.inputArr[amp.index+1]
			amp.inputArr[targetIndex] = val

			amp.index += 2
			break
		case 4:
			index1 := amp.inputArr[amp.index+1]
			val := amp.inputArr[index1]
			amp.outputQueue = append(amp.outputQueue, val)

			amp.index += 2
			return
		case 5:
			index1 := amp.inputArr[amp.index+1]
			mode1 := instruction % 10
			instruction /= 10
			val1 := index1
			if mode1 == 0 {
				val1 = amp.inputArr[index1]
			}

			if val1 != 0 {
				index2 := amp.inputArr[amp.index+2]
				mode2 := instruction % 10
				newInstructionPointer := index2
				if mode2 == 0 {
					newInstructionPointer = amp.inputArr[index2]
				}
				amp.index = newInstructionPointer
			} else {
				amp.index += 3
			}
			break
		case 6:
			index1 := amp.inputArr[amp.index+1]
			mode1 := instruction % 10
			instruction /= 10
			val1 := index1
			if mode1 == 0 {
				val1 = amp.inputArr[index1]
			}

			if val1 == 0 {
				index2 := amp.inputArr[amp.index+2]
				mode2 := instruction % 10
				newInstructionPointer := index2
				if mode2 == 0 {
					newInstructionPointer = amp.inputArr[index2]
				}
				amp.index = newInstructionPointer
			} else {
				amp.index += 3
			}
			break
		case 7:
			index1 := amp.inputArr[amp.index+1]
			mode1 := instruction % 10
			val1 := index1
			if mode1 == 0 {
				val1 = amp.inputArr[index1]
			}
			instruction /= 10

			index2 := amp.inputArr[amp.index+2]
			mode2 := instruction % 10
			val2 := index2
			if mode2 == 0 {
				val2 = amp.inputArr[index2]
			}
			instruction /= 10

			targetIndex := amp.inputArr[amp.index+3]

			if val1 < val2 {
				amp.inputArr[targetIndex] = 1
			} else {
				amp.inputArr[targetIndex] = 0
			}

			amp.index += 4
			break
		case 8:
			index1 := amp.inputArr[amp.index+1]
			mode1 := instruction % 10
			val1 := index1
			if mode1 == 0 {
				val1 = amp.inputArr[index1]
			}
			instruction /= 10

			index2 := amp.inputArr[amp.index+2]
			mode2 := instruction % 10
			val2 := index2
			if mode2 == 0 {
				val2 = amp.inputArr[index2]
			}
			instruction /= 10

			targetIndex := amp.inputArr[amp.index+3]

			if val1 == val2 {
				amp.inputArr[targetIndex] = 1
			} else {
				amp.inputArr[targetIndex] = 0
			}

			amp.index += 4
			break
		default:
			fmt.Println("Error. Got opcode ", opcode, "on instruction ", amp.index)
			return
		}
	}
	return
}

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

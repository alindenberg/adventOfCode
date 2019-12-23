package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := "59719811742386712072322509550573967421647565332667367184388997335292349852954113343804787102604664096288440135472284308373326245877593956199225516071210882728614292871131765110416999817460140955856338830118060988497097324334962543389288979535054141495171461720836525090700092901849537843081841755954360811618153200442803197286399570023355821961989595705705045742262477597293974158696594795118783767300148414702347570064139665680516053143032825288231685962359393267461932384683218413483205671636464298057303588424278653449749781937014234119757220011471950196190313903906218080178644004164122665292870495547666700781057929319060171363468213087408071790"
	arr := getInputArr(input)

	res := runPhases(100, arr)
	for i := 0; i < 8; i++ {
		fmt.Print(res[i])
	}
	fmt.Println()
}

func runPhases(count int, input []int) []int {
	for i := 0; i < count; i++ {
		output := []int{}
		for outputIndex := 0; outputIndex < len(input); outputIndex++ {
			sum := 0
			for inputIndex := 0; inputIndex < len(input); inputIndex++ {
				multiplier := getPatternMultiplier(outputIndex, inputIndex)
				sum += input[inputIndex] * multiplier
			}
			output = append(output, int(math.Abs(float64(sum%10))))
		}
		input = output
	}
	return input
}

func getPatternMultiplier(outputIndex int, inputIndex int) int {
	if inputIndex < outputIndex {
		return 0
	}

	location := inputIndex - outputIndex
	index := location % (4 * (outputIndex + 1))
	actualPosition := math.Floor(float64(index / (outputIndex + 1)))
	if actualPosition == 0 {
		return 1
	}
	if actualPosition == 2 {
		return -1
	}
	return 0
}

func getInputArr(input string) []int {
	result := []int{}
	for i := 0; i < len(input); i++ {
		val, _ := strconv.Atoi(string(input[i]))
		result = append(result, val)
	}
	return result
}

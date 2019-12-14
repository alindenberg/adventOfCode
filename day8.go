package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// https://adventofcode.com/2019/day/8
func main() {
	file, err := ioutil.ReadFile("day8input.txt")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	height := 6
	width := 25
	layers := getLayers(file, height*width)

	decodedImage := getDecodedImage(layers, width*height)
	for i := 0; i < height*width; i++ {
		fmt.Print(decodedImage[i])
		if i < (height*width)-1 && (i+1)%width == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func getLayers(fileBytes []byte, layerLength int) [][]int {
	arr := strings.Split(string(fileBytes), "")
	arrIndex := 0
	layers := [][]int{}
	for arrIndex < len(arr) {
		layer := []int{}
		for i := 0; i < layerLength; i++ {
			val, _ := strconv.Atoi(arr[arrIndex])
			layer = append(layer, val)
			arrIndex++
		}
		layers = append(layers, layer)
	}
	return layers
}

func getDecodedImage(layers [][]int, totalCount int) []string {
	arr := []string{}
	arrIndex := 0
	for arrIndex < totalCount {
		for _, layer := range layers {
			if layer[arrIndex] == 2 {
				continue
			}

			val := " "
			if layer[arrIndex] == 1 {
				val = "#"
			}
			arr = append(arr, val)
			arrIndex++
			break
		}
	}

	return arr
}

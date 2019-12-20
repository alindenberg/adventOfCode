package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("day14input.txt")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	inputStrings := strings.Split(string(file), "\n")
	chemicalMap := *getChemicalMap(inputStrings)

	baseChemicals := make(map[string]int)
	for _, chemical := range chemicalMap {
		if len(*chemical.formula) == 1 && (*chemical.formula)["ORE"] > 0 {
			baseChemicals[chemical.name] = (*chemical.formula)["ORE"]
		}
	}

	amountUsed := make(map[Name]Amount)
	surplus := make(map[Name]Amount)

	fuelNeeded := 1
	ore := 0
	result := *getCostInBaseChemicals("FUEL", fuelNeeded, &chemicalMap, &baseChemicals, &amountUsed, &surplus)
	for _, amountNeeded := range result {
		ore += amountNeeded
	}
	fmt.Println("Part 1 : ", ore)

	// part 2 - brute forced by trying a few different values for fuelNeeded, until highest return
	// value that was under a trillion. Ideally would implement a binary search to run on that between two
	// given fuelNeeded values
	// fmt.Println("Part 2 - ", 13108426)
}

func getCostInBaseChemicals(
	name Name, multiple Amount,
	chemicalMap *map[Name]*Chemical, baseChemicals *map[Name]Amount,
	amountUsed *map[Name]Amount, surplus *map[Name]Amount) *map[Name]Amount {
	if productionAmount, ok := (*baseChemicals)[name]; ok {
		return &(map[Name]Amount{name: (multiple * productionAmount)})
	}

	baseCostMap := make(map[Name]Amount)
	for chemicalName, chemicalAmount := range *((*chemicalMap)[name]).formula {
		productionAmount := (*(*chemicalMap)[chemicalName]).quantity
		amountNeeded := chemicalAmount * multiple

		initialValue := 0
		if (*surplus)[chemicalName] <= amountNeeded {
			initialValue = (*surplus)[chemicalName]
			(*surplus)[chemicalName] = 0
		} else {
			initialValue = (*surplus)[chemicalName]
			(*surplus)[chemicalName] -= initialValue
		}

		productionsNeeded := 0
		for initialValue < amountNeeded {
			productionsNeeded++
			initialValue += productionAmount
		}

		if chemicalName == "HKGWZ" && name == "FUEL" {
			fmt.Println("Production amount is ", productionAmount)
			fmt.Println("Productions needed is ", productionsNeeded)
		}

		(*surplus)[chemicalName] = initialValue - amountNeeded

		updateMap(&baseCostMap, getCostInBaseChemicals(chemicalName, productionsNeeded, chemicalMap, baseChemicals, amountUsed, surplus))
	}
	return &baseCostMap
}

func updateMap(originalMap *map[Name]Amount, newMap *map[Name]Amount) {
	for name, amount := range *newMap {
		(*originalMap)[name] += amount
	}
}

func getChemicalMap(inputStrings []string) *map[Name]*Chemical {
	chemicals := make(map[Name]*Chemical)
	for _, str := range inputStrings {
		strComponents := strings.Split(str, "=>")

		inputs := strings.Split(strings.Trim(strComponents[0], " "), ",")
		formula := make(map[Name]Amount)
		for _, input := range inputs {
			inputComponents := strings.Split(strings.Trim(input, " "), " ")
			quantity, _ := strconv.Atoi(inputComponents[0])
			name := inputComponents[1]

			formula[name] = quantity
		}
		output := strings.Split(strings.Trim(strComponents[1], " "), " ")
		name := output[1]
		quantity, _ := strconv.Atoi(output[0])

		chemicals[name] = chemical(name, quantity, &formula)
	}
	return &chemicals
}

func chemical(name Name, quantity Amount, formula *map[Name]Amount) *Chemical {
	return &Chemical{
		formula,
		quantity,
		name,
	}
}

type Chemical struct {
	formula  *map[Name]Amount
	quantity Amount
	name     Name
}

type Name = string
type Amount = int

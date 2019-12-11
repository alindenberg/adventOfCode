package main

import (
	"fmt"
)

func main() {
	input1 := 264360
	input2 := 746325

	validNumbers := []int{}
	for i := input1; i <= input2; i++ {
		val := i
		lastDigit := -1
		isValid := true
		hasDouble := false
		matchingDigits := []int{}
		for val > 9 {

			digit := val % 10
			val /= 10

			if lastDigit == -1 {
				matchingDigits = append(matchingDigits, digit)
				lastDigit = digit
				continue
			}
			if digit > lastDigit {
				isValid = false
				break
			}
			if digit == lastDigit {
				matchingDigits = append(matchingDigits, digit)
			} else {
				if len(matchingDigits) == 2 {
					hasDouble = true
				}
				matchingDigits = []int{digit}
			}
			lastDigit = digit
		}
		if val > lastDigit {
			isValid = false
		} else if val == lastDigit {
			if len(matchingDigits) == 1 {
				hasDouble = true
			}
		} else {
			if len(matchingDigits) == 2 {
				hasDouble = true
			}
		}
		if isValid && hasDouble {
			validNumbers = append(validNumbers, i)
		}
	}
	fmt.Println("COUNT:  ", len(validNumbers))
}

package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"strconv"
)

func day3Part1() {
	lines, err := utils.ReadLines("input/day3.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		total := 0
		for _, line := range lines {
			joltage := 0
			for i := 0; i < len(line); i++ {
				digit := line[i]
				rest := line[(i + 1):]
				for _, char := range rest {
					value, _ := strconv.Atoi(string(digit) + string(char))
					if value > joltage {
						joltage = value
					}
				}
			}
			fmt.Printf("In %s, you can make the largest joltage possible %d.\n", line, joltage)
			total += joltage
		}
		fmt.Printf("The total output joltage is %d.\n", total)
	}
}

func day3Part2() {
	lines, err := utils.ReadLines("input/day3.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		limit := 12
		total := 0
		for _, line := range lines {
			var joltage [12]int
			startPosition := 0
			for i := 0; i < limit; i++ {
				index := 0
				maxDigit := 0
				subline := line[startPosition : len(line)-(limit-1-i)]

				for j, c := range subline {
					digit, _ := strconv.Atoi(string(c))
					if maxDigit < digit {
						maxDigit = digit
						index = j
					}
				}

				joltage[i] = maxDigit
				startPosition += (index + 1)
			}

			joltageValue := arrayToInt(joltage)
			fmt.Printf("In %s, you can make the largest joltage possible %d.\n", line, joltageValue)
			total += joltageValue
		}
		fmt.Printf("The total output joltage is %d.\n", total)
	}
}

func arrayToInt(a [12]int) int {
	result := 0
	for _, digit := range a {
		result = result*10 + digit
	}
	return result
}

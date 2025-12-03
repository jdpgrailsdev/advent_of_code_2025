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
			batteryIndex := 0
			for i := 0; i < len(line); i++ {
				digit, _ := strconv.Atoi(string(line[i]))
				if joltage[batteryIndex] <= digit {
					if len(line)-i < (limit - batteryIndex) {
						batteryIndex = min((limit - 1), batteryIndex+1)
						joltage[batteryIndex] = digit
					} else {
						joltage[batteryIndex] = digit
					}
				} else if batteryIndex == limit-1 {
					if joltage[batteryIndex] <= digit {
						joltage[batteryIndex] = digit
					}
				} else {
					batteryIndex++
					joltage[batteryIndex] = digit
				}
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

package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"strconv"
	"strings"
	"unicode"
)

type Problem struct {
	Data     []int
	Operator string
}

type Problem2 struct {
	Data     []string
	Operator string
}

func day6Part1() {
	lines, err := utils.ReadLines("input/day6.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		rowCount := len(lines)
		colCount := len(strings.Fields(lines[0]))
		problems := make([]Problem, colCount)
		data := make([][]int, len(problems))
		for i, line := range lines {
			columns := strings.Fields(line)
			if columns[0] == "+" || columns[0] == "*" {
				for o, operator := range columns {
					problems[o] = Problem{
						Data:     data[o],
						Operator: operator,
					}
				}
			} else {
				for c, value := range columns {
					intValue, _ := strconv.Atoi(value)
					if len(data[c]) == 0 {
						data[c] = make([]int, rowCount-1)
					}
					data[c][i] = intValue
				}
			}
		}

		total := 0
		for _, problem := range problems {
			if problem.Operator == "+" {
				sum := 0
				for _, value := range problem.Data {
					sum += value
				}
				total += sum
			} else {
				product := problem.Data[0]
				for _, value := range problem.Data[1:] {
					product *= value
				}
				total += product
			}
		}

		fmt.Printf("The grand total of all of the answers is %d.\n", total)
	}
}

func day6Part2() {
	lines, err := utils.ReadLines("input/day6.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		total := 0
		maxLength := 0
		for _, line := range lines {
			if len(line) > maxLength {
				maxLength = len(line)
			}
		}
		numbers := []int{}
		for i := maxLength - 1; i >= 0; i-- {
			currentNum := 0
			for j := 0; j < len(lines)-1; j++ {
				line := lines[j]
				if i < len(line) {
					currentChar := line[i]
					if !unicode.IsSpace(rune(currentChar)) {
						number, _ := strconv.Atoi(string(currentChar))
						currentNum = (currentNum * 10) + number
					}
				}
			}
			if currentNum != 0 {
				numbers = append(numbers, currentNum)
			}
			if i < len(lines[len(lines)-1]) {
				operator := lines[len(lines)-1][i]
				if operator == '+' {
					sum := 0
					for _, value := range numbers {
						sum += value
					}
					//                     fmt.Printf("Performing operation + on values %v = %d\n", numbers, sum)
					total += sum
					numbers = []int{}
				} else if operator == '*' {
					product := numbers[0]
					for _, operand := range numbers[1:] {
						product *= operand
					}
					//                     fmt.Printf("Performing operation * on values %v = %d\n", numbers, product)
					total += product
					numbers = []int{}
				}
			}
		}

		fmt.Printf("The grand total of all of the answers is %d.\n", total)
	}
}

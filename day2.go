package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"strconv"
	"strings"
)

func day2() {
	lines, err := utils.ReadLines("input/day2.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		total := 0
		ranges := strings.Split(lines[0], ",")
		for _, r := range ranges {
			values := strings.Split(r, "-")
			low, _ := strconv.Atoi(values[0])
			high, _ := strconv.Atoi(values[1])

			for i := low; i <= high; i++ {
				s := strconv.Itoa(i)
				l := len(s)
				if l%2 == 0 {
					left := s[0:(l / 2)]
					right := s[(l / 2):l]

					if left == right {
						fmt.Printf("%d is an invalid product ID.\n", i)
						total += i
					}
				}
			}
		}

		fmt.Printf("Sum of invalid product ID's is %d.\n", total)
	}
}

func day2Part2() {
	lines, err := utils.ReadLines("input/day2.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		total := 0
		ranges := strings.Split(lines[0], ",")
		for _, r := range ranges {
			values := strings.Split(r, "-")
			low, _ := strconv.Atoi(values[0])
			high, _ := strconv.Atoi(values[1])

			for i := low; i <= high; i++ {
				s := strconv.Itoa(i)
				l := len(s)
				accum := ""
				for _, char := range s {
					accum += string(char)
					if len(accum) != l && len(strings.Replace(s, accum, "", -1)) == 0 {
						fmt.Printf("%d is an invalid product ID.\n", i)
						total += i
						break
					}
				}
			}
		}

		fmt.Printf("Sum of invalid product ID's is %d.\n", total)
	}
}

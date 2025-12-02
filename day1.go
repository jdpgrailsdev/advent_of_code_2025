package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"regexp"
	"strconv"
)

func day1() {
	lines, err := utils.ReadLines("input/day1.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		pattern := regexp.MustCompile(`(L|R)(\d+)`)
		position := 50
		count := 0
		fmt.Printf("The dial starts by pointing at %d\n", position)
		for _, line := range lines {
			match := pattern.FindStringSubmatch(line)
			direction := match[1]
			clicks, parseErr := strconv.Atoi(match[2])
			if parseErr != nil {
				fmt.Printf("Error parsing line input %s: %v\n", line, err)
			} else {
				previousPosition := position
				fullRotations := 0
				if (clicks / 100) > 0 {
					fullRotations = clicks / 100
				}

				netClicks := clicks
				// A full rotation puts the dial back to the same location, so
				// we only care about rotations that are net beyond that
				if (clicks % 100) > 0 {
					netClicks = clicks % 100
				}

				rotationCount := fullRotations

				if direction == "L" {
					position -= netClicks
				} else {
					position += netClicks
				}

				if position > 99 {
					position = position - 100
					if position != 0 {
						rotationCount++
					}
				} else if position < 0 {
					position = position + 100
					if previousPosition != 0 && position != 0 {
						rotationCount++
					}
				}

				if position == 0 {
					rotationCount++
				}

				count += rotationCount

				fmt.Printf("The dial is rotated %s%d to point at %d; During this rotation, it points at zero %d time(s).\n", direction, clicks, position, rotationCount)
			}
		}
		fmt.Printf("Pass-count: %d\n", count)
	}
}

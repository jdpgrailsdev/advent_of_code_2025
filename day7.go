package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"slices"
	"strings"
)

func day7Part1() {
	lines, err := utils.ReadLines("input/day7.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		splitCount := 0
		streamPos := []int{}
		for _, line := range lines {
			if len(streamPos) == 0 {
				pos := strings.IndexRune(line, 'S')
				streamPos = append(streamPos, pos)
			} else {
				newStreams := []int{}
				for _, pos := range streamPos {
					if line[pos] == '.' {
						newStreams = append(newStreams, pos)
					} else if line[pos] == '^' {
						splitCount++
						newStreams = append(newStreams, pos-1, pos+1)
					}
				}

				slices.Sort(newStreams)
				streamPos = slices.Compact(newStreams)
				fmt.Printf("Stream positions = %v\n", streamPos)
			}
		}

		fmt.Printf("The beam has been split a total of %d times.\n", splitCount)
	}
}

type Position struct {
	Row int
	Col int
}

func day7Part2() {
	lines, err := utils.ReadLines("input/day7.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		line := lines[0]
		pos := strings.IndexRune(line, 'S')
		position := Position{Row: 1, Col: pos}
		cache := make(map[Position]int)
		totalPaths := search(lines, position, cache)

		fmt.Printf("The beam has been split a total of %d timelines.\n", totalPaths)
	}
}

func search(lines []string, position Position, cache map[Position]int) int {
	if position.Row == len(lines)-1 {
		return 1
	}

	if cached, true := cache[position]; true {
		return cached
	}

	count := 0
	if lines[position.Row][position.Col] == '.' {
		count = search(lines, Position{Row: position.Row + 1, Col: position.Col}, cache)
	} else if lines[position.Row][position.Col] == '^' {
		count = search(lines, Position{Row: position.Row + 1, Col: position.Col - 1}, cache) + search(lines, Position{Row: position.Row + 1, Col: position.Col + 1}, cache)
	}

	cache[position] = count
	return count

}

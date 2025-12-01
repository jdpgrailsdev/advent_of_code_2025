package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	day1()
}

func day1() {
	var pattern = regexp.MustCompile(`(L|R)(\d+)`)
	var lines, err = readLines("input/day1.txt")
	var position = 50
	var count = 0
	if err == nil {
		fmt.Printf("The dial starts by pointing at %d\n", position)
		for _, line := range lines {
			var match = pattern.FindStringSubmatch(line)
			var direction = match[1]
			var clicks, parseErr = strconv.Atoi(match[2])
			if parseErr == nil {
				var netClicks = clicks
				// A full rotation puts the dial back to the same location, so
				// we only care about rotations that are net beyond that
				if (clicks % 100) > 0 {
					netClicks = clicks % 100
				}

				if direction == "L" {
					position -= netClicks
				} else {
					position += netClicks
				}

				if position > 99 {
					position = position - 100
				} else if position < 0 {
					position = position + 100
				}

				fmt.Printf("The dial is rotated %s%d to point at %d.\n", direction, clicks, position)

				if position == 0 {
					count++
				}
			}
		}
	}

	fmt.Printf("Pass-count: %d", count)
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after the function returns

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error during scanning: %w", err)
	}

	return lines, nil
}

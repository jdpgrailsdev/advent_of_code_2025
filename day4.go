package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
)

func day4Part1() {
	lines, err := utils.ReadLines("input/day4.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		matrix := buildMatrix(lines)

		total := 0

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == '@' {
					neighborCount := checkNeighbors(i, j, matrix)
					if neighborCount < 4 {
						total++
					}
				}
			}
		}

		fmt.Printf("A total of %d paper rolls of paper can be accessed by a forklift.\n", total)
	}
}

func day4Part2() {
	lines, err := utils.ReadLines("input/day4.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		matrix := buildMatrix(lines)
		total := scanMatrix(matrix, 0)
		fmt.Printf("A total of %d paper rolls of paper can be accessed by a forklift.\n", total)
	}
}

func scanMatrix(matrix [][]rune, total int) int {
	removed := false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '@' {
				neighborCount := checkNeighbors(i, j, matrix)
				if neighborCount < 4 {
					matrix[i][j] = 'X'
					removed = true
					total++
				}
			}
		}
	}
	if removed == true {
		total = scanMatrix(matrix, total)
	}

	return total

}

func checkNeighbors(x int, y int, matrix [][]rune) int {
	count := 0
	relativeX := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	relativeY := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		newX, newY := x+relativeX[i], y+relativeY[i]

		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
			if matrix[newX][newY] == '@' {
				count++
			}
		}
	}

	return count
}

func buildMatrix(lines []string) [][]rune {
	matrix := [][]rune{}
	for _, line := range lines {
		row := []rune(line)
		matrix = append(matrix, row)
	}
	return matrix
}

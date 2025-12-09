package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) toString() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

type Rectangle struct {
	X1 int // min
	X2 int // max
	Y1 int // min
	Y2 int // max
}

func (r Rectangle) area() int {
	return (r.X2 - r.X1 + 1) * (r.Y2 - r.Y1 + 1)
}

func (r Rectangle) overlaps(other Rectangle) bool {
	return r.X1 < other.X2 && r.X2 > other.X1 && r.Y1 < other.Y2 && r.Y2 > other.Y1
}

func (r Rectangle) toString() string {
	return fmt.Sprintf("(%d,%d),(%d,%d)", r.X1, r.Y1, r.X2, r.Y2)
}

func day9Part1() {
	lines, err := utils.ReadLines("input/day9.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		coordinates := []Coordinate{}

		for _, line := range lines {
			coords := strings.Split(line, ",")
			y, _ := strconv.Atoi(coords[0])
			x, _ := strconv.Atoi(coords[1])
			coordinates = append(coordinates, Coordinate{X: x, Y: y})
		}

		sort.Slice(coordinates, func(i, j int) bool {
			if coordinates[i].Y != coordinates[j].Y {
				return coordinates[i].Y < coordinates[j].Y
			}
			return coordinates[i].X < coordinates[j].X
		})

		maxArea := findMaxArea(coordinates[0], 0, coordinates)

		fmt.Printf("The largest area of any rectangle is %d\n", maxArea)
	}
}

func day9Part2() {
	lines, err := utils.ReadLines("input/day9.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		coordinates := []Coordinate{}

		for _, line := range lines {
			coords := strings.Split(line, ",")
			y, _ := strconv.Atoi(coords[0])
			x, _ := strconv.Atoi(coords[1])
			coordinates = append(coordinates, Coordinate{X: x, Y: y})
		}

		maxArea := 0
		edges := []Rectangle{}
		for i, coordinate := range coordinates {
			r := buildRectangle(coordinate, coordinates[(i+1)%len(coordinates)])
			edges = append(edges, r)
		}

		rectangles := buildRectangles(coordinates)

		for _, r := range rectangles {
			overlaps := false
			for _, e := range edges {
				if r.overlaps(e) == true {
					overlaps = true
					break
				}
			}

			if overlaps == false {
				maxArea = max(maxArea, r.area())
			}
		}

		fmt.Printf("The largest area of any rectangle is %d\n", maxArea)
	}
}

func findMaxArea(anchor Coordinate, currentMaxArea int, coordinates []Coordinate) int {
	maxArea := currentMaxArea

	for i := 0; i < len(coordinates); i++ {
		coordinate := coordinates[i]
		if coordinate.X != anchor.X && coordinate.Y != anchor.Y {
			maxArea = max(maxArea, (coordinate.X-anchor.X+1)*(coordinate.Y-anchor.Y+1))
		}
	}

	nextIndex := slices.Index(coordinates, anchor) + 1

	if nextIndex < len(coordinates) {
		maxArea = max(maxArea, findMaxArea(coordinates[nextIndex], maxArea, coordinates))
	}

	return maxArea
}

func buildRectangles(coordinates []Coordinate) []Rectangle {
	rectangles := []Rectangle{}

	for i := 0; i < len(coordinates); i++ {
		for j := 0; j < len(coordinates); j++ {
			if i != j {
				r := buildRectangle(coordinates[i], coordinates[j])
				rectangles = append(rectangles, r)
			}
		}
	}

	return rectangles
}

func buildRectangle(coordinate1 Coordinate, coordinate2 Coordinate) Rectangle {
	x1 := min(coordinate1.X, coordinate2.X)
	y1 := min(coordinate1.Y, coordinate2.Y)
	x2 := max(coordinate1.X, coordinate2.X)
	y2 := max(coordinate1.Y, coordinate2.Y)
	return Rectangle{X1: x1, X2: x2, Y1: y1, Y2: y2}
}

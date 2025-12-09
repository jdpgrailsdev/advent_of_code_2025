package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	ID int
	X  int
	Y  int
	Z  int
}

func (j JunctionBox) toString() string {
	return fmt.Sprintf("(%.0f,%.0f,%.0f)", j.X, j.Y, j.Z)
}

func (j JunctionBox) compare(other JunctionBox) bool {
	return j == other
}

type Connection struct {
	Left  JunctionBox
	Right JunctionBox
}

type Circuit struct {
	Boxes []JunctionBox
}

func (c Circuit) compare(other Circuit) bool {
	return slices.Equal(c.Boxes, other.Boxes)
}

func (c Circuit) print() {
	for _, b := range c.Boxes {
		fmt.Printf("%s\t", b.toString())
	}
	fmt.Println()
}

func day8Part1() {
	input := "input/day8.txt"
	lines, err := utils.ReadLines(input)
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		maxPairs := 1000
		if strings.Contains(input, "example") {
			maxPairs = 10
		}
		junctionBoxes := []JunctionBox{}
		for i, line := range lines {
			coords := strings.Split(line, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			z, _ := strconv.Atoi(coords[2])
			junctionBoxes = append(junctionBoxes, JunctionBox{ID: i, X: x, Y: y, Z: z})
		}

		connections := computeDistances(junctionBoxes)
		circuits := findCircuits(connections, junctionBoxes, maxPairs)
		sort.Slice(circuits, func(i, j int) bool {
			return len(circuits[i].Boxes) > len(circuits[j].Boxes)
		})

		total := 1
		for _, c := range circuits[0:3] {
			fmt.Printf("Using circuit with length %d in total product...\n", len(c.Boxes))
			total *= len(c.Boxes)
		}

		fmt.Printf("The product of the three largest circuits is %d\n", total)
	}
}

func day8Part2() {
	// lines, err := utils.ReadLines("input/day8-example.txt")
	//
	//	if err != nil {
	//		fmt.Printf("Error reading from input file: %v\n", err)
	//	} else {
	//
	// }
}

func sortConnections(connections map[int]Connection) []int {
	keys := []int{}
	for k := range connections {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return keys
}

func findCircuits(connections map[int]Connection, boxes []JunctionBox, maxPairs int) []Circuit {
	circuits := []Circuit{}
	pairs := 0
	sortedKeys := sortConnections(connections)

	for _, key := range sortedKeys {
		connection := connections[key]
		left := connection.Left
		right := connection.Right

		if left != right {
			circuit, index := findCircuit(circuits, left, right)

			if circuit == nil {
				circuit = &Circuit{Boxes: []JunctionBox{}}
			}

			if !slices.Contains(circuit.Boxes, left) || !slices.Contains(circuit.Boxes, right) {
				circuit.Boxes = appendUnique(circuit.Boxes, left)
				circuit.Boxes = appendUnique(circuit.Boxes, right)
				if index > -1 {
					circuits[index] = *circuit
				} else {
					circuits = append(circuits, *circuit)
				}
				pairs++
			}
		}

		if pairs == maxPairs {
			break
		}
	}

	return circuits
}

func findCircuit(circuits []Circuit, box1 JunctionBox, box2 JunctionBox) (*Circuit, int) {
	for i, c := range circuits {
		if slices.Contains(c.Boxes, box1) || slices.Contains(c.Boxes, box2) {
			return &c, i
		}
	}

	return nil, -1

}

func computeDistance(box1 JunctionBox, box2 JunctionBox) int {
	// Avoid sqrt for performance
	dx := box1.X - box2.X
	dy := box1.Y - box2.Y
	dz := box1.Z - box2.Z
	return (dx * dx) + (dy * dy) + (dz * dz)
}

func computeDistances(boxes []JunctionBox) map[int]Connection {
	points := len(boxes)
	connections := make(map[int]Connection)

	for i := 0; i < points; i++ {
		for j := 0; j < points; j++ {
			if i != j {
				left := boxes[i]
				right := boxes[j]
				distance := computeDistance(left, right)
				_, ok := connections[distance]
				if !ok {
					connections[distance] = Connection{Left: left, Right: right}
				}
			}
		}
	}

	return connections
}

func appendUnique(slice []JunctionBox, element JunctionBox) []JunctionBox {
	if slices.Contains(slice, element) {
		return slice
	} else {
		return append(slice, element)
	}
}

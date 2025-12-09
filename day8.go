package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	X float64
	Y float64
	Z float64
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

func day8Part1() {
	lines, err := utils.ReadLines("input/day8.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		maxPairs := 1000
		junctionBoxes := []JunctionBox{}
		for _, line := range lines {
			coords := strings.Split(line, ",")
			x, _ := strconv.ParseFloat(coords[0], 64)
			y, _ := strconv.ParseFloat(coords[1], 64)
			z, _ := strconv.ParseFloat(coords[2], 64)
			junctionBoxes = append(junctionBoxes, JunctionBox{X: x, Y: y, Z: z})
		}

		connections := computeDistances(junctionBoxes)
		circuits := findCircuits(connections, junctionBoxes, maxPairs)
		sort.Slice(circuits, func(i, j int) bool {
			return len(circuits[i].Boxes) > len(circuits[j].Boxes)
		})

		totalPairs := 0
		for _, c := range circuits[0:3] {
			totalPairs += len(c.Boxes)
		}
		fmt.Printf("Total junction pairs present in circuits: %d\n", totalPairs)

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

func sortConnections(connections map[float64]Connection) []float64 {
	keys := []float64{}
	for k := range connections {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	return keys
}

func findCircuits(connections map[float64]Connection, boxes []JunctionBox, maxPairs int) []Circuit {
	circuits := []Circuit{}
	pairs := 0
	sortedKeys := sortConnections(connections)

	for _, key := range sortedKeys {
		connection := connections[key]
		left := connection.Left
		right := connection.Right

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
		} else {
			fmt.Printf("Circuit already contains %s and %s: \n", left.toString(), right.toString())
			for _, b := range circuit.Boxes {
				fmt.Printf("%s ", b.toString())
			}
			fmt.Println()
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

func computeDistance(box1 JunctionBox, box2 JunctionBox) float64 {
	return math.Pow(box1.X-box2.X, 2) + math.Pow(box1.Y-box2.Y, 2) + math.Pow(box1.Z-box2.Z, 2)
}

func computeDistances(boxes []JunctionBox) map[float64]Connection {
	points := len(boxes)
	connections := make(map[float64]Connection)

	for i := 0; i < points; i++ {
		for j := 0; j < points; j++ {
			if i != j {
				left := boxes[i]
				right := boxes[j]
				distance := computeDistance(left, right)
				connections[distance] = Connection{Left: left, Right: right}
			}
		}
	}

	return connections
}

func appendUnique(slice []JunctionBox, element JunctionBox) []JunctionBox {
	for _, box := range slice {
		if box.compare(element) {
			return slice
		}
	}
	return append(slice, element)
}

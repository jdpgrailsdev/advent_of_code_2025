package main

import (
	"fmt"
	"github.com/jdpgrailsdev/utils"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func day5Part1() {
	lines, err := utils.ReadLines("input/day5.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		ranges := []Range{}
		index := 0
		fresh := 0
		for i, line := range lines {
			if line == "" {
				index = i + 1
				break
			}
			values := strings.Split(line, "-")
			min, _ := strconv.Atoi(values[0])
			max, _ := strconv.Atoi(values[1])
			ranges = append(ranges, Range{
				Min: min,
				Max: max,
			})
		}

		for _, line := range lines[index:] {
			id, _ := strconv.Atoi(line)
			for _, r := range ranges {
				if id >= r.Min && id <= r.Max {
					fresh++
					break
				}
			}
		}

		fmt.Printf("There are %d fresh ingredient IDs in the inventory management system.\n", fresh)
	}
}

type Node struct {
	Value Range
	Left  *Node
	Right *Node
}

func day5Part2() {
	lines, err := utils.ReadLines("input/day5.txt")
	if err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
	} else {
		ranges := []Range{}
		for _, line := range lines {
			if line == "" {
				break
			}
			values := strings.Split(line, "-")
			min, _ := strconv.Atoi(values[0])
			max, _ := strconv.Atoi(values[1])
			ranges = append(ranges, Range{
				Min: min,
				Max: max,
			})
		}

		// Sort the ranges by starting/min value to make tree insertion work correctly
		sort.SliceStable(ranges, func(i, j int) bool {
			return ranges[i].Min < ranges[j].Min
		})

		// Build the range tree, which will merge any overlapping ranges
		root := &Node{Value: ranges[0]}
		for _, r := range ranges[1:] {
			root.Insert(r)
		}

		freshIds := countFresh(root)
		fmt.Printf("There are %d possible fresh ingredient IDs in the inventory management system.\n", freshIds)
	}
}

func (n *Node) Insert(value Range) *Node {
	if n == nil {
		return &Node{Value: value}
	}

	if value.Max >= n.Value.Min && value.Min <= n.Value.Max {
		if value.Min >= n.Value.Min && value.Max <= n.Value.Max {
			// If the value is completely contained within the current node's value, skip it
			return n
		}
		n.Value.Min = min(n.Value.Min, value.Min)
		n.Value.Max = max(n.Value.Max, value.Max)
		return n
	}

	if value.Min < n.Value.Min {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left = n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right = n.Right.Insert(value)
		}
	}
	return n
}

func countFresh(node *Node) int {
	if node == nil {
		return 0
	}

	count := countFresh(node.Left)
	count += (node.Value.Max + 1) - node.Value.Min
	count += countFresh(node.Right)
	return count
}

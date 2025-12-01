package main

import (
	"aoc-2025-go/utils"
	"fmt"
)

func part1(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	fmt.Printf("Processed %v lines", len(lines))

	return 99

}

func part2(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	fmt.Printf("Processed %v lines", len(lines))
	return 99

}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}

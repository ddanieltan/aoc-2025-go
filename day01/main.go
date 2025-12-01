package main

import (
	"aoc-2025-go/utils"
	"fmt"
	"strconv"
)

type Instruction struct {
	Direction rune
	Steps     int
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction
	for _, line := range lines {
		a, b := line[0], line[1:]
		steps, err := strconv.Atoi(string(b))
		if err != nil {
			return nil
		}
		instruct := Instruction{rune(a), steps}
		instructions = append(instructions, instruct)
	}
	return instructions
}

func part1(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	instructions := parseInstructions(lines)
	curr := 50
	var cnt int
	for _, i := range instructions {
		if i.Direction == 'L' {
			curr -= i.Steps
		} else if i.Direction == 'R' {
			curr += i.Steps
		}
		curr = curr % 100
		if curr < 0 {
			curr += 100
		}
		if curr == 0 {
			cnt++
		}
	}

	return cnt
}

func part2(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	instructions := parseInstructions(lines)
	curr := 50
	var cnt int
	for _, i := range instructions {
		if i.Direction == 'L' {
			for j := 0; j < i.Steps; j++ {
				curr--
				if curr < 0 {
					curr = 99
				}
				if curr == 0 {
					cnt++
				}
			}
		} else if i.Direction == 'R' {
			for j := 0; j < i.Steps; j++ {
				curr++
				if curr > 99 {
					curr = 0
				}
				if curr == 0 {
					cnt++
				}
			}
		}
	}
	return cnt

}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}

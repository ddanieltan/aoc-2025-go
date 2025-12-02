package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := part1("example.txt")
	expected := 1227775554
	if result != int64(expected) {
		t.Errorf("Part 1 Example: result=%d, expected=%d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2("example.txt")
	expected := int64(4174379265)
	if result != expected {
		t.Errorf("Part 2 Example: result=%d, expected=%d", result, expected)
	}
}

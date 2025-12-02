package main

import (
	"aoc-2025-go/utils"
	"fmt"
	"strconv"
	"strings"
)

type Id struct {
	first int64
	last  int64
}

func parseIds(line string) []Id {
	commas := strings.Split(line, ",")
	var ids []Id
	for _, s := range commas {
		dash := strings.Split(s, "-")

		// we are working with big numbers here -> int64
		first, err := strconv.ParseInt(dash[0], 10, 64)
		if err != nil {
			panic(err)
		}
		last, err := strconv.ParseInt(dash[1], 10, 64)
		if err != nil {
			panic(err)
		}
		id := Id{first, last}
		ids = append(ids, id)
	}
	return ids
}

func repeatsOnce(num int64) bool {
	numStr := strconv.FormatInt(num, 10)
	// Number is valid if it has odd number of chars
	if len(numStr)%2 != 0 {
		return false
	}
	midpoint := len(numStr) / 2
	frontStr, backStr := numStr[:midpoint], numStr[midpoint:]
	if string(frontStr) == string(backStr) {
		return true
	}
	return false
}

func part1(path string) int64 {
	line, err := utils.ReadString(path)
	if err != nil {
		return 0
	}
	ids := parseIds(line)
	var sum int64
	for _, id := range ids {
		for i := id.first; i <= id.last; i++ {
			if repeatsOnce(i) {
				sum += i
			}
		}
	}
	return sum
}

func repeatsN(num int64, n int) bool {
	numStr := strconv.FormatInt(num, 10)

	if len(numStr) == n {
		return false
	}
	if len(numStr)%n == 0 {
		var repeatedStr string
		if n == 1 {
			repeatedStr = strings.Repeat(string(numStr[:1]), len(numStr))
		} else {
			cutoff := len(numStr) / n
			repeatedStr = strings.Repeat(string(numStr[:cutoff]), n)
		}
		return repeatedStr == numStr
	}
	return false
}

func part2(path string) int64 {
	line, err := utils.ReadString(path)
	if err != nil {
		return 0
	}
	ids := parseIds(line)
	var seen = make(map[int64]bool)
	for _, id := range ids {
		for i := id.first; i <= id.last; i++ {
			// I peeked at input set, max len of string is 10
			// So max no. of chunks is 5
			for _, n := range []int{1, 2, 3, 4, 5} {
				if repeatsN(i, n) {
					seen[i] = true
				}
			}
		}
	}

	var sum int64
	for k, _ := range seen {
		sum += k
	}
	return sum

}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}

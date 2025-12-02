# Learnings from Day 2 Solution

This document provides a critique and suggestions for the Go solution to the Advent of Code 2025 Day 2 problem.

## General Observations

The code is well-structured, with responsibilities separated into different functions. The use of `int64` is appropriate for the large numbers in the problem. The input parsing is handled correctly.

## Critiques and Suggestions

### `part1` Function

The `repeatsOnce` function is clear and correctly implements the logic for part 1. No major issues were found.

### `part2` Function

The `repeatsN` function has some areas for improvement:

1.  **Hardcoded Repetition Checks**: The solution iterates through a hardcoded list of repetitions `[]int{1, 2, 3, 4, 5}`. This is based on an assumption about the input data ("I peeked at input set, max len of string is 10"). This approach is not robust and might fail with different inputs.

    *   **Suggestion**: A more general solution would be to find all possible lengths of the repeating sequence. For a number with `L` digits, the possible lengths of the repeating unit are the divisors of `L` (excluding `L` itself).

2.  **Incorrect Starting Point for Repetitions**: The problem states that an ID is invalid if it is "made only of some sequence of digits repeated **at least twice**." The code starts checking for repetitions from `n=1`, which is incorrect.

    *   **Suggestion**: The loop for checking repetitions should start from `n=2`.

3.  **Redundant Logic in `repeatsN`**: The logic inside `repeatsN` has a special case for `n == 1`. This can be simplified. The `cutoff` can be calculated as `len(numStr) / n` for all `n`, and `strings.Repeat` will work correctly.

    *   **Suggestion**: Unify the logic for all values of `n` to make the function cleaner.

### Refined `repeatsN` and `part2`

Here is an example of how `repeatsN` and `part2` could be improved:

```go
// isValid checks if a number is formed by a repeating sequence of digits.
// The sequence must repeat at least twice.
func isValid(num int64) bool {
	numStr := strconv.FormatInt(num, 10)
	length := len(numStr)

	// Check for all possible lengths of the repeating unit.
	// The length of the unit must be a divisor of the total length.
	for unitLen := 1; unitLen <= length/2; unitLen++ {
		if length%unitLen == 0 {
			repeats := length / unitLen
			if repeats >= 2 {
				unit := numStr[:unitLen]
				if strings.Repeat(unit, repeats) == numStr {
					return true
				}
			}
		}
	}
	return false
}

func part2(path string) int64 {
	line, err := utils.ReadString(path)
	if err != nil {
		return 0
	}
	ids := parseIds(line)
	var sum int64
	for _, id := range ids {
		for i := id.first; i <= id.last; i++ {
			if isValid(i) {
				sum += i
			}
		}
	}
	return sum
}
```

In this revised version:
-   A single function `isValid` checks for any valid repetition.
-   The logic correctly checks for all possible lengths of the repeating unit.
-   It ensures that the number of repetitions is at least 2.
-   The `part2` function is simplified as it no longer needs the `seen` map (since `isValid` is called once per number).

By making these changes, the code becomes more robust, correct, and easier to understand.

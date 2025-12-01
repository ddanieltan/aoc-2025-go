# Learnings from Day 01

Hello! Here is a review of your Go solution for Advent of Code 2025, Day 1. The solution is correct and gets the right answer, which is the most important part of AoC! The following are some suggestions and observations that might be helpful for writing even more robust and efficient Go code in the future.

## 1. Algorithmic Efficiency in Part 2

In `part2`, the solution simulates each individual "click" of the dial.

```go
// From part2 function
for j := 0; j < i.Steps; j++ {
    curr--
    if curr < 0 {
        curr = 99
    }
    if curr == 0 {
        cnt++
    }
}
```

While this works perfectly for the given input, it could be inefficient if the number of steps were extremely large (e.g., billions). The puzzle hints at this with the `R1000` example.

A more efficient, mathematical approach is to calculate how many times the dial would pass `0` without simulating every step.

For any given rotation, the number of times the dial passes `0` is the number of full `360` degree rotations, plus one extra time if the remaining rotation crosses `0`.

Here is how you could calculate the number of times `0` is passed for a given rotation, `i`:
- The number of full rotations is `i.Steps / 100`. Each full rotation passes `0` once.
- The remaining rotation is `i.Steps % 100`.
- For a right turn ('R'), we cross `0` an additional time if the dial starts at `curr` and moves `i.Steps % 100` positions, passing `100` (which is `0` on the dial). This happens if `curr + (i.Steps % 100) >= 100`.
- For a left turn ('L'), we cross `0` an additional time if `curr - (i.Steps % 100) < 0`.

This approach avoids loops and performs the calculation in constant time for each instruction.

## 2. Handling Modulo with Negative Numbers

In `part1`, you correctly handle the circular nature of the dial with this code:

```go
// From part1 function
curr = curr % 100
if curr < 0 {
    curr += 100
}
```

This is clear and works well. There is a common Go idiom to handle this in a single, slightly more concise line. The trick is to add the base before the second modulo operation, which ensures the result is always positive.

```go
// A more concise way to handle positive and negative modulo
curr = (curr%100 + 100) % 100
```
This is not necessarily "better," but it's a useful pattern to recognize and have in your toolkit.

## 3. Using Constants for "Magic Numbers"

The numbers `100` and `99` appear several times in the code. In software engineering, we often call these "magic numbers" because they appear without a clear explanation of what they represent.

It's a good practice to define these as constants at the top of your package. This improves readability and makes the code easier to maintain if the value ever needs to change.

```go
const dialSize = 100

// ... inside a function
curr = (curr%dialSize + dialSize) % dialSize
```

## 4. Robust Input Parsing

The `parseInstructions` function uses string slicing to separate the direction from the number:

```go
// From parseInstructions function
a, b := line[0], line[1:]
```

This works for the puzzle input, but it's a bit fragile. It assumes the direction is always a single byte and the rest of the line is the number. A more robust approach, especially if the format could vary, might involve using functions from the `strings` package, like `strings.Fields` or `strings.SplitN`.

However, for the context of Advent of Code, your direct approach is often the fastest and perfectly acceptable.

## Conclusion

Great job on solving the puzzle! These suggestions are about refining the solution and exploring different Go idioms and performance considerations. Keep up the great work!

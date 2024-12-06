package main

import (
	"fmt"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/advent_of_code"
)

func main() {
	input, err := aoc.GetInputFile("1")
	std.Error(err, 1)

	nums, err := std.ExtractNumbersByLine(input)
	std.Error(err, 2)

	fmt.Println("Day 2 - Part 1:", NumberOfSafeLines(nums))
	fmt.Println("Day 2 - Part 2:", NumberOfSafeLinesWithDampening(nums))
}

func createCombinationsWithOneMissing(line []int) [][]int {
	// Use Map to create combinations
	return std.Map(line, func(i int) []int {
		return append(append([]int(nil), line[:i]...), line[i+1:]...)
	})
}

func isSequence(line []int, crescent bool) bool {
	// Optimized sequence check logic
	for i := 1; i < len(line); i++ {
		if std.Abs(line[i-1]-line[i]) > 3 || line[i-1] == line[i] ||
			(crescent && line[i-1] > line[i]) || (!crescent && line[i-1] < line[i]) {
			return false
		}
	}
	return true
}

func NumberOfSafeLines(lines [][]int) int {
	return std.Reduce(lines, 0, func(acc int, line []int) int {
		if len(line) == 1 || isSequence(line, line[0] < line[1]) {
			return acc + 1
		}
		return acc
	})
}

func NumberOfSafeLinesWithDampening(lines [][]int) int {
	return std.Reduce(lines, 0, func(acc int, line []int) int {
		if len(line) == 1 || isSequence(line, line[0] < line[1]) {
			return acc + 1
		}

		dampened := createCombinationsWithOneMissing(line)
		for _, newLine := range dampened {
			if len(newLine) == 1 || isSequence(newLine, newLine[0] < newLine[1]) {
				return acc + 1
			}
		}

		return acc
	})
}

package main

import (
	"fmt"
	"sort"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/aoc"
)

func main() {
	input, err := aoc.GetInputFile("1")
	std.Error(err, 1)

	nums, err := std.ExtractNumbers(input)
	std.Error(err, 2)

	left, right := std.SliceSplit(nums)

	fmt.Println("Day 1 - Part 1:", calculateErrorSum(left, right))
	fmt.Println("Day 1 - Part 2:", calculateSimilarityScore(left, right))
}

func calculateErrorSum(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	return std.Reduce(left, 0, func(acc int, l int) int {
		return acc + std.Abs(l - right[acc])
	})
}

func calculateSimilarityScore(left, right []int) int {
	return std.Reduce(left, 0, func(acc int, l int) int {
		return acc + l * std.CountOccurrences(right, l)
	})
}

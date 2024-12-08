package main

import (
	"fmt"
	"os"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/aoc"
)

func main() {
	raw, err := aoc.GetInputFile("4")
	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	matrix := utils.ExtractByLine(raw)

	fmt.Println("Day 4 - Part 1:", utils.FindXmas(matrix))
	fmt.Println("Day 4 - Part 2:", utils.FindX_mas(matrix))
}


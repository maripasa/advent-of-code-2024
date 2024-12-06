package main

import (
	"advent_of_code_2024/utils/advent_of_code"
	"fmt"
	"os"
)

func main() {
  if err := aoc.ManageAdventOfCodePuzzles(2024); err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
  }
}

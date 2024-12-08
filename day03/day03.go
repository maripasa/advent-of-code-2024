package main

import (
	"fmt"
	"os"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/aoc"
)

func main() {
  input, err := aoc.GetInputFile("3")
  std.Error(err, 1)

  nums, err := utils.ExtractMuls(raw)
  if err != nil {
    fmt.Println(err)
    os.Exit(2)
  }

  result := 0
  for _, couple := range nums {
    result += couple[0] * couple[1]
  }

	fmt.Println("Day 3 - Part 1:", result)

  raw = utils.RemoveDontDo(raw)
    
  nums, err = utils.ExtractMuls(raw)
  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }

  result = 0
  for _, couple := range nums {
    result += couple[0] * couple[1]
  }

	fmt.Println("Day 3 - Part 2:", result)
  
}



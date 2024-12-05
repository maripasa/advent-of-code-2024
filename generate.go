package main

import (
	"advent_of_code_2024/utils"
	"fmt"
	"os"
)

func main() {
  err := utils.ManageAdventOfCodePuzzles(2024)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

package main

import (
	"fmt"
	"os"
)

func Day3() {
  raw, err := GetInputHTTP("3")
  if err != nil {
    os.Exit(1)
  }
  
  nums, err := ExtractMuls(raw)
  if err != nil {
    os.Exit(2)
  }

  result := 0
  for _, couple := range nums {
    result += couple[0] * couple[1]
  }

	fmt.Println("Day 3 - Part 1:", result)

  raw = RemoveDontDo(raw)
    
  nums, err = ExtractMuls(raw)
  if err != nil {
    os.Exit(2)
  }

  result = 0
  for _, couple := range nums {
    result += couple[0] * couple[1]
  }

	fmt.Println("Day 3 - Part 2:", result)
  
}



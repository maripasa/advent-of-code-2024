package main

import (
	"fmt"
	"strings"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/aoc"
)

func main() {
	input, err := aoc.GetInputFile("8")
	std.Error(err, 1)
  
	fmt.Println("Day 8 - Part 1:", len(calculateAntiNodes(parseInput(input))))
	fmt.Println("Day 8 - Part 2:", len(calculateAntiNodesConsideringHarmonics(parseInput(input))))
}

func parseInput(input string) ([]string, map[rune][]std.Vector2[int]) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	antennas := make(map[rune][]std.Vector2[int])
	for y, line := range lines {
		for x, value := range line {
			if value != '.' { antennas[value] = append(antennas[value], std.Vector2[int]{X: x, Y: y}) }
		}
	}
	return lines, antennas
}

func calculateAntiNodes(input []string, antennas map[rune][]std.Vector2[int]) std.Set[std.Vector2[int]] {
  antinodes := std.Set[std.Vector2[int]]{}

	for _, values := range antennas {
		generatingPairs := std.CombinationPairs(values)
		for _, pair := range generatingPairs {
      distance := pair[0].Subtract(pair[1])
			if aoc.IsWithinBounds(pair[0].Add(distance), input) { antinodes.Add(pair[0].Add(distance)) }
			if aoc.IsWithinBounds(pair[1].Subtract(distance), input) { antinodes.Add(pair[1].Subtract(distance)) }
		}
	}
	return antinodes
}

func calculateAntiNodesConsideringHarmonics(input []string, antennas map[rune][]std.Vector2[int]) std.Set[std.Vector2[int]] {
  antinodes := std.Set[std.Vector2[int]]{}

	for _, values := range antennas {
		pairs := std.CombinationPairs(values)
		for _, pair := range pairs {
      antinodes.Add(pair[0])
      antinodes.Add(pair[1])
      distance := pair[0].Subtract(pair[1])
      for i := 1 ; ; i++ {
			  antinode1 := pair[0].Add(distance.Scale(i))
        if !aoc.IsWithinBounds(antinode1, input) { break }
        antinodes.Add(antinode1)
      }
      for i := 1 ; ; i++ {
			  antinode2 := pair[1].Subtract(distance.Scale(i))
        if !aoc.IsWithinBounds(antinode2, input) { break }
        antinodes.Add(antinode2)
      }
		}
	}
	return antinodes
}

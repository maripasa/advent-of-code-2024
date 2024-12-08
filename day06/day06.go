package main

import (
	"fmt"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/aoc"
)

func main() {
	input, err := aoc.GetInputFile("6")
	std.Error(err, 1)

	room := std.ExtractByLine(input)

	start, startDirection, err := findGuard(room)
	std.Error(err, 1)
  
  visited := simulateGuard(room, start, startDirection)

	fmt.Println("Day 6 - Part 1: ", len(visited))
  fmt.Println("Day 6 - Part 2: ", countObstructionPositions(visited ,room, start, startDirection))
}

var directions = []std.Vector2[int]{
  {X: 0, Y: -1}, // Up
  {X: 1, Y: 0},  // Right
  {X: 0, Y: 1},  // Down
  {X: -1, Y: 0}, // Left
}

func findGuard(room []string) (std.Vector2[int], int, error) {
	for y, row := range room {
		for x, ch := range row {
			switch ch {
			case '^':
        return std.Vector2[int]{X: x, Y: y}, 0, nil
			case '>':
        return std.Vector2[int]{X: x, Y: y}, 1, nil
			case 'v':
        return std.Vector2[int]{X: x, Y: y}, 2, nil
			case '<':
        return std.Vector2[int]{X: x, Y: y}, 3, nil
			}
		}
	}
	return std.Vector2[int]{}, 0, fmt.Errorf("No guard starting position found in input.")
}

func simulateGuard(room []string, start std.Vector2[int], direction int) std.Set[std.Vector2[int]] {
	rows, cols := len(room), len(room[0])
	visited := make(std.Set[std.Vector2[int]])
	guard := start

	isWithinBounds := func(pos std.Vector2[int]) bool {
		return pos.Y >= 0 && pos.Y < rows && pos.X >= 0 && pos.X < cols
	}

	visited.Add(guard)
	for {
    next := std.Vector2[int].Add(guard, directions[direction])
		if !isWithinBounds(next) { break }
		if room[next.Y][next.X] != '#' {
			guard = next
			visited.Add(guard)
			continue
		}
		direction = (direction + 1) % 4
	}
	return visited
}

func causesLoop(room []string, start std.Vector2[int], direction int) bool {
	rows, cols := len(room), len(room[0])
	visited := make(std.Set[struct {
		std.Vector2[int]
		direction int
	}])

	isWithinBounds := func(pos std.Vector2[int]) bool {
		return pos.Y >= 0 && pos.X >= 0 && pos.Y < rows && pos.X < cols
	}

	guard := start

	for {
		state := struct {
		  std.Vector2[int]	
			direction int
		}{guard, direction}

		if _, found := visited[state]; found {
			return true
		}

		visited.Add(state)
    next := std.Vector2[int].Add(guard, directions[direction])
		if !isWithinBounds(next) { break }
		if room[next.Y][next.X] != '#' {
			guard = next
			continue
		}

		direction = (direction + 1) % 4
	}
	return false
}

func countObstructionPositions(visitedLocations std.Set[std.Vector2[int]], room []string, start std.Vector2[int], startDirection int) int {
	loopCausingPositions := 0

  for possible := range visitedLocations {
    if possible.X == start.X && possible.Y == start.Y { continue }
    original := room[possible.Y][possible.X] 
    room[possible.Y] = std.Replace(room[possible.Y], possible.X, "#")
    if causesLoop(room, start, startDirection) { loopCausingPositions++ }
    room[possible.Y] = std.Replace(room[possible.Y], possible.X, string(original))
  }
	return loopCausingPositions
}


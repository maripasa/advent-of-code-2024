package main

import (
	"fmt"
	"advent_of_code_2024/utils/std"
 	"advent_of_code_2024/utils/aoc"
  "strings"
  "strconv"
)

func main() {
	input, err := aoc.GetInputFile("7")
	std.Error(err, 1)

 	equations := parseInput(input)
	total := 0

	for _, eq := range equations {
		if canMatchTarget(eq.target, eq.values, []string{"*", "+"}) {
			total += eq.target
		}
	} 

	fmt.Println("Day 7 - Part 1:", total)

	total = 0

	for i, eq := range equations {
		if canMatchTarget(eq.target, eq.values, []string{"*", "+", "||"}) {
      fmt.Println("Correct: ", i)
			total += eq.target
		}
	} 

	fmt.Println("Day 7 - Part 2:", total)
}

func parseInput(input string) []struct {
	target int
	values []int
} {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	equations := []struct {
		target int
		values []int
	}{}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
    target, _ := strconv.Atoi(parts[0])
		valuesStr := strings.Split(parts[1], " ")
		values := []int{}
		for _, v := range valuesStr {
			num, _ := strconv.Atoi(v)
			values = append(values, num)
		}
		equations = append(equations, struct {
			target int
			values []int
		}{target, values})
	}
	return equations
}

func generateOperators(n int, operators []string) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	combinations := [][]string{{}}

	for i := 0; i < n; i++ {
		newCombinations := [][]string{}
		for _, comb := range combinations {
			for _, op := range operators {
				newCombinations = append(newCombinations, append(append([]string{}, comb...), op))
			}
		}
		combinations = newCombinations
	}
	return combinations
}

func evaluateExpression(values []int, operators []string) int {
	result := values[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+": result += values[i+1]
		case "*": result *= values[i+1]
    case "||":
			newNumber, _ := strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(values[i+1]))
      result = newNumber
		}
	}
	return result
}

func canMatchTarget(target int, values []int, operators []string) bool {
	if len(values) == 1 {
		return values[0] == target
	}

	operatorCombinations := generateOperators(len(values) - 1, operators)
	for _, ops := range operatorCombinations {
		if evaluateExpression(values, ops) == target {
			return true
		}
	}
	return false
}


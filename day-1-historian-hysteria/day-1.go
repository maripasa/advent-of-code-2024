package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sort"
)

//To use must export session
//
//Done with:
//export SESSION=<your_advent_of_code_session>

func main() {
	input, err := getInputHTTP()
	if err != nil {
		os.Exit(1)
	}
	nums, err := extractNumbers(input)

	if err != nil {
		os.Exit(2)
	}

	left, right := make([]int, 0, len(nums)/2), make([]int, 0, len(nums)/2)

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			left = append(left, nums[i])
			continue
		}
		right = append(right, nums[i])
	}

	fmt.Println("Part 1:", calculateErrorSum(left, right))
	fmt.Println("Part 2:", calculateSimilarityScore(left, right))
}

func calculateErrorSum(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	
	errorSum := 0
	for i := range left {
		errorSum += abs(left[i] - right[i])
	}
	return errorSum
}

func calculateSimilarityScore(left, right []int) int {
	score := 0
	for index := range left {
		score += left[index] * count(right, left[index])
	}
	return score
}

func extractNumbers(input string) ([]int, error) {
	nums := regexp.MustCompile(`\d+`).FindAllString(input, -1)
	extracted := make([]int, len(nums))
	
	for i, num := range nums {
		result, err := strconv.Atoi(num)
    		if err != nil {
			return extracted, err
    		}
    		extracted[i] = result
	}
	return extracted, nil 
}

func getInputHTTP() (string, error) {
	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/1/input", nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: os.Getenv("SESSION"),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func count(slice []int, val int) int {
	count := 0
	for index := range slice {
		if slice[index] == val {
			count++
		}
	}
	return count
}

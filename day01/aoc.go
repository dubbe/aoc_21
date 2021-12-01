package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getSolutionPart1(input []int) int {
	increases := 0
	lastValue := 0
	for _, value := range input {
		if value > lastValue {
			increases++
		}
		lastValue = value
	}
	return increases-1
}

func getSolutionPart2(input []int) int {
	increases := 0
	lastWindow := 0
	for i, _ := range input {
		if i < 3 {
			continue
		}
		currentWindow := input[i] + input[i-1] + input[i-2]
		if currentWindow > lastWindow {
			increases++
		}
		lastWindow = currentWindow
	}
	return increases
}

func parseInput(input string) ([]int, error) {
	var ints []int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input, err := parseInput(string(inputBytes))
	if err != nil {
		panic("couldn't parse input")
	}

	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}

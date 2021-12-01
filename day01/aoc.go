package main

import (
	"fmt"
	"os"
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

func main() {
	input, err := readInts("input.txt")
	if err != nil {
		panic("could not read file")
	}

	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}

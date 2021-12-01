package main

import (
	"fmt"
	"os"
)

func getSolutionPart1(input []int) (increases int) {
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}
	return
}

func getSolutionPart2(input []int) (increases int) {
	lastWindow := input[2] + input[1] + input[0]
	for i := 3; i < len(input); i++ {
		currentWindow := input[i] + input[i-1] + input[i-2]
		if currentWindow > lastWindow {
			increases++
		}
		lastWindow = currentWindow
	}
	return
}

func getSolutionPart2Tobe(input []int) (increases int) {
	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			increases++
		}
	}
	return
}

func main() {
	if input, err := readInts("input.txt"); err != nil {
		panic("could not read file")
	} else {
		if os.Getenv("part") == "part2" {
			fmt.Println(getSolutionPart2(input))
		} else {
			fmt.Println(getSolutionPart1(input))
		}
	}
}

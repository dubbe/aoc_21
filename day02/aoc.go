package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSolutionPart1(input []string) int {
	depth := 0
	horizontal := 0
	for _, v := range input {
		operations := strings.Split(v, " ")
		i, err := strconv.Atoi(operations[1])
		if err != nil {
			panic("error")
		}
		switch operations[0] {
		case "forward":
			horizontal += i
		case "down":
			depth += i
		case "up":
			depth -= i
		}

	}

	return depth * horizontal
}

func getSolutionPart2(input []string) (increases int) {
	aim := 0
	horizontal := 0
	depth := 0
	for _, v := range input {
		operations := strings.Split(v, " ")
		i, err := strconv.Atoi(operations[1])
		if err != nil {
			panic("error")
		}
		switch operations[0] {
		case "forward":
			horizontal += i
			depth += aim*i
		case "down":
			aim += i
		case "up":
			aim -= i
		}

	}

	return depth * horizontal
}

func main() {
	if input, err := readStrings("input.txt"); err != nil {
		panic("could not read file")
	} else {
		if os.Getenv("part") == "part2" {
			fmt.Println(getSolutionPart2(input))
		} else {
			fmt.Println(getSolutionPart1(input))
		}
	}
}

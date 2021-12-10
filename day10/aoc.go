package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseLine(input string) string {

	validChunks := []string{"()", "[]", "{}", "<>"}

	prevLen := len(input)
	for {
		for _, chunk := range validChunks {
			input = strings.ReplaceAll(input, chunk, "")			
		}
		if prevLen == len(input) {
			break
		}
		prevLen = len(input)
	}

	for _, chunk := range []string{"(", "[", "{", "<"} {
		input = strings.ReplaceAll(input, chunk, "")			
	}
	if len(input) == 0 {
		return ""
	}

	return string(input[0])
}

func parseLine2(input string) string {

	validChunks := []string{"()", "[]", "{}", "<>"}

	prevLen := len(input)
	for {
		for _, chunk := range validChunks {
			input = strings.ReplaceAll(input, chunk, "")			
		}
		if prevLen == len(input) {
			break
		}

		prevLen = len(input)
	}

	missingClosing := input
	for _, chunk := range []string{"(", "[", "{", "<"} {
		input = strings.ReplaceAll(input, chunk, "")			
	}

	if len(input) != 0 {
		return ""
	}

	missingClosing = Reverse(missingClosing)

	missingClosing = strings.ReplaceAll(missingClosing, "(", ")")	
	missingClosing = strings.ReplaceAll(missingClosing, "{", "}")	
	missingClosing = strings.ReplaceAll(missingClosing, "[", "]")	
	missingClosing = strings.ReplaceAll(missingClosing, "<", ">")	
	return missingClosing
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func calculateScore(s string) int {
	sum := 0

	for _, char := range s {
		sum *= 5
		switch char {
			case ')': 
			sum += 1
			case ']': 
			sum += 2
			case '}': 
			sum += 3
			case '>':
			sum += 4
		}
	}

	return sum
}

func getSolutionPart1(input []string) int{
	sum := 0
	for _, i := range input {
		char := parseLine(i)
		switch char {
			case ")": 
			sum += 3
			case "]": 
			sum += 57
			case "}": 
			sum += 1197
			case ">":
			sum += 25137
		}
	}
	return sum
}

func getSolutionPart2(input []string) int {

	scores := []int{}
	for _, i := range input {
		missing := parseLine2(i)
		if missing == "" {
			continue
		}
		scores = append(scores, calculateScore(missing))
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	if input, err := readInput("input.txt"); err != nil {
		panic("could not read file")
	} else {
		if os.Getenv("part") == "part2" {
			fmt.Println(getSolutionPart2(input))
		} else {
			fmt.Println(getSolutionPart1(input))
		}
	}
}

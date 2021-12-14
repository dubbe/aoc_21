package main

import (
	"fmt"
	"os"
)

func parseInsertionRules(input []string) map[string]rune {
	rules := map[string]rune{}
	for _, row := range input {
		var pair string
		var element rune
		fmt.Sscanf(row, "%s -> %c", &pair, &element)
		rules[pair] = element
	}
	return rules
}

func getPairs(instr string) map[string]uint64 {
	pairs := map[string]uint64{}
	for i := range instr {
		if i == len(instr)-1 {
			break
		}
		pairs[instr[i:i+2]]++
	}

	return pairs
}

func solve(pairs map[string]uint64, rules map[string]rune, iterations int) map[string]uint64 {
	for i := 0; i < iterations; i++ {
		next := map[string]uint64{}

		for pair, count := range pairs {
			next[string(pair[0])+string(rules[pair])] += count
			next[string(rules[pair])+string(pair[1])] += count
		}
		pairs = next
	}

	counted := map[string]uint64{}
	for pair, c := range pairs {
		counted[string(pair[0])] += c
	}

	return counted
}

func maxMin(m map[string]uint64) (uint64, uint64) {
	heighest := uint64(0)
	lowest := ^uint64(0)
	for _, v := range m {
		if v > heighest {
			heighest = v
		}
		if v < lowest {
			lowest = v
		}
	}
	return heighest, lowest
}

func getSolutionPart1(input []string) uint64 {
	instructions := input[0]
	pairs := getPairs(instructions)
	rules := parseInsertionRules(input[2:])
	polymer := solve(pairs, rules, 10)
	polymer[string(instructions[len(instructions)-1])]++

	max, min := maxMin(polymer)
	return max - min
}

func getSolutionPart2(input []string) uint64 {
	instructions := input[0]
	pairs := getPairs(instructions)
	rules := parseInsertionRules(input[2:])
	polymer := solve(pairs, rules, 40)
	polymer[string(instructions[len(instructions)-1])]++

	max, min := maxMin(polymer)
	return max - min
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

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string

func init() {
	input, _ = readInput("input-test.txt")
}

func TestParseLine(t *testing.T) {
	line := "{([(<{}[<>[]}>{[]{[(<()>"
	assert.Equal(t, "}", parseLine(line))

	line = "[({(<(())[]>[[{[]{<()<>>"
	assert.Equal(t, "", parseLine(line))
}

func TestParseLine2(t *testing.T) {
	line := "[({(<(())[]>[[{[]{<()<>>"
	assert.Equal(t, "}}]])})]", parseLine2(line))
}

func TestCalculateScore(t *testing.T) {
	line := "])}>"
	assert.Equal(t, 294, calculateScore(line))

	line = "}}]])})]"
	assert.Equal(t, 288957, calculateScore(line))

	line = ")}>]})"
	assert.Equal(t, 5566, calculateScore(line))

	line = "}}>}>))))"
	assert.Equal(t, 1480781, calculateScore(line))
}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 26397

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 288957

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		realInput, _ := readInput("input.txt")
		getSolutionPart1(realInput)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		realInput, _ := readInput("input.txt")
		getSolutionPart2(realInput)
	}
}

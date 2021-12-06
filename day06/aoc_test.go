package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input string
func init() {
	input, _ = readInput("input-test.txt")
}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 5934

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 26984457539

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
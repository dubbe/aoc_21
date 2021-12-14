package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string

func init() {
	input, _ = readInput("input-test.txt")
}
func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := uint64(1588)
	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart1actual(t *testing.T) {
	realInput, _ := readInput("input.txt")
	expectedSolution := uint64(3406)
	actualSolution := getSolutionPart1(realInput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2actual(t *testing.T) {
	realInput, _ := readInput("input.txt")
	expectedSolution := uint64(3941782230241)
	actualSolution := getSolutionPart2(realInput)
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

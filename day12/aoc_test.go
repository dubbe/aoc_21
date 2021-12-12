package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string
func init() {
	input, _ = readInput("input-test.txt")
}

func TestSmallExample(t *testing.T) {
	smallInput := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
	expectedSolution := 10
	actualSolution := getSolutionPart1(smallInput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestSmallExample2(t *testing.T) {
	smallInput := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
	expectedSolution := 36
	actualSolution := getSolutionPart2(smallInput)
	assert.Equal(t, expectedSolution, actualSolution)
}
func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 226
	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 3509

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
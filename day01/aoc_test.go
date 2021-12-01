package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAOC_getSolutionPart1(t *testing.T) {
	input, err := readInts("input_test.txt")
	assert.Nil(t, err)

	expectedSolution := 7

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	input, err := readInts("input_test.txt")
	assert.Nil(t, err)

	expectedSolution := 5

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2Tobe(t *testing.T) {
	input, err := readInts("input_test.txt")
	assert.Nil(t, err)

	expectedSolution := 5

	actualSolution := getSolutionPart2Tobe(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input, _ := readInts("input.txt")
		getSolutionPart1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input, _ := readInts("input.txt")
		getSolutionPart2(input)
	}
}

func BenchmarkPart2Tobe(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input, _ := readInts("input.txt")
		getSolutionPart2Tobe(input)
	}
}
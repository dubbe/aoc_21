package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAOC_getSolutionPart1(t *testing.T) {
	input, err := readStrings("input-test.txt")
	assert.Nil(t, err)

	expectedSolution := 150

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	input, err := readStrings("input-test.txt")
	assert.Nil(t, err)

	expectedSolution := 900

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

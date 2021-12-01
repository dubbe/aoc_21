package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAOC_getSolutionPart1(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("input_test.txt")
	assert.Nil(t, err)
	
	input, err := parseInput(string(inputBytes))
	assert.Nil(t, err)

	expectedSolution := 7

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("input_test.txt")
	assert.Nil(t, err)
	
	input, err := parseInput(string(inputBytes))
	assert.Nil(t, err)

	expectedSolution := 5

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

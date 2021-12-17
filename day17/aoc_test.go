package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string

func TestTestTrajectory(t *testing.T) {
	tr, _ := testTrajectory(7, 2, 20, 30, -5, -10)
	assert.True(t, tr)

	tr, _ = testTrajectory(6, 3, 20, 30, -5, -10)
	assert.True(t, tr)

	tr, _ = testTrajectory(9, 0, 20, 30, -5, -10)
	assert.True(t, tr)

	tr, _ = testTrajectory(17, -4, 20, 30, -5, -10)
	assert.False(t, tr)
}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 45
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 112
	actualSolution := getSolutionPart2(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getSolutionPart1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getSolutionPart2(input)
	}
}

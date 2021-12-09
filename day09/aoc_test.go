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
	expectedSolution := 15

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 1134

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestGetBasinTopRight(t *testing.T) {
	pane := parsePane(input)
	foundPoints := []Point{}
	sum := getBasin(Point{x:9, y:0}, pane, &foundPoints)
	assert.Equal(t, 9, sum)
}

func TestGetBasinTopLeft(t *testing.T) {
	pane := parsePane(input)
	foundPoints := []Point{}
	sum := getBasin(Point{x:1, y:0}, pane, &foundPoints)
	assert.Equal(t, 3, sum)
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
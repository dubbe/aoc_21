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
	line := parseLines("1,1 -> 1,3")
	assert.Len(t, line.points, 3)

	line = parseLines("9,7 -> 7,7")
	assert.Len(t, line.points, 3)
}

func TestParseLines2(t *testing.T) {
	line := parseLines("9,7 -> 7,7")
	assert.Len(t, line.points, 3)

	line = parseLines("1,1 -> 3,3")
	assert.Len(t, line.points, 3)
	assert.Equal(t, line.points[0], Point{x: 1, y: 1})
	assert.Equal(t, line.points[1], Point{x: 2, y: 2})
	assert.Equal(t, line.points[2], Point{x: 3, y: 3})
}
func TestParseLines3(t *testing.T) {
	line := parseLines("9,7 -> 7,9")
	assert.Len(t, line.points, 3)
	assert.Equal(t, Point{x: 9, y: 7}, line.points[0])
	assert.Equal(t, Point{x: 8, y: 8}, line.points[1])
	assert.Equal(t, Point{x: 7, y: 9}, line.points[2])
}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 5

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 12

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

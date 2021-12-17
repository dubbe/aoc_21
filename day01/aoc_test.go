package main

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testInput string
var ints []int

func init() {
	ints = []int{}
	inp := strings.Split(input, "\n")
	for _, s := range inp {
    i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 7

	actualSolution := getSolutionPart1(ints)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 5

	actualSolution := getSolutionPart2(ints)
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

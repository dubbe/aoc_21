package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string

func init() {
	input, _ = readInput("input-test.txt")
}

func TestParseInput(t *testing.T) {
	str := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	signalPatterns, _ := parseLine(str)

	parsedSignalPatterns := parseSignalPatterns(signalPatterns)
	assert.Equal(t, SortStringByCharacter("cagedb"), parsedSignalPatterns[0])
	assert.Equal(t, SortStringByCharacter("ab"), parsedSignalPatterns[1])
	assert.Equal(t, SortStringByCharacter("gcdfa"), parsedSignalPatterns[2])
	assert.Equal(t, SortStringByCharacter("fbcad"), parsedSignalPatterns[3])
	assert.Equal(t, SortStringByCharacter("eafb"), parsedSignalPatterns[4])
	assert.Equal(t, SortStringByCharacter("cdfbe"), parsedSignalPatterns[5])
	assert.Equal(t, SortStringByCharacter("cdfgeb"), parsedSignalPatterns[6])
	assert.Equal(t, SortStringByCharacter("dab"), parsedSignalPatterns[7])
	assert.Equal(t, SortStringByCharacter("acedgfb"), parsedSignalPatterns[8])
	assert.Equal(t, SortStringByCharacter("cefabd"), parsedSignalPatterns[9])

}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 26

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 61229

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

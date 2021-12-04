package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	testInput := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
 		"8  2 23  4 24",
		"21  9 14 16  7",
 		"6 10  3 18  5",
		"1 12 20 15 19",
	}
	numbers, boards := parseInput(testInput)
	assert.Len(t, numbers, 27)
	assert.Len(t, boards, 1)
}

func TestParseBoard(t *testing.T) {
	testInput := []string{
		"22 13 17 11  0",
 		"8  2 23  4 24",
		"21  9 14 16  7",
 		"6 10  3 18  5",
		"1 12 20 15 19",
	}
	board := parseBoard(testInput)
	assert.NotNil(t, board)
	horizontalRow := []string{"22", "13", "17", "11", "0"}
	verticalRow := []string{"22", "8", "21", "6", "1"}
	assert.Equal(t, horizontalRow, board.horizontalRows[0].numbers)
	assert.Equal(t, verticalRow, board.verticalRows[0].numbers)
}

func TestAOC_getSolutionPart1(t *testing.T) {
	input, err := readStrings("input-test.txt")
	assert.Nil(t, err)

	expectedSolution := 4512

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	input, err := readStrings("input-test.txt")
	assert.Nil(t, err)

	expectedSolution := 1924

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

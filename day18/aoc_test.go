package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string


func TestAddition(t *testing.T) {
	node1 :="[[[[4,3],4],4],[7,[[8,4],9]]]"
	node2 := "[1,1]"
	added := addition(node1, node2)
	assert.Equal(t, "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", added)
}


func TestFindFirstToExplode(t *testing.T) {
	snailfish := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	found, index, str := findFirstToExplode(snailfish)
	assert.Equal(t, true, found)
	assert.Equal(t, []int{4,9}, index)
	assert.Equal(t, "[4,3]", str)

	snailfish = "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]"
	found, index, str = findFirstToExplode(snailfish)
	assert.Equal(t, true, found)
	assert.Equal(t, []int{16,21}, index)
	assert.Equal(t, "[8,4]", str)
}

func TestFindFirstToSplit(t *testing.T) {
	snailfish := "[[[[0,7],4],[15,[0,13]]],[1,1]]"
	found, index, i := findFirstToSplit(snailfish)
	assert.Equal(t, true, found)
	assert.Equal(t, []int{13,15}, index)
	assert.Equal(t, 15, i)
}

func TestXY(t *testing.T) {
	snail := "[18,99]"
	x,y := getXY(snail)
	assert.Equal(t, 18, x)
	assert.Equal(t, 99, y)
}

func TestExplode(t *testing.T) {
	example :=  "[[[[[9,8],1],2],3],4]"
	expectedSolution := "[[[[0,9],2],3],4]"
	actual, _ := explode(example)
	assert.Equal(t, expectedSolution, actual)
}

func TestSpit(t *testing.T) {
	example := "[[[[0,7],4],[15,[0,13]]],[1,1]]"
	expectedSolution := "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"
	split1, _ := split(example)
	assert.Equal(t, expectedSolution, split1)
}


func TestDoStuff(t *testing.T) {
	actual :=  "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	expectedSolution :="[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	assert.Equal(t, expectedSolution, doStuff(actual))
}

func TestCalculateMagnitude(t *testing.T) {
	
	actual :=  "[[1,2],[[3,4],5]]"
	expectedSolution := 143
	assert.Equal(t, expectedSolution, calculateMagnitude(actual))
}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 4140
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 3993
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

package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := int(474140)
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

// func TestSwitchCube(t *testing.T) {
// 	x := []int{10, 12}
// 	y := []int{10, 12}
// 	z := []int{10, 12}
// 	cubes := map[Cube]bool{}
// 	cubes = reboot(true, x, y, z, cubes)
// 	assert.Len(t, cubes, 27)

// 	x = []int{11, 13}
// 	y = []int{11, 13}
// 	z = []int{11, 13}
// 	cubes = reboot(true, x, y, z, cubes)
// 	assert.Len(t, cubes, 27+19)

// 	x = []int{9, 11}
// 	y = []int{9, 11}
// 	z = []int{9, 11}
// 	cubes = reboot(false, x, y, z, cubes)
// 	assert.Len(t, cubes, int(27+19-8))

// 	x = []int{10, 10}
// 	y = []int{10, 10}
// 	z = []int{10, 10}
// 	cubes = reboot(true, x, y, z, cubes)
// 	assert.Equal(t, cubes, int(39))
// }

func TestSimpleExample(t *testing.T) {
	str := "on x=10..12,y=10..12,z=10..12\non x=11..13,y=11..13,z=11..13\noff x=9..11,y=9..11,z=9..11\non x=10..10,y=10..10,z=10..10"

	actualSolution := getSolutionPart1(str)
	assert.Equal(t, int(39), actualSolution)

}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := int64(2758514936282235)
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

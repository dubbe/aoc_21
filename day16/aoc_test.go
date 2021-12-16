package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string

// func TestParseHexToBit(t *testing.T) {
// 	actual := parseHexToBit("D2FE28")
// 	expected := "110100101111111000101000"
// 	assert.Equal(t, expected, actual)
// 	actualInt, left := parsePacket(actual)
// 	assert.Equal(t, int64(6), actualInt)
// 	assert.Equal(t, "000", left)
// }

// func TestParsePacket(t *testing.T) {
// 	bit := parseHexToBit("EE00D40C823060")
// 	actual, _ := parsePacket(bit)
// 	expected := "110100101111111000101000"
// 	assert.Equal(t, expected, actual)
// }


func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := int64(16)
	actualSolution := getSolutionPart1("8A004A801A8002F478")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(12)
	actualSolution = getSolutionPart1("620080001611562C8802118E34")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(23)
	actualSolution = getSolutionPart1("C0015000016115A2E0802F182340")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(31)
	actualSolution = getSolutionPart1("A0016C880162017C3686B18A3D4780")
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := int64(3)
	actualSolution := getSolutionPart2("C200B40A82")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(54)
	actualSolution = getSolutionPart2("04005AC33890")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(7)
	actualSolution = getSolutionPart2("880086C3E88112")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(9)
	actualSolution = getSolutionPart2("CE00C43D881120")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(1)
	actualSolution = getSolutionPart2("D8005AC2A8F0")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(0)
	actualSolution = getSolutionPart2("F600BC2D8F")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(0)
	actualSolution = getSolutionPart2("9C005AC2F8F0")
	assert.Equal(t, expectedSolution, actualSolution)

	expectedSolution = int64(1)
	actualSolution = getSolutionPart2("9C0141080250320F1802104A08")
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

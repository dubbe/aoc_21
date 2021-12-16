package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
)

//go:embed input.txt
var input string



func parseHexToBit(input string) string  {
	conversion := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	bits := ""

	for _, r := range input {
		bits += conversion[r]
	}

	return bits
}

func parsePacket(input string) (int64, string, int64) {
	version, _ := strconv.ParseInt(input[0:3],2,64)
	packetType, _ := strconv.ParseInt(input[3:6],2,64)
	value := int64(0)
	left := ""
	switch packetType {
	case 0: 
		v, l, values := parseOperator(input[6:])
		left = l
		version += v
		value = sum(values)
	case 1:
		v, l, values := parseOperator(input[6:])
		left = l
		version += v
		value = product(values)
	case 2:
		v, l, values := parseOperator(input[6:])
		left = l
		version += v
		value = min(values)
	case 3:
		v, l, values := parseOperator(input[6:])
		left = l
		version += v
		value = max(values)
	case 4: 
		left, value = parseLiteral(input[6:])
	case 5:
		v, l, values := parseOperator(input[6:])
		left = l
		version += v
		if values[0] > values[1] {
			value = 1
		} else {
			value = 0
		}
	case 6:
		v, l, values := parseOperator(input[6:])
		left = l
		version += v
		if values[0] < values[1] {
			value = 1
		} else {
			value = 0
		}
	case 7:
		v, l, values := parseOperator(input[6:])
		left = l
		version += v
		if values[0] == values[1] {
			value = 1
		} else {
			value = 0
		}
	default:
		v, l, _ := parseOperator(input[6:])
		left = l
		version += v
	}

	return version, left, value
}

func sum(array []int64) int64 {  
	result := int64(0)  
	for _, v := range array {  
	 result += v  
	}  
	return result  
 }  

 func product(array []int64) int64 {  
	result := array[0]  
	for _, v := range array[1:] {  
	 result *= v  
	}  
	return result  
 } 

 func min(array []int64) int64 {  
	result := array[0]    
	for _, v := range array[1:] {  
	 if v < result {
		 result = v
	 }
	}  
	return result  
 } 

 func max(array []int64) int64 {  
	result := array[0]    
	for _, v := range array[1:] {  
	 if v > result {
		 result = v
	 }
	}  
	return result  
 } 

func parseOperator(input string) (int64, string, []int64) {
	lengthTypeId := input[0:1]
	version := int64(0)
	left := ""
	values := []int64{}
	switch lengthTypeId {
	case "0":
		subPacketLength, _ :=  strconv.ParseInt(input[1:16],2,64)
		// fmt.Printf("subPacketLength: %d\n", subPacketLength)
		// fmt.Printf("what: %s\n", input[16:])
		newInput := input[16:]
		readLength := 0
		
		for {
			v, inputLeft, value := parsePacket(newInput)
			values = append(values, value)
			version += v
			readLength += len(newInput)-len(inputLeft)
			newInput = inputLeft
			if readLength >= int(subPacketLength) {
				left = inputLeft
				break
			}
		}
	case "1":
		subPacketLength, _ :=  strconv.ParseInt(input[1:12],2,64)
		newInput := input[12:]
		i := 0
		for {
			v, inp, value := parsePacket(newInput)
			values = append(values, value)
			version += v
			newInput = inp

			i++

			if i == int(subPacketLength) {
				left = inp
				break;
			}
			
			
		}
	}

	return version, left, values
}

func parseLiteral(input string) (string, int64) {
	i:=0 
	value := ""
	for {
		start := input[i:i+1]
		value += input[i+1:i+5]

		i+=5
		if start == "0" {
			break;
		}
		
	}

	literalValue, _ := strconv.ParseInt(value,2,64)

	return input[i:], literalValue
}

func getSolutionPart1(input string) int64 {
	bit := parseHexToBit(input)
	actualInt, _, _ := parsePacket(bit)
	return actualInt
}

func getSolutionPart2(input string) int64 {
	bit := parseHexToBit(input)
	_, _, sum := parsePacket(bit)
	return sum
}

func main() {

	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}

}

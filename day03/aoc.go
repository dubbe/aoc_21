package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSolutionPart1(input []string) int {
	
	flippedStrings := map[int]string{}

	for x, values := range input {
		for i, v := range values {
			if x == 0 {
				flippedStrings[i] = string(v)
			} else {
				flippedStrings[i] = flippedStrings[i] + string(v)
			}
		}
	}

	gammaString := ""
	episolonString := ""
	n := 0
	for n < len(flippedStrings) {
		bit := flippedStrings[n]
		ones := strings.Count(bit, "1")
		zeroes := strings.Count(bit, "0")

		if ones > zeroes {
			gammaString += "1"
			episolonString += "0"
		} else {
			episolonString += "1"
			gammaString += "0"
		}
		n++
	}

	gamma, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		fmt.Println(err)
	} 

	episolon, err := strconv.ParseInt(episolonString, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	
	return int(gamma * episolon)
}

func getSolutionPart2(input []string) (increases int) {
	flippedStrings := map[int]string{}

	for x, values := range input {
		for i, v := range values {
			if x == 0 {
				flippedStrings[i] = string(v)
			} else {
				flippedStrings[i] = flippedStrings[i] + string(v)
			}
		}
	}

	oxygen := input
	
	for i, _ := range input[0] {
		newList := []string{}
		most := getMostInColumn(oxygen, i)
		
		for _, in := range oxygen {
			if most == string(in[i]) {
				newList = append(newList, in)
			}
		}
		oxygen = newList
	}

	co2 := input
	
	for i, _ := range input[0] {
		most := getMostInColumn(co2, i)
		newList := []string{}

		for _, in := range co2 {
			if most != string(in[i]) {
				newList = append(newList, in)
			}
		}
		co2 = newList
		if len(co2) == 1 {
			break;
		}
	}

	o, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	c, err := strconv.ParseInt(co2[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	
	return int(o * c)

}

func getMostInColumn(inputs []string, position int) string {
	zeroes := 0
	ones := 0
	for _, input := range inputs {
		if input[position] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if ones >= zeroes {
		return "1"
	}
	return "0"
}

func main() {
	if input, err := readStrings("input.txt"); err != nil {
		panic("could not read file")
	} else {
		if os.Getenv("part") == "part2" {
			fmt.Println(getSolutionPart2(input))
		} else {
			fmt.Println(getSolutionPart1(input))
		}
	}
}

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
		fmt.Printf("%s \n", bit)
		ones := strings.Count(bit, "1")
		zeroes := strings.Count(bit, "0")

		fmt.Printf("ones: %d, zeroes: %d\n", ones, zeroes)

		if ones > zeroes {
			gammaString += "1"
			episolonString += "0"
		} else {
			episolonString += "1"
			gammaString += "0"
		}
		n++
	}
	fmt.Printf("%s, %s \n", gammaString, episolonString)
	gamma, err := strconv.ParseInt(gammaString, 2, 64)
	
	if err != nil {
		fmt.Println(err)
	} 

	episolon, err := strconv.ParseInt(episolonString, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%d, %d\n", gamma, episolon)
	result := gamma * episolon
	return int(result)
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
		most := getMostInColumn(oxygen, i)
		fmt.Printf("most: %s \n", most)
		newList := []string{}
		fmt.Printf("input[%d]: %s, most: %s \n", i, input[i], most)
		for _, in := range oxygen {
			fmt.Printf("most: %s, in: %s, i: %d, in[%d]: %s\n", most, in, i, i, string(in[i]))
			if most == string(in[i]) {
				newList = append(newList, in)
			}
		}
		oxygen = newList
	}

	fmt.Printf("oxygen: %v \n", oxygen)


	co2 := input
	
	for i, _ := range input[0] {
		most := getMostInColumn(co2, i)
		fmt.Printf("most: %s \n", most)
		newList := []string{}
		fmt.Printf("input[%d]: %s, most: %s \n", i, input[i], most)
		for _, in := range co2 {
			fmt.Printf("most: %s, in: %s, i: %d, in[%d]: %s\n", most, in, i, i, string(in[i]))
			if most != string(in[i]) {
				newList = append(newList, in)
			}
		}
		co2 = newList
		if len(co2) == 1 {
			break;
		}
	}

	fmt.Printf("oxygen: %v \n", oxygen)
	fmt.Printf("co2: %v \n", co2)

	// fmt.Printf("mostInColumn: %s \n", mostInColumn)

	o, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	c, err := strconv.ParseInt(co2[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%d, %d\n", o, c)
	result := o * c
	return int(result)

}

func getMostInColumn(inputs []string, position int) string {
	zeroes := 0
	ones := 0
	for _, input := range inputs {

		// fmt.Printf("input: %s, input[%d]: %s\n", input, position, string(input[position]))

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

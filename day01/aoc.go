package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getSolutionPart1(input []int) (increases int) {
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}
	return
}

func getSolutionPart2(input []int) (increases int) {
	for i := 3; i < len(input); i++ { // kudos tobe!
		if input[i] > input[i-3] {
			increases++
		}
	}
	return
}

func main() {
	inp := strings.Split(input, "\n")
	ints := []int{}
	for i, s := range inp {
    ints[i], _ = strconv.Atoi(s)
	}
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(ints))
	} else {
		fmt.Println(getSolutionPart1(ints))
	}

}

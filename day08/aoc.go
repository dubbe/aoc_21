package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func parseLine(input string) (map[int]string, []string)  {
	inp := strings.Split(input, " | ")
	signalPatterns := strings.Split(inp[0], " ")
	sort.Slice(signalPatterns, func(i, j int) bool {
		return len(signalPatterns[i]) < len(signalPatterns[j])
	})
	signalPatternsReturn := map[int]string{}
	for i, sp := range signalPatterns {
		signalPatternsReturn[i] = SortStringByCharacter(sp)
	}

	outputValues := strings.Split(inp[1], " ")
	newOutputValues := []string{}
	for _, s := range outputValues {
		newOutputValues = append(newOutputValues, SortStringByCharacter(s))
	}
	return signalPatternsReturn, newOutputValues
}

func diff(a, b string) []string {
	temp := map[rune]int{}
	for _, s := range a {
			temp[s]++
	}
	for _, s := range b {
			temp[s]--
	}

	var result []string
	for s, v := range temp {
			if v != 0 {
					result = append(result, string(s))
			}
	}
	return result
}

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
      var r []rune
      for _, runeValue := range s {
              r = append(r, runeValue)
      }
      return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

func parseSignalPatterns(signalPatterns map[int]string) map[int]string {
	var c, f string
	
	ints := map[int]string{}
	ints[1] = signalPatterns[0]
	ints[7] = signalPatterns[1]
	ints[4] = signalPatterns[2]
	ints[8] = signalPatterns[9]

	letterCount := 0
	for _, s := range signalPatterns {
		if strings.Contains(s, string(ints[1][0])) {
			letterCount++
		}
	}
	if letterCount == 8 {
		c = string(ints[1][0])
		f = string(ints[1][1])
	} else {
		c = string(ints[1][1])
		f = string(ints[1][0])
	}

	for _, s := range signalPatterns {
		
		if !strings.Contains(s, c) {
			if len(s) == 5 {
				ints[5] = s
			} else {
				ints[6] = s
			}
		}
	}

	for _, s := range signalPatterns {
		if !strings.Contains(s, f) {
			ints[2] = s
		}
	}
	for _, s := range signalPatterns {
		
		if len(s) == 6 && s != ints[6] {
			diff := diff(s, ints[5])
			if len(diff) == 1 && diff[0] == c {
				ints[9] = s
			} else {
				ints[0] = s
			}
		}
	}

	for _, s := range signalPatterns {
		found := false
		for _, i := range ints {
			if s == i {
				found = true
				continue
			}
			
		}
		if !found {
			ints[3] = s
		}
	}

	return ints
}

func twistIt(signalPatterns map[int]string) (newSignalPatterns map[string]int) {
	newSignalPatterns = map[string]int{}
	for i, s := range signalPatterns {
		newSignalPatterns[s] = i
	}
	return newSignalPatterns
}

// func parseOutputValues(outputValues map[int]string, parsedSignalPatterns map[int]string), map[string]int {

// }

func getSolutionPart1(input []string) int{
	sum := 0
	for _, line := range input {
		_, outputValues := parseLine(line)
		for _, ov := range outputValues {
			if len(ov) == 2 || len(ov) == 4 || len(ov) == 3 || len(ov) == 7 {
				sum++
			}
		}
	}

	return sum
}

func getSolutionPart2(input []string) int {
	sum := 0
	for _, line := range input {
		signalPatterns, outputValues := parseLine(line)
		parsedSignalPatterns := twistIt(parseSignalPatterns(signalPatterns))
		parsedSum := ""
		for _, ov := range outputValues {
			parsedSum = fmt.Sprintf("%s%d", parsedSum, parsedSignalPatterns[ov])
		}
		iSum, _ := strconv.Atoi(parsedSum) 
		sum += iSum
	}

	return sum

}

func main() {
	if input, err := readInput("input.txt"); err != nil {
		panic("could not read file")
	} else {
		if os.Getenv("part") == "part2" {
			fmt.Println(getSolutionPart2(input))
		} else {
			fmt.Println(getSolutionPart1(input))
		}
	}
}

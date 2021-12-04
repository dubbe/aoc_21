package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	numbers []string
	found int
}

type Board struct{
	foundNumbers []string
	verticalRows [5]Row
	horizontalRows [5]Row
	won bool 
}

func parseInput(input []string) (numbers []string, boards []Board) {
	numbers = strings.Split(input[0], ",")
	boards = []Board{}

	boardInput := []string{}
	for i:=2; i<len(input); i++ {
		if input[i] == "" {
			boards = append(boards, parseBoard(boardInput))
			boardInput = []string{}
			continue
		}
		boardInput = append(boardInput, input[i])
		
	}
	boards = append(boards, parseBoard(boardInput))
	boardInput = []string{}
	return numbers, boards
}

func parseBoard(input []string) Board {
	board := Board{
		verticalRows: [5]Row{}, 
		horizontalRows: [5]Row{}, 
	}
	for x, inputRow := range(input) {
		
		numbers := strings.Fields(inputRow)
		for y, n := range numbers {
			board.horizontalRows[x].numbers = append(board.horizontalRows[x].numbers, n)
			board.verticalRows[y].numbers = append(board.verticalRows[y].numbers, n)
		}
	}

	return board
}

func contains(s []string, e string) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}

func markBoard(board *Board, num string) (won bool) {
	won = false
	found := false
	
	for y, horRows := range board.horizontalRows {
		
		if contains(horRows.numbers, num) {
			found = true
			board.foundNumbers = append(board.foundNumbers, num)
			board.horizontalRows[y].found ++
			if board.horizontalRows[y].found == 5 {
				won = true
				board.won = true
			}
		}
	}
	if found == true {
		for y, verRows := range board.verticalRows {
			if contains(verRows.numbers, num) {
				board.verticalRows[y].found ++
				if board.verticalRows[y].found == 5 {
					won = true
					board.won = true
				}
			}
		}
	}
	return
}

func getSum(board Board, num string) int {
	unmarkedNumbersSum := 0
	for _, rows := range board.horizontalRows {
		for _, n := range rows.numbers {
			if !contains(board.foundNumbers, n) {
				convertedN, _ := strconv.Atoi(n)
				unmarkedNumbersSum += convertedN
			}
		}
	}

	convertedNum, _ := strconv.Atoi(num)
	return (unmarkedNumbersSum-convertedNum) * convertedNum
}

func getSolutionPart1(input []string) int {
	numbers, boards := parseInput(input)

	for _, num := range numbers {
		for i, board := range boards {
			won := markBoard(&boards[i], num)
			if won == true {
				return getSum(board, num)
			}	


		}
	}
 	return 0
}



func getSolutionPart2(input []string) int {
	numbers, boards := parseInput(input)
	
	for _, num := range numbers {
		boardsLeft := 0
		won := false
		winningBoard := Board{} 

		for i, board := range boards {
			if board.won {
				continue
			}
			boardsLeft++
			won = markBoard(&boards[i], num)
			if won {
				winningBoard = board
			}
			
		}
		if boardsLeft == 1 && won {
			return getSum(winningBoard, num)
		}
	}
 	return 0

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

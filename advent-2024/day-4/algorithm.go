package day4

import (
	"fmt"
	"os"
	"strings"
)

type RuneMatrix = [][]rune

func ReadInput(filename string) (RuneMatrix, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return [][]rune{}, err
	}

	lines := strings.Split(string(content), "\n")
	var matrix = make([][]rune, len(lines))
	for index, line := range lines {
		for _, char := range line {
			matrix[index] = append(matrix[index], char)
		}
	}

	return matrix, nil
}

var directions = [][]int{
	{-1, 0, 1},
	{-1, 0, 1},
	{-1, 0, 1},
}

func FindWordInstances(word string, input RuneMatrix) int {
	wordChars := make([]rune, len(word))
	for _, char := range word {
		wordChars = append(wordChars, char)
	}

	for _, inputLine := range input {
		for _, inputChar := range inputLine {
			if inputChar != wordChars[0] {
				continue
			}

			// for _, line := range directions {
			// 	for _, pos := range line {
			// 		IsCorrectWord(1, wordChars, input, pos)
			// 	}
			// }
		}
	}

	return 0
}

type IsCorrectWordParams struct {
	Index         int
	Word          []rune
	Input         RuneMatrix
	XPos          int
	YPos          int
	CurrentXIndex int
	CurrentYIndex int
}

func IsCorrectWord(params IsCorrectWordParams) {
	yIndex := Clamp(params.CurrentYIndex+params.YPos, 0, len(params.Input))
	xIndex := Clamp(params.CurrentXIndex+params.XPos, 0, len(params.Input[0]))

	fmt.Println(string(params.Input[yIndex][xIndex]))
}

func Clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

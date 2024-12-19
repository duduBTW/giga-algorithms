package day4

import (
	"errors"
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

type Direction struct {
	X func(position, amount int) int
	Y func(position, amount int) int
}

func Previous(position, amount int) int {
	return position - amount
}

func Next(position, amount int) int {
	return position + amount
}

func Neutral(position, amount int) int {
	return position
}

var directions = map[string]Direction{
	"up-left":  {X: Previous, Y: Previous},
	"up":       {X: Neutral, Y: Previous},
	"up-right": {X: Next, Y: Previous},

	"down-left":  {X: Previous, Y: Next},
	"down":       {X: Neutral, Y: Next},
	"down-right": {X: Next, Y: Next},

	"left":  {X: Previous, Y: Neutral},
	"right": {X: Next, Y: Neutral},
}

func FindWordInstances(word string, input RuneMatrix) []Highlight {
	highlights := []Highlight{}

	wordChars := make([]rune, len(word))
	for index, char := range word {
		wordChars[index] = char
	}

	for yIndex, inputLine := range input {
		for xIndex, inputChar := range inputLine {
			// if first char is correct
			if inputChar != wordChars[0] {
				continue
			}

			// if the rest of the word matches
			for _, direction := range directions {
				if !CanGoToDirection(CanGoToDirectionParams{
					Input:     input,
					Direction: direction,
					X:         xIndex,
					Y:         yIndex,
				}) {
					continue
				}

				initialPosition := Position{
					X: xIndex,
					Y: yIndex,
				}
				isCorrectWord, endPosition := CheckIsCorrectWord(IsCorrectWordParams{
					Position:        1,
					Word:            wordChars,
					Input:           input,
					InitialPosition: initialPosition,
					Direction:       direction,
				})

				if isCorrectWord {
					highlights = append(highlights, Highlight{
						Start: initialPosition,
						End:   endPosition,
					})
				}
			}
		}
	}

	return highlights
}

type Position struct {
	X int
	Y int
}

type Highlight struct {
	Start Position
	End   Position
}

type IsCorrectWordParams struct {
	Position        int
	Word            []rune
	Input           RuneMatrix
	InitialPosition Position
	Direction       Direction
}

func CheckIsCorrectWord(params IsCorrectWordParams) (bool, Position) {
	position := Position{}

	yIndex, err := Clamp(params.Direction.Y(params.InitialPosition.Y, params.Position), -1, len(params.Input)-1)
	position.Y = yIndex
	if err != nil {
		return false, position
	}

	xIndex, err := Clamp(params.Direction.X(params.InitialPosition.X, params.Position), -1, len(params.Input[yIndex])-1)
	position.X = xIndex
	if err != nil {
		return false, position
	}

	if len(params.Input[yIndex]) == 0 {
		return false, position
	}

	char := params.Input[yIndex][xIndex]
	wordChar := params.Word[params.Position]
	if char == wordChar && params.Position == len(params.Word)-1 {
		return true, position
	}

	if char != wordChar {
		return false, position
	}

	params.Position = params.Position + 1
	return CheckIsCorrectWord(params)
}

type CanGoToDirectionParams struct {
	Input     RuneMatrix
	Direction Direction
	X         int
	Y         int
}

func CanGoToDirection(params CanGoToDirectionParams) bool {
	yIndex, err := Clamp(params.Direction.Y(params.Y, 0), -1, len(params.Input)-1)
	if err != nil {
		return false
	}

	_, err = Clamp(params.Direction.X(params.X, 0), -1, len(params.Input[yIndex])-1)
	if err != nil {
		return false
	}

	return true
}

func Clamp(value, min, max int) (int, error) {
	if value <= min || value > max {
		return 0, errors.New("invalid position")
	}

	return value, nil
}

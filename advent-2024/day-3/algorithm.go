package day3

import (
	"os"
	"regexp"
	"strconv"
)

func ReadInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func RegexPart1(content string) int {
	result := 0

	pattern := `mul\((\d+),(\d+)\)`
	regexResult := regexp.MustCompile(pattern)

	matches := regexResult.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		firstNumber, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}
		secondNumber, err := strconv.Atoi(match[2])
		if err != nil {
			continue
		}

		result += firstNumber * secondNumber
	}

	return result
}

type Expression struct {
	Name            string
	OnFound         func(vals []string)
	ValidateArgChar func(rune) bool
}

func ExpressionFinder(input string, expressions []Expression) {
	runningExpressions := make([]int, len(expressions))
	var completedExpression Expression
	var args = [][]rune{}
	currentArgIndex := 0
	var opened bool
	ops := false

	for _, inputChar := range input {
		if completedExpression.Name != "" {
			if opened && inputChar == ',' {
				currentArgIndex++
				continue
			}

			if opened && inputChar == ')' {
				argStrings := make([]string, len(args))
				for i, runes := range args {
					argStrings[i] = string(runes) // Convert each []rune to a string
				}

				completedExpression.OnFound(argStrings)
				completedExpression = Expression{}
				opened = false
				currentArgIndex = 0
				args = [][]rune{}
				ops = false
				continue
			}

			if opened {
				if !completedExpression.ValidateArgChar(inputChar) {
					opened = false
					currentArgIndex = 0
					args = [][]rune{}
					ops = true
				} else {
					if len(args) <= currentArgIndex {
						args = append(args, []rune{})
					}

					args[currentArgIndex] = append(args[currentArgIndex], inputChar)
					continue
				}
			}

			if !ops && !opened && inputChar == '(' {
				opened = true
				currentArgIndex = 0
				args = [][]rune{}
				continue
			}

			completedExpression = Expression{}
			ops = false
		}

		// fmt.Println(string(inputChar), runningExpressions)
		for index, expression := range expressions {
			runes := []rune(expression.Name)
			foundChar := runes[runningExpressions[index]]
			if foundChar == inputChar {
				if len(runes) == runningExpressions[index]+1 {
					completedExpression = expression
					runningExpressions[index] = 0
				} else {
					runningExpressions[index]++
				}
			} else {
				runningExpressions[index] = 0
			}
		}
	}
}

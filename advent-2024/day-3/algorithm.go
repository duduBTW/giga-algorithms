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
	ValidateArgChar func(rune) bool
}

type ExpressionFoundPosition struct {
	Start int
	End   int
	Name  string
	Args  []string
}

type ProcessExpression struct {
	Completed       bool
	Expression      Expression
	Args            [][]rune
	CurrentArgIndex int
	Opened          bool
	StartIndex      int
	FoundPosition   ExpressionFoundPosition
}

type ProcessResult = int

const (
	Success   ProcessResult = 0
	Completed ProcessResult = 1
	Abort     ProcessResult = 2
)

func (process *ProcessExpression) divideArgs() ProcessResult {
	if len(process.Args) < process.CurrentArgIndex+1 {
		return Abort
	}

	process.CurrentArgIndex++
	return Success
}
func (process *ProcessExpression) closeFunc(charIndex int) ProcessResult {
	argStrings := make([]string, len(process.Args))
	for i, runes := range process.Args {
		argStrings[i] = string(runes)
	}

	process.FoundPosition = ExpressionFoundPosition{
		Start: process.StartIndex,
		End:   charIndex,
		Name:  process.Expression.Name,
		Args:  argStrings,
	}
	return Completed
}
func (process *ProcessExpression) functionArg(inputChar rune) ProcessResult {
	if !process.Expression.ValidateArgChar(inputChar) {
		return Abort
	}

	// Add new arg if necessary
	if len(process.Args) <= process.CurrentArgIndex {
		process.Args = append(process.Args, []rune{})
	}
	// Append char to arg
	process.Args[process.CurrentArgIndex] = append(process.Args[process.CurrentArgIndex], inputChar)
	return Success
}
func (process *ProcessExpression) openFunc() ProcessResult {
	process.Opened = true
	return Success
}
func (process *ProcessExpression) Process(charIndex int, inputChar rune) ProcessResult {
	if process.Opened && inputChar == ',' {
		return process.divideArgs()
	}
	if process.Opened && inputChar == ')' {
		return process.closeFunc(charIndex)
	}
	if process.Opened {
		return process.functionArg(inputChar)
	}
	if !process.Opened && inputChar == '(' {
		return process.openFunc()
	}

	// Aborts if the function was not opened
	return Abort
}

func ExpressionFinder(input string, expressions []Expression) []ExpressionFoundPosition {
	runningExpressions := make([]int, len(expressions))
	completedExpression := ProcessExpression{}
	result := []ExpressionFoundPosition{}

	for charIndex, inputChar := range input {
		if completedExpression.Completed {
			processResult := completedExpression.Process(charIndex, inputChar)
			switch processResult {
			case Success:
				continue
			case Completed:
				result = append(result, completedExpression.FoundPosition)
				completedExpression = ProcessExpression{}
				continue
			case Abort:
				completedExpression = ProcessExpression{}
			}
		}

		for index, expression := range expressions {
			runes := []rune(expression.Name)
			foundChar := runes[runningExpressions[index]]
			if foundChar == inputChar {
				if len(runes) == runningExpressions[index]+1 {
					completedExpression.Expression = expression
					completedExpression.Completed = true
					completedExpression.StartIndex = charIndex - len(expression.Name)
					runningExpressions[index] = 0
				} else {
					runningExpressions[index]++
				}
			} else {
				runningExpressions[index] = 0
			}
		}
	}

	return result
}

func MultiplyArgs(args []string) int {
	result := 1
	for _, val := range args {
		numericVal, err := strconv.Atoi(val)
		if err != nil {
			continue
		}

		result *= numericVal
	}

	return result
}

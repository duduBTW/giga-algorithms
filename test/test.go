package main

import (
	"fmt"
	"strconv"

	day3 "github.com/dudubtw/giga-algorithms/advent-2024/day-3"
)

func main() {
	content, err := day3.ReadInput("advent-2024/day-3/data-part-1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	resultMath := 0
	active := true

	day3.ExpressionFinder(content, []day3.Expression{
		{
			Name: "mul",
			OnFound: func(vals []string) {
				if !active {
					return
				}

				if len(vals) != 2 {
					return
				}
				currentMultiplication := 1
				for _, val := range vals {
					numericVal, err := strconv.Atoi(val)
					if err != nil {
						return
					}

					currentMultiplication *= numericVal
				}
				resultMath += currentMultiplication
			},
			ValidateArgChar: func(r rune) bool {
				_, err := strconv.Atoi(string(r))
				return err == nil
			},
		},
		{
			Name: "do",
			OnFound: func(vals []string) {
				active = true
			},
			ValidateArgChar: func(r rune) bool {
				return false
			},
		},
		{
			Name: "don't",
			OnFound: func(vals []string) {
				active = false
			},
			ValidateArgChar: func(r rune) bool {
				return false
			},
		},
	})

	fmt.Println(resultMath)
}

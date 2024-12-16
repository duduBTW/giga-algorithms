package main

import (
	day4 "github.com/dudubtw/giga-algorithms/advent-2024/day-4"
)

func main() {
	// matrix, err := day4.ReadInput("advent-2024/day-4/data-part-1.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	day4.IsCorrectWord(day4.IsCorrectWordParams{
		Index: 0,
		Word:  []rune{'X', 'M', 'A', 'S'},
		Input: [][]rune{
			{'T', 'E', 'S', 'T'},
			{'T', 'E', 'S', 'T'},
			{'T', 'E', 'S', 'T'},
			{'T', 'E', 'S', 'T'},
		},
		XPos:          1,
		YPos:          0,
		CurrentXIndex: 0,
		CurrentYIndex: 0,
	})
}

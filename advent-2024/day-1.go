package advent2024

import (
	"os"
	"strconv"
	"strings"
	"sort"
)

func SplitLine(line string) []int {
	hasFlipped := false
	var leftString string
	var rightString string

	for _, char := range line {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			hasFlipped = true
			continue
		}

		if hasFlipped {
			rightString += string(char)
		} else {
			leftString += string(char)
		}
	}

	left, _ := strconv.Atoi(leftString)
	right, _ := strconv.Atoi(rightString)
	return []int{left, right}
}



func ReadDay1Input(filename string) ([][]int, error) {
	content, err := os.ReadFile(DATA_DIR + filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var left []int
	var right []int
	for _, line := range lines {
		if line == "" {
			continue
		}

		splitedLine := SplitLine(line)
		left = append(left, splitedLine[0])
		right = append(right, splitedLine[1])
	}

	return [][]int{left, right, make([]int, len(left))}, nil
}


// func Day11(list [][]int) int {
// 	sort.Slice(list[0], func(i, j int) bool {
// 		return list[0][i] < list[0][j]
// 	})
// 	sort.Slice(list[1], func(i, j int) bool {
// 		return list[1][i] < list[1][j]
// 	})

// 	total := 0
// 	for index, _ := range list[0] {

// 		left := list[0][index]
// 		right := list[1][index]

// 		total += absDiffInt(right, left)
// 	}

// 	return total
// }

// func absDiffInt(x, y int) int {
// 	if x < y {
// 		return y - x
// 	}
// 	return x - y
// }


func SortLines(list [][]int) [][]int {
	newList := make([][]int, 3)

	newList[0] = make([]int, len(list[0]))
	newList[1] = make([]int, len(list[1]))
	newList[2] = make([]int, len(list[0]))

	copy(newList[0], list[0])
	copy(newList[1], list[1])

	sort.Slice(newList[0], func(i, j int) bool {
		return newList[0][i] < newList[0][j]
	})
	sort.Slice(newList[1], func(i, j int) bool {
		return newList[1][i] < newList[1][j]
	})

	return newList
}

func CalculateAll(list [][]int) (int, [][]int) {
	total := 0
	listWithDiffs := list
	for index, _ := range list[0] {
		diff := Day11Line(list, index)
		total += diff
		listWithDiffs[2][index] = diff
	}
	return total, listWithDiffs
}

func Day11Line(list [][]int, index int) int {
	left := list[0][index]
	right := list[1][index]
	return absDiffInt(right, left)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
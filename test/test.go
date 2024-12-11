package main

import (
	"fmt"

	advent2024 "github.com/dudubtw/giga-algorithms/advent-2024"
)

func main() {
	lines, err := advent2024.ReadDay2Input("day-2-1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(advent2024.Day2Part1(lines))
	fmt.Println(advent2024.Day2Part2(lines))
}

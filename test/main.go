package main

import (
	"fmt"

	day4 "github.com/dudubtw/giga-algorithms/advent-2024/day-4"
)

func main() {
	matrix, err := day4.ReadInput("D:/Peronal/giga-algorithms/advent-2024/day-4/data-part-1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(day4.FindWordInstances("XMAS", matrix))
}

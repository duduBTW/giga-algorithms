package main

import (
	"fmt"

	day5 "github.com/dudubtw/giga-algorithms/advent-2024/day-5"
)

func main() {
	manual, _ := day5.ReadInput()
	total := day5.FixManual(manual)
	fmt.Println(total)
}

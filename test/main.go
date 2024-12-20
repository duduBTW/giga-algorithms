package main

import (
	"fmt"
	"strconv"

	day4 "github.com/dudubtw/giga-algorithms/advent-2024/day-4"
)

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	a := day4.Highlight{
		Start: day4.Position{
			X: 1,
			Y: 6,
		},
		End: day4.Position{
			X: 3,
			Y: 4,
		},
	}

	b := day4.Highlight{
		Start: day4.Position{
			X: 3,
			Y: 5,
		},
		End: day4.Position{
			X: 1,
			Y: 3,
		},
	}

	middlePoints := make(map[string]day4.Highlight)

	for _, highlight := range []day4.Highlight{a, b} {
		maxY := float64(MaxInt(highlight.Start.Y, highlight.End.Y))
		minY := float64(MinInt(highlight.Start.Y, highlight.End.Y))

		maxX := float64(MaxInt(highlight.Start.X, highlight.End.X))
		minX := float64(MinInt(highlight.Start.X, highlight.End.X))

		aa := strconv.FormatFloat(((minY - maxY) / 2), 'f', -1, 64)
		bb := strconv.FormatFloat(((minX + maxX) / 2), 'f', -1, 64)
		fmt.Println(aa, bb)

		middlePoint := strconv.FormatFloat(((minY-maxY)/2), 'f', -1, 64) + strconv.FormatFloat(((minX+maxX)/2), 'f', -1, 64)

		cHighlight, ok := middlePoints[middlePoint]
		if ok {
			fmt.Println(highlight, cHighlight)
			delete(middlePoints, middlePoint)
			continue
		}

		middlePoints[middlePoint] = highlight
	}

}

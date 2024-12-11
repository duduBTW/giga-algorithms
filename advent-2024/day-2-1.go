package advent2024

import (
	"os"
	"strconv"
	"strings"
)

const (
	Day2Part2ModeIncrease  = 1
	Day2Part2ModeDecrease  = -1
	Day2Part2ModeUndefined = 0
)

func ReadDay2Input(filename string) ([][]int, error) {
	content, err := os.ReadFile(DATA_DIR + filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var reports [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}

		report := []int{}
		for _, num := range strings.Split(line, " ") {
			reportNum, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}

			report = append(report, reportNum)
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func Day2Part2(reports [][]int) int {
	safeCount := 0

	for _, report := range reports {
		reportCopy := append([]int{}, report...)
		for i := 0; i <= len(reportCopy); i++ {
			unsafeIndex := Day2Part2IsReportUnsafeIndex(reportCopy)
			if unsafeIndex == NotFoundDay2Part2Index {
				safeCount++
				break
			}

			if i == len(reportCopy) {
				break
			}

			reportCopy = append(reportCopy[:i], reportCopy[i+1:]...)
		}
	}

	return safeCount
}

func Day2Part1(reports [][]int) int {
	safeCount := 0
	for _, report := range reports {
		if Day2Part2IsReportUnsafeIndex(report) == NotFoundDay2Part2Index {
			safeCount++
		}
	}

	return safeCount
}

const NotFoundDay2Part2Index = -1

func Day2Part2IsReportUnsafeIndex(report []int) int {
	if len(report) < 2 {
		return NotFoundDay2Part2Index
	}

	// getMode returns the mode of the report
	// 0 = undefined
	// 1 - increase
	// -1 - decrease
	mode := Day2Part2Mode(report[1] - report[0])
	if mode == Day2Part2ModeUndefined {
		return 1
	}

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		absDiiff := absInt(diff)
		changedMode := mode != Day2Part2Mode(diff)
		if Day2Part2IsDiffNotInRange(absDiiff) || changedMode {
			return i
		}
	}

	return NotFoundDay2Part2Index
}

func Day2Part2IsDiffNotInRange(diff int) bool {
	const maxDiff = 3
	const minDiff = 1
	return diff > maxDiff || diff < minDiff
}

func Day2Part2Mode(diff int) int {
	if diff == 0 {
		return Day2Part2ModeUndefined
	}

	if diff > 0 {
		return Day2Part2ModeIncrease
	}

	return Day2Part2ModeDecrease
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

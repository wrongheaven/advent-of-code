package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 2)
	day.PrintResults(part1, part2)
}

func part1(lines []string) (int, error) {
	reports := [][]int{}
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		report := []int{}
		for _, level := range lineSplit {
			if level == "" {
				continue
			}

			levelInt, err := strconv.Atoi(level)
			if err != nil {
				return 0, err
			}
			report = append(report, levelInt)
		}

		reports = append(reports, report)
	}

	numSafe := 0
	for _, report := range reports {
		if isSafe1(report) {
			numSafe++
		}
	}

	return numSafe, nil
}

func isSafe1(report []int) bool {
	ascending := false
	descending := false

	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]

		if math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff == 0 {
			return false
		} else if diff < 0 {
			if descending {
				return false
			}
			ascending = true
		} else {
			if ascending {
				return false
			}
			descending = true
		}
	}

	return true
}

func part2(lines []string) (int, error) {
	reports := [][]int{}
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")

		report := []int{}
		for _, level := range lineSplit {
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				return 0, err
			}
			report = append(report, levelInt)
		}

		reports = append(reports, report)
	}

	numSafe := 0
	for _, report := range reports {
		if isSafe2(report) {
			numSafe++
		}
	}

	return numSafe, nil
}

func isSafe2(report []int) bool {
	if isSafe1(report) {
		return true
	}

	for i := range report {
		newReport := slices.Delete(slices.Clone(report), i, i+1)
		if isSafe1(newReport) {
			return true
		}
	}

	return false
}

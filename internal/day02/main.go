package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// data, err := os.ReadFile("./internal/inputs/day02.txt")
	data, err := os.ReadFile("./internal/day02/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans1, err := part1(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: %d\n", ans1)

	ans2, err := part2(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2: %d\n", ans2)
}

func part1(data []byte) (int, error) {
	reports := [][]int{}
	for _, line := range strings.Split(string(data), "\n") {
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

func part2(data []byte) (int, error) {
	reports := [][]int{}
	for _, line := range strings.Split(string(data), "\n") {
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
	numAscending := 0
	numDescending := 0

	numBadReports := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]

		if math.Abs(float64(diff)) > 3 {
			numBadReports++
		}

		if diff == 0 {
			numBadReports++
		} else if diff < 0 {
			numAscending++
		} else {
			numDescending++
		}
	}

	if numBadReports > 1 {
		return false
	}

	if numAscending > 1 && numDescending > 1 {
		return false
	}

	if numBadReports > 0 && numAscending > 0 && numDescending > 0 {
		return false
	}

	return true
}

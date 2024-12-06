package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 5)

	ex1, ans1, ex2, ans2 := day.Run(part1, part2)

	fmt.Println("Example 1:", ex1)
	fmt.Println("Input   1:", ans1)
	fmt.Println("Example 2:", ex2)
	fmt.Println("Input   2:", ans2)
}

func parseInput(lines []string) (orderingRules [][2]int, updates [][]int) {
	parseRules := true
	for _, line := range lines {
		if line == "" {
			parseRules = false
			continue
		}

		if parseRules {
			strSplit := strings.Split(line, "|")
			n1, _ := strconv.Atoi(strSplit[0])
			n2, _ := strconv.Atoi(strSplit[1])
			orderingRules = append(orderingRules, [2]int{n1, n2})
		} else {
			update := util.Map(strings.Split(line, ","), func(s string) int {
				n, _ := strconv.Atoi(s)
				return n
			})

			updates = append(updates, update)
		}
	}

	return orderingRules, updates
}

func part1(lines []string) (int, error) {
	orderingRules, updates := parseInput(lines)

	correctOrder := [][]int{}
	for _, update := range updates {
		if isCorrectOrder(orderingRules, update) {
			correctOrder = append(correctOrder, update)
		}
	}

	sumOfMiddleNumbers := util.Reduce(correctOrder, getMiddleNumber, 0)

	return sumOfMiddleNumbers, nil
}

func isCorrectOrder(orderingRules [][2]int, update []int) bool {
	for _, rule := range orderingRules {
		idx1 := slices.Index(update, rule[0])
		idx2 := slices.Index(update, rule[1])

		if idx1 != -1 && idx2 != -1 && idx1 > idx2 {
			return false
		}
	}

	return true
}

func getMiddleNumber(nums []int) int {
	return nums[len(nums)/2]
}

func part2(lines []string) (int, error) {
	orderingRules, updates := parseInput(lines)

	incorrectOrder := [][]int{}
	for _, update := range updates {
		if !isCorrectOrder(orderingRules, update) {
			var err error
			update, err = fixOrder(orderingRules, update)
			if err != nil {
				return 0, err
			}
			incorrectOrder = append(incorrectOrder, update)
		}
	}

	sumOfMiddleNumbers := util.Reduce(incorrectOrder, getMiddleNumber, 0)

	return sumOfMiddleNumbers, nil
}

func fixOrder(orderingRules [][2]int, update []int) ([]int, error) {
	tries := 0
	for {
		for _, rule := range orderingRules {
			idx1 := slices.Index(update, rule[0])
			idx2 := slices.Index(update, rule[1])

			if idx1 != -1 && idx2 != -1 && idx1 > idx2 {
				update[idx1], update[idx2] = update[idx2], update[idx1]
			}
		}

		if isCorrectOrder(orderingRules, update) {
			break
		}

		tries++
		if tries >= 100+len(orderingRules) {
			return nil, fmt.Errorf("timed out")
		}
	}

	return update, nil
}

package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 1)
	day.PrintResults(part1, part2)
}

func part1(lines []string) (int, error) {
	var list1 []int = []int{}
	var list2 []int = []int{}

	for _, line := range lines {
		nums := strings.Split(line, " ")
		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		n2, err := strconv.Atoi(nums[len(nums)-1])
		if err != nil {
			return 0, err
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	totalDiff := 0
	for i := range len(list1) {
		totalDiff += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return totalDiff, nil
}

func part2(lines []string) (int, error) {
	hashmap := make(map[int]int)

	var list1 []int = []int{}
	var list2 []int = []int{}

	for _, line := range lines {
		nums := strings.Split(line, " ")
		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		n2, err := strconv.Atoi(nums[len(nums)-1])
		if err != nil {
			return 0, err
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	for _, num := range list2 {
		hashmap[num]++
	}

	similarityScore := 0
	for _, num := range list1 {
		similarityScore += num * hashmap[num]
	}

	return similarityScore, nil
}

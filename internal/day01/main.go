package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./internal/inputs/day01.txt")
	// data, err := os.ReadFile("./internal/day01/example.txt")
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
	var list1 []int = []int{}
	var list2 []int = []int{}

	for _, line := range strings.Split(string(data), "\n") {
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

func part2(data []byte) (int, error) {
	hashmap := make(map[int]int)

	var list1 []int = []int{}
	var list2 []int = []int{}

	for _, line := range strings.Split(string(data), "\n") {
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

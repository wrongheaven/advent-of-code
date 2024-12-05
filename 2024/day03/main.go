package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("./2024/inputs/day03.txt")
	// data, err := os.ReadFile("./2024/day03/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans1, err := part1(string(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: %d\n", ans1)

	ans2, err := part2(string(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2: %d\n", ans2)
}

func part1(data string) (int, error) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	multSum := 0
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		multSum += n1 * n2
	}

	return multSum, nil
}

func part2(data string) (int, error) {
	// const (
	// 	DO_MATCH   = 4
	// 	DONT_MATCH = 7
	// )

	// re := regexp.MustCompile(`(do\(\)|don\'t\(\))`)
	// matches := re.FindAllStringIndex(data, -1)

	// removeIdxs := [][]int{}
	// for i := 0; i < len(matches)-1; i++ {
	// 	if matches[i][1]-matches[i][0] == DONT_MATCH {
	// 		for j := i + 1; j < len(matches); j++ {
	// 			if matches[j][1]-matches[j][0] == DO_MATCH {
	// 				removeIdxs = append(removeIdxs, []int{matches[i][0], matches[j][1]})
	// 				i = j + 1
	// 			}
	// 		}
	// 	}
	// }

	// for i := len(removeIdxs) - 1; i >= 0; i-- {
	// 	p := data[:removeIdxs[i][0]]
	// 	q := data[removeIdxs[i][1]:]
	// 	data = p + q
	// }

	// return part1(data)

	return 0, nil
}

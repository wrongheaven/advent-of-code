package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 3)

	ex1, ans1, ex2, ans2 := day.Run(part1, part2)

	fmt.Println("Example 1:", ex1)
	fmt.Println("Input   1:", ans1)
	fmt.Println("Example 2:", ex2)
	fmt.Println("Input   2:", ans2)
}

func part1(lines []string) (int, error) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(strings.Join(lines, "\n"), -1)

	multSum := 0
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		multSum += n1 * n2
	}

	return multSum, nil
}

func part2(lines []string) (int, error) {
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

package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 3)
	day.PrintResults(part1, part2)
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

	str := strings.Join(lines, "\n")

	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i:], "don't()") {
			str, i = redactStringPart(str, i)
		}
	}

	return part1([]string{str})
}

func redactStringPart(str string, idx int) (newStr string, newIdx int) {
	for i := idx; i < len(str); i++ {
		if strings.HasPrefix(str[i:], "do()") {
			return str, i
		}

		str = str[:i] + "x" + str[i+1:]
	}

	return str, len(str)
}

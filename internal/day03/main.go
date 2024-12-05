package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("./internal/inputs/day03.txt")
	// data, err := os.ReadFile("./internal/day03/example.txt")
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
	// re := regexp.MustCompile(`don\'t\(\).*do\(\)`)
	// matches := re.FindAllString(data, -1)

	// for _, match := range matches {
	// 	idx := strings.Index(match, "do()")
	// 	if idx < len(match)-1-len("do()") {
	// 		match = match[:idx]
	// 	}

	// 	data = strings.ReplaceAll(data, match, " ")
	// }

	return part1(data)
}

package main

import (
	"fmt"
	"strconv"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 9)
	day.PrintResults(part1, part2)
}

func part1(lines []string) (int, error) {
	if lines[0] == "" {
		return 0, fmt.Errorf("no lines")
	}

	blocks := getBlockRep(lines[0])
	blocks = moveFileBlocks(blocks)

	checksum := calcChecksum(blocks)

	return checksum, nil
}

func part2(lines []string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func getBlockRep(input string) []int {
	result := []int{}
	idx := 0
	for i, r := range input {
		size, _ := strconv.Atoi(string(r))
		isFile := i%2 == 0

		if isFile {
			for range size {
				result = append(result, idx)
			}
			idx++
		} else {
			for range size {
				result = append(result, -1)
			}
		}
	}

	return result
}

func moveFileBlocks(input []int) []int {
	j := len(input) - 1
	for i := 0; i < len(input); i++ {
		if input[i] != -1 {
			continue
		}

		for ; ; j-- {
			if input[j] != -1 {
				break
			}
		}

		if j <= i {
			break
		}

		input[i], input[j] = input[j], input[i]
	}
	return input
}

func calcChecksum(blocks []int) int {
	sum := 0
	for i, n := range blocks {
		if n == -1 {
			break
		}

		sum += i * n
	}
	return sum
}

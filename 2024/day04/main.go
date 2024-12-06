package main

import (
	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 4)
	day.PrintResults(part1, part2)
}

func part1(lines []string) (int, error) {
	totalWords := 0
	for row := range lines {
		for col := range lines[row] {
			if lines[row][col] == 'X' {
				totalWords += countWordsFound(lines, row, col)
			}
		}
	}

	return totalWords, nil
}

func countWordsFound(lines []string, row, col int) int {
	wordsFound := 0
	// north
	if row >= 3 &&
		lines[row-1][col] == 'M' &&
		lines[row-2][col] == 'A' &&
		lines[row-3][col] == 'S' {
		wordsFound++
	}
	// south
	if row < len(lines)-3 &&
		lines[row+1][col] == 'M' &&
		lines[row+2][col] == 'A' &&
		lines[row+3][col] == 'S' {
		wordsFound++
	}
	// east
	if col < len(lines[row])-3 &&
		lines[row][col+1] == 'M' &&
		lines[row][col+2] == 'A' &&
		lines[row][col+3] == 'S' {
		wordsFound++
	}
	// west
	if col >= 3 &&
		lines[row][col-1] == 'M' &&
		lines[row][col-2] == 'A' &&
		lines[row][col-3] == 'S' {
		wordsFound++
	}
	// ne
	if row >= 3 && col < len(lines[row])-3 &&
		lines[row-1][col+1] == 'M' &&
		lines[row-2][col+2] == 'A' &&
		lines[row-3][col+3] == 'S' {
		wordsFound++
	}
	// nw
	if row >= 3 && col >= 3 &&
		lines[row-1][col-1] == 'M' &&
		lines[row-2][col-2] == 'A' &&
		lines[row-3][col-3] == 'S' {
		wordsFound++
	}
	// se
	if row < len(lines)-3 && col < len(lines[row])-3 &&
		lines[row+1][col+1] == 'M' &&
		lines[row+2][col+2] == 'A' &&
		lines[row+3][col+3] == 'S' {
		wordsFound++
	}
	// sw
	if row < len(lines)-3 && col >= 3 &&
		lines[row+1][col-1] == 'M' &&
		lines[row+2][col-2] == 'A' &&
		lines[row+3][col-3] == 'S' {
		wordsFound++
	}

	return wordsFound
}

func part2(lines []string) (int, error) {
	totalWords := 0
	for row := range lines {
		for col := range lines[row] {
			if lines[row][col] == 'M' {
				totalWords += countMAS(lines, row, col)
			}
		}
	}

	return totalWords, nil
}

func countMAS(lines []string, row, col int) int {
	wordsFound := 0

	// se
	if row < len(lines)-2 && col < len(lines[row])-2 &&
		lines[row+1][col+1] == 'A' &&
		lines[row+2][col+2] == 'S' {
		if lines[row][col+2] == 'M' && lines[row+2][col] == 'S' ||
			lines[row+2][col] == 'M' && lines[row][col+2] == 'S' {
			wordsFound++
		}
	}
	// nw
	if row >= 2 && col >= 2 &&
		lines[row-1][col-1] == 'A' &&
		lines[row-2][col-2] == 'S' {
		if lines[row][col-2] == 'M' && lines[row-2][col] == 'S' ||
			lines[row-2][col] == 'M' && lines[row][col-2] == 'S' {
			wordsFound++
		}
	}

	return wordsFound
}

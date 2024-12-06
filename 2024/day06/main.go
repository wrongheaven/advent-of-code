package main

import (
	"fmt"
	"slices"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

type Dir int

const (
	UP    Dir = 1
	RIGHT Dir = 2
	DOWN  Dir = 3
	LEFT  Dir = 4
)

type Guard struct {
	Position  [2]int
	Direction Dir
}

func main() {
	day := util.NewDay(2024, 6)
	day.PrintResults(part1, part2)
}

func part1(lines []string) (int, error) {
	guard := findGuard(lines)

	visited := make(map[[2]int]bool)
	c := 0
	for {
		c++
		visited[guard.Position] = true
		var outOfBounds bool
		guard, outOfBounds = guardAdvance(guard, lines)
		if outOfBounds {
			break
		}

		// break if we're stuck in an infinite loop (should not happen in part 1)
		if c == len(lines)*len(lines[0]) {
			break
		}
	}

	return len(visited), nil
}

func findGuard(lines []string) Guard {
	guard := Guard{}

	for row, line := range lines {
		for col, c := range line {
			switch c {
			case '^':
				return Guard{Position: [2]int{col, row}, Direction: UP}
			case 'v':
				return Guard{Position: [2]int{col, row}, Direction: DOWN}
			case '<':
				return Guard{Position: [2]int{col, row}, Direction: LEFT}
			case '>':
				return Guard{Position: [2]int{col, row}, Direction: RIGHT}
			}
		}
	}

	return guard
}

func guardAdvance(guard Guard, lines []string) (newGuard Guard, outOfBounds bool) {
	if guard.Direction == UP && guard.Position[1] == 0 ||
		guard.Direction == DOWN && guard.Position[1] == len(lines)-1 ||
		guard.Direction == LEFT && guard.Position[0] == 0 ||
		guard.Direction == RIGHT && guard.Position[0] == len(lines[0])-1 {
		return Guard{}, true
	}

	switch guard.Direction {
	case UP:
		if lines[guard.Position[1]-1][guard.Position[0]] == '#' {
			guard.Direction = (guard.Direction)%4 + 1
		} else {
			guard.Position[1]--
		}
	case DOWN:
		if lines[guard.Position[1]+1][guard.Position[0]] == '#' {
			guard.Direction = (guard.Direction)%4 + 1
		} else {
			guard.Position[1]++
		}
	case LEFT:
		if lines[guard.Position[1]][guard.Position[0]-1] == '#' {
			guard.Direction = (guard.Direction)%4 + 1
		} else {
			guard.Position[0]--
		}
	case RIGHT:
		if lines[guard.Position[1]][guard.Position[0]+1] == '#' {
			guard.Direction = (guard.Direction)%4 + 1
		} else {
			guard.Position[0]++
		}
	}

	return guard, false
}

// painfully slow, but works
func part2(lines []string) (int, error) {
	guard := findGuard(lines)

	loops := 0
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] != '.' {
				continue
			}

			newLines := slices.Clone(lines)
			newLines[row] = newLines[row][:col] + "#" + newLines[row][col+1:]
			if stuckInLoop(guard, newLines) {
				loops++
			}

			fmt.Println("row:", row, "col:", col)
		}
	}

	return loops, nil
}

func stuckInLoop(guard Guard, lines []string) bool {
	visited := make(map[[2]int][]Dir)

	for {
		// move/turn
		var outOfBounds bool
		guard, outOfBounds = guardAdvance(guard, lines)

		// check if we're out of bounds
		if outOfBounds {
			return false
		}

		// check if we've been in this location with this orientation before
		if slices.Contains(visited[guard.Position], guard.Direction) {
			return true
		}

		// record current position and orientation
		visited[guard.Position] = append(visited[guard.Position], guard.Direction)
	}
}

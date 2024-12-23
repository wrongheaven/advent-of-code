package main

import (
	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 8)
	day.PrintResults(part1, part2)
}

func part1(lines []string) (int, error) {
	antennaLocations := make(map[rune][][2]int)

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			char := rune(lines[row][col])
			if char != '.' {
				antennaLocations[char] =
					append(antennaLocations[char], [2]int{row, col})
			}
		}
	}

	getAntinodeLocations := func(locs [][2]int) [][2]int {
		antinodeLocations := [][2]int{}
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j < len(locs); j++ {
				newRow1 := 2*locs[i][0] - locs[j][0]
				newCol1 := 2*locs[i][1] - locs[j][1]
				newRow2 := 2*locs[j][0] - locs[i][0]
				newCol2 := 2*locs[j][1] - locs[i][1]

				if newRow1 >= 0 && newRow1 < len(lines[0]) &&
					newCol1 >= 0 && newCol1 < len(lines) {
					antinodeLocations = append(antinodeLocations, [2]int{newRow1, newCol1})
				}
				if newRow2 >= 0 && newRow2 < len(lines[0]) &&
					newCol2 >= 0 && newCol2 < len(lines) {
					antinodeLocations = append(antinodeLocations, [2]int{newRow2, newCol2})
				}
			}
		}

		return antinodeLocations
	}

	allAntinodeLocations := make(map[[2]int]bool)
	for _, locs := range antennaLocations {
		for _, loc := range getAntinodeLocations(locs) {
			allAntinodeLocations[loc] = true
		}
	}

	return len(allAntinodeLocations), nil
}

func part2(lines []string) (int, error) {
	antennaLocations := make(map[rune][][2]int)

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			char := rune(lines[row][col])
			if char != '.' {
				antennaLocations[char] =
					append(antennaLocations[char], [2]int{row, col})
			}
		}
	}

	getAntinodeLocations := func(locs [][2]int) [][2]int {
		antinodeLocations := [][2]int{}
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j < len(locs); j++ {
				locA := locs[i]
				locB := locs[j]

				deltaX := locB[1] - locA[1]
				deltaY := locB[0] - locA[0]

				// find starting coords.
				startX := locA[1]
				startY := locA[0]
				for {
					startX -= deltaX
					startY -= deltaY
					if startX < 0 || startY < 0 ||
						startX >= len(lines[0]) || startY >= len(lines) {
						// outside canvas, backtrack
						startX += deltaX
						startY += deltaY
						break
					}
				}

				// starting coord found
				for {
					antinodeLocations = append(antinodeLocations, [2]int{startY, startX})

					startX += deltaX
					startY += deltaY

					if startX < 0 || startY < 0 ||
						startX >= len(lines[0]) || startY >= len(lines) {
						break
					}
				}
			}
		}

		return antinodeLocations
	}

	allAntinodeLocations := make(map[[2]int]bool)
	for _, locs := range antennaLocations {
		for _, loc := range getAntinodeLocations(locs) {
			allAntinodeLocations[loc] = true
		}
	}

	return len(allAntinodeLocations), nil
}

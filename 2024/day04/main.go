package main

import (
	"fmt"
	"log"

	"github.com/wrongheaven/advent-of-code-2024/util"
)

func main() {
	day := util.NewDay(2024, 4)
	ans1, err := part1(day)
	if err != nil {
		log.Fatal(err)
	}
	ans2, err := part2(day)
	if err != nil {
		log.Fatal(err)
	}

	// // data, err := os.ReadFile("./2024/inputs/day04.txt")
	// data, err := os.ReadFile("./2024/day04/example.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ans1, err := part1(string(data))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("Part 1: %d\n", ans1)

	// ans2, err := part2(string(data))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("Part 2: %d\n", ans2)
}

func part1(day util.Day) (int, error) {
	return 0, nil
}

func part2(day util.Day) (int, error) {
	return 0, nil
}

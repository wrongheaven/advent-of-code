package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// data, err := os.ReadFile("./internal/inputs/day04.txt")
	data, err := os.ReadFile("./internal/day04/example.txt")
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
	return 0, nil
}

func part2(data string) (int, error) {
	return 0, nil
}

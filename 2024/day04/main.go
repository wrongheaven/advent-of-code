package main

import (
	"fmt"
	"log"
	"os"
)

type Day struct {
	Number int
}

func NewDay(dayNumber int) Day {
	return Day{
		Number: dayNumber,
	}
}

func main() {
	// day := NewDay(4)

	// data, err := os.ReadFile("./2024/inputs/day04.txt")
	data, err := os.ReadFile("./2024/day04/example.txt")
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

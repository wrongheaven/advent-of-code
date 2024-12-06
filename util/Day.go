package util

import (
	"fmt"
	"os"
	"strings"
)

type PartFunc func(lines []string) (ans int, err error)

type Day struct {
	Year    int
	Day     int
	example string
	input   string
}

func NewDay(year, day int) Day {
	newDay := Day{Year: year, Day: day}

	exampleString := ""
	exampleData, err := os.ReadFile(fmt.Sprintf("./%d/day%02d/example.txt", year, day))
	if err == nil {
		exampleString = string(exampleData)
	}
	newDay.example = exampleString

	inputString := ""
	inputData, err := os.ReadFile(fmt.Sprintf("./%d/inputs/day%02d.txt", year, day))
	if err == nil {
		inputString = string(inputData)
	}
	newDay.input = inputString

	return newDay
}

func (d Day) Run(part1, part2 PartFunc) (ex1, ans1, ex2, ans2 int) {
	exampleLines := d.exampleLines()
	inputLines := d.inputLines()

	ex1, err := part1(exampleLines)
	if err != nil {
		ex1 = -1
	}
	ans1, err = part1(inputLines)
	if err != nil {
		ans1 = -1
	}

	ex2, err = part2(exampleLines)
	if err != nil {
		ex2 = -1
	}
	ans2, err = part2(inputLines)
	if err != nil {
		ans2 = -1
	}

	return ex1, ans1, ex2, ans2
}

func (d Day) exampleLines() []string {
	return strings.Split(d.example, "\n")
}

func (d Day) inputLines() []string {
	return strings.Split(d.input, "\n")
}

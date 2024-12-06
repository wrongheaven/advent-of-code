package help

import (
	"fmt"
	"strings"
)

func PrintGrid(grid []string) {
	fmt.Println(strings.Repeat("-", len(grid)+4))
	for _, line := range grid {
		fmt.Println("  " + line)
	}
	fmt.Println(strings.Repeat("-", len(grid)+4))
}

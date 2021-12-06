package main

import (
	"fmt"

	"github.com/jirojo2/adventofcode-2021/challenges/day1"
	"github.com/jirojo2/adventofcode-2021/utils"
)

func main() {
	measurements := utils.ReadInputInts("day1")

	counter1 := day1.CalculatePart1(measurements)
	counter2 := day1.CalculatePart2(measurements)

	fmt.Printf("Part 1 solution: %v\n", counter1)
	fmt.Printf("Part 2 solution: %v\n", counter2)
}

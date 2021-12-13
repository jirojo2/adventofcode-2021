package main

import (
	"log"

	"github.com/jirojo2/adventofcode-2021/challenges/day2"
	"github.com/jirojo2/adventofcode-2021/utils"
)

func main() {
	input := utils.ReadInput("day2")

	pos, err := day2.CalculatePositionPart1(input)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Result (part 1): %d\n", pos)

	pos, err = day2.CalculatePositionPart2(input)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Result (part 2): %d\n", pos)
}

package main

import (
	"log"

	"github.com/jirojo2/adventofcode-2021/challenges/day3"
	"github.com/jirojo2/adventofcode-2021/utils"
)

func main() {
	input := utils.ReadInput("day3")

	value := day3.CalculatePowerRate(input)
	log.Printf("Result (part 1): %d\n", value)
}

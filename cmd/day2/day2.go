package main

import (
	"log"

	"github.com/jirojo2/adventofcode-2021/challenges/day2"
	"github.com/jirojo2/adventofcode-2021/utils"
)

func main() {
	input := utils.ReadInput("day2")
	pos, err := day2.CalculatePosition(input)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Result: %d\n", pos)
}

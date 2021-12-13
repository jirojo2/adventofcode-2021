package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessCommandPart1(cmd string) (pos int, depth int, err error) {
	tokens := strings.Split(cmd, " ")
	op := tokens[0]
	value, err := strconv.Atoi(tokens[1])

	switch op {
	case "forward":
		return value, 0, err
	case "up":
		return 0, -value, err
	case "down":
		return 0, value, err
	}
	return 0, 0, fmt.Errorf("unknown command: %s", cmd)
}

func CalculatePositionPart1(input []string) (int, error) {
	pos := 0
	depth := 0
	for i := range input {
		deltaPos, deltaDepth, err := ProcessCommandPart1(input[i])
		if err != nil {
			return 0, err
		}

		pos += deltaPos
		depth += deltaDepth
	}
	return pos * depth, nil
}

func ProcessCommandPart2(cmd string, aim int) (deltaPos int, deltaDepth int, deltaAim int, err error) {
	tokens := strings.Split(cmd, " ")
	op := tokens[0]
	value, err := strconv.Atoi(tokens[1])

	switch op {
	case "forward":
		return value, aim * value, 0, err
	case "up":
		return 0, 0, -value, err
	case "down":
		return 0, 0, value, err
	}
	return 0, 0, 0, fmt.Errorf("unknown command: %s", cmd)
}

func CalculatePositionPart2(input []string) (int, error) {
	aim := 0
	pos := 0
	depth := 0
	for i := range input {
		deltaPos, deltaDepth, deltaAim, err := ProcessCommandPart2(input[i], aim)
		if err != nil {
			return 0, err
		}

		pos += deltaPos
		depth += deltaDepth
		aim += deltaAim
	}
	return pos * depth, nil
}

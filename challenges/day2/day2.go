package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessCommand(cmd string) (pos int, depth int, err error) {
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

func CalculatePosition(input []string) (int, error) {
	pos := 0
	depth := 0
	for i := range input {
		deltaPos, deltaDepth, err := ProcessCommand(input[i])
		if err != nil {
			return 0, err
		}

		pos += deltaPos
		depth += deltaDepth
	}
	return pos * depth, nil
}

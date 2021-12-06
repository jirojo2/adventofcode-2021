package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	measurements := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		depth, err := strconv.Atoi(line)
		if err == nil {
			measurements = append(measurements, depth)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return measurements
}

func CalculatePart1(measurements []int) int {
	counter := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			counter++
		}
	}
	return counter
}

func CalculatePart2(measurements []int) int {
	counter := 0
	for i := 5; i < len(measurements); i += 3 {
		windowA := measurements[i-3] + measurements[i-4] + measurements[i-5]
		windowB := measurements[i] + measurements[i-1] + measurements[i-2]
		if windowB > windowA {
			counter++
		}
	}
	return counter
}

func main() {
	measurements := readInput()

	counter1 := CalculatePart1(measurements)
	counter2 := CalculatePart2(measurements)

	fmt.Printf("Part 1 solution: %v\n", counter1)
	fmt.Printf("Part 2 solution: %v\n", counter2)
}

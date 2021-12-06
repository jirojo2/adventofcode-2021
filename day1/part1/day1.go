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

func main() {
	measurments := readInput()

	counter := 0
	for i := 1; i < len(measurments); i++ {
		if measurments[i] > measurments[i-1] {
			counter++
		}
	}

	fmt.Println(counter)
}

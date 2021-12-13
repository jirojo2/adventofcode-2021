package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInput(challenge string) []string {
	var file *os.File
	if _, err := os.Stat("input.txt"); os.IsNotExist(err) {
		file, err = os.Open(fmt.Sprintf("challenges/%s/input.txt", challenge))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()

	input := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func ReadInputInts(challenge string) []int {
	input := ReadInput(challenge)
	values := make([]int, len(input))
	for i := range input {
		value, err := strconv.Atoi(input[i])
		if err == nil {
			values[i] = value
		}
	}
	return values
}

package day3

import "math"

func CalculatePowerRate(input []string) int {
	if len(input) <= 0 {
		return 0
	}
	gamma := 0
	epsilon := 0
	counter0 := make([]int, len(input[0]))
	counter1 := make([]int, len(input[0]))
	for i := 0; i < len(counter0); i++ {
		for j := range input {
			bit := input[j][i]
			if bit == '0' {
				counter0[i]++
			} else if bit == '1' {
				counter1[i]++
			}
		}
		pos := len(counter0) - i - 1
		if counter1[i] > counter0[i] {
			gamma += int(math.Pow(2, float64(pos)))
		} else {
			epsilon += int(math.Pow(2, float64(pos)))
		}
	}
	return gamma * epsilon
}

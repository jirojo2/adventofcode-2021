package day1

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
	for i := 3; i < len(measurements); i++ {
		windowA := measurements[i-1] + measurements[i-2] + measurements[i-3]
		windowB := measurements[i] + measurements[i-1] + measurements[i-2]
		if windowB > windowA {
			counter++
		}
	}
	return counter
}

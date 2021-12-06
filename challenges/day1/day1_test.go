package day1

import (
	"testing"

	"github.com/jirojo2/adventofcode-2021/utils"
)

func TestPart1(t *testing.T) {
	data := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	counter := CalculatePart1(data)
	expected := 7
	if counter != expected {
		t.Errorf("There are %d measurements that are larger than the previous measurement. %d were expected", counter, expected)
	}
}

func TestPart2(t *testing.T) {
	data := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	counter := CalculatePart2(data)
	expected := 5
	if counter != expected {
		t.Errorf("There are %d sums that are larger than the previous sum. %d were expected", counter, expected)
	}

	measurements := utils.ReadInputInts("day1")
	counter = CalculatePart2(measurements)
	expected = 1797
	if counter != expected {
		t.Errorf("There are %d sums that are larger than the previous sum. %d were expected", counter, expected)
	}
}

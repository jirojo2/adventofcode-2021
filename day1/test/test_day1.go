package main

import "testing"

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
	if counter1 != expected {
		t.Errorf("There are %d measurements that are larger than the previous measurement. %d were expected", counter, expected)
	}
}

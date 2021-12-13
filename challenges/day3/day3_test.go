package day3

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	value := CalculatePowerRate(input)
	expected := 198
	if value != expected {
		t.Errorf("So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198. Instead, got %v", value)
	}
}

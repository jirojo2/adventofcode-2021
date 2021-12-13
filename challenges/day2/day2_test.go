package day2

import "testing"

func TestPart1(t *testing.T) {
	commands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	pos, err := CalculatePositionPart1(commands)
	expected := 150
	if err != nil {
		t.Error(err)
	}
	if pos != expected {
		t.Errorf("Got %d (horizontal position multiplied by depth). %d was expected", pos, expected)
	}
}

func TestPart2(t *testing.T) {
	commands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	pos, err := CalculatePositionPart2(commands)
	expected := 900
	if err != nil {
		t.Error(err)
	}
	if pos != expected {
		t.Errorf("Got %d (horizontal position multiplied by depth). %d was expected", pos, expected)
	}
}

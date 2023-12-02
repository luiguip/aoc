package main

import "testing"

func TestPart1(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	expected := 142
	actual := part1(input)
	if actual != expected {
		t.Errorf("expected: %d actual: %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	expected := 281
	actual := part2(input)
	if actual != expected {
		t.Errorf("expected: %d actual: %d", expected, actual)
	}
}

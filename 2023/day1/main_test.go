package main

import (
	"testing"

	"github.com/luiguip/aoc/internal"
)

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

func TestFullPart1(t *testing.T) {
	input := internal.ReadFile("./input.txt")
	expected := 54927
	actual := part1(input)
	if actual != expected {
		t.Errorf("expected: %d actual: %d", expected, actual)
	}
}

func TestFullPart2(t *testing.T) {
	input := internal.ReadFile("./input.txt")
	expected := 54581
	actual := part2(input)
	if actual != expected {
		t.Errorf("expected: %d actual: %d", expected, actual)
	}
}

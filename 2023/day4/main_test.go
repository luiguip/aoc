package main

import (
	"testing"

	"github.com/luiguip/aoc/internal"
)

var inputs = []string{`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`}

var expectedsPart1 = []int{13}

var expectedsPart2 = []int{30}

func TestPart1(t *testing.T) {
	for i := range inputs {
		actual := part1(inputs[i])
		if actual != expectedsPart1[i] {
			t.Errorf("expected: %d actual %d", expectedsPart1[i], actual)
		}
	}
}

func TestPart2(t *testing.T) {
	for i := range inputs {
		actual := part2(inputs[i])
		if actual != expectedsPart2[i] {
			t.Errorf("expected: %d actual %d", expectedsPart2[i], actual)
		}
	}
}

func TestFullPart1(t *testing.T) {
	expected := 24706
	actual := part1(internal.ReadFile("./input.txt"))
	if actual != expected {
		t.Errorf("expected: %d actual %d", expected, actual)
	}
}

func TestFullPart2(t *testing.T) {
	expected := 13114317
	actual := part2(internal.ReadFile("./input.txt"))
	if actual != expected {
		t.Errorf("expected: %d actual %d", expected, actual)
	}
}

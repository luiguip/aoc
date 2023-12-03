package main

import (
	"testing"

	"github.com/luiguip/aoc/internal"
)

var inputs = []string{`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
	`$..
.11
.11
$..
..$
11.
11.
..$`,
	`12.......*..
+.........34
.......-12..
..78........
..*....60...
78..........
.......23...
....90*12...
............
2.2......12.
.*.........*
1.1.......56`,
	`12.......*..
+.........34
.......-12..
..78........
..*....60...
78.........9
.5.....23..$
8...90*12...
............
2.2......12.
.*.........*
1.1..503+.56`}

var expectedPart1 = []int{4361, 44, 413, 925}
var expectedPart2 = []int{467835, 0, 6756, 6756}

func TestPart1(t *testing.T) {
	for i := range inputs {
		actual := part1(inputs[i])
		if actual != expectedPart1[i] {
			t.Errorf("expected: %d actual %d", expectedPart1[i], actual)
		}
	}
}

func TestPart2(t *testing.T) {
	for i := range inputs {
		actual := part2(inputs[i])
		if actual != expectedPart2[i] {
			t.Errorf("expected: %d actual %d", expectedPart2[i], actual)
		}
	}
}

func TestFullPart1(t *testing.T) {
	expected := 512794
	actual := part1(internal.ReadFile("./input.txt"))
	if actual != expected {
		t.Errorf("expected: %d actual %d", expected, actual)
	}
}

func TestFullPart2(t *testing.T) {
	expected := 67779080
	actual := part2(internal.ReadFile("./input.txt"))
	if actual != expected {
		t.Errorf("expected: %d actual %d", expected, actual)
	}
}

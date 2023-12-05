package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/luiguip/aoc/internal"
)

func main() {
	text := internal.ReadFile("./2023/day4/input.txt")
	fmt.Println(part1(text))
	fmt.Println(part2(text))
}

func part1(text string) int {
	lines := strings.Split(text, "\n")
	sum := 0
	for _, l := range lines {
		sum += toScore(l)
	}
	return sum
}

func toScore(line string) int {
	cleanedLine := strings.Split(line, ": ")[1]
	rawWinning, rawActual := parse(cleanedLine)
	count := -1
	for _, a := range rawActual {
		if slices.Contains(rawWinning, a) {
			count++
		}
	}
	if count == -1 {
		return 0
	}
	return int(math.Pow(2, float64(count)))
}

func parse(line string) ([]string, []string) {
	splitted := strings.Split(line, " | ")
	rawWinning := strings.Fields(splitted[0])
	rawActual := strings.Fields(splitted[1])
	return rawWinning, rawActual
}

func part2(text string) int {
	lines := strings.Split(text, "\n")
	sum := 0
	copies := initCopies(len(lines))
	for i, l := range lines {
		copies = toScore2(l, copies, i)
		sum += copies[i]
	}
	fmt.Println(copies)
	return sum
}

func toScore2(line string, copies map[int]int, i int) map[int]int {
	cleanedLine := strings.Split(line, ": ")[1]
	rawWinning, rawActual := parse(cleanedLine)
	count := 0
	for _, a := range rawActual {
		if slices.Contains(rawWinning, a) {
			count++
			copies[i+count] += copies[i]
		}
	}
	return copies
}

func initCopies(size int) map[int]int {
	copies := make(map[int]int, size)
	for i := 0; i < size; i++ {
		copies[i] = 1
	}
	return copies
}

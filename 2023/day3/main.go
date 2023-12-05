package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/luiguip/aoc/internal"
)

func main() {
	text := internal.ReadFile("./2023/day3/input.txt")
	fmt.Println(part1(text))
	fmt.Println(part2(text))
}

func part1(text string) int {
	matrix := textToMatrix(text)
	sum := 0
	for j := range matrix {
		sum += sumPartNumbers(matrix, j)
	}
	return sum
}

func textToMatrix(text string) [][]rune {
	lines := strings.Split(text, "\n")
	matrix := [][]rune{}
	for _, l := range lines {
		runes := []rune(l)
		matrix = append(matrix, runes)
	}
	return matrix
}

func sumPartNumbers(matrix [][]rune, j int) int {
	sum := 0
	rawNumber := []rune{}
	symbolNearby := false
	for i := range matrix[j] {
		if unicode.IsDigit(matrix[j][i]) {
			rawNumber = append(rawNumber, matrix[j][i])
			if !symbolNearby {
				symbolNearby = hasSymbolNearby(matrix, i, j)
			}
		} else {
			if symbolNearby {
				number, err := strconv.Atoi(string(rawNumber))
				internal.PanicOnError(err)
				sum += number
			}
			symbolNearby = false
			rawNumber = []rune{}
		}
	}
	if symbolNearby {
		number, err := strconv.Atoi(string(rawNumber))
		internal.PanicOnError(err)
		sum += number
	}
	return sum
}

func hasSymbolNearby(matrix [][]rune, i int, j int) bool {
	for jj := j - 1; jj <= j+1; jj++ {
		for ii := i - 1; ii <= i+1; ii++ {
			if jj >= 0 && ii >= 0 && jj < len(matrix) && ii < len(matrix[j]) && !unicode.IsDigit(matrix[jj][ii]) && matrix[jj][ii] != '.' {
				return true
			}
		}
	}
	return false
}

func part2(text string) int {
	matrix := textToMatrix(text)
	sum := 0
	for j := range matrix {
		sum += sumGearLine(matrix, j)
	}
	return sum
}

func sumGearLine(matrix [][]rune, j int) int {
	sum := 0
	for i := range matrix[j] {
		if matrix[j][i] == '*' {
			sum += sumGear(matrix, j, i)
		}
	}
	return sum
}

func sumGear(matrix [][]rune, j int, i int) int {
	numbers := findNumbers(matrix, j, i)
	if len(numbers) != 2 {
		return 0
	}
	return numbers[0] * numbers[1]
}

func findNumbers(matrix [][]rune, j int, i int) []int {
	numbers := []int{}
	for jj := j - 1; jj <= j+1; jj++ {
		for ii := i - 1; ii <= i+1; ii++ {
			if jj >= 0 && ii >= 0 && jj < len(matrix) && ii < len(matrix[j]) && unicode.IsDigit(matrix[jj][ii]) {
				var rawNumber []rune
				rawNumber, ii = extractRawNumber(matrix[jj], ii)
				number, err := strconv.Atoi(string(rawNumber))
				internal.PanicOnError(err)
				numbers = append(numbers, number)
			}
		}

	}
	return numbers
}

func extractRawNumber(line []rune, ii int) ([]rune, int) {
	for ii > 0 && unicode.IsDigit(line[ii]) {
		ii--
	}
	if !unicode.IsDigit(line[ii]) {
		ii++
	}
	rawNumber := []rune{}
	for ii < len(line) && unicode.IsDigit(line[ii]) {
		rawNumber = append(rawNumber, line[ii])
		ii++
	}
	return rawNumber, ii
}

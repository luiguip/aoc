package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var numMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	data, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	text := string(data)
	fmt.Println(part1(text))
	fmt.Println(part2(text))
}

func part1(text string) int {
	lines := strings.Split(text, "\n")
	sum := 0
	for _, l := range lines {
		number := extractNumber(l)
		sum += number
	}
	return sum
}

func part2(text string) int {
	lines := strings.Split(text, "\n")
	sum := 0
	for _, l := range lines {
		line := formatLine(l)
		reversedLine := formatReversedLine(l)
		number := extractNumber2(line, reversedLine)
		sum += number
	}
	return sum
}

func extractNumber(line string) int {
	number := 0
	last := '0'
	for _, r := range line {
		if unicode.IsDigit(r) {
			if number == 0 {
				rNum := r - '0'
				number = 10 * int(rNum)
			}
			last = r
		}
	}
	rNum := last - '0'
	number += int(rNum)
	return number
}

func extractNumber2(formatted string, reverseFormatted string) int {
	number := 0
	last := '0'
	for _, r := range formatted {
		if unicode.IsDigit(r) {
			rNum := r - '0'
			number = 10 * int(rNum)
			break
		}
	}
	rf := []rune(reverseFormatted)
	for i := len(rf) - 1; i >= 0; i-- {
		if unicode.IsDigit(rf[i]) {
			last = rf[i]
			break
		}
	}
	number += int(last - '0')
	return number
}

func formatLine(line string) string {
	formattedLine := ""
	for _, r := range line {
		formattedLine += string(r)
		formattedLine = replaceNumberName(formattedLine)
	}
	return formattedLine
}

func formatReversedLine(line string) string {
	formattedLine := ""
	rs := []rune(line)
	for i := len(rs) - 1; i >= 0; i-- {
		formattedLine = string(rs[i]) + formattedLine
		formattedLine = replaceNumberName(formattedLine)
	}
	return formattedLine
}

func replaceNumberName(s string) string {
	replaced := s
	for k, v := range numMap {
		if strings.Contains(s, k) {
			replaced = strings.Replace(replaced, k, v, 1)
		}
	}
	return replaced
}

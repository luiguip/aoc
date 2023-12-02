package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luiguip/aoc/internal"
)

type color string

const (
	red   color = "red"
	blue  color = "blue"
	green color = "green"
)

type record map[color]int

type game struct {
	id      int
	records []record
}

func main() {
	text := internal.ReadFile("./2023/day2/input.txt")
	fmt.Println(part1(text))
	fmt.Println(part2(text))
}

func part1(text string) int {
	count := 0
	for _, line := range strings.Split(text, "\n") {
		game := parseGame(line)
		if gameIsValid(game) {
			count += game.id
		}
	}
	return count
}

func gameIsValid(game *game) bool {
	for _, record := range game.records {
		if !recordIsValid(record) {
			return false
		}
	}
	return true
}

func recordIsValid(record record) bool {
	return record[blue] <= 14 && record[green] <= 13 && record[red] <= 12
}

func part2(text string) int {
	count := 0
	for _, line := range strings.Split(text, "\n") {
		game := parseGame(line)
		count += powerMinimumCubes(game.records)
	}
	return count
}

func powerMinimumCubes(records []record) int {
	b, g, r := 0, 0, 0
	for _, record := range records {
		b = max(b, record[blue])
		g = max(g, record[green])
		r = max(r, record[red])
	}
	return b * g * r
}

func parseGame(line string) *game {
	id, info := parseGameId(line)
	records := parseRecords(info)
	return &game{id: id, records: records}
}

func parseGameId(line string) (int, string) {
	splitted := strings.Split(line, ":")
	rawId := strings.Trim(strings.Split(splitted[0], " ")[1], " ")
	id, err := strconv.Atoi(rawId)
	internal.PanicOnError(err)
	return id, splitted[1]
}

func parseRecords(info string) []record {
	rawRecords := strings.Split(info, ";")
	records := []record{}
	for _, rs := range rawRecords {
		records = append(records, parseRecord(rs))
	}
	return records
}

func parseRecord(rawSet string) record {
	rawRecord := strings.Split(rawSet, ",")
	record := record{}
	for _, cubeCount := range rawRecord {
		color, count := parseCubeCount(cubeCount)
		record[color] = count
	}
	return record
}

func parseCubeCount(rawCubeCount string) (color, int) {
	cc := strings.Split(strings.Trim(rawCubeCount, " "), " ")
	c := color(cc[1])
	count, err := strconv.Atoi(cc[0])
	internal.PanicOnError(err)
	return c, count
}

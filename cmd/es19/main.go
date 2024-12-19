package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

var Cache = make(map[string]int64)

func Matches(row []byte, components [][]byte) int64 {
	v, ok := Cache[string(row)]
	if ok {
		return v
	}
	v = MatchesNoCache(row, components)
	Cache[string(row)] = v
	return v
}

// Return the number of different ways to produce the string "row"
// given the set of strings identified by "components" list
func MatchesNoCache(row []byte, components [][]byte) int64 {
	if len(row) == 0 {
		return 1
	}
	sum := int64(0)
	for _, item := range components {
		// Discard items that doesn't fit the string
		if len(row) < len(item) {
			continue
		}
		if bytes.Equal(row[0:len(item)], item) {
			sum += Matches(row[len(item):], components)
		}
	}
	return sum
}

func Part2(fileName string) {
	scanner := utility.ScanFile(fileName)
	scanner.Scan()
	first := scanner.Text()
	items := strings.Split(first, ", ")
	components := make([][]byte, len(items))
	for i, item := range items {
		components[i] = []byte(item)
	}
	scanner.Scan()
	scanner.Text()
	count := int64(0)
	for scanner.Scan() {
		row := scanner.Text()
		res := Matches([]byte(row), components)
		// log.Println(row, res)
		count += res
	}
	log.Println(count)
}

func Part1(fileName string) {
	scanner := utility.ScanFile(fileName)
	scanner.Scan()
	first := scanner.Text()
	items := strings.Split(first, ", ")
	components := make([][]byte, len(items))
	for i, item := range items {
		components[i] = []byte(item)
	}
	scanner.Scan()
	scanner.Text()
	count := 0
	for scanner.Scan() {
		row := scanner.Text()
		if Matches([]byte(row), components) > 0 {
			count++
		}
	}
	log.Println(count)
}

func main() {
	Part1("cmd/es19/test.txt")
	Part1("cmd/es19/input.txt")
	Part2("cmd/es19/test.txt")
	Part2("cmd/es19/input.txt")
}

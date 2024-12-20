package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

// Return the number of different ways to produce the string "row"
// given the set of strings identified by "components" list
func Matches(row []byte, components [][]byte, cache map[string]int) int {
	if len(row) == 0 {
		return 1
	}
	v, ok := cache[string(row)]
	if ok {
		return v
	}
	v = 0
	for _, item := range components {
		if len(row) < len(item) {
			continue
		}
		if bytes.Equal(row[0:len(item)], item) {
			v += Matches(row[len(item):], components, cache)
		}
	}
	cache[string(row)] = v
	return v
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
	count := 0
	cache := make(map[string]int)
	for scanner.Scan() {
		row := scanner.Text()
		res := Matches([]byte(row), components, cache)
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
	cache := make(map[string]int)
	for scanner.Scan() {
		row := scanner.Text()
		if Matches([]byte(row), components, cache) > 0 {
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

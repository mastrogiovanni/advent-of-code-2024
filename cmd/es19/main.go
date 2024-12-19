package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

var Cache = make(map[string]int)

func Matches(row []byte, components [][]byte) int {
	v, ok := Cache[string(row)]
	if ok {
		return v
	}
	v = MatchesNoCache(row, components)
	Cache[string(row)] = v
	log.Println(len(Cache))
	return v
}

func MatchesNoCache(row []byte, components [][]byte) int {
	if len(row) == 0 {
		return 1
	}
	sum := 0
	for _, item := range components {
		if len(row) < len(item) {
			continue
		}
		// log.Println(string(row[0:len(item)]), string(item), bytes.Equal(row[0:len(item)], item))
		if bytes.Equal(row[0:len(item)], item) {
			// log.Println(string(row[0:len(item)]), string(item), bytes.Equal(row[0:len(item)], item))
			res := Matches(row[len(item):], components)
			sum += res
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
	count := 0
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
	regexpString := fmt.Sprintf("^(%s)+$", strings.Join(items, "|"))
	// log.Println(regexpString)
	r := regexp.MustCompile(regexpString)
	scanner.Scan()
	scanner.Text()
	count := 0
	for scanner.Scan() {
		row := scanner.Text()
		if r.Match([]byte(row)) {
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

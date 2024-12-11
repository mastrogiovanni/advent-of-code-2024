package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Normalize(value string) string {
	v, _ := strconv.Atoi(value)
	return fmt.Sprintf("%d", v)
}

func Rule3(value string) string {
	v, _ := strconv.Atoi(value)
	return fmt.Sprintf("%d", v*2024)
}

var Cache = make(map[string]int)

func CachedCountElements(item string, blinks int) int {
	key := fmt.Sprintf("%s-%d", item, blinks)
	if v, ok := Cache[key]; ok {
		return v
	}
	result := CountElements(item, blinks)
	Cache[key] = result
	return result
}

func CountElements(item string, blinks int) int {
	if blinks == 0 {
		return 1
	}
	if item == "0" {
		result := CachedCountElements("1", blinks-1)
		return result
	}

	if len(item)%2 == 0 {
		return CachedCountElements(Normalize(item[0:len(item)/2]), blinks-1) + CachedCountElements(Normalize(item[len(item)/2:]), blinks-1)
	}

	return CachedCountElements(Rule3(item), blinks-1)
}

func Part1(fileName string, blinks int) {
	b, err := os.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	components := strings.Split(str, " ")

	count := 0
	for _, component := range components {
		count += CachedCountElements(component, blinks)
	}
	log.Println(count)
}

func main() {
	Part1("cmd/es11/test.txt", 25)
	Part1("cmd/es11/input.txt", 25)
	Part1("cmd/es11/test.txt", 75)
	Part1("cmd/es11/input.txt", 75)
}

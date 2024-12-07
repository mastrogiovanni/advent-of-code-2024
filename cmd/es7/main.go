package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Analyze1(target int, accumulated int, items []int) bool {
	if accumulated > target {
		return false
	}
	if len(items) == 0 {
		return target == accumulated
	}
	sum := Analyze1(target, accumulated+items[0], items[1:])
	if sum {
		return sum
	}
	product := Analyze1(target, accumulated*items[0], items[1:])
	return product
}

func Analyze2(target int, accumulated int, items []int) bool {
	if accumulated > target {
		return false
	}
	if len(items) == 0 {
		return target == accumulated
	}
	v, _ := strconv.Atoi(fmt.Sprintf("%d%d", accumulated, items[0]))
	concat := Analyze2(target, v, items[1:])
	if concat {
		return concat
	}
	sum := Analyze2(target, accumulated+items[0], items[1:])
	if sum {
		return sum
	}
	product := Analyze2(target, accumulated*items[0], items[1:])
	return product
}

func Part1(scanner *bufio.Scanner) {
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		tot_and_list := strings.Split(line, ":")
		tot, _ := strconv.Atoi(tot_and_list[0])
		items_string := strings.Split(tot_and_list[1][1:], " ")
		items := utility.StrToIntList(items_string)
		if Analyze1(tot, items[0], items[1:]) {
			log.Println(line)
			result += tot
		}
	}
	log.Println(result)
}

func Part2(scanner *bufio.Scanner) {
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		tot_and_list := strings.Split(line, ":")
		tot, _ := strconv.Atoi(tot_and_list[0])
		items_string := strings.Split(tot_and_list[1][1:], " ")
		items := utility.StrToIntList(items_string)
		if Analyze2(tot, items[0], items[1:]) {
			log.Println(line)
			result += tot
		}
	}
	log.Println(result)
}

func main() {
	Part1(utility.ScanFile("cmd/es7/test.txt"))
	Part1(utility.ScanFile("cmd/es7/input.txt"))
	Part2(utility.ScanFile("cmd/es7/test.txt"))
	Part2(utility.ScanFile("cmd/es7/input.txt"))
}

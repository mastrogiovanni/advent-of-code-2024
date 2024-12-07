package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Increasing(items []int) bool {
	for i := 0; i < len(items)-1; i++ {
		if items[i] < items[i+1] && items[i+1]-items[i] < 4 {
			continue
		}
		return false
	}
	return true
}

func Decreasing(items []int) bool {
	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] && items[i]-items[i+1] < 4 {
			continue
		}
		return false
	}
	return true
}

func Correct(items []int) bool {
	if Increasing(items) {
		return true
	}
	if Decreasing(items) {
		return true
	}
	return false
}

func SafeCorrect(items []int) bool {
	if Correct(items) {
		return true
	}
	if Correct(items[1:]) {
		return true
	}
	if Correct(items[0 : len(items)-1]) {
		return true
	}
	for i := 1; i < len(items)-1; i++ {
		cp := make([]int, len(items))
		copy(cp, items)
		if Correct(append(cp[:i], cp[i+1:]...)) {
			return true
		}
	}
	return false
}

func Part1(scanner *bufio.Scanner) {
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " ")
		items := utility.StrToIntList(elements)
		if Correct(items) {
			count++
		}
	}
	log.Println(count)

}

func Part2(scanner *bufio.Scanner) {
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " ")
		items := utility.StrToIntList(elements)
		if SafeCorrect(items) {
			count++
		}
	}
	log.Println(count)

}

func main() {
	Part1(utility.ScanFile("cmd/es2/test.txt"))
	Part1(utility.ScanFile("cmd/es2/input.txt"))
	Part2(utility.ScanFile("cmd/es2/test.txt"))
	Part2(utility.ScanFile("cmd/es2/input.txt"))
}

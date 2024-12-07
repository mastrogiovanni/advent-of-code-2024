package main

import (
	"bufio"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Dist(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func Part2(scanner *bufio.Scanner) {
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, "   ")
		v1, _ := strconv.Atoi(pair[0])
		v2, _ := strconv.Atoi(pair[1])
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}
	sort.Ints(l1)
	sort.Ints(l2)

	result := 0
	for i := 0; i < len(l1); i++ {
		count := 0
		for _, x := range l2 {
			if x == l1[i] {
				count = count + 1
			}
		}
		result += l1[i] * count
	}
	log.Println(result)
}

func Part1(scanner *bufio.Scanner) {
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, "   ")
		v1, _ := strconv.Atoi(pair[0])
		v2, _ := strconv.Atoi(pair[1])
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}
	sort.Ints(l1)
	sort.Ints(l2)
	result := 0
	for i := 0; i < len(l1); i++ {
		result += Dist(l1[i], l2[i])
	}
	log.Println(result)
}

func main() {
	Part1(utility.ScanFile("cmd/es1/test.txt"))
	Part1(utility.ScanFile("cmd/es1/input.txt"))
	Part2(utility.ScanFile("cmd/es1/test.txt"))
	Part2(utility.ScanFile("cmd/es1/input.txt"))
}

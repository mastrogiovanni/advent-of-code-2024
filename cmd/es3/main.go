package main

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Eval(mul string) int {
	numbers := strings.Split(mul[4:len(mul)-1], ",")
	v1, _ := strconv.Atoi(numbers[0])
	v2, _ := strconv.Atoi(numbers[1])
	return v1 * v2
}

func Part1(buffer bytes.Buffer) {
	var validID = regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0
	for _, item := range validID.FindAll(buffer.Bytes(), -1) {
		sum += Eval(string(item))
	}
	log.Println(sum)
}

func Part2(buffer bytes.Buffer) {
	var validID = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	sum := 0
	enabled := true
	for _, item := range validID.FindAll(buffer.Bytes(), -1) {
		s := string(item)
		if s == "do()" {
			enabled = true
		} else {
			if s == "don't()" {
				enabled = false
			} else {
				if enabled {
					sum += Eval(string(item))
				}
			}
		}
	}
	log.Println(sum)
}

func FullBytes(fileName string) bytes.Buffer {
	scanner := utility.ScanFile(fileName)
	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}
	return buffer
}

func main() {

	Part1(FullBytes("cmd/es3/test.txt"))
	Part1(FullBytes("cmd/es3/input.txt"))
	Part2(FullBytes("cmd/es3/test.txt"))
	Part2(FullBytes("cmd/es3/input.txt"))

}

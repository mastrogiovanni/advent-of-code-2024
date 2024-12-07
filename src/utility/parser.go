package utility

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ScanFile(fileName string) *bufio.Scanner {
	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func StrToIntList(elements []string) []int {
	result := make([]int, len(elements))
	for i := 0; i < len(elements); i++ {
		v, _ := strconv.Atoi(elements[i])
		result[i] = v
	}
	return result
}

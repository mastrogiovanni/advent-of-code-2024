package utility

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func StrToIntList(elements []string) []int {
	result := make([]int, len(elements))
	for i := 0; i < len(elements); i++ {
		v, _ := strconv.Atoi(elements[i])
		result[i] = v
	}
	return result
}

func IntToStrList(elements []int) []string {
	result := make([]string, len(elements))
	for i := 0; i < len(elements); i++ {
		result[i] = fmt.Sprintf("%d", elements[i])
	}
	return result
}

func UIntToStrList(elements []uint) []string {
	result := make([]string, len(elements))
	for i := 0; i < len(elements); i++ {
		result[i] = fmt.Sprintf("%d", elements[i])
	}
	return result
}

func UInt32ToStrList(elements []uint32) []string {
	result := make([]string, len(elements))
	for i := 0; i < len(elements); i++ {
		result[i] = fmt.Sprintf("%d", elements[i])
	}
	return result
}

func UInt64ToStrList(elements []uint64) []string {
	result := make([]string, len(elements))
	for i := 0; i < len(elements); i++ {
		result[i] = fmt.Sprintf("%d", elements[i])
	}
	return result
}

func ScanFile(fileName string) *bufio.Scanner {
	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func FullBytes(fileName string) bytes.Buffer {
	scanner := ScanFile(fileName)
	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}
	return buffer
}

func FileLines(fileName string) []string {
	scanner := ScanFile(fileName)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	return fileLines
}

func FileLinesFromScanner(scanner *bufio.Scanner) []string {
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	return fileLines
}

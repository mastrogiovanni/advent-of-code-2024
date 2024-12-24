package main

import (
	"log"
	"strconv"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func Secret(secret int) int {
	secret = prune(mix(secret<<6, secret))  // secret * 64
	secret = prune(mix(secret>>5, secret))  // secret / 32
	secret = prune(mix(secret<<11, secret)) // secret * 2048
	return secret
}

func Part1(fileName string) {
	lines := utility.FileLines(fileName)
	sum := 0
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		for i := 0; i < 2000; i++ {
			secret = Secret(secret)
		}
		sum += secret
	}
	log.Println(sum)
}

func Unit(value int) int {
	s := strconv.Itoa(value)
	v, _ := strconv.Atoi(string(s[len(s)-1]))
	return v
}

func GetSequence(secret int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < 2000; i++ {
		beforeUnit := Unit(secret)
		after := Secret(secret)
		afterUnit := Unit(after)
		result = append(result, []int{afterUnit, afterUnit - beforeUnit})
		secret = after
	}
	return result
}

func GetAllSequences(secrets []int) {
	result := make([][][]int, 0)
	for _, secret := range secrets {
		result = append(result, GetSequence(secret))
	}
	max := 0

	for selected := 0; selected < len(result); selected++ {

		// i: start of the sequence of 4
		for i := 0; i < len(result[selected])-4; i++ {

			current := 0                        // current
			current += result[selected][i+3][0] // sasdas

			// log.Println("0", i, result[0][i+3][0])

			// seller j
			for j := 1; j < len(result); j++ {

				if j == selected {
					continue
				}

				// start index k
				for k := 0; k < len(result[j])-4; k++ {

					found := true

					// shift z
					for z := 0; z < 4; z++ {
						if result[j][k+z][1] != result[selected][i+z][1] {
							found = false
							break
						}
					}

					if found {
						current += result[j][k+3][0]
						break
					}

				}

			}

			if current > max {
				max = current
				log.Println(max, result[selected][i:i+4])
			}

		}
	}
}

func Part2(fileName string) {
	lines := utility.FileLines(fileName)
	secrets := make([]int, 0)
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		secrets = append(secrets, secret)
	}
	GetAllSequences(secrets)
}

func main() {
	// Part1("cmd/es22/test.txt")
	// Part1("cmd/es22/input.txt")
	// Part2("cmd/es22/test2.txt")
	Part2("cmd/es22/input.txt")

	// log.Println(GetSequence(123))
}

// 1801 basso

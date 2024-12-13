package main

import (
	"log"
	"regexp"
	"strconv"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func GetPairWithRegex(text string, regexp *regexp.Regexp) (int, int) {
	r := regexp.FindStringSubmatch(text)
	x, _ := strconv.Atoi(r[1])
	y, _ := strconv.Atoi(r[2])
	return x, y
}

func Resolve(fileInput string, stride utility.Point) {

	result := 0

	regButtonA := regexp.MustCompile(`Button A: X\+(?P<X>\d+), Y\+(?P<Y>\d+)`)
	regButtonB := regexp.MustCompile(`Button B: X\+(?P<X>\d+), Y\+(?P<Y>\d+)`)
	regTarget := regexp.MustCompile(`Prize: X=(?P<X>\d+), Y=(?P<Y>\d+)`)

	scanner := utility.ScanFile(fileInput)
	for scanner.Scan() {

		xA, yA := GetPairWithRegex(scanner.Text(), regButtonA)
		scanner.Scan()

		xB, yB := GetPairWithRegex(scanner.Text(), regButtonB)
		scanner.Scan()

		xT, yT := GetPairWithRegex(scanner.Text(), regTarget)

		xT += stride.X
		yT += stride.Y

		beta := (xA*yT - yA*xT) / (xA*yB - xB*yA)
		alpha := (xT - beta*xB) / xA

		if xT == alpha*xA+beta*xB && yT == alpha*yA+beta*yB {
			result += alpha*3 + beta
		}

		if !scanner.Scan() {
			break
		}
		scanner.Text()

	}
	log.Println(result)
}

func main() {
	Resolve("cmd/es13/test.txt", utility.Point{
		X: 0,
		Y: 0,
	})
	Resolve("cmd/es13/input.txt", utility.Point{
		X: 0,
		Y: 0,
	})
	Resolve("cmd/es13/test.txt", utility.Point{
		X: 10000000000000,
		Y: 10000000000000,
	})
	Resolve("cmd/es13/input.txt", utility.Point{
		X: 10000000000000,
		Y: 10000000000000,
	})

}

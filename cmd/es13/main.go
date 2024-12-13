package main

import (
	"errors"
	"log"
	"regexp"
	"strconv"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Track(start utility.Point, a, b utility.Point, target utility.Point) utility.Point {
	delta := max(target.X/b.X, target.Y/b.Y) + 1

	for dB := delta; dB >= 0; dB = dB - 1 {
		log.Println(dB)
		finalX := target.X - dB*b.X
		finalY := target.Y - dB*b.Y
		if finalX < 0 {
			dX := (dB*b.X - target.X) / a.X
			if finalY < 0 {
				dY := (dB*b.Y - target.Y) / a.Y
				dB -= min(dX, dY)
			} else {
				dB -= dX
			}
		} else {
			if finalY < 0 {
				dY := (dB*b.Y - target.Y) / a.Y
				dB -= dY
			}
		}

		if finalX%a.X != 0 || finalY%a.Y != 0 {
			continue
		}
		dA := finalX / a.X
		if finalY/a.Y == dA {
			return utility.Point{
				X: dA,
				Y: dB,
			}
		}
	}
	return utility.Point{
		X: -1,
		Y: -1,
	}
}

// solveSystem2x2 solves the system Ax = y for a 2x2 matrix A and a vector y
func solveSystem2x2(A [2][2]int, y [2]int) ([2]int, error) {
	// Calculate the determinant of A
	det := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	if det == 0 {
		return [2]int{}, errors.New("matrix A is singular and cannot be inverted")
	}

	// Solve for x by multiplying invA with y
	x := [2]int{
		(A[1][1]*y[0] - A[0][1]*y[1]) / det,
		(-A[1][0]*y[0] + A[0][0]*y[1]) / det,
	}

	return x, nil
}

func Part1(fileInput string, stride utility.Point) {

	result := 0

	regButtonA := regexp.MustCompile(`Button A: X\+(?P<X>\d+), Y\+(?P<Y>\d+)`)
	regButtonB := regexp.MustCompile(`Button B: X\+(?P<X>\d+), Y\+(?P<Y>\d+)`)
	regTarget := regexp.MustCompile(`Prize: X=(?P<X>\d+), Y=(?P<Y>\d+)`)

	scanner := utility.ScanFile(fileInput)
	for scanner.Scan() {
		mA := regButtonA.FindStringSubmatch(scanner.Text())
		xA, _ := strconv.Atoi(mA[1])
		yA, _ := strconv.Atoi(mA[2])
		scanner.Scan()
		mB := regButtonB.FindStringSubmatch(scanner.Text())
		xB, _ := strconv.Atoi(mB[1])
		yB, _ := strconv.Atoi(mB[2])
		scanner.Scan()
		mT := regTarget.FindStringSubmatch(scanner.Text())
		xT, _ := strconv.Atoi(mT[1])
		yT, _ := strconv.Atoi(mT[2])

		xT += stride.X
		yT += stride.Y

		beta := (xA*yT - yA*xT) / (xA*yB - xB*yA)
		alpha := (xT - beta*xB) / xA

		x := []int{
			alpha,
			beta,
		}

		if xT == alpha*xA+beta*xB && yT == alpha*yA+beta*yB {
			result += x[0]*3 + x[1]
		}
		if !scanner.Scan() {
			break
		}
		scanner.Text()
	}
	log.Println(result)
}

func main() {
	Part1("cmd/es13/test.txt", utility.Point{
		X: 0,
		Y: 0,
	})
	Part1("cmd/es13/input.txt", utility.Point{
		X: 0,
		Y: 0,
	})
	Part1("cmd/es13/test.txt", utility.Point{
		X: 10000000000000,
		Y: 10000000000000,
	})
	Part1("cmd/es13/input.txt", utility.Point{
		X: 10000000000000,
		Y: 10000000000000,
	})

}

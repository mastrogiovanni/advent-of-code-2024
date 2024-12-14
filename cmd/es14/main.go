package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Part1(fileName string, w, h int, steps int) {
	scanner := utility.ScanFile(fileName)

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	cx := (w - 1) / 2
	cy := (h - 1) / 2

	for scanner.Scan() {
		line := scanner.Text()
		posAndSpeed := strings.Split(line, " ")
		pos := strings.Split(posAndSpeed[0][2:], ",")
		speed := strings.Split(posAndSpeed[1][2:], ",")

		px, _ := strconv.Atoi(pos[0])
		py, _ := strconv.Atoi(pos[1])
		vx, _ := strconv.Atoi(speed[0])
		vy, _ := strconv.Atoi(speed[1])

		px = (px + steps*vx) % w
		py = (py + steps*vy) % h

		for px < 0 {
			px += w
		}

		for py < 0 {
			py += h
		}

		if px < cx {
			if py < cy {
				q1 = q1 + 1
			} else if py > cy {
				q2 = q2 + 1
			}
		} else if px > cx {
			if py < cy {
				q3 = q3 + 1
			} else if py > cy {
				q4 = q4 + 1
			}
		}
		log.Println("Q", q1, q2, q3, q4)

	}
	log.Println(q1 * q2 * q3 * q4)
}

func Part2(fileName string, w, h int) {
	scanner := utility.ScanFile(fileName)

	// cx := (w - 1) / 2
	// cy := (h - 1) / 2

	robots := make([]utility.Point, 0)

	speeds := make([]utility.Point, 0)

	for scanner.Scan() {
		line := scanner.Text()
		posAndSpeed := strings.Split(line, " ")
		pos := strings.Split(posAndSpeed[0][2:], ",")
		speed := strings.Split(posAndSpeed[1][2:], ",")

		px, _ := strconv.Atoi(pos[0])
		py, _ := strconv.Atoi(pos[1])
		vx, _ := strconv.Atoi(speed[0])
		vy, _ := strconv.Atoi(speed[1])

		robots = append(robots, utility.Point{X: px, Y: py})
		speeds = append(speeds, utility.Point{X: vx, Y: vy})
	}

	count := 0
	steps := 1 // 180 * 55

	for ; ; count += steps {

		upLeft := image.Point{0, 0}
		lowRight := image.Point{w, h}
		img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

		cyan := color.RGBA{0, 0, 0, 0xff}

		// busy := make(map[utility.Point]bool)
		for i := 0; i < len(robots); i++ {
			px := (robots[i].X + steps*speeds[i].X) % w
			py := (robots[i].Y + steps*speeds[i].Y) % h
			for px < 0 {
				px += w
			}
			for py < 0 {
				py += h
			}
			robots[i].X = px
			robots[i].Y = py
			img.Set(px, py, cyan)
			// busy[utility.Point{X: px, Y: py}] = true
		}

		f, _ := os.Create(fmt.Sprintf("pippo/%06d.png", count))
		png.Encode(f, img)

		// // Check symmetry
		// symmetric := true
		// for i := 0; i < len(robots); i++ {
		// 	px := robots[i].X
		// 	py := robots[i].Y
		// 	sx := px
		// 	if px < cx {
		// 		sx = cx + (cx - px)
		// 	}
		// 	if px > cx {
		// 		sx = cx - (px - cx)
		// 	}
		// 	if _, ok := busy[utility.Point{X: sx, Y: py}]; !ok {
		// 		symmetric = false
		// 		break
		// 	}
		// }
		// if symmetric {
		// 	break
		// }

	}

	log.Println(count + 1)

}

func main() {
	// Part1("cmd/es14/test.txt", 11, 7, 100)
	// Part1("cmd/es14/input.txt", 101, 103, 100)

	Part2("cmd/es14/input.txt", 101, 103)

}

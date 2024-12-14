package main

import (
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
	}
	log.Println(q1 * q2 * q3 * q4)
}

func ComponentSize(point utility.Point, globalVisited map[utility.Point]bool, width, height int) int {
	visited := make(map[utility.Point]bool)
	size := 1
	for toVisit := []utility.Point{point}; len(toVisit) > 0; {
		point := toVisit[0]
		toVisit = toVisit[1:]
		if _, ok := visited[point]; ok {
			continue
		}
		visited[point] = true
		for direction := 0; direction < 4; direction++ {
			if !(point.X+utility.Directions[direction][0] >= 0 &&
				point.X+utility.Directions[direction][0] < width &&
				point.Y+utility.Directions[direction][1] >= 0 &&
				point.Y+utility.Directions[direction][1] < height) {
				continue
			}
			adjacent := utility.Point{
				X: point.X + utility.Directions[direction][0],
				Y: point.Y + utility.Directions[direction][1],
			}
			if _, ok := globalVisited[adjacent]; !ok {
				// Not colored
				continue
			}
			if _, ok := visited[adjacent]; ok {
				// Already visited
				continue
			}
			toVisit = append(toVisit, adjacent)
			size++
		}
	}
	return size
}

func MakePicture(globalVisited map[utility.Point]bool, width, height int) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	black := color.RGBA{0, 0, 0, 0xff}
	for point := range globalVisited {
		img.Set(point.X, point.Y, black)
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func Part2(fileName string, w, h int) {
	scanner := utility.ScanFile(fileName)

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

	for ; ; count++ {

		globalVisited := make(map[utility.Point]bool)
		for i := 0; i < len(robots); i++ {
			px := (robots[i].X + speeds[i].X) % w
			py := (robots[i].Y + speeds[i].Y) % h
			for px < 0 {
				px += w
			}
			for py < 0 {
				py += h
			}
			robots[i].X = px
			robots[i].Y = py
			globalVisited[utility.Point{X: px, Y: py}] = true
		}

		maxSize := 0
		for p := range globalVisited {
			size := ComponentSize(p, globalVisited, w, h)
			if size > maxSize {
				maxSize = size
			}
		}
		if maxSize > 200 {
			log.Println(count)
			MakePicture(globalVisited, w, h)
			break
		}

	}

}

func main() {
	Part1("cmd/es14/test.txt", 11, 7, 100)
	Part1("cmd/es14/input.txt", 101, 103, 100)
	Part2("cmd/es14/input.txt", 101, 103)
}

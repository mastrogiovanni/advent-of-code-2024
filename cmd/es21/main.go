package main

import (
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

var Keyboard1 = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var Keyboard2 = [][]rune{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
}

func FindPosition(symbol rune, keyboard [][]rune) utility.Point {
	for y := 0; y < len(keyboard); y++ {
		for x := 0; x < len(keyboard[y]); x++ {
			if symbol == keyboard[y][x] {
				return utility.Point{X: x, Y: y}
			}
		}
	}
	return utility.Point{}
}

type Context struct {
	Current     utility.Point
	A           utility.Point
	Empty       utility.Point
	Next        utility.Point
	Sequence    string
	Index       int
	Keyboard    [][]rune
	CurrentPath string
	Path        []string
}

func (c *Context) Move(direction utility.Direction) {
	c.Current.Move(direction)
	switch direction {
	case utility.North:
		c.CurrentPath += "^"
	case utility.East:
		c.CurrentPath += ">"
	case utility.West:
		c.CurrentPath += "<"
	case utility.South:
		c.CurrentPath += "v"
	}
}

func NewContext(sequence string, keyboard [][]rune) *Context {
	A := FindPosition('A', keyboard)
	Empty := FindPosition(' ', keyboard)
	return &Context{
		Current:     A,
		A:           A,
		Empty:       Empty,
		Next:        FindPosition(rune(sequence[0]), keyboard),
		Sequence:    sequence,
		Index:       0,
		Keyboard:    keyboard,
		CurrentPath: "",
		Path:        make([]string, 0),
	}
}

func IsInSegment(point, a, b utility.Point) bool {
	if a.X == b.X {
		if point.X != a.X {
			return false
		}
		return point.Y >= min(a.Y, b.Y) && point.Y <= max(a.Y, b.Y)
	}
	if a.Y == b.Y {
		if point.Y != a.Y {
			return false
		}
		return point.X >= min(a.X, b.X) && point.X <= max(a.X, b.X)
	}
	log.Println("Error")
	return false
}

func FindSequence(context *Context) {
	if context.Current == context.Empty {
		return
	}
	if context.Current == context.Next {
		context.CurrentPath += "A"
		context.Index = context.Index + 1
		if context.Index >= len(context.Sequence) {
			context.Path = append(context.Path, context.CurrentPath)
			return
		}
		context.Next = FindPosition(rune(context.Sequence[context.Index]), context.Keyboard)
		FindSequence(context)
		return
	}
	if context.Current.Y == context.Next.Y {
		if context.Current.X < context.Next.X {
			context.Move(utility.West)
		} else {
			context.Move(utility.East)
		}
		FindSequence(context)
		return
	}
	if context.Current.X == context.Next.X {
		if context.Current.Y < context.Next.Y {
			context.Move(utility.South)
		} else {
			context.Move(utility.North)
		}
		FindSequence(context)
		return
	}
	if context.Current.X < context.Next.X && context.Current.Y < context.Next.Y {
		angle1 := utility.Point{X: context.Current.X, Y: context.Next.Y}
		angle2 := utility.Point{X: context.Next.X, Y: context.Current.Y}
		if IsInSegment(context.Empty, context.Current, angle1) || IsInSegment(context.Empty, angle1, context.Next) {
			// Must use angle2
		} else {
			// can use angle1
		}

	}
	if context.Current.X < context.Next.X && context.Current.Y > context.Next.Y {

	}
	if context.Current.X > context.Next.X && context.Current.Y < context.Next.Y {

	}
	if context.Current.X > context.Next.X && context.Current.Y > context.Next.Y {

	}
}

func main() {
	rows := utility.FileLines("cmd/es21/test.txt")
	for _, row := range rows {
		log.Println(row)
		for _, symbol := range row {
			log.Println(string(symbol), FindPosition(symbol, Keyboard1))
		}
		empty := FindPosition(' ', keyboard)
		log.Println(FindSequence(row, Keyboard1))

	}

}

package main

import (
	"log"
	"slices"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type Registers struct {
	A int
	B int
	C int
}

func (r *Registers) GetComboOperand(index int) int {
	if index < 4 {
		return index
	}
	if index == 4 {
		return r.A
	}
	if index == 5 {
		return r.B
	}
	if index == 6 {
		return r.C
	}
	return 0
}

func Processor(program []int, registers *Registers) []int {
	t := registers.A
	output := make([]int, 0)
	for pc := 0; pc < len(program); {
		instruction := program[pc]
		operand := program[pc+1]
		combo := registers.GetComboOperand(operand)
		pc = pc + 2
		switch instruction {
		case 0: // adv: OK
			registers.A >>= combo
		case 1: // bxl
			registers.B ^= operand
		case 2: // bst
			registers.B = combo % 8
		case 3: // jnz
			if registers.A != 0 {
				pc = operand
				continue
			}
		case 4: // bxc
			registers.B ^= registers.C
		case 5: // out
			output = append(output, combo%8)
		case 6: // bdv
			registers.B = (registers.A >> combo)
		case 7: // cdv
			registers.C = (registers.A >> combo)
		}
	}
	log.Println(t, output)
	return output
}

func Evaluate(a, b, c int, program []int) {
	log.Println()
	log.Println("Evaluate")
	registers := Registers{
		A: a,
		B: b,
		C: c,
	}
	output := Processor(program, &registers)
	outputStringList := utility.IntToStrList(output)
	outputString := strings.Join(outputStringList, ",")
	log.Println("Output:", outputString)
	log.Println("Registers:", registers)
}

func Part2(b, c int, program []int) {
	a := 0
	for i := len(program) - 1; i >= 0; i-- {
		// make room for 3 bits by shifting to the left
		a <<= 3
		// check incrementally only the latest bits,
		// until we find the right value
		registers := Registers{
			A: a,
			B: b,
			C: c,
		}
		for !slices.Equal(Processor(program, &registers), program[i:]) {
			a++
			registers = Registers{
				A: a,
				B: b,
				C: c,
			}
		}
	}
	log.Println(a)
}

func main() {
	// Evaluate(0, 0, 9, []int{2, 6})
	// Evaluate(10, 0, 0, []int{5, 0, 5, 1, 5, 4})
	// Evaluate(2024, 0, 0, []int{0, 1, 5, 4, 3, 0})
	// Evaluate(0, 29, 0, []int{1, 7})
	// Evaluate(0, 2024, 43690, []int{4, 0})
	// Evaluate(729, 0, 0, []int{0, 1, 5, 4, 3, 0})

	// Part 1
	// Evaluate(30899381, 0, 0, []int{2, 4, 1, 1, 7, 5, 4, 0, 0, 3, 1, 6, 5, 5, 3, 0})

	// part2
	// FindCopy(0, 0, 0, []int{0, 3, 5, 4, 3, 0})
	Part2(0, 0, []int{2, 4, 1, 1, 7, 5, 4, 0, 0, 3, 1, 6, 5, 5, 3, 0})

	// 4: a
	// 5: b
	// 6: c

	// b = a 				// 2, 4
	// b = b ^ 1			// 1, 1
	// c = a >> b 			// 7, 5
	// b = b ^ c 			// 4, 0,
	// a = a >> 3  			// 0, 3,
	// b = b ^ 6  			// 1, 6,
	// output(b%8)			// 5, 5,
	// if a != 0 return 0 	// 3, 0
}

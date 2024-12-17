package main

import (
	"log"
	"reflect"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type Registers struct {
	A uint64
	B uint64
	C uint64
}

func (r *Registers) GetComboOperand(index uint64) uint64 {
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
	log.Fatalf("Error")
	return 0
}

func Processor(program []uint64, registers *Registers, limit int) []uint64 {
	output := make([]uint64, 0)
	var pc uint64 = 0
	for int(pc) < len(program) {

		instruction := program[pc]
		// log.Println("Instruction", instruction)
		operand := program[pc+1]
		pc = pc + 2

		switch instruction {
		case 0: // adv: OK
			registers.A = (registers.A >> registers.GetComboOperand(operand))
		case 1: // bxl
			registers.B = (registers.B ^ operand)
		case 2: // bst
			registers.B = (registers.GetComboOperand(operand) % 8)
		case 3: // jnz
			if registers.A != 0 {
				pc = operand
			}
		case 4: // bxc
			registers.B = (registers.B ^ registers.C)
		case 5: // out
			output = append(output, registers.GetComboOperand(operand)%8)
			if limit > 0 && len(output) > limit {
				return []uint64{}
			}
		case 6: // bdv
			registers.B = (registers.A >> registers.GetComboOperand(operand))
		case 7: // cdv
			registers.C = (registers.A >> registers.GetComboOperand(operand))
		}
	}
	return output
}

// Register A: 729
// Register B: 0
// Register C: 0

// Program: 0,1,5,4,3,0

func Evaluate(a, b, c uint64, program []uint64) {
	log.Println()
	log.Println("Evaluate")
	registers := Registers{
		A: a,
		B: b,
		C: c,
	}
	output := Processor(program, &registers, -1)
	outputStringList := utility.UInt64ToStrList(output)
	outputString := strings.Join(outputStringList, ",")
	log.Println("Output:", outputString)
	log.Println("Registers:", registers)
}

func FindCopy(a, b, c uint64, program []uint64) {
	i := uint64(2.8e+14)
	for ; ; i++ {
		registers := Registers{
			A: uint64(i),
			B: b,
			C: c,
		}
		output := Processor(program, &registers, len(program))
		// log.Println(uint64(i), len(program), len(output))
		if len(output) != len(program) {
			// log.Println("Skip")
			continue
		}
		if reflect.DeepEqual(output, program) {
			// log.Println("Found", i)
			break
		}
	}
}

func main() {
	// Evaluate(0, 0, 9, []uint64{2, 6})
	// Evaluate(10, 0, 0, []uint64{5, 0, 5, 1, 5, 4})
	// Evaluate(2024, 0, 0, []uint64{0, 1, 5, 4, 3, 0})
	// Evaluate(0, 29, 0, []uint64{1, 7})
	// Evaluate(0, 2024, 43690, []uint64{4, 0})
	// Evaluate(729, 0, 0, []uint64{0, 1, 5, 4, 3, 0})

	// // part 1
	// // 2,4,1,1,7,5,4,0,0,3,1,6,5,5,3,0
	// Evaluate(30899381, 0, 0, []uint64{2, 4, 1, 1, 7, 5, 4, 0, 0, 3, 1, 6, 5, 5, 3, 0})

	// part2
	// FindCopy(0, 0, 0, []uint64{0, 3, 5, 4, 3, 0})
	FindCopy(30899381, 0, 0, []uint64{2, 4, 1, 1, 7, 5, 4, 0, 0, 3, 1, 6, 5, 5, 3, 0})

	// b = a 				// 2, 4
	// b = b ^ 1			// 1, 1
	// c = a >> b 			// 7, 5
	// b = b ^ c 			// 4, 0,
	// a = a >> 3  		// 0, 3,
	// b = b ^ 6  			// 1, 6,
	// output(b%8)			// 5, 5,
	// if a != 0 return 0 	// 3, 0
}

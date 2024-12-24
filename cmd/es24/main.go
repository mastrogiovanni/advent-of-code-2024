package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
	"golang.org/x/exp/maps"
)

type Port struct {
	In1       string
	In2       string
	Operation string
	Out       string
	Level     int
}

// e.g.: ntg XOR fgs -> mjb
func ParsePort(row string) *Port {
	port := Port{}
	items := strings.Split(row, " -> ")
	port.Out = items[1]
	items = strings.Split(items[0], " ")
	port.In1 = items[0]
	port.Operation = items[1]
	port.In2 = items[2]
	return &port
}

// Compute output: -1 if no all inputs are provided
func Compute(values map[string]int, port *Port) bool {
	_, okOut := values[port.Out]
	if okOut {
		return false
	}
	v1, ok1 := values[port.In1]
	if !ok1 {
		return false
	}
	v2, ok2 := values[port.In2]
	if !ok2 {
		return false
	}
	if port.Operation == "AND" {
		values[port.Out] = v1 & v2
		return true
	}
	if port.Operation == "OR" {
		values[port.Out] = v1 | v2
		return true
	}
	if port.Operation == "XOR" {
		values[port.Out] = v1 ^ v2
		return true
	}
	return false
}

func BinaryToNumber(binary string) int64 {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return i
}

func NumberToBinaryString(value int64) string {
	return strconv.FormatInt(value, 2)
}

func ExtractVariable(variable byte, values map[string]int) string {
	result := make([]string, 0)
	for key, value := range values {
		if key[0] != variable {
			continue
		}
		result = append(result, key+fmt.Sprintf("%s,%d", key, value))
	}
	slices.Sort(result)
	response := ""
	for _, v := range result {
		response = strings.Split(v, ",")[1] + response
	}
	return response
}

func Parse(fileName string) (map[string]int, map[string]*Port) {
	rows := utility.FileLines(fileName)
	init := make(map[string]int)
	ports := make(map[string]*Port)
	initializying := true
	for _, row := range rows {
		if row == "" {
			initializying = false
			continue
		}
		if initializying {
			items := strings.Split(row, ": ")
			value, _ := strconv.Atoi(items[1])
			init[items[0]] = value
		} else {
			p := ParsePort(row)
			ports[p.Out] = p
		}
	}
	return init, ports
}

func CompleteCompute(init map[string]int, ports map[string]*Port) {
	for {
		changed := false
		for _, port := range ports {
			if Compute(init, port) {
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}

func CountOnes(binary string) int {
	result := 0
	for _, i := range binary {
		if i == '1' {
			result++
		}
	}
	return result
}

func Part1(fileName string) {
	init, ports := Parse(fileName)
	CompleteCompute(init, ports)
	log.Println(BinaryToNumber(ExtractVariable('z', init)))
}

func Invert(i int, ports map[string]*Port) (string, string) {

	for _, port := range ports {
		if port.Operation != "XOR" {
			continue
		}
		if i > 0 {
			i--
			continue
		}
		log.Println("Inverting", port)
		tmp := port.In2
		port.In2 = port.In1
		port.In1 = tmp
		return port.In2, port.In1
	}
	return "", ""
}

func Part2(fileName string) {
	init, ports := Parse(fileName)
	x := ExtractVariable('x', init)
	y := ExtractVariable('y', init)
	targetZ := NumberToBinaryString(BinaryToNumber(x) + BinaryToNumber(y))
	onesTarget := CountOnes(targetZ)
	log.Println(x + " + " + y + " = " + targetZ)

	inverted := 0
	invlist := make([]string, 0)
	z := ""
	for {
		variables := make(map[string]int)
		for k, v := range init {
			variables[k] = v
		}
		CompleteCompute(variables, ports)
		z = ExtractVariable('z', variables)
		z = z[len(z)-len(targetZ):]
		log.Println(targetZ, z)
		if z == targetZ {
			break
		}
		log.Println(targetZ, z, CountOnes(z), onesTarget)
		if CountOnes(z) == onesTarget {
			break
		}
		a, b := Invert(inverted, ports)
		if a == "" && b == "" {
			break
		}
		inverted++
		invlist = append(invlist, a, b)
	}
	// if z == targetZ {
	// 	log.Println(invlist)
	// }
	log.Println(targetZ, z, invlist)

	// log.Println("Target", targetZ, "Obtained", z)
}

func Mermaid(fileName string) {
	_, ports := Parse(fileName)
	output := "flowchart RL\n"
	pports := maps.Values(ports)
	for i, port := range pports {
		if port.In1[0] == 'x' || port.In1[0] == 'y' {
			output += fmt.Sprintf("\t%v --> %v(%s-%v)\n", port.In1, i, port.Operation, i)
		}
		if port.In2[0] == 'x' || port.In2[0] == 'y' {
			output += fmt.Sprintf("\t%v --> %v(%s-%v)\n", port.In2, i, port.Operation, i)
		}
		if port.Out[0] == 'z' {
			output += fmt.Sprintf("\t%v(%s-%v) --> %s\n", i, port.Operation, i, port.Out)
		}
		for j, portj := range pports {
			if portj.Out == port.In1 || portj.Out == port.In2 {
				output += fmt.Sprintf("\t%v(%s-%v) --> %v(%s-%v)\n", j, portj.Operation, j, i, port.Operation, i)
			}
		}
	}
	fmt.Println(output)
}

func main() {
	// Part1("cmd/es24/test.txt")
	// Part1("cmd/es24/input.txt")
	Mermaid("cmd/es24/input.txt")
	// Part2("cmd/es24/test.txt")
}

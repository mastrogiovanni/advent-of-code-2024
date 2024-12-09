package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type Slot struct {
	File   int // -1 means free
	Length int
	Next   *Slot
	Prev   *Slot
}

func GetSize(ch byte) int {
	switch ch {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return -1
}

func CreateDisk(diskMap []byte, fileIndex int, isFile bool) *Slot {
	if len(diskMap) == 0 {
		return nil
	}
	var next *Slot = nil
	if isFile {
		next = CreateDisk(diskMap[1:], fileIndex+1, !isFile)
	} else {
		next = CreateDisk(diskMap[1:], fileIndex, !isFile)
	}
	size := GetSize(diskMap[0])
	if isFile {
		slot := &Slot{
			File:   fileIndex,
			Length: size,
			Next:   next,
		}
		if next != nil {
			next.Prev = slot
		}
		return slot
	} else {
		slot := &Slot{
			File:   -1,
			Length: size,
			Next:   next,
		}
		if next != nil {
			next.Prev = slot
		}
		return slot
	}
}

func Last(slot *Slot) *Slot {
	last := slot
	prev := slot
	for ; last != nil; last = last.Next {
		prev = last
	}
	return prev
}

func FirstFree(slot *Slot) *Slot {
	free := slot
	for ; free != nil && free.File != -1; free = free.Next {
	}
	return free
}

func Dump(slot *Slot) string {
	var buffer bytes.Buffer
	for item := slot; item != nil; item = item.Next {
		if item.File < 0 {
			for i := 0; i < item.Length; i++ {
				buffer.WriteString(".")
			}
		} else {
			for i := 0; i < item.Length; i++ {
				buffer.WriteString(fmt.Sprintf("%d", item.File))
			}
		}
	}
	return buffer.String()
}

func Compact(slot *Slot) {
	last := Last(slot)
	free := FirstFree(slot)

	for free != nil {
		// log.Println(Dump(slot))
		// Skip to next free
		if free.File != -1 {
			free = free.Next
			continue
		}

		// Copy bytes
		if last.Length < free.Length {
			// Copy all slot into the free
			// Create another free slot with remaining
			newFree := &Slot{
				File:   -1,
				Length: free.Length - last.Length,
				Next:   free.Next,
			}
			if free.Next != nil {
				free.Next.Prev = newFree
			}
			free.File = last.File
			free.Length = last.Length
			free.Next = newFree
			newFree.Prev = free
			last.File = -1
		} else if last.Length == free.Length {
			free.File = last.File
			last.File = -1
		} else {
			// The slot doesn't fit the free: copy and reduce
			free.File = last.File
			last.Length = last.Length - free.Length
		}
		// Move to next free
		for free = free.Next; free != nil && free.File != -1; free = free.Next {
			if free == last {
				free = nil
				break
			}
		}
		if free == nil {
			break
		}
		// Move to previous last
		for ; last != nil && last.File == -1; last = last.Prev {
			if free == last {
				last = nil
				break
			}
		}
		if last == nil {
			break
		}
	}
}

func MergeFree(slot *Slot) {
	for item := slot; item != nil; item = item.Next {
		if item.File >= 0 {
			continue
		}
		if item.Next == nil {
			continue
		}
		if item.Next.File < 0 {
			item.Length = item.Length + item.Next.Length
			item.Next = item.Next.Next
			if item.Next != nil {
				item.Next.Prev = item
			}
			item = item.Prev
		}
	}
}

func Compact2(slot *Slot) {

	for last := Last(slot); last != nil; last = last.Prev {

		if last.File < 0 {
			continue
		}

		free := slot
		for ; free != nil && free.File != -1; free = free.Next {
			if free == last {
				free = nil
				break
			}
		}
		if free == nil {
			continue
		}

		for ; ; free = free.Next {
			if free == last {
				free = nil
				break
			}
			if free == nil {
				break
			}
			if free.File >= 0 {
				continue
			}
			if free.Length < last.Length {
				continue
			}
			break
		}
		if free == nil {
			continue
		}

		// Copy bytes
		if last.Length < free.Length {
			// Copy all slot into the free
			// Create another free slot with remaining
			newFree := &Slot{
				File:   -1,
				Length: free.Length - last.Length,
				Next:   free.Next,
			}
			if free.Next != nil {
				free.Next.Prev = newFree
			}
			free.File = last.File
			free.Length = last.Length
			free.Next = newFree
			newFree.Prev = free
			last.File = -1
		} else {
			free.File = last.File
			last.File = -1
		}
		MergeFree(slot)
	}
}

func CheckSum(slot *Slot) int {
	index := 0
	sum := 0
	for item := slot; item != nil; item = item.Next {
		if item.File < 0 {
			index = index + item.Length
			continue
		}
		for j := 0; j < item.Length; j++ {
			sum = sum + (index+j)*item.File
		}
		index = index + item.Length
	}
	return sum
}

func GrabInput(fileName string) *Slot {
	scanner := utility.ScanFile(fileName)
	scanner.Scan()
	row := scanner.Text()
	return CreateDisk([]byte(row), 0, true)
}

func Part1(fileName string) {
	slot := GrabInput(fileName)
	Compact(slot)
	log.Println(CheckSum(slot))
}

func Part2(fileName string) {
	slot := GrabInput(fileName)
	Compact2(slot)
	log.Println(CheckSum(slot))
}

func main() {
	Part1("cmd/es9/test.txt")
	Part1("cmd/es9/input.txt")
	Part2("cmd/es9/test.txt")
	Part2("cmd/es9/input.txt")
}

package main

import (
	"fmt"
)

func main() {
	program := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 6, 19, 1, 19, 5, 23, 2, 13, 23, 27, 1, 10, 27, 31, 2, 6, 31, 35, 1, 9, 35, 39, 2, 10, 39, 43, 1, 43, 9, 47, 1, 47, 9, 51, 2, 10, 51, 55, 1, 55, 9, 59, 1, 59, 5, 63, 1, 63, 6, 67, 2, 6, 67, 71, 2, 10, 71, 75, 1, 75, 5, 79, 1, 9, 79, 83, 2, 83, 10, 87, 1, 87, 6, 91, 1, 13, 91, 95, 2, 10, 95, 99, 1, 99, 6, 103, 2, 13, 103, 107, 1, 107, 2, 111, 1, 111, 9, 0, 99, 2, 14, 0, 0}
	result := program[0]
	var noun, verb int
loop:
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			result = runProgram(restoreProgam(program, noun, verb))[0]
			if result == 19690720 {
				fmt.Printf("NOUN %d and VERB %d => %d", noun, verb, 100*noun+verb)
				break loop
			}
		}
	}
}

func restoreProgam(program []int, noun int, verb int) []int {
	newProgram := make([]int, len(program))
	copy(newProgram, program)
	newProgram[1] = noun
	newProgram[2] = verb
	return newProgram
}

func runProgram(program []int) []int {
	instructionPointer := 0

	for program[instructionPointer] != 99 {
		p1 := program[instructionPointer+1]
		p2 := program[instructionPointer+2]
		p3 := program[instructionPointer+3]
		switch program[instructionPointer] {
		case 1:
			program[p3] = program[p1] + program[p2]
		case 2:
			program[p3] = program[p1] * program[p2]
		}
		instructionPointer += 4
	}

	return program
}

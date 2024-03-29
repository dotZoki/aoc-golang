package main

import (
	"reflect"
	"testing"
)

func TestRunProgram(t *testing.T) {
	programs := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, program := range programs {
		if !reflect.DeepEqual(runProgram(program.input), program.output) {
			t.Errorf("%v should be %v but it is %v", program.input, program.output, runProgram(program.input))
		}
	}
}
